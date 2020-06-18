package validator

import (
	"github.com/go-playground/validator/v10"
	"strconv"
	"unicode/utf8"
)

func New() (validate *validator.Validate, err error) {
	validate = validator.New()
	if err = validate.RegisterValidation("minLength", minLengthValidator); err != nil { return }
	if err = validate.RegisterValidation("maxLength", maxLengthValidator); err != nil { return }
	if err = validate.RegisterValidation("length", lengthValidator); err != nil { return }

	return
}

func minLengthValidator(fl validator.FieldLevel) bool {
	paramLen, err := strconv.Atoi(fl.Param())
	fieldLen := utf8.RuneCountInString(fl.Field().String())
	if err != nil {
		return false
	}

	return fieldLen > paramLen
}

func maxLengthValidator(fl validator.FieldLevel) bool {
	paramLen, err := strconv.Atoi(fl.Param())
	fieldLen := utf8.RuneCountInString(fl.Field().String())
	if err != nil {
		return false
	}

	return fieldLen < paramLen
}

func lengthValidator(fl validator.FieldLevel) bool {
	fieldLen := len(strconv.Itoa(int(fl.Field().Int())))
	paramLen, err := strconv.Atoi(fl.Param())
	if err != nil {
		return false
	}

	return fieldLen == paramLen
}