package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lz *zap.SugaredLogger

func initLogger() (*zap.Logger, error) {
	conf := zap.NewDevelopmentConfig()
	conf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := conf.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func InitLogger() error {
	logger, err := initLogger()
	if err != nil {
		return fmt.Errorf("cannot init a logger, error: %w", err)
	}
	lz = logger.Sugar()
	return nil
}

func Log() *zap.SugaredLogger {
	return lz
}
