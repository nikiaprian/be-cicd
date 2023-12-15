package utils

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func MessageErrorByValidation(err error) string {

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		firstError := ve[0]
		fmt.Println(firstError.Tag())

		if firstError.Tag() == "regexp" {
			if firstError.Field() == "password" {
				return "Password must one lowercase, one uppercase, one number, and one special character!"
			}
		} else if firstError.Tag() == "required" {
			return fmt.Sprintf("%s is required!", firstError.Field())
		}
		return firstError.Error()
	}

	return "Something went wrong"
}

func NewValidator() *validator.Validate {
	var validate *validator.Validate
	validate = validator.New()
	return validate
}

func PasswordValidator(password string) error {

	regex, _ := regexp.Compile(`^(.+?[a-z])(.+?[A-Z])(.+?\d)(.+?[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	isMatch := regex.MatchString(password)

	if isMatch == false {
		return errors.New("Password must one lowercase, one uppercase, one number, and one special character!")
	}
	return nil
}
