package user

import (
	"net/http"

	"github.com/labstack/echo"
)

type handler struct {
	service *service
}

func NewHandler(s *service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) CreateUser(c echo.Context) error {
	h.service.CreateUser()
	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})

}

func (h *handler) DeleteUser(c echo.Context) error {
	h.service.DeleteUser()
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
