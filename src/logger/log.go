package logger

import (
	"github.com/singhkshitij/golang-rest-service-starter/src/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"strings"
)

var logger *zap.Logger

func Must(logger *zap.Logger, err error) *zap.Logger {
	if err != nil {
		panic(err)
	}
	return logger
}

func NewLogger(logFile string) (*zap.Logger, error) {
	var err error

	loggerConfig := zap.Config{
		Level:    zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "timestamp",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    "func",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths: []string{"stdout"}, //add filepath if req
	}

	logLevel := config.LogLevel()
	level := loggerConfig.Level.Level()
	err = level.Set(logLevel)
	if err != nil {
		log.Fatal("Error while setting log level in logger ", err.Error())
	}
	loggerConfig.Level = zap.NewAtomicLevelAt(level)

	if strings.ToLower(logLevel) == "debug" {
		loggerConfig.Development = true
	}

	logger, err = loggerConfig.Build()
	if err != nil {
		log.Fatal("Error while initialising logger", err.Error())
	}
	return logger, err
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func KV(key string, val interface{}) zap.Field {
	return zap.Any(key, val)
}
