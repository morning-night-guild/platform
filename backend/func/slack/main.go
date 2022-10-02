package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	"github.com/bufbuild/connect-go"
	articlev1 "github.com/morning-night-guild/platform/pkg/connect/article/v1"
	"github.com/morning-night-guild/platform/pkg/connect/article/v1/articlev1connect"
	"github.com/pkg/errors"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const (
	shutdownTime      = 10
	readHeaderTimeout = 30 * time.Second
)

func main() {
	secret := os.Getenv("SLACK_SIGNING_SECRET")
	if len(secret) == 0 {
		log.Fatal("secret no set")
	}

	coreServiceURL := os.Getenv("CORE_SERVICE_URL")
	if len(coreServiceURL) == 0 {
		log.Fatal("service url no set")
	}

	apiKey := os.Getenv("API_KEY")
	if len(apiKey) == 0 {
		log.Fatal("api key no set")
	}

	mux := http.NewServeMux()
	hc := http.DefaultClient
	ac := articlev1connect.NewArticleServiceClient(hc, coreServiceURL)

	mux.Handle("/", SlackHandler{
		Secret: secret,
		ArticleServiceClient: ArticleServiceClient{
			APIKey: apiKey,
			Client: ac,
		},
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	s := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: readHeaderTimeout,
	}

	log.Printf("Server running on %s", s.Addr)

	go func() {
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln("Server closed with error:", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTime*time.Second)

	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Println("Failed to gracefully shutdown:", err)
	}

	log.Println("HTTPServer shutdown")
}

type SlackHandler struct {
	Secret               string
	ArticleServiceClient ArticleServiceClient
}

func (s SlackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// @see https://github.com/slack-go/slack/blob/master/examples/eventsapi/events.go
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := s.verify(r.Header, body, s.Secret); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		s.challenge(w, body)

		return
	}

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent

		log.Printf("receved event type is %s", innerEvent.Type)

		if err := s.handleSlackEvent(r.Context(), innerEvent); err != nil {
			log.Printf("error occurred %v", err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}
}

func (s SlackHandler) handleSlackEvent(ctx context.Context, event slackevents.EventsAPIInnerEvent) error {
	switch ev := event.Data.(type) {
	// @see https://api.slack.com/events/link_shared
	// link_shareのイベントは発火しなかったため一旦断念
	// @see https://api.slack.com/events/message
	case *slackevents.MessageEvent:
		if len(ev.Text) == 0 {
			log.Println("message is empty")

			return nil
		}

		r := regexp.MustCompile(`http(.*)://([a-zA-Z0-9/\-\_\.]*)`)
		u := r.FindString(ev.Text)

		if _, err := url.Parse(u); err != nil {
			return err
		}

		if _, err := s.ArticleServiceClient.Share(ctx, u); err != nil {
			return err
		}
	default:
		// errorを返すとslackがリトライしてくるため
		log.Printf("undefined event %+v", ev)

		return nil
	}

	return nil
}

func (s SlackHandler) verify(header http.Header, body []byte, secret string) error {
	sv, err := slack.NewSecretsVerifier(header, secret)
	if err != nil {
		return errors.Wrap(err, "failed new secrets verify")
	}

	if _, err := sv.Write(body); err != nil {
		return errors.Wrap(err, "failed write body")
	}

	if err := sv.Ensure(); err != nil {
		return errors.Wrap(err, "failed ensure")
	}

	return nil
}

func (s SlackHandler) challenge(w http.ResponseWriter, body []byte) {
	var r *slackevents.ChallengeResponse

	if err := json.Unmarshal(body, &r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "text")

	_, _ = w.Write([]byte(r.Challenge))
}

type ArticleServiceClient struct {
	APIKey string
	Client articlev1connect.ArticleServiceClient
}

func (a ArticleServiceClient) Share(
	ctx context.Context,
	url string,
) (*connect.Response[articlev1.ShareResponse], error) {
	ar := articlev1.ShareRequest{
		Url: url,
	}

	request := connect.Request[articlev1.ShareRequest]{
		Msg: &ar,
	}

	request.Header().Set("X-API-KEY", a.APIKey)

	return a.Client.Share(ctx, &request)
}
