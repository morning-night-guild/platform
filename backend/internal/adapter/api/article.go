package api

import (
	"encoding/json"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	articlev1 "github.com/morning-night-guild/platform/pkg/connect/proto/article/v1"
	"github.com/morning-night-guild/platform/pkg/openapi"
)

func (api API) V1ListArticles(w http.ResponseWriter, r *http.Request, params openapi.V1ListArticlesParams) {
	pageToken := ""
	if params.PageToken != nil {
		pageToken = *params.PageToken
	}

	req := &articlev1.ListRequest{
		PageToken:   pageToken,
		MaxPageSize: uint32(params.MaxPageSize),
	}

	res, err := api.connect.Article.List(r.Context(), connect.NewRequest(req))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	articles := make([]openapi.Article, len(res.Msg.Articles))

	for i, article := range res.Msg.Articles {
		uid, _ := uuid.Parse(article.Id)

		articles[i] = openapi.Article{
			Id:          &uid,
			Title:       &article.Title,
			Url:         &article.Url,
			Description: &article.Description,
			Thumbnail:   &article.Thumbnail,
			Tags:        &article.Tags,
		}
	}

	rs := openapi.ListArticleResponse{
		Articles:      &articles,
		NextPageToken: &res.Msg.NextPageToken,
	}

	if err := json.NewEncoder(w).Encode(rs); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
