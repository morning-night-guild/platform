package log

import "go.uber.org/zap"

func Log() *zap.Logger {
	logger, _ := zap.NewProduction()

	return logger
}
