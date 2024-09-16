package main

import (
	"personal-portfolio/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Serve static files
	e.Static("/css", "public/css")
	e.Static("/js", "assets/js")

	// Register routes
	routes.Register(e)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
