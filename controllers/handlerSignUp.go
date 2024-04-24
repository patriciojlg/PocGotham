package controllers

import (
	views "PocGotham/views"
	"net/http"
)

func HandlerSignUpWrap(w http.ResponseWriter, r *http.Request) {
	content := views.SignUpForm()
	component := views.BaseTemplate(content)
	component.Render(r.Context(), w)
}
