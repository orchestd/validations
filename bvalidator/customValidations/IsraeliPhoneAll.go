package customValidations

import (
	"github.com/go-playground/validator"
	"regexp"
)

// validate phone number "031234567" | "0501234567"
func ValidateIsraeliPhoneAll(fl validator.FieldLevel) bool {
	re := regexp.MustCompile("^[0][5][0|1|2|3|4|5|6|7|8|9]{1}[0-9]{7}$")
	fv := fl.Field().String()

	// allow empty string if fields is not a required field
	if re.MatchString(fv) || fv == "" {
		return true
	}
	return false
}
