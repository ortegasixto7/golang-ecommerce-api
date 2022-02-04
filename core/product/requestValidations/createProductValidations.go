package requestValidations

import (
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/requestModels"
	"github.com/ortegasixto7/echo-go-supermarket-api/external/validations/customErrors"
)

type CreateProductRequestValidation struct{}

func (item *CreateProductRequestValidation) Validate(request requestModels.CreateProductRequest) (requestError customErrors.RequestErrorEnum, errorCode customErrors.ErrorCodeEnum) {
	if request.Name == "" {
		return customErrors.BAD_REQUEST, customErrors.NAME_IS_MANDATORY
	}
	if request.Price == 0 {
		return customErrors.BAD_REQUEST, customErrors.PRICE_IS_MANDATORY
	}
	return "", ""
}
