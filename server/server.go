package server

import (
	"personal-portfolio/routes"

	"github.com/labstack/echo/v4"
)

type Server struct {
	// add dependencies initialiser if needed
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {

	app := echo.New()

	// Register routes
	routes.Register(app)

	return app.Start(":8080")
}
