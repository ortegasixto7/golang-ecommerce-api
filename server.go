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
	server.PUT("/products/add-quantity", presentation.ProductRouter{}.AddQuantity)
	server.PUT("/products", presentation.ProductRouter{}.Update)
	server.POST("/products", presentation.ProductRouter{}.Create)
	server.GET("/products", presentation.ProductRouter{}.GetAll)
	server.GET("/products/:id", presentation.ProductRouter{}.GetById)
	server.POST("/admin/create-admin-user", presentation.AdminRouter{}.CreateAdminUser)
	server.POST("/auth/login", presentation.AuthRouter{}.Login)
	server.Logger.Fatal(server.Start(":1323"))
}
