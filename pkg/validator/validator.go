package validator

import (
	"github.com/dxckboi/hugeman-exam/pkg/errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validatorgo "github.com/go-playground/validator/v10"
	enTranslators "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	validate   *validatorgo.Validate
	translator ut.Translator
}

func NewValidator() *Validator {
	validate := validatorgo.New()

	english := en.New()
	uni := ut.New(english, english)
	trans, found := uni.GetTranslator("en")
	if !found {
		panic("locale language not found")
	}

	_ = enTranslators.RegisterDefaultTranslations(validate, trans)

	return &Validator{validate: validate, translator: trans}
}

func (c Validator) Struct(s interface{}) error {
	err := c.validate.Struct(s)
	return c.translate(err)
}

func (c Validator) Var(field interface{}, tag string) error {
	err := c.validate.Var(field, tag)
	return c.translate(err)
}

func (c Validator) translate(err error) error {
	if err == nil {
		return nil
	}

	return c.translates(err)[0]
}

func (c Validator) translates(err error) []error {
	if err == nil {
		return nil
	}

	validatorgoErrs := err.(validatorgo.ValidationErrors)
	var errs []error
	for _, e := range validatorgoErrs {
		parsed := errors.BadRequest(e.Translate(c.translator))
		errs = append(errs, parsed)
	}

	return errs
}
