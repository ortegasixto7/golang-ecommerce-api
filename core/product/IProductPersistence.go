package product

type IProductPersistence interface {
	Save(*Product)
}
