package controller

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	me "github.com/morning-night-guild/platform/internal/model/errors"
	"github.com/morning-night-guild/platform/pkg/log"
)

var errInternal = errors.New("internal server")

var ErrInternal = connect.NewError(
	connect.CodeInternal,
	errInternal,
)

var errInvalidArgument = errors.New("bad request")

var ErrInvalidArgument = connect.NewError(
	connect.CodeInvalidArgument,
	errInvalidArgument,
)
var errUnauthorized = errors.New("unauthorized")

var ErrUnauthorized = connect.NewError(
	connect.CodeUnauthenticated,
	errUnauthorized,
)

// handleError 発生したエラーを対応するgrpcのステータス込みのエラーに変換する関数.
func handleError(ctx context.Context, err error) error {
	logger := log.GetLogCtx(ctx)

	switch {
	case
		asValidationError(err),
		asURLError(err):
		logger.Warn(err.Error())

		return ErrInvalidArgument
	default:
		logger.Error(err.Error())

		return ErrInternal
	}
}

func asValidationError(err error) bool {
	var target me.ValidationError

	return errors.As(err, &target)
}

func asURLError(err error) bool {
	var target me.URLError

	return errors.As(err, &target)
}
