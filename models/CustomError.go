package models

import "errors"

// ErrorPersonalizado es un tipo de error personalizado que incluye una descripción del error
type CustomError struct {
	Message error
}

func NewCustomError(message string) *CustomError {
	return &CustomError{
		Message: errors.New(message),
	}

}
