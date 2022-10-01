package loggers

import (
	"genesis_test_case/src/pkg/domain/logger"
	"genesis_test_case/src/pkg/types/filemodes"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger(logPath string) logger.Logger {
	core := setupDefaultZapCore(logPath)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger.Sugar()
}

func setupDefaultZapCore(logPath string) zapcore.Core {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	fileMode := os.ModeDir | (filemodes.OS_USER_RW | filemodes.OS_ALL_R)
	logFile, _ := os.OpenFile(logPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		fileMode,
	)
	writer := zapcore.AddSync(logFile)

	return zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, zapcore.DebugLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)
}
