package validator

import (
	"fmt"
	apperror "template/utils/error"

	v "github.com/go-playground/validator/v10"
)

func Validate(s interface{}) error {
	validate := v.New()
	err := validate.Struct(s)

	var errMsg []string

	if err != nil {
		for _, err := range err.(v.ValidationErrors) {
			errMsg = append(errMsg, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
	}

	if len(errMsg) > 0 {
		return apperror.NewWithContext(errMsg, "validation error")
	}

	return err
}
