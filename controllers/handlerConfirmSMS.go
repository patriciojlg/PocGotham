package controllers

import (
	views "PocGotham/views"
	"net/http"
)

func HandlerConfirmSMS(w http.ResponseWriter, r *http.Request) {
	content := views.ConfirmSMS()
	component := views.BaseTemplate(content)
	component.Render(r.Context(), w)
}
