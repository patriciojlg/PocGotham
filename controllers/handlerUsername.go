package controllers

import (
	models "PocGotham/models"
	validators "PocGotham/validators"
	views "PocGotham/views"
	"net/http"
)

var PostValidators []func(string) error = []func(string) error{validators.IsEmailValid, validators.IsEmpty}
var GetValidators []func(string) error = []func(string) error{}

func usernameInputTextdefaultStatus() *models.InputTextModel {
	username := models.NewInputTextModel("username")
	username.SetLabelText("Email")
	username.SetRequired(true)
	username.SetId("username")
	username.SetName("username")
	username.SetType("email")
	username.SetAutocomplete("email")
	username.SetHXTrigger("change submit")
	username.SetHXEndpoint("/partial-username-input-text")
	username.SetClassLabelByStatus("default")
	username.SetClassInputTextByStatus("default")
	return username
}
func getValidatorsByMethod(method string) []func(string) error {
	switch {
	case method == "GET":
		return GetValidators
	case method == "POST":
		return PostValidators
	default:
		return nil
	}

}

func MainControllerUsernameInputText(w http.ResponseWriter, r *http.Request) {
	modelUsername := usernameInputTextdefaultStatus()
	valueUsername := r.FormValue("username")
	modelUsername.SetValue(valueUsername)

	validatorByMethod := getValidatorsByMethod(r.Method)
	modelUsername.Validator.AddArrayFuncsValidators(validatorByMethod)

	if valueUsername == "" && r.Method == "GET" {
		modelUsername.SetStatusByString("default", "")
		component := views.InputTextViews(modelUsername)
		component.Render(r.Context(), w)
		return
	}
	modelUsername.ValidateNow()

	component := views.InputTextViews(modelUsername)
	component.Render(r.Context(), w)
}
