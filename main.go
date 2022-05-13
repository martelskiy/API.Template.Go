package main

import (
	"context"

	"github.com/api-template-go/features/healthcheck"
	"github.com/api-template-go/features/shared/api/host"
	"github.com/api-template-go/features/shared/api/route"
	"github.com/api-template-go/features/shared/lifecycle"
	"github.com/api-template-go/features/shared/logger"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           API Swagger
// @version         1.0
func main() {
	context := context.Background()
	log := logger.Initialize()

	router := route.NewRouter()
	router.
		WithRoute(route.NewRoute("/swagger/", httpSwagger.WrapHandler)).
		WithRoute(route.NewRoute("/status", healthcheck.GetStatus))

	host := host.NewHost("80", router)
	err := host.RunAsync()
	if err != nil {
		lifecycle.StopApplication("error running web host")
	}

	lifecycle.ListenForApplicationShutDown(func() {
		log.Info("shutdown signal received. Terminating the web host")
		host.Terminate(context)
	})
}
