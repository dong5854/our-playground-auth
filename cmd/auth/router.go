package main

import (
	"github.com/labstack/echo/v4"

	"github.com/Team-OurPlayground/our-playground-auth/internal/auth/controller"
	"github.com/Team-OurPlayground/our-playground-auth/internal/auth/repository/entgo"
	"github.com/Team-OurPlayground/our-playground-auth/internal/auth/service"
	"github.com/Team-OurPlayground/our-playground-auth/internal/config"
)

func SetupApp() *echo.Echo {
	echoInstance := echo.New()

	registerRoute(echoInstance)

	return echoInstance
}

func registerRoute(e *echo.Echo) {
	userRepository := entgo.NewUserRepository(config.GetEntClient())
	authService := service.NewAuthService(userRepository)
	authController := controller.NewAuthController(authService)

	userGroup := e.Group("/users")
	setUserGroup(userGroup, authController)
}

func setUserGroup(group *echo.Group, controller *controller.Auth) {
	group.POST("", controller.SignUp)
	group.POST("/sign-in", controller.SignIn)
}
