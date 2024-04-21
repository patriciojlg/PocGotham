package validators

import (
	"errors"
	"unicode"
)

// Función validadora que verifica si un número es positivo
func ValidateStrongPassowrd(s string) error {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasMinLen {
		return errors.New("la contraseña debiera tener al menos 7 carácteres")
	}
	if !hasUpper {
		return errors.New("la contraseña debiera tener al menos una letra mayúscula")
	}
	if !hasLower {
		return errors.New("la contraseña debiera tener al menos una letra minúscula")
	}
	if !hasNumber {
		return errors.New("la contraseña debiera tener al menos un número")
	}
	if !hasSpecial {
		return errors.New("la contraseña debiera tener al menos un carácter espcial (*/_°! etc)")
	}
	return nil
}
