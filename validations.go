package validations

import "bitbucket.org/HeilaSystems/servicereply"

type CustomValidation interface {
	GetValidationTag() string
	GetValidationFunction() interface{}
}

type Builder interface {
	AddCustomValidations(...CustomValidation) Builder
	Build() (Validations,error)
}
type Validations interface {
	Validate(req interface{}) servicereply.ServiceReply
}

