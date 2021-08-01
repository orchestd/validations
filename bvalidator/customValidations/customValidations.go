package customValidations

import (
	"bitbucket.org/HeilaSystems/validations"
	"github.com/go-playground/validator/v10"
)

type customValidation struct {
	validationTag      string
	validationFunction func(fl validator.FieldLevel) bool
}

func (v *customValidation) GetValidationTag() string {
	return v.validationTag
}

func (v *customValidation) GetValidationFunction() interface{} {
	return v.validationFunction
}

func NewCustomValidation(validationTag string, validationFunction func(fl validator.FieldLevel) bool) func() validations.CustomValidation {
	return func() validations.CustomValidation {
		return &customValidation{validationTag: validationTag, validationFunction: validationFunction}
	}
}
