package logger

import (
	"log"

	"go.uber.org/zap"
)

type ILogger interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
}

func NewLogger() ILogger {
	newLogger := logger{}

	logger, _ := zap.NewProduction()
	newLogger.zapSugar = logger.Sugar()

	return &newLogger
}

type logger struct {
	zapSugar *zap.SugaredLogger
}

func (logger logger) Debug(args ...any) {
	logger.zapSugar.Debug(args)
}

func (logger logger) Info(args ...any) {
	logger.zapSugar.Info(args)
}

func (logger logger) Warn(args ...any) {
	logger.zapSugar.Warn(args)
}

func (logger logger) Error(args ...any) {
	logger.zapSugar.Error(args)
}

func (logger logger) FlushLog() {
	err := logger.zapSugar.Sync()
	if err != nil {
		nativePrint(err)
	}
}

func nativePrint(msg ...any) {
	log.Println(msg)
}
