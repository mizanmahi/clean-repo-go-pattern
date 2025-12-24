package main

import (
	"repository-go/internal/config"
	"repository-go/internal/server"

	"github.com/labstack/echo"
)

func main() {
	cfg := config.LoadEnv()
	db := config.DatabaseConnection(cfg)
	e := echo.New()

	server.StartHTTPServer(e, db)
}
