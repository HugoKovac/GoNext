package validator

import (
	"github.com/go-playground/validator/v10"
)

// RegisterCustomValidators registers all custom validators with the validator instance
func RegisterCustomValidators(v *validator.Validate) {
	v.RegisterValidation("password", validatePassword)
}

// validatePassword is a validator.Func for password validation
func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return ValidatePassword(password) == nil
}
