package controllers

import (
	models "PocGotham/models"
	validators "PocGotham/validators"
	views "PocGotham/views"
	"net/http"
)

func usernameInputTextdefaultStatus() *models.InputTextModel {
	username := models.NewInputTextModel("username")
	username.SetLabelText("Email")
	username.SetRequired(true)
	username.SetId("username")
	username.SetName("username")
	username.SetType("email")
	username.SetAutocomplete("email")
	username.SetHXTrigger("change")
	username.SetHXEndpoint("/partial-username-input-text")
	username.SetClassLabelByStatus("default")
	username.SetClassInputTextByStatus("default")
	return username
}

func getStatus(valueUsername string) (string, string) {
	if validators.IsEmpty(valueUsername) {
		return "empty", ""

	}
	if !validators.IsEmailValid(valueUsername) {
		return "error", "Formato de email no válio"
	}

	return "success", ""
}

func MainControllerUsernameInputText(w http.ResponseWriter, r *http.Request) {
	modelUsername := usernameInputTextdefaultStatus()
	valueUsername := r.FormValue("username")
	modelUsername.SetValue(valueUsername)

	status, errorMessage := getStatus(valueUsername)
	if errorMessage != "" {
		modelUsername.SetErrorMessage(errorMessage)
	}
	modelUsername.SetClassLabelByStatus(status)
	modelUsername.SetClassInputTextByStatus(status)

	component := views.InputTextViews(modelUsername)

	component.Render(r.Context(), w)
}

func renderEmptyError(modelUsername *models.InputTextModel, w http.ResponseWriter, r *http.Request) {
	modelUsername.SetErrorMessage("Campo requerido")
	status := "error"
	modelUsername.SetClassLabelByStatus(status)
	modelUsername.SetClassInputTextByStatus(status)

	component := views.InputTextViews(modelUsername)

	component.Render(r.Context(), w)
}
func FromHookUsernameInputText(w http.ResponseWriter, r *http.Request) {
	var status string

	modelUsername := usernameInputTextdefaultStatus()
	valueUsername := r.FormValue("username")
	modelUsername.SetValue(valueUsername)
	if valueUsername == "" {
		renderEmptyError(modelUsername, w, r)
		return

	}
	status, errorMessage := getStatus(valueUsername)
	modelUsername.SetClassLabelByStatus(status)
	modelUsername.SetClassInputTextByStatus(status)
	modelUsername.SetErrorMessage(errorMessage)
	component := views.InputTextViews(modelUsername)

	component.Render(r.Context(), w)

}
