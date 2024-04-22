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
	router.Get(settings.USERNAME_PARTIAL_ENDOINT, controllers.MainControllerUsernameInputText)
	router.Post(settings.USERNAME_PARTIAL_ENDOINT, controllers.MainControllerUsernameInputText)
	router.Get(settings.PASSWORD_PARTIAL_ENDOINT, controllers.MainControllerPasswordInputText)
	router.Post(settings.PASSWORD_PARTIAL_ENDOINT, controllers.MainControllerPasswordInputText)
	router.Post(settings.REMEMBER_PARTIAL_ENDOINT, controllers.HandlerCheckboxRemember)

	router.Get(settings.SIGNUP_PARTIAL_ENDOINT, controllers.HandlerSignUpButton)
	router.Post(settings.SIGNUP_PARTIAL_ENDOINT, controllers.HandlerSignUpButton)

	slog.Info("Server starting", "port", settings.PORT)

	err := http.ListenAndServe(settings.SERVER, router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
