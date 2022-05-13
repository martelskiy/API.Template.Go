package main

import (
	"context"

	"github.com/api-template-go/features/healthcheck"
	"github.com/api-template-go/features/shared/api/host"
	"github.com/api-template-go/features/shared/api/route"
	"github.com/api-template-go/features/shared/lifecycle"
	"github.com/api-template-go/features/shared/logger"
)

func main() {
	context := context.Background()
	log := logger.Initialize()

	router := route.NewRouter()
	router.
		WithAPIDocumentation().
		WithRoute(route.NewRoute("/status", healthcheck.GetStatus))

	host := host.NewHost("8080", router)
	err := host.RunAsync()
	if err != nil {
		lifecycle.StopApplication("error running web host")
	}

	lifecycle.ListenForApplicationShutDown(func() {
		log.Info("shutdown signal received. Terminating the web host")
		host.Terminate(context)
	})
}
