package controller

import (
	"errors"

	me "github.com/morning-night-guild/platform/model/errors"
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
func handleError(err error) error {
	switch {
	case asValidationError(err):
		return ErrInvalidArgument
	default:
		return ErrInternal
	}
}

func asValidationError(err error) bool {
	var target me.ValidationError

	return errors.As(err, &target)
}
