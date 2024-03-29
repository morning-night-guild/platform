package api

import (
	"net/http"

	"github.com/bufbuild/connect-go"
	healthv1 "github.com/morning-night-guild/platform/pkg/connect/proto/health/v1"
	"github.com/morning-night-guild/platform/pkg/log"
)

func (api *API) V1HealthAPI(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func (api *API) V1HealthCore(w http.ResponseWriter, r *http.Request) {
	ctx := log.SetLogCtx(r.Context())

	req := &healthv1.CheckRequest{}

	if _, err := api.connect.Health.Check(ctx, connect.NewRequest(req)); err != nil {
		log.GetLogCtx(ctx).Error("failed to check health core", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}
