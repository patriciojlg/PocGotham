package main

import (
	controllers "PocGotham/controllers"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", controllers.HandlerBase)

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

	port := os.Getenv("PORT")
	slog.Info("Server starting", "port", port)

	// Start the server with error handling
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
