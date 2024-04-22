package main

import (
	controllers "PocGotham/controllers"
	settings "PocGotham/settings"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// Router
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)

	// Routes
	router.Get("/", controllers.HandlerBase)
	router.Get("/partial-username-input-text", controllers.MainControllerUsernameInputText)
	router.Post("/partial-username-input-text", controllers.MainControllerUsernameInputText)
	router.Get("/partial-password-input-text", controllers.MainControllerPasswordInputText)
	router.Post("/partial-password-input-text", controllers.MainControllerPasswordInputText)
	router.Post("/partial-remember-checkbox", controllers.HandlerCheckboxRemember)

	router.Get("/submit-signup", controllers.HandlerSignUpButton)
	router.Post("/submit-signup", controllers.HandlerSignUpButton)

	slog.Info("Server starting", "port", settings.PORT)

	err := http.ListenAndServe(settings.SERVER, router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
