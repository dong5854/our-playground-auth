package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Auth struct {
}

func NewAuthController() *Auth {
	return &Auth{}
}

func (a *Auth) HealthCheck(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Hello, Auth!")
}
