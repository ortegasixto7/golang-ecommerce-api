package validations

import (
	"errors"

	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/requests"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/validations/customErrors"
)

type CreateProductRequestValidation struct{}

func (item CreateProductRequestValidation) Validate(request *requests.CreateProductRequest) error {
	if request.Name == "" {
		return errors.New(customErrors.NAME_IS_REQUIRED)
	}
	if request.Price == 0 {
		return errors.New(customErrors.PRICE_IS_REQUIRED)
	}
	if request.Categories == nil {
		request.Categories = []string{}
	}
	return nil
}
