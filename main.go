package main

import (
	controllers "PocGotham/controllers"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandlerBase(w, r)
	})

	// Configurar el manejador para renderizar el HTML
	router.Post("/partial-username-input-text", func(w http.ResponseWriter, r *http.Request) {
		controllers.MainControllerUsernameInputText(w, r)
	})

	router.Post("/partial-password-input-text", func(w http.ResponseWriter, r *http.Request) {
		controllers.MainControllerPasswordInputText(w, r)
	})

	router.Post("/partial-remember-checkbox", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandlerCheckboxRemember(w, r)
	})

	router.Post("/submit-signup", func(w http.ResponseWriter, r *http.Request) {
		controllers.FromHookUsernameInputText(w, r)
		controllers.HandlerSignUpButton(w, r)
	})

	router.Get("/submit-signup", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandlerSignUpButton(w, r)
	})

	log.Printf("Server starting on port 8080")

	// Start the server with error handling
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
