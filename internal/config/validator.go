package config

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var trans ut.Translator

func NewValidator() *validator.Validate {
	v := validator.New()

	locale := en.New()
	uni := ut.New(locale, locale)
	t, _ := uni.GetTranslator("en")
	trans = t

	// using indonesian local translation
	// localeID := id.New()
	// uni := ut.New(localeID, localeID)
	// trans, _ := uni.GetTranslator("id")
	// translate_id.RegisterDefaultTranslations(v, trans)

	//change field name into json tag
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		tag := fld.Tag.Get("json")
		if tag == "-" || tag == "" {
			return fld.Name // fallback ke nama field struct
		}
		return strings.Split(tag, ",")[0]
	})

	// Default translations EN
	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		panic("failed to register default EN translations: " + err.Error())
	}

	if err := RegisterSymbolValidation(v); err != nil {
		panic("failed to register symbol validation: " + err.Error())
	}

	if err := RegisterSymbolTranslation(v, trans); err != nil {
		panic("failed to register symbol translation: " + err.Error())
	}
	return v
}

func RegisterSymbolValidation(v *validator.Validate) error {
	return v.RegisterValidation("symbol", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()

		symbolRegex := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]{};':"\\|,.<>\/?]`)
		return symbolRegex.MatchString(value)
	})
}

func RegisterSymbolTranslation(v *validator.Validate, trans ut.Translator) error {
	return v.RegisterTranslation("symbol", trans,
		func(ut ut.Translator) error {
			return ut.Add("symbol", "{0} must contain at least one symbol", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("symbol", fe.Field())
			return t
		},
	)
}

func GetTranslator() ut.Translator {
	return trans
}
