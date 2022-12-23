package main

const port = "8080"

func main() {
	echoInstance := SetupApp()

	echoInstance.Logger.Fatal(echoInstance.Start(":" + port))
}
