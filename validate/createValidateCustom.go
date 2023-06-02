package validate

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func CreateCustomValidations(validate *validator.Validate) {
	validate.RegisterValidation("name", validateNameOfUser)
	validate.RegisterValidation("description", validateDescription)
}

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

