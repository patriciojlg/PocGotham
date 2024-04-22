package controllers

import (
	models "PocGotham/models"
	validators "PocGotham/validators"
	views "PocGotham/views"
	"net/http"
)

func passwordInputTextStatus() *models.InputTextModel {
	password := models.NewInputTextModel("password")
	password.SetLabelText("Contraseña")
	password.SetRequired(true)
	password.SetId("password")
	password.SetName("password")
	password.SetType("password")
	password.SetAutocomplete("password")
	password.SetHXTrigger("change submit")
	password.SetHXEndpoint("/partial-password-input-text")
	password.SetClassLabelByStatus("default")
	password.SetClassInputTextByStatus("default")
	return password
}

func getPasswordStatus(valuePassword string) (string, string) {
	if validators.IsEmpty(valuePassword) {
		return "empty", ""

	}
	errPassword := validators.ValidateStrongPassowrd(valuePassword)
	if errPassword != nil {
		return "error", errPassword.Error()
	}

	return "success", ""
}

func MainControllerPasswordInputText(w http.ResponseWriter, r *http.Request) {
	modelPassword := passwordInputTextStatus()
	valuePassword := r.FormValue("password")
	modelPassword.SetValue(valuePassword)

	status, errorMessage := getPasswordStatus(valuePassword)
	if errorMessage != "" {
		modelPassword.SetErrorMessage(errorMessage)
	}
	modelPassword.SetClassLabelByStatus(status)
	modelPassword.SetClassInputTextByStatus(status)

	component := views.InputTextViews(modelPassword)

	component.Render(r.Context(), w)
}

func ControllerPassowrdInputFromHook(w http.ResponseWriter, r *http.Request) {
	modelPassword := passwordInputTextStatus()
	valuePassword := r.FormValue("password")
	modelPassword.SetValue(valuePassword)

	status, errorMessage := getPasswordStatus(valuePassword)
	if errorMessage != "" {
		modelPassword.SetErrorMessage(errorMessage)
	}
	modelPassword.SetClassLabelByStatus(status)
	modelPassword.SetClassInputTextByStatus(status)

	component := views.InputTextViews(modelPassword)

	component.Render(r.Context(), w)
}
