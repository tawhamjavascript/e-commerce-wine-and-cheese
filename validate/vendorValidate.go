package validate

import (
	"e-commerce/model"

	"github.com/go-playground/validator/v10"
)



func ValidateRegisterVendor(vendor *model.RegisterVendor) (err error) {
	validate := validator.New()
	CreateCustomValidations(validate)
	err = validate.Struct(vendor)
	return err
}

func ValidateLoginVendor(vendor *model.SignInVendor) (err error) {
	validate := validator.New()
	err = validate.Struct(vendor)
	return err
}
