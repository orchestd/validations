package customValidations

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

/*
regex to allow the user to enter only numbers, +, -, whitespace and ()
It respects the parenthesis balance and there is always a number after a symbol
*/
func ValidatePhoneNumberAll(fl validator.FieldLevel) bool {
	// allow empty string if fields is not a required field
	if fv := fl.Field().String(); fv == "" {
		return true
	} else {
		return isValidPhoneNumberAll(fv)
	}
}

func isValidPhoneNumberAll(val string) bool {
	if len(val) > 15 {
		return false
	}
	if re := regexp.MustCompile(`^([+]?[\s0-9]+)?(\d{3}|[(]?[0-9]+[)])?([-]?[\s]?[0-9]+(?:\.[0-9]+)?(?:-[0-9]+(?:\.[0-9]+)?)?)$`); re.MatchString(val) {
		return true
	}
	return false
}
