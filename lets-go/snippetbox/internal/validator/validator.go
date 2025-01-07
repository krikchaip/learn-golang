package validator

import (
	"slices"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors map[string]string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

// adds an error message to the FieldErrors map if not exists
func (v *Validator) AddFieldError(key, message string) {
	// NOTE: keep in mind that the zero value for 'map' type is a pointer
	if v.FieldErrors == nil {
		// v.FieldErrors = make(map[string]string)
		v.FieldErrors = map[string]string{}
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

func (v *Validator) CheckField(validated bool, key, message string) {
	if !validated {
		v.AddFieldError(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func PermittedValues[T comparable](value T, permittedValues []T) bool {
	return slices.Contains(permittedValues, value)
}
