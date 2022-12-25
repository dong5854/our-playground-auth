package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Team-OurPlayground/our-playground-auth/internal/auth/controller/dto"
	"github.com/Team-OurPlayground/our-playground-auth/internal/auth/service"
)

type Auth struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *Auth {
	return &Auth{
		authService: authService,
	}
}

func (a *Auth) SignUp(c echo.Context) (err error) {
	req := new(dto.SignUpRequest)
	if err = c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err = a.authService.SignUp(req)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.String(http.StatusOK, "OK")
}

func (a *Auth) SignIn(c echo.Context) (err error) {
	return nil
}
