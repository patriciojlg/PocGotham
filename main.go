package main

import (
	controllers "PocGotham/controllers"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandlerBase(w, r)
	})
	// Configurar el manejador para renderizar el HTML
	r.Post("/partial-username-input-text", func(w http.ResponseWriter, r *http.Request) {

		controllers.FromHookUsernameInputText(w, r)

	})
	r.Get("/partial-username-input-text", func(w http.ResponseWriter, r *http.Request) {

		controllers.FromHookUsernameInputText(w, r)
	})
	r.Post("/partial-password-input-text", func(w http.ResponseWriter, r *http.Request) {
		controllers.MainControllerPasswordInputText(w, r)

	})
	r.Post("/partial-remember-checkbox", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandlerCheckboxRemember(w, r)

	})
	r.Post("/submit-signup", func(w http.ResponseWriter, r *http.Request) {

		controllers.HandlerSignUpButton(w, r)
	})
	r.Get("/submit-signup", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandlerSignUpButton(w, r)
	})
	http.ListenAndServe(":8080", r)
}
