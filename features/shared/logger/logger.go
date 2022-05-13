package logger

import "go.uber.org/zap"

var sugaredLogger *zap.SugaredLogger
var isInitialized = false

func Initialize() {
	log, _ := zap.NewProduction()
	sugaredLogger = log.Sugar()
	isInitialized = true
}

func Get() *zap.SugaredLogger {
	if !isInitialized {
		logger, _ := zap.NewProduction()
		sugaredLogger = logger.Sugar()
		isInitialized = true
	}
	return sugaredLogger
}

func Dispose() {
	if sugaredLogger != nil {
		defer sugaredLogger.Sync()
	}
}
