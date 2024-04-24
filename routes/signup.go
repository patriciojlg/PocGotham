package routes

import (
	"PocGotham/controllers"
	"PocGotham/settings"

	"github.com/go-chi/chi"
)

func SingUpRoutes(router chi.Router) {
	router.Get("/sign-up", controllers.HandlerSignUpWrap)
	router.Get(settings.USERNAME_PARTIAL_ENDOINT, controllers.MainControllerUsernameInputText)
	router.Post(settings.USERNAME_PARTIAL_ENDOINT, controllers.MainControllerUsernameInputText)
	router.Get(settings.PASSWORD_PARTIAL_ENDOINT, controllers.MainControllerPasswordInputText)
	router.Post(settings.PASSWORD_PARTIAL_ENDOINT, controllers.MainControllerPasswordInputText)
	router.Post(settings.REMEMBER_PARTIAL_ENDOINT, controllers.HandlerCheckboxRemember)
	router.Get(settings.SIGNUP_PARTIAL_ENDOINT, controllers.HandlerSignUpButton)
	router.Post(settings.SIGNUP_PARTIAL_ENDOINT, controllers.HandlerSignUpButton)

}
