package requests

type CreateProductRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	PhotoUrl    string   `json:"photoUrl"`
	Price       float64  `json:"price"`
}
