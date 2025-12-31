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
	var user User

	// 1️⃣ Bind request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
		})
	}

	// 2️⃣ Call service with data
	result, err := h.service.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create user",
		})
	}

	// 3️⃣ Return response
	return c.JSON(http.StatusCreated, result)

}

func (h *handler) DeleteUser(c echo.Context) error {
	h.service.DeleteUser()
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
