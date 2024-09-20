package server

import (
	"personal-portfolio/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	// add dependencies initialiser if needed
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {

	app := echo.New()

	app.Use(middleware.Gzip())

	// Register routes
	routes.RegisterRoutes(app)

	return app.Start(":8080")
}
