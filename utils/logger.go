package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapConfig() zap.Config {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.MessageKey = "message"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return config
}
