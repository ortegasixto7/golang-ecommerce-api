package requests

type AddProductQuantityRequest struct {
	Id       string `json:"id"`
	Quantity int32  `json:"quantity"`
}
