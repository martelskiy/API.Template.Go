package lifecycle

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/api-template-go/features/shared/logger"
)

var log = logger.Get()

func ListenForApplicationShutDown(shutdownFunc func()) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	sig := <-signalChannel
	switch sig {
	case os.Interrupt, syscall.SIGTERM:
		shutdownFunc()
	}
}

func StopApplication(message string) {
	log.With(message).Error("stopping application")
	os.Exit(-1)
}
