package handlers

import (
	"personal-portfolio/views"

	"github.com/labstack/echo/v4"
)

func Home(ctx echo.Context) error {
	return EchoRender(ctx, views.Home("go-echo-htmx-templ", Count))
}
