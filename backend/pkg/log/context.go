package log

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type key struct{}

func SetLogCtx(ctx context.Context) context.Context {
	logger, _ := zap.NewProduction()

	rid := uuid.NewString()

	log := logger.With(zap.String("rid", rid))

	return context.WithValue(ctx, key{}, log)
}

func GetLogCtx(ctx context.Context) *zap.Logger {
	v := ctx.Value(key{})

	log, ok := v.(*zap.Logger)

	if !ok {
		logger, _ := zap.NewProduction()

		return logger
	}

	return log
}
