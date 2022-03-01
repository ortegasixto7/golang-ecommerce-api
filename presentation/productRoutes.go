package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/requests"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/dependencyInjector"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/validations/customErrors"
)

type ProductRouter struct{}

func (this ProductRouter) AddQuantity(c echo.Context) error {
	request := new(requests.AddProductQuantityRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	controller := new(dependencyInjector.ContainerBuilder).GetProductController()
	requestError, errorCode := controller.AddQuantity(request)
	if requestError != "" {
		return this.buildInvalidResponse(c, requestError, errorCode)
	}
	return c.JSON(http.StatusOK, request)
}

func (this ProductRouter) Update(c echo.Context) error {
	request := new(requests.UpdateProductRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	controller := new(dependencyInjector.ContainerBuilder).GetProductController()
	requestError, errorCode := controller.Update(request)
	if requestError != "" {
		return this.buildInvalidResponse(c, requestError, errorCode)
	}
	return c.JSON(http.StatusOK, request)
}

func (this ProductRouter) Create(c echo.Context) error {
	request := new(requests.CreateProductRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	controller := new(dependencyInjector.ContainerBuilder).GetProductController()
	requestError, errorCode := controller.Create(request)
	if requestError != "" {
		return this.buildInvalidResponse(c, requestError, errorCode)
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

func (this ProductRouter) buildInvalidResponse(c echo.Context, requestError string, errorCode string) error {
	if requestError == string(customErrors.BAD_REQUEST) {
		return c.JSON(http.StatusBadRequest, customErrors.CustomResponse{ErrorCode: string(errorCode)})
	}
	if requestError == string(customErrors.NOT_FOUND) {
		return c.JSON(http.StatusNotFound, customErrors.CustomResponse{ErrorCode: string(errorCode)})
	}
	return c.JSON(http.StatusNotFound, customErrors.CustomResponse{ErrorCode: string(errorCode)})
}