package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// Validatephone 校验手机号
func Validatephone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, phone)
	if !ok {
		return false
	}
	return true
}
