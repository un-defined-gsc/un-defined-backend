package main

import (
	"github.com/un-defined-gsc/un-defined-backend/internal/app"
	"github.com/un-defined-gsc/un-defined-backend/internal/config"
)

// @title API Service
// @description API Service for ProjectMonWeb
// @version v1
// @host 127.0.0.1:8080
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in cookie
// @name session_id
func main() {
	err := config.InitializeConfig("config/config.yaml")
	if err != nil {
		panic(err)
	}
	cfg := config.GetConfig()
	app.Run(&cfg)
}
