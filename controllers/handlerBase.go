package controllers

import (
	views "PocGotham/views"
	"net/http"
)

func HandlerBase(w http.ResponseWriter, r *http.Request) {
	component := views.BaseTemplate()
	component.Render(r.Context(), w)
}
