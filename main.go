package main

import (
	"github.com/AdamSiekierski/go-adam/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api.MainGroup(e)

	e.Logger.Fatal(e.Start(":1337"))
}
