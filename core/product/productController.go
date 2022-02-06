package product

import (
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/requests"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/validations"
)

type ProductController struct {
	ProductService ProductService
}

func (this ProductController) GetById(id string) Product {
	return this.ProductService.GetById(id)
}

func (this ProductController) GetAll() []Product {
	return this.ProductService.GetAll()
}

func (this ProductController) Create(request *requests.CreateProductRequest) (requestError string, errorCode string) {
	requestError, errorCode = new(validations.CreateProductRequestValidation).Validate(request)
	if requestError != "" {
		return requestError, errorCode
	}
	product := Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Quantity:    0}
	this.ProductService.Save(product)
	return requestError, errorCode
}
