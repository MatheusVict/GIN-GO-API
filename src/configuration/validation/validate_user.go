package validation

import (
	"encoding/json"
	"errors"
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/gin-gonic/gin/binding"
	en2 "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en2.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(
	validation_err error,
) *errorsHandle.ErrorsHandle {
	var jsonErr *json.UnmarshalTypeError
	var jsonvalidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return errorsHandle.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonvalidationError) {
		errorsCauses := []errorsHandle.Causes{}

		for _, fieldError := range validation_err.(validator.ValidationErrors) {
			cause := errorsHandle.Causes{
				Message: fieldError.Translate(transl),
				Field:   fieldError.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return errorsHandle.NewBadRequestValidationError("Some fields are invalids", errorsCauses)
	} else {
		return errorsHandle.NewBadRequestError("Error typing to convert fields")
	}

}
