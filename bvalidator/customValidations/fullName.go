package customValidations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateFullName(fl validator.FieldLevel) bool {
	// Allow empty string if the field is not a required field.
	// The "required" validation should handle requiredness separately.
	if fv := fl.Field().String(); fv == "" {
		return true
	}

	return isValidFullName(fl.Field().String())
}

// Full name must contain at least two parts.
// Each part must contain at least two Unicode letters.
// Additional name parts are allowed.
//
// Examples:
// "John Smith"           -> valid
// "John Michael Smith"   -> valid
// "José García"          -> valid
// "J Smith"              -> invalid
// "John S"               -> invalid
// "John"                 -> invalid
// "John123 Smith"        -> invalid
func isValidFullName(val string) bool {
	re := regexp.MustCompile(`^[\p{L}]{2,}(?:\s+[\p{L}]{2,})+$`)
	return re.MatchString(val)
}
