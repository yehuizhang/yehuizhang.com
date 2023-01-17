package logger

import "go.uber.org/zap"

type Logger = zap.SugaredLogger

func InitLogger() *Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	return sugar
}
