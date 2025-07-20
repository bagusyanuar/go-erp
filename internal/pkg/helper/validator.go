package helper

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translate_id "github.com/go-playground/validator/v10/translations/id"
)

func Validate(v *validator.Validate, request any) (messages map[string][]string, err error) {
	err = v.Struct(request)
	if err != nil {
		localeID := id.New()
		uni := ut.New(localeID, localeID)
		trans, _ := uni.GetTranslator("id")
		translate_id.RegisterDefaultTranslations(v, trans)

		tmpMap := make(map[string][]string)

		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			f, _ := reflect.TypeOf(request).Elem().FieldByName(field)
			jsonName, _ := f.Tag.Lookup("json")
			translated := strings.ToLower(e.Translate(trans))
			tmpMap[jsonName] = append(tmpMap[jsonName], translated)
		}
		messages = tmpMap
	}
	return messages, err
}
