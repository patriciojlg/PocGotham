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
	// Router
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)

	// Routes
	router.Get("/", controllers.HandlerBase)

	router.Post("/partial-username-input-text", controllers.MainControllerUsernameInputText)
	router.Post("/partial-password-input-text", controllers.MainControllerPasswordInputText)
	router.Post("/partial-remember-checkbox", controllers.HandlerCheckboxRemember)

	router.Get("/submit-signup", controllers.HandlerSignUpButton)
	router.Post("/submit-signup", func(w http.ResponseWriter, r *http.Request) {
		controllers.FromHookUsernameInputText(w, r)
		controllers.HandlerSignUpButton(w, r)
	})

	// Server
	port := os.Getenv("PORT")
	slog.Info("Server starting", "port", port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
