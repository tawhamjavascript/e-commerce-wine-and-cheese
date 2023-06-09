package validate

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)



func CreateCustomValidations(validate *validator.Validate) {
	validate.RegisterValidation("nameManga", validateNameManga)
	validate.RegisterValidation("description", validateDescription)
	validate.RegisterValidation("general", validateGeneral)
	validate.RegisterValidation("name", validateNameUser)
}

func validateNameManga(fl validator.FieldLevel) bool {
	nameManga := fl.Field().String()
	regex := regexp.MustCompile(`^[a-zA-ZÀ-ÿ ]+$`)
	return regex.MatchString(nameManga)
}
func validateNameUser(fl validator.FieldLevel) bool {
	nameUser := fl.Field().String()
	regex := regexp.MustCompile(`^[a-zA-ZÀ-ÿ ]+$`)
	return regex.MatchString(nameUser)
}


func validateDescription(fl validator.FieldLevel) bool {
	description := fl.Field().String()
	regex := regexp.MustCompile(`^[a-zA-ZÀ-ÿ0-9 ,()]+$`)
	return regex.MatchString(description)
}

func validateGeneral(fl validator.FieldLevel) bool {
	general := fl.Field().String()
	regex := regexp.MustCompile(`^[a-zA-ZÀ-ÿ ]+$`)
	return regex.MatchString(general)
}

