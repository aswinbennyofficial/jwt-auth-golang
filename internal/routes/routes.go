package routes

import (
	"net/http"

	"github.com/aswinbennyofficial/jwt-auth-golang/internal/controllers"
	"github.com/aswinbennyofficial/jwt-auth-golang/internal/middleware"
)


func Routes(){
	http.HandleFunc("/health",controllers.HandleHealth)

	http.Handle("/welcome",middleware.Authorize(http.HandlerFunc(controllers.HandleWelcome)))
	

	http.HandleFunc("/signin",controllers.HandleSignin)
	http.HandleFunc("/refresh", controllers.HandleRefresh)
	http.HandleFunc("/logout", controllers.HandleLogout)

}