package routes

import(
	
	"net/http"
	"github.com/aswinbennyofficial/jwt-auth-golang/internal/controllers"
)


func Routes(){
	http.HandleFunc("/health",controllers.HandleHealth)
}