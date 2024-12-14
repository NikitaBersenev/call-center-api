package myerr

import (
	"call-center/pkg/tools"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidatorCallMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "phonenumber":
		return "Invalid phone number format"
	case "required":
		return "The '" + fe.Field() + "' field is required"

	}
	return fe.Error()
}

func ValidatorError(err error) []error {
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]error, len(ve))
			for i, fe := range ve {
				out[i] = fmt.Errorf(ValidatorCallMsg(fe))
			}
			return out
		}
	}
	return nil
}

func ValidatorErrorStr(err error) string {
	errorsVal := tools.Map(ValidatorError(err), func(el error) string {
		return el.Error()
	})
	return strings.Join(errorsVal, ", ")
}
