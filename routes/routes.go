package routes

import (
	"net/http"
	"personal-portfolio/handlers"
	"personal-portfolio/views"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	// Handle static assets:
	assetHandler := http.FileServer(views.GetPublicAssetsFileSystem())
	e.GET("/public/*", echo.WrapHandler(http.StripPrefix("/public/", assetHandler)))

	// Define routes
	e.GET("/", handlers.Home)
	e.POST("/count/inc", handlers.CountIncrement)
	e.POST("/count/dec", handlers.CountDecrement)
}
