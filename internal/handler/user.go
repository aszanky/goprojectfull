package handler

import (
	"net/http"

	"github.com/aszanky/goprojectfull/internal/models"
	"github.com/labstack/echo/v4"
)

// Handler
func User(c echo.Context) error {
	data := &models.User{
		Name:  "Asaduddin",
		Email: "asaduddin@yahoo.com",
	}

	return c.JSON(http.StatusOK, data)
}
