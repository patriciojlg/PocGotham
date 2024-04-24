package routes

import (
	"PocGotham/controllers"

	"github.com/go-chi/chi"
)

func ConfirmSMSRoutes(router chi.Router) {
	router.Get("/confirm-sms", controllers.HandlerConfirmSMS)

}
