package customValidations

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

/*
match international and local phone number formats:
^(\+\d{1,2}\s?)?1?\-?\.?\s?\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$

also match the specific format 03-0000000 or variations with space or dot separators:
^04[\s.-]?\d{7}$
*/
func ValidatePhoneNumberAll(fl validator.FieldLevel) bool {
	// allow empty string if fields is not a required field
	if fv := fl.Field().String(); fv == "" {
		return true
	} else if re := regexp.MustCompile("^(\\+\\d{1,2}\\s?)?1?\\-?\\.?\\s?\\(?\\d{3}\\)?[\\s.-]?\\d{3}[\\s.-]?\\d{4}$|^04[\\s.-]?\\d{7}$"); re.MatchString(fv) {
		return true
	}
	return false
}
