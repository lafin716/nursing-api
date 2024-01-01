package validates

import "github.com/go-playground/validator/v10"

type ValidateEntity struct {
	Name      string
	Validator func(fl validator.FieldLevel) bool
}

type ValidateRegister interface {
	Execute(*validator.Validate)
}

func NewValidator() *validator.Validate {
	validate := validator.New()
	return validate
}

func NewValidatorWithCustomValidators(entities *[]ValidateEntity) (*validator.Validate, error) {
	validate := validator.New()

	for _, entity := range *entities {
		err := validate.RegisterValidation(entity.Name, entity.Validator)
		if err != nil {
			return nil, err
		}
	}

	return validate, nil
}

func ValidateErrors(err error) map[string]string {
	fields := map[string]string{}

	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}
