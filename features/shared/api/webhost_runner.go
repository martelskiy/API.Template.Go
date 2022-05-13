package api

import (
	"net/http"

	"github.com/api-template-go/features/healthcheck"
	"github.com/api-template-go/features/shared/logger"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	port = ":8080"
)

func RunWebHost() {
	log := logger.Get()
	router := mux.NewRouter().StrictSlash(true)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.HandleFunc("/status", healthcheck.GetStatus)

	log.Infof("web server listening at: '%s'", port)

	http.ListenAndServe(port, router)
}
