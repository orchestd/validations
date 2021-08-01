package customValidations

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// validate phone number "0501234567"
func ValidateIsraeliPhoneMobile(fl validator.FieldLevel) bool {
	// allow empty string if fields is not a required field
	if fv := fl.Field().String(); fv == "" {
		return true
	} else if re := regexp.MustCompile("^[0][5][0|1|2|3|4|5|6|7|8|9]{1}[0-9]{7}$"); re.MatchString(fv) {
		return true
	}
	return false
}
