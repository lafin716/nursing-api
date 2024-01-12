package api

import "github.com/go-playground/validator/v10"

func passwordValidator(fl validator.FieldLevel) bool {
	return fl.Field().Len() >= 4 && fl.Field().Len() <= 20
}

func keywordValidator(fl validator.FieldLevel) bool {
	return fl.Field().Len() >= 2
}
