package validator

import (
	"slices"
	"strings"
	"unicode/utf8"
)

type FormValidator struct {
	FormErrors map[string]string
}

func (formValidator *FormValidator) Valid() bool {
	return len(formValidator.FormErrors) == 0
}

func (formValidator *FormValidator) AddFormError(field string, msg string) {
	if formValidator.FormErrors == nil {
		formValidator.FormErrors = make(map[string]string)
	}

	if _, exists := formValidator.FormErrors[field]; !exists {
		formValidator.FormErrors[field] = msg
	}
}

func (formValidator *FormValidator) CheckField(ok bool, key, message string) {
	if !ok {
		formValidator.AddFormError(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MaxChars Returns true if under char count inclusive
func MaxChars(value string, limit int) bool {
	return utf8.RuneCountInString(value) <= limit
}

// PermittedValue Returns true if a value is in a list of specific permitted values.
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}
