package logger

import (
	"go.uber.org/zap"
	"yehuizhang.com/go-webapp-gin/pkg/flags"
)

type Logger = zap.SugaredLogger

func InitLogger(flags *flags.FlagParser) *Logger {

	var logger *zap.Logger
	switch flags.Env {
	case "prod", "production":
		logger, _ = zap.NewProduction()
	case "test":
		config := zap.NewDevelopmentConfig()
		config.Level.SetLevel(zap.PanicLevel)
		config.EncoderConfig.TimeKey = ""
		logger, _ = config.Build()
	default:
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	return sugar
}
