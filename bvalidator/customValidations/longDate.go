package customValidations

import (
	"github.com/go-playground/validator"
	"regexp"
)

// validate datetime format "2009-11-10 23:00"
func ValidateLongDate(fl validator.FieldLevel) bool {
	re := regexp.MustCompile("[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) (2[0-3]|[01][0-9]):[0-5][0-9]")
	fv := fl.Field().String()

	// allow empty string if fields is not a required field
	if fv == "" || re.MatchString(fv)  {
		return true
	}
	return false
}
