package controller

import (
	"errors"
	"log"

	"github.com/bufbuild/connect-go"
	me "github.com/morning-night-guild/platform/app/core/model/errors"
)

var ErrInternal = connect.NewError(
	connect.CodeInternal,
	errors.New("internal server"),
)

var ErrInvalidArgument = connect.NewError(
	connect.CodeInvalidArgument,
	errors.New("bad request"),
)

var ErrUnauthorized = connect.NewError(
	connect.CodeUnauthenticated,
	errors.New("unauthorized"),
)

// handleError 発生したエラーを対応するgrpcのステータス込みのエラーに変換する関数.
func handleError(err error) error {
	log.Printf("error: %v", err)

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
