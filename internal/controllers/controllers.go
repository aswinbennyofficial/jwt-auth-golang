package controllers

import (
	"net/http"
	"log"
	"fmt"
	//"github.com/aswinbennyofficial/jwt-auth-golang/internal/models"
)


func HandleWelcome(w http.ResponseWriter, r *http.Request){
	// Parse and validate JWT from request
	claims, err := ParseAndValidateJWT(r)
	if err != nil {
		log.Println("ERROR WHILE PARSING/VALIDATING JWT: ", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println("ACCESS APPROVED TO /home")
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))

		
}

func HandleHealth(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello worlds"))
}