package validation

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

func GetField(fe validator.FieldError, t int) string {
	var builder strings.Builder

	for i, char := range fe.Field() {
		if i > 0 && char >= 'A' && char <= 'Z' {
			if t == 1 {
				builder.WriteRune('_')
			} else {
				builder.WriteRune(' ')
			}
		}
		builder.WriteRune(char)
	}

	return strings.ToLower(builder.String())
}

func GetError(err error, ve validator.ValidationErrors) any {
	if errors.As(err, &ve) {
		out := make(map[string]string)
		for _, fe := range ve {
			out[GetField(fe, 1)] = GetErrorMsg(fe)
		}

		return out
	}

	return nil
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("The %s field must not be left blank.", GetField(fe, 0))
	case "eqfield":
		return "Password and confirm password doesn't match"
	case "min":
		return fmt.Sprintf("The %s length should be greater than %s", GetField(fe, 0), fe.Param())
	case "max":
		return fmt.Sprintf("The %s length should be less than %s", GetField(fe, 0), fe.Param())
	case "email":
		return "The email field must be an email"
	default:
		return fmt.Sprintf("Validation failed on the %s field", GetField(fe, 0))
	}
}
