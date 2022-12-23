package main

import (
	"github.com/labstack/echo/v4"

	"github.com/Team-OurPlayground/our-playground-auth/internal/auth/controller"
)

func SetupApp() *echo.Echo {
	echoInstance := echo.New()

	registerRoute(echoInstance)

	return echoInstance
}

func registerRoute(e *echo.Echo) {
	authController := controller.NewAuthController()

	authGroup := e.Group("/auth")
	setAuthGroup(authGroup, authController)
}

func setAuthGroup(group *echo.Group, controller *controller.Auth) {
	group.GET("", controller.HealthCheck)
}
