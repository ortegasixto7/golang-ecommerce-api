package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ortegasixto7/echo-go-supermarket-api/persistence/mongoDb"
	"github.com/ortegasixto7/echo-go-supermarket-api/presentation"
)

func main() {
	server := echo.New()
	mongoDb.Setup()
	server.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "API is working!")
	})
	server.PUT("/products/add-quantity", func(c echo.Context) error {
		return presentation.ProductRouter{}.AddQuantity(c)
	})
	server.PUT("/products", func(c echo.Context) error {
		return presentation.ProductRouter{}.Update(c)
	})
	server.POST("/products", func(c echo.Context) error {
		return presentation.ProductRouter{}.Create(c)
	})
	server.GET("/products", func(c echo.Context) error {
		return presentation.ProductRouter{}.GetAll(c)
	})
	server.GET("/products/:id", func(c echo.Context) error {
		return presentation.ProductRouter{}.GetById(c)
	})
	server.Logger.Fatal(server.Start(":1323"))
}
