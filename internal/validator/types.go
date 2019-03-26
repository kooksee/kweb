package validator

import (
	"reflect"
)

// IsNumber is the validation function for validating if the current field's value is a valid number.
func (t *KValidator) IsNumber() bool {
	switch t.data.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
		return true
	default:
		return numberRegex.MatchString(t.data.String())
	}
}

// IsNumeric is the validation function for validating if the current field's value is a valid numeric value.
func (t *KValidator) IsNumeric() bool {
	switch t.data.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
		return true
	default:
		return numericRegex.MatchString(t.data.String())
	}
}

// IsAlphanum is the validation function for validating if the current field's value is a valid alphanumeric value.
func (t *KValidator) IsAlphanum() bool {
	return alphaNumericRegex.MatchString(t.data.String())
}

// IsAlpha is the validation function for validating if the current field's value is a valid alpha value.
func (t *KValidator) IsAlpha() bool {
	return alphaRegex.MatchString(t.data.String())
}

// IsAlphanumUnicode is the validation function for validating if the current field's value is a valid alphanumeric unicode value.
func (t *KValidator) IsAlphanumUnicode() bool {
	return alphaUnicodeNumericRegex.MatchString(t.data.String())
}

// IsAlphaUnicode is the validation function for validating if the current field's value is a valid alpha unicode value.
func (t *KValidator) isAlphaUnicode() bool {
	return alphaUnicodeRegex.MatchString(t.data.String())
}
