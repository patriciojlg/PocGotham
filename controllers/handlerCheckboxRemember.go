package controllers

import (
	models "PocGotham/models"
	settings "PocGotham/settings"
	views "PocGotham/views"
	"net/http"
)

func HandlerCheckboxRemember(w http.ResponseWriter, r *http.Request) {
	remember := models.NewCheckbox("remember")
	remember.SetLabelText("Recuérdame")
	remember.SetId("remember")
	remember.SetName("remember")
	remember.SetChecked(false)
	remember.SetHXPostEndpoint(settings.PASSWORD_PARTIAL_ENDOINT)
	remember.SetHXSwap("outerHTML")
	remember.SetLabelCSSClassByDefault()
	remember.SetInputCSSClassByStatus()
	remember.SetFor("remember")

	if r.FormValue("remember") == "on" {
		remember.SetChecked(true)
	}
	component := views.CheckboxView(remember)

	component.Render(r.Context(), w)
}
