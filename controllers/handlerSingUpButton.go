package controllers

import (
	models "PocGotham/models"
	views "PocGotham/views"
	"net/http"
)

var classSignupButton = "flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"

func HandlerSignUpButton(w http.ResponseWriter, r *http.Request) {
	modelButton := models.NewButton("signup-button", "Crear cuenta")
	modelButton.SetDisabled(false)
	modelButton.SetHXPostEndpoint("/submit-signup")
	modelButton.SetHXTrigger("click")
	modelButton.SetHXSwap("outerHTML")
	modelButton.SetCSSClassButton(classSignupButton)
	component := views.ButtonTempl(modelButton)
	component.Render(r.Context(), w)
	/*
		hook := []templ.Component{}
		if r.FormValue("hook") != "" {
			hook = append(hook, views.NewHook(r.FormValue("hook")))
		}
		//component.Render(r.Context(), w)

		for _, comp := range repeat {
			comp.Render(r.Context(), w)
		}
	*/
}
