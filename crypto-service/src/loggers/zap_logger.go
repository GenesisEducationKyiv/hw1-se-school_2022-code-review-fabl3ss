package loggers

import (
	"log"

	"go.uber.org/zap"
)

type zapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger(logPath string) *zapLogger {
	logger, err := setupDefaultZapLogger(logPath)
	if err != nil {
		log.Println(err.Error(), "unable to create new zap logger")
		return nil
	}

	return &zapLogger{
		logger: logger.Sugar(),
	}
}

func (l *zapLogger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *zapLogger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *zapLogger) Error(msg string) {
	l.logger.Error(msg)
}

func setupDefaultZapLogger(logPath string) (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{logPath}
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
