package product

import "fmt"

type IProductService interface {
	Save(Product)
	Update(*Product)
	Delete(id string)
	GetAll() []Product
	GetById(id string) Product
}

type ProductService struct {
	Repository IProductPersistence
}

func NewProductService(repository IProductPersistence) *ProductService {
	return &ProductService{Repository: repository}
}

func (this ProductService) Save(product Product) {
	this.Repository.Save(product)
}

func (this ProductService) Update(product Product) {
	fmt.Print("Product Updated")
}

func (this ProductService) Delete(id string) {
	fmt.Print("Product Deleted")
}

func (this ProductService) GetAll() []Product {
	return []Product{}
}

func (this ProductService) GetById(id string) Product {
	return Product{}
}
