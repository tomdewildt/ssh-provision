package validate

import (
	"errors"

	validator "github.com/go-playground/validator/v10"
)

// Schema is used to validate a input map based on a schema that
// contains go-playground/validator tags. It takes a schema in map
// form as the first parameter and an input as the second parameter.
// It returns nil if the schema passes validation or an error if the
// schema doesn't pass validation.
func Schema(schema map[string]string, input map[string]interface{}) error {
	if schema == nil {
		return errors.New("invalid schema")
	}
	if input == nil {
		return errors.New("invalid input")
	}

	validate := validator.New()
	validate.RegisterValidation("schema", func(fl validator.FieldLevel) bool {
		m, ok := fl.Field().Interface().(map[string]interface{})
		if !ok {
			return false
		}

		for k, v := range schema {
			if validate.Var(m[k], v) != nil {
				return false
			}
		}

		return true
	})

	if err := validate.Var(input, "schema"); err != nil {
		return errors.New("input does not match schema")
	}
	return nil
}
