package validate

import (
	"e-commerce/model"

	"github.com/go-playground/validator/v10"
)



func ValidateRegisterProduct(product *model.RegisterProduct) (err error) {
	validate := validator.New()
	CreateCustomValidations(validate)
	err = validate.Struct(product)
	return err
}