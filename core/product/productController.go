package product

import (
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/requests"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/validations"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	ProductService ProductService
}

func (this ProductController) AddQuantity(request *requests.AddProductQuantityRequest) (errorCode string) {
	// requestError, errorCode = new(validations.UpdateProductRequest).Validate(request)
	// if requestError != "" {
	// 	return requestError, errorCode
	// }
	productResult := this.ProductService.GetById(request.Id)
	productResult.Quantity += request.Quantity
	this.ProductService.Update(productResult)
	return errorCode
}

func (this ProductController) GetById(id string) Product {
	return this.ProductService.GetById(id)
}

func (this ProductController) GetAll() []Product {
	return this.ProductService.GetAll()
}

func (this ProductController) Update(request *requests.UpdateProductRequest) (errorCode string) {
	// requestError, errorCode = new(validations.UpdateProductRequest).Validate(request)
	// if requestError != "" {
	// 	return requestError, errorCode
	// }
	productResult := this.ProductService.GetById(request.Id)
	productResult.Name = request.Name
	productResult.Description = request.Description
	productResult.Price = request.Price
	productResult.Categories = request.Categories
	productResult.PhotoUrl = request.PhotoUrl
	this.ProductService.Update(productResult)
	return errorCode
}

func (this ProductController) Create(request *requests.CreateProductRequest) (errorCode string) {
	errorCode = new(validations.CreateProductRequestValidation).Validate(request)
	if errorCode != "" {
		return errorCode
	}
	product := Product{
		Id:          primitive.NewObjectID().Hex(),
		Name:        request.Name,
		Description: request.Description,
		Categories:  request.Categories,
		PhotoUrl:    request.PhotoUrl,
		Price:       request.Price,
		Quantity:    0}
	this.ProductService.Save(product)
	return errorCode
}
