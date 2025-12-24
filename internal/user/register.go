package user

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	e.POST("/users", handler.CreateUser)
	e.DELETE("/users/:id", handler.DeleteUser)
}
