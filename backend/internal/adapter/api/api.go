package api

import "github.com/morning-night-guild/platform/pkg/openapi"

var _ openapi.ServerInterface = (*API)(nil)

type API struct {
	connect *Connect
}

func New(
	connect *Connect,
) *API {
	return &API{
		connect: connect,
	}
}
