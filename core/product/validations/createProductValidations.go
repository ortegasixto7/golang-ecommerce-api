package validations

import (
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/requests"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/validations/customErrors"
)

type CreateProductRequestValidation struct{}

func (item CreateProductRequestValidation) Validate(request *requests.CreateProductRequest) (errorCode string) {
	if request.Name == "" {
		return string(customErrors.NAME_IS_MANDATORY)
	}
	if request.Price == 0 {
		return string(customErrors.PRICE_IS_MANDATORY)
	}
	return ""
}
