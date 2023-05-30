package vendorValidate

import (
	"e-commerce/model"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateRegisterVendor(vendor *model.RegisterVendor) (err error) {
	validate = validator.New()
	_ = validate.RegisterValidation("name", validateNameOfUser)
	_ = validate.RegisterValidation("description", validateDescription)
	err = validate.Struct(vendor)
	return err
}

func ValidateLoginVendor(vendor *model.SignInVendor) (err error) {
	err = validate.Struct(vendor)

	return
}
