package routes

import(
	
	"net/http"
	"github.com/aswinbennyofficial/jwt-auth-golang/internal/controllers"
)


func Routes(){
	http.HandleFunc("/health",controllers.HandleHealth)
	http.HandleFunc("/signin",controllers.HandleSignin)
	http.HandleFunc("/welcome", controllers.HandleWelcome)
	http.HandleFunc("/refresh", controllers.HandleRefresh)
	http.HandleFunc("/logout", controllers.HandleLogout)

}