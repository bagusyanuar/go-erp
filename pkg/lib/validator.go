package lib

import (
	"strings"

	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/go-playground/validator/v10"
)

func Validate(v *validator.Validate, request any) (messages map[string][]string, err error) {
	err = v.Struct(request)
	if err != nil {
		trans := config.GetTranslator()
		tmpMap := make(map[string][]string)

		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()

			// using default field name
			// f, _ := reflect.TypeOf(request).Elem().FieldByName(field)
			// jsonName, _ := f.Tag.Lookup("json")
			// tmpMap[jsonName] = append(tmpMap[jsonName], translated)

			translated := strings.ToLower(e.Translate(trans))
			tmpMap[field] = append(tmpMap[field], translated)
		}
		messages = tmpMap
	}
	return messages, err
}
