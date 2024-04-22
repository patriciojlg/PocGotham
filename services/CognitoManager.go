package services

import (
	models "PocGotham/models"
	"fmt"
)

func CreateNewUser(username, password string) error {
	cognito := models.NewCognito()
	err := cognito.CreateUser(username, password)
	if err != nil {
		fmt.Println("Error al crear el nuevo usuario:", err)
		return err
	}
	// HOOK WITH MAILGUN SERVICE (SEND CONFIRMATION EMAIL)
	return nil

}
