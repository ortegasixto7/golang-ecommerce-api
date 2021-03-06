package product

import (
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/requests"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/validations"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	ProductPersistence IProductPersistence
}

func (this ProductController) AddQuantity(request *requests.AddProductQuantityRequest) error {
	// requestError, errorCode = new(validations.UpdateProductRequest).Validate(request)
	// if requestError != "" {
	// 	return requestError, errorCode
	// }
	productResult := this.ProductPersistence.GetById(request.Id)
	productResult.Quantity += request.Quantity
	this.ProductPersistence.Update(productResult)
	return nil
}

func (this ProductController) GetById(id string) Product {
	return this.ProductPersistence.GetById(id)
}

func (this ProductController) GetAll() []Product {
	return this.ProductPersistence.GetAll()
}

func (this ProductController) Update(request *requests.UpdateProductRequest) error {
	// requestError, errorCode = new(validations.UpdateProductRequest).Validate(request)
	// if requestError != "" {
	// 	return requestError, errorCode
	// }
	productResult := this.ProductPersistence.GetById(request.Id)
	productResult.Name = request.Name
	productResult.Description = request.Description
	productResult.Price = request.Price
	productResult.Categories = request.Categories
	productResult.PhotoUrl = request.PhotoUrl
	this.ProductPersistence.Update(productResult)
	return nil
}

func (this ProductController) Create(request *requests.CreateProductRequest) error {
	errorCode := new(validations.CreateProductRequestValidation).Validate(request)
	if errorCode != nil {
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
	this.ProductPersistence.Save(product)
	return errorCode
}
