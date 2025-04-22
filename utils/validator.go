package utils

import (
	"strings"

	"github.com/gin-gonic/gin/binding"
	en "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var Trans ut.Translator

func InitTransValidator() {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		ut := ut.New(en, en)
		Trans, _ = ut.GetTranslator("en")
		_ = en_translations.RegisterDefaultTranslations(validate, Trans)
	}
}

func ExtractValidationErrors(err error) map[string]string {
	errFields := make(map[string]string)

	validatorErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		errFields["error"] = "incorrect usage format"
		return errFields
	}

	for _, errField := range validatorErrors {
		errFields[strings.ToLower(errField.Field())] = errField.Translate(Trans)
	}

	return errFields
}
