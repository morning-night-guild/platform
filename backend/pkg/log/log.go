package log

import "go.uber.org/zap"

func Log() *zap.Logger {
	// デフォルト値でロギング設定をしているので
	// errorは発生しない
	// -> errorはにぎりつぶしても問題ない
	logger, _ := zap.NewProduction()

	return logger
}
