package main

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"github.com/tediferous/wordmark/html"
)

func main() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
		return html.Page().Render(c.Request().Context(), c.Response())
	})

	if err := e.Start(":1717"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
