package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aswinbennyofficial/jwt-auth-golang/internal/models"

)


func HandleWelcome(w http.ResponseWriter, r *http.Request){
	
	claims, ok := r.Context().Value("claims").(*models.Claims)
	if !ok {
		log.Println("Claims not found in context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("ACCESS APPROVED TO /WELCOME")
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))

		
}

func HandleHealth(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello worlds"))
}