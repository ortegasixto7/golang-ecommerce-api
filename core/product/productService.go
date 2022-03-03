package product

type IProductService interface {
	Save(Product)
	Update(Product)
	Delete(id string)
	GetAll() []Product
	GetById(id string) Product
}

type ProductService struct {
	Persistence IProductPersistence
}

func NewProductService(persistence IProductPersistence) *ProductService {
	return &ProductService{Persistence: persistence}
}

func (this ProductService) Save(product Product) {
	this.Persistence.Save(product)
}

func (this ProductService) Update(product Product) {
	this.Persistence.Update(product)
}

func (this ProductService) Delete(id string) {
	this.Persistence.Delete(id)
}

func (this ProductService) GetAll() []Product {
	return this.Persistence.GetAll()
}

func (this ProductService) GetById(id string) Product {
	return this.Persistence.GetById(id)
}
