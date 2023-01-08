package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Team-OurPlayground/our-playground-auth/internal/auth/controller/dto"
	"github.com/Team-OurPlayground/our-playground-auth/internal/auth/service"
	"github.com/Team-OurPlayground/our-playground-auth/internal/util/jwt"
)

type Auth struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *Auth {
	return &Auth{
		authService: authService,
	}
}

func (a *Auth) GetPublicKey(c echo.Context) (err error) {
	return c.String(http.StatusOK, string(jwt.GetRawPublicKey()))
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
	if err != nil {
		return c.String(http.StatusInternalServerError, "GetToken error: SignIn")
	}
	return c.JSON(http.StatusOK, resp)
}

func (a *Auth) Refresh(c echo.Context) (err error) {
	req := new(dto.RefreshRequest)
	if err = c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "request binding error")
	}

	resp, err := a.authService.Refresh(req)
	if err != nil {
		return c.String(http.StatusInternalServerError, "authService.Refresh error: Refresh")
	}

	return c.JSON(http.StatusOK, resp)
}
