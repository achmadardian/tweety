package validate

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	id "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

func InitTranslator() {
	id := id.New()
	uni = ut.New(id, id)

	trans, _ = uni.GetTranslator("id")
	id_translations.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans)
}

func ExtractValidationErrors(req any, err error) map[string]string {
	errFields := make(map[string]string)

	errVal, ok := err.(validator.ValidationErrors)
	if !ok {
		errFields["error"] = "invalid error type"
		return errFields
	}

	t := reflect.TypeOf(req)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for _, e := range errVal {
		field, found := t.FieldByName(e.Field())
		if !found {
			continue
		}

		jsonTag := field.Tag.Get("json")
		jsonKey := strings.Split(jsonTag, ",")[0]
		if jsonKey == "" || jsonKey == "-" {
			jsonKey = strings.ToLower(e.Field())
		}

		errFields[jsonKey] = e.Translate(trans)
	}

	return errFields
}
