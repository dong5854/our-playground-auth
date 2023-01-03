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
		return c.String(http.StatusBadRequest, "request binding error")
	}

	err = a.authService.SignUp(req)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.String(http.StatusOK, "OK")
}

func (a *Auth) SignIn(c echo.Context) (err error) {
	req := new(dto.SignInRequest)
	if err = c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "request binding error")
	}

	ok, err := a.authService.SignIn(req)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if !ok {
		return c.String(http.StatusUnauthorized, "failed to sign in")
	}
	resp, err := a.authService.GetToken(req.Email)
	print(resp)
	return c.JSON(http.StatusOK, resp)
}
