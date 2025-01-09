package validator

import (
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"
)

// parsing this pattern once at startup and storing the compiled
// result in a variable is more performant than re-parsing
// the pattern each time we need it
// ref: https://html.spec.whatwg.org/multipage/input.html#valid-e-mail-address
var EmailRegex = regexp.MustCompile(
	"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
)

type Validator struct {
	FieldErrors map[string]string

	// hold any validation errors which are not related to a specific form field
	NonFieldErrors []string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
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

func (v *Validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}

func (v *Validator) CheckField(validated bool, key, message string) {
	if !validated {
		v.AddFieldError(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func PermittedValues[T comparable](value T, permittedValues []T) bool {
	return slices.Contains(permittedValues, value)
}

func Matches(value string, r *regexp.Regexp) bool {
	return r.MatchString(value)
}
