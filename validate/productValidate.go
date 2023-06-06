package validate

import (
	"e-commerce/model"

	"github.com/go-playground/validator/v10"
)



func ValidateRegisterProduct(product *model.RegisterUpdateProduct) (err error) {
	validate := validator.New()
	CreateCustomValidations(validate)
	err = validate.Struct(product)
	return err
}

func ValidateUpdateProduct(product *model.RegisterUpdateProduct) (err error) {
	validate := validator.New()
	CreateCustomValidations(validate)
	return validate.Struct(product)

}