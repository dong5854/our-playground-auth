package main

import (
	"github.com/Team-OurPlayground/our-playground-auth/internal/config"
)

const port = "8080"

func main() {
	config.EntMysqlInitialize()
	config.InitJWTKeys()
	client := config.GetEntClient()
	defer client.Close()

	echoInstance := SetupApp()

	echoInstance.Logger.Fatal(echoInstance.Start(":" + port))
}
