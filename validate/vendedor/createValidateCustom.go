package vendorValidate

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func validateNameOfUser(fl validator.FieldLevel) bool {
	nome := fl.Field().String()
	regex := regexp.MustCompile(`^[a-zA-ZÀ-ÿ ]+$`)
	return regex.MatchString(nome)
}

func validateDescription(fl validator.FieldLevel) bool {
	nome := fl.Field().String()
	regex := regexp.MustCompile(`^[a-zA-ZÀ-ÿ ,()]+$`)
	return regex.MatchString(nome)
}
