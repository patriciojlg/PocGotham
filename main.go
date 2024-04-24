package main

import (
	routes "PocGotham/routes"
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
	routes.SingUpRoutes(router)
	routes.ConfirmSMSRoutes(router)
	// Routes

	slog.Info("Server starting", "port", settings.PORT)

	err := http.ListenAndServe(settings.SERVER, router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
