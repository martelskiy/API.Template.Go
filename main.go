package main

import (
	"github.com/api-template-go/features/shared/api"
	"github.com/api-template-go/features/shared/logger"
)

// @title           API Swagger
// @version         1.0
func main() {
	logger.Initialize()
	api.RunWebHost()
}
