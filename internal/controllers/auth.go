package controllers

import (
	"encoding/json"
	// "errors"
	"log"
	"net/http"
	// "os"
	"time"

	"github.com/aswinbennyofficial/jwt-auth-golang/internal/database"
	"github.com/aswinbennyofficial/jwt-auth-golang/internal/models"
	"github.com/aswinbennyofficial/jwt-auth-golang/internal/utility"
	// "github.com/golang-jwt/jwt/v5"
	// "github.com/joho/godotenv"
)








func HandleSignin(w http.ResponseWriter, r *http.Request){

	var creds models.Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the expected password from our in memory map
	expectedPasswordHash,err := database.GetPasswordHashFromDb(creds.Username)
	if err != nil {
		log.Println("Error while getting password from database: ",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if utility.CheckPasswordHash(creds.Password, expectedPasswordHash) == false{
		log.Println("Incorrect password")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create a new JWT token
	signedToken, err := utility.GenerateToken(creds.Username)
	if err != nil {
		log.Println("ERROR OCCURRED WHILE CREATING JWT TOKEN: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("JWT GENERATED FOR %s",creds.Username)

	// Setting expiration time for cookie
	expirationTime := time.Now().Add(5 * time.Minute)

	http.SetCookie(w, &http.Cookie{
		Name:    "JWtoken",
		Value:   signedToken,
		Expires: expirationTime,
	})


}

// TODO ENV refresh time
func HandleRefresh(w http.ResponseWriter, r *http.Request){
	// Parse and validate JWT from request
	claims, err := utility.ParseAndValidateJWT(r)
	if err != nil {
		log.Println("ERROR WHILE PARSING/VALIDATING JWT: ", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}


	// Generate a new JWT token
	signedToken, err := utility.GenerateToken(claims.Username)
	if err != nil {
		log.Println("ERROR OCCURRED WHILE CREATING JWT TOKEN: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Setting JWT claims
	expirationTime := time.Now().Add(5 * time.Minute)

	http.SetCookie(w, &http.Cookie{
		Name:    "JWtoken",
		Value:   signedToken,
		Expires: expirationTime,
	})

	log.Println("TOKEN REFRESH SUCCESSFUL")
}


func HandleLogout(w http.ResponseWriter, r *http.Request){
	log.Println("LOGOUT SUCCESSFUL")
	http.SetCookie(w, &http.Cookie{
		Name:    "JWtoken",
		Expires: time.Now(),
	})
}


func HandleSignup(w http.ResponseWriter, r *http.Request){
	var user models.NewUser
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isUserExist,err:= database.DoesUserExist(user.Username)
	if(err!=nil){
		log.Println("Error while checking if user exists: ",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if isUserExist{
		log.Println("User already exists")
		w.Write([]byte("User already exists"))
		w.WriteHeader(http.StatusConflict)
		return
	}
	log.Println("User does not exist")
	w.WriteHeader(http.StatusOK)

	// Hashing the password with the default cost of 10
	hashedPassword, err := utility.HashPassword(user.Password)
	if err != nil {
		log.Println("Error while hashing password: ",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Replacing existing password with hashed password
	user.Password=hashedPassword

	// Adding user to database
	err = database.AddUserToDb(user)
	if err != nil {
		log.Println("Error while adding user to database: ",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("User added to database")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User added to database"))

	
}