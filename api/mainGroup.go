package api

import (
	"github.com/AdamSiekierski/go-adam/api/handlers"
	"github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
	e.GET("/", handlers.Home)
}