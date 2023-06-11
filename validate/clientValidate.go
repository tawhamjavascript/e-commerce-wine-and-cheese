package validate

import (
	"e-commerce/model"

	"github.com/go-playground/validator/v10"
)


func ValidateRegisterClient(vendor *model.SignUpClient) (err error) {
	validate := validator.New()
	CreateCustomValidations(validate)
	err = validate.Struct(vendor)
	return err
}

func ValidateLoginClient(vendor *model.SignInClient) (err error) {
	validate := validator.New()
	err = validate.Struct(vendor)
	return err
}
