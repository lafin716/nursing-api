package validates

import "github.com/go-playground/validator/v10"

type ValidateEntity struct {
	Name      string
	Validator func(fl validator.FieldLevel) bool
}

type ValidateError struct {
	Field string
	Tag   string
	Param string
}

type ValidateRegister interface {
	Execute(*validator.Validate)
}

func NewRuleGroup(rules ...*ValidateEntity) []*ValidateEntity {
	ruleGroup := []*ValidateEntity{}
	for _, rule := range rules {
		ruleGroup = append(ruleGroup, rule)
	}
	return ruleGroup
}

func NewRule(name string, rule func(fl validator.FieldLevel) bool) *ValidateEntity {
	return &ValidateEntity{
		Name:      name,
		Validator: rule,
	}
}

func NewValidator() *validator.Validate {
	validate := validator.New()
	return validate
}

func NewValidatorWithCustomValidators(entities []*ValidateEntity) (*validator.Validate, error) {
	validate := validator.New()

	for _, entity := range entities {
		err := validate.RegisterValidation(entity.Name, entity.Validator)
		if err != nil {
			return nil, err
		}
	}

	return validate, nil
}

func ValidateErrors(err error) map[string]ValidateError {
	fields := map[string]ValidateError{}

	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = ValidateError{
			Field: err.Field(),
			Tag:   err.Tag(),
			Param: err.Param(),
		}
	}

	return fields
}
