package server

import (
	"repository-go/internal/user"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// root composition can be done here or in main.go based on the complexity of the application.

func StartHTTPServer(e *echo.Echo, db *gorm.DB) {

	user.RegisterRoutes(e, db)

	e.Logger.Fatal(e.Start(":5000"))

}
