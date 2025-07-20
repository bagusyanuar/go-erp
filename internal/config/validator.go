package config

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	v := validator.New()
	if err := RegisterSymbolValidation(v); err != nil {
		panic("failed to register symbol validation: " + err.Error())
	}
	return v
}

func RegisterSymbolValidation(v *validator.Validate) error {
	return v.RegisterValidation("symbol", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()

		// Regex untuk simbol: semua karakter non-alfanumerik
		symbolRegex := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]{};':"\\|,.<>\/?]`)
		return symbolRegex.MatchString(value)
	})
}
