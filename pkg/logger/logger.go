package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{logger: logger}
}

func (l *Logger) Info(message string, fields ...zap.Field) {
	l.logger.Info(message, fields...)
}

func (l *Logger) Debug(message string, fields ...zap.Field) {
	l.logger.Debug(message, fields...)
}

func (l *Logger) Error(message string, fields ...zap.Field) {
	l.logger.Error(message, fields...)
}
