package product

type ProductPersistence interface {
	Save(product Product)
	Update(*Product)
	Delete(id string)
	GetAll() []Product
	GetById(id string) Product
}
