package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main2() {
	server := echo.New()
	server.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello, World!")
	})
	server.Logger.Fatal(server.Start(":1323"))
}
