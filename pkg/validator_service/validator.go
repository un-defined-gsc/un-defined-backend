package validator_service

import (
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/tr"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	tr_translations "github.com/go-playground/validator/v10/translations/tr"
)

type ValidatorService struct {
	validate *validator.Validate
}

type keyErr struct {
	Key string `json:"field"`
	Err string `json:"error"`
}

var transTr *ut.Translator
var transEn *ut.Translator

// func GetTranslator(lang string) *ut.Translator {
// 	switch lang {
// 	case "tr":
// 		return transTr
// 	case "en":
// 		return transEn
// 	default:
// 		return transEn
// 	}
// }

// NewValidator func for create a new validator for model fields.

func NewValidatorService() *ValidatorService {
	trLang := tr.New()
	enLang := en.New()
	uni := ut.New(enLang, trLang, enLang)
	transtr, _ := uni.GetTranslator("tr")
	transen, _ := uni.GetTranslator("en")
	transTr = &transtr
	transEn = &transen
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})
	if err := tr_translations.RegisterDefaultTranslations(validate, *transTr); err != nil {
		log.Fatalf("error on register default tr translations")
	}
	if err := en_translations.RegisterDefaultTranslations(validate, *transEn); err != nil {
		log.Fatalf("error on register default en translations")
	}
	return &ValidatorService{
		validate: validate,
	}
}

func (vs *ValidatorService) ValidateStruct(s any) error {
	return vs.validate.Struct(s)
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error, lang string) (result []keyErr) {
	var trans *ut.Translator
	switch lang {
	case "tr":
		trans = transTr
	case "en":
		trans = transEn
	default:
		trans = transEn
	}
	for _, err := range err.(validator.ValidationErrors) {
		result = append(result, keyErr{
			Key: err.Field(),
			Err: err.Translate(*trans),
		})
	}
	return result
}
