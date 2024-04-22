package validators

import (
	"errors"
	"regexp"
)

func IsEmailValid(email string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("formato de email no válido")
	}
	return nil
}

func IsEmpty(value string) error {
	if value == "" {
		return errors.New("campo requerido")
	}
	return nil

}
