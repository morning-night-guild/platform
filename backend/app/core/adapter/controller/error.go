package controller

import (
	"context"
	"errors"

	me "github.com/morning-night-guild/platform/app/core/model/errors"
	"github.com/morning-night-guild/platform/pkg/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrInternal = status.New(
	codes.Internal,
	"internal server",
).Err()

var ErrInvalidArgument = status.New(
	codes.InvalidArgument,
	"bad request",
).Err()

var ErrUnauthorized = status.New(
	codes.Unauthenticated,
	"unauthorized",
).Err()

// handleError 発生したエラーを対応するgrpcのステータス込みのエラーに変換する関数.
func handleError(ctx context.Context, err error) error {
	log := log.GetLogCtx(ctx)

	switch {
	case asValidationError(err):
		log.Warn(err.Error())

		return ErrInvalidArgument
	default:
		log.Error(err.Error())

		return ErrInternal
	}
}

func asValidationError(err error) bool {
	var target me.ValidationError

	return errors.As(err, &target)
}
