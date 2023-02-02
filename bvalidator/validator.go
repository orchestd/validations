package bvalidator

import (
	"errors"
	"github.com/orchestd/servicereply"
	"github.com/orchestd/validations"
)

type validatorDep struct {
	validator *validator.Validate
}

type validationError struct {
	FieldName      string
	Value          interface{}
	ValidationRule string
	Error          string
}

/*
Validate use the validator tags in oreder to validate the request struct
for example:

for request struct that looks like this:

type SendOtpRequest struct {
	Recipient   string    `json:"recipient" validate:"required,len=10"`
}
(recipient is mandatory field and should its length should be exactly 10)

if the client send this json:
{
	"recipient":  "stavgayer@gmail.com",
}

the returned error will look like this:
{
    "status": "rejected",
    "message": {
        "id": "validation",
        "values": {
            "Recipient": {
                "expectedParam": "10",
                "key": "len"
            }
        }
    }
}
*/
func (v validatorDep) Validate(req interface{}) servicereply.ServiceReply {
	err := v.validator.Struct(req)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return servicereply.NewInternalServiceError(err)
		}
		valueMap := servicereply.ValuesMap{}
		for _, err := range err.(validator.ValidationErrors) {
			val := make(map[string]string)
			val["key"] = err.Tag()
			if len(err.Param()) > 0 {
				val["expectedParam"] = err.Param()
			}
			valueMap[err.Field()] = val
		}
		return servicereply.NewRejectedReply("validation").WithReplyValues(valueMap)
	}
	return nil
}

// NewValidator validatorDep constructor
func NewValidator(cfg *validatorConfig) (validations.Validations, error) {
	validate := validator.New()
	for _, validation := range cfg.customValidations {
		if f, ok := validation.GetValidationFunction().(func(fl validator.FieldLevel) bool); !ok {
			return nil, errors.New(validation.GetValidationTag() + " is not a valid custom validator interface ")
		} else if err := validate.RegisterValidation(validation.GetValidationTag(), f); err != nil {
			return nil, err
		}
	}
	return validatorDep{validate}, nil
}
