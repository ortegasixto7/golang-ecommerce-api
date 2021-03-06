package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/requests"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/dependencyInjector"
)

type ProductRouter struct{}

func (this ProductRouter) AddQuantity(c echo.Context) error {
	request := new(requests.AddProductQuantityRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	controller := new(dependencyInjector.ContainerBuilder).GetProductController()
	errorCode := controller.AddQuantity(request)
	if errorCode != nil {
		return BuildInvalidResponse(c, errorCode)
	}
	return c.JSON(http.StatusOK, request)
}

func (this ProductRouter) Update(c echo.Context) error {
	request := new(requests.UpdateProductRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	controller := new(dependencyInjector.ContainerBuilder).GetProductController()
	errorCode := controller.Update(request)
	if errorCode != nil {
		return BuildInvalidResponse(c, errorCode)
	}
	return c.JSON(http.StatusOK, request)
}

func (this ProductRouter) Create(c echo.Context) error {
	request := new(requests.CreateProductRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	controller := new(dependencyInjector.ContainerBuilder).GetProductController()
	errorCode := controller.Create(request)
	if errorCode != nil {
		return BuildInvalidResponse(c, errorCode)
	}
	return c.JSON(http.StatusOK, request)
}

func (this ProductRouter) GetAll(c echo.Context) error {
	controller := new(dependencyInjector.ContainerBuilder).GetProductController()
	products := controller.GetAll()
	return c.JSON(http.StatusOK, products)
}

func (this ProductRouter) GetById(c echo.Context) error {
	controller := new(dependencyInjector.ContainerBuilder).GetProductController()
	id := c.Param("id")
	product := controller.GetById(id)
	return c.JSON(http.StatusOK, product)
}
