package customValidations

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// regular espression
func ValidateRegexp(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	if field == "" {
		return true
	}
	regexString := fl.Param()
	regex := regexp.MustCompile(regexString)
	match := regex.MatchString(field)
	return match
}
