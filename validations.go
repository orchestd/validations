package validations

import (
	. "github.com/orchestd/servicereply"
)

type CustomValidation interface {
	GetValidationTag() string
	GetValidationFunction() interface{}
}

type Builder interface {
	AddCustomValidations(...CustomValidation) Builder
	Build() (Validations, error)
}
type Validations interface {
	Validate(req interface{}) ServiceReply
}

type Validator interface {
	Validate() ServiceReply
	GetId() string
	GetName() string
	GetIsEnabledByDefault() bool
}

type CacheValidator struct {
	Id        string
	SortOrder int
	Enabled   bool
}

func NewValidatorCont() ValidatorCont {
	return ValidatorCont{}
}

type ValidatorCont struct {
	Validators []Validator
}

type ValidatorRunner interface {
	Validate(c context.Context, validators ValidatorCont) ServiceReply
}
