package main

import (
	"net/http"

	"github.com/galenchurch/hmi/routes"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func main() {
	app := echo.New()

	app.Use(mw.LoggerWithConfig(mw.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	app.Use(mw.StaticWithConfig(mw.StaticConfig{
		Root: "public",
	}))

	app.GET("/ping", func(cxt echo.Context) error {
		return cxt.String(http.StatusOK, "pong")
	})

	app.POST("/led/:led", routes.LedHandler)

	err := app.Start(":9000")
	if err != nil {
		app.Logger.Fatal(err)
	}
}
