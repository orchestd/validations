package validations

import (
	. "bitbucket.org/HeilaSystems/servicereply"
	"context"
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
}

type CacheValidator struct {
	Id        string
	SortOrder int
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
