package validation

import (
	"github.com/go-playground/validator/v10"
)

func Validate(ent interface{}) error {
	validate := validator.New()
	return validate.Struct(ent)
}
