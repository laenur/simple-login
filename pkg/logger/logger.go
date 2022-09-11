package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Setup() {
	configLogger := zap.NewDevelopmentConfig()
	configLogger.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	configLogger.DisableStacktrace = true
	logger, _ := configLogger.Build()
	zap.ReplaceGlobals(logger)
}

func Info(msg string, keysAndValues ...interface{}) {
	zap.S().Infow(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	zap.S().Errorw(msg, keysAndValues...)
}
