package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aswinbennyofficial/jwt-auth-golang/internal/database"
	"github.com/aswinbennyofficial/jwt-auth-golang/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// TODO
func HashPassword(password string) string {
	return password
}

func GetJWTKey() []byte{
	// Getting environment variables
	err := godotenv.Load(".env")
	if err != nil {
        log.Printf("Error loading environment variables file in controllers.HandleSignin()")
		return nil
    }

	JWT_KEY:=os.Getenv("JWT_KEY")
	if JWT_KEY==""{
		log.Println("Error loading JWT_KEY in controllers.HandleSignin()")
		return nil
	}
	return []byte(JWT_KEY)
}


// Parse and validate JWT from request
func ParseAndValidateJWT(r *http.Request) (*models.Claims, error) {
	cookieVar, err := r.Cookie("JWtoken")
	if err != nil {
		return nil, err
	}

	// Get JWT string from cookie
	JWTstring := cookieVar.Value

	// Initialize a new instance of Claims
	claims := &models.Claims{}

	// Parse the JWT string and store in claims
	tokenVar, err := jwt.ParseWithClaims(JWTstring, claims, func(token *jwt.Token) (interface{}, error) {
		return GetJWTKey(), nil
	})

	if err != nil {
		return nil, err
	}

	if !tokenVar.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}



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
	}

	if expectedPasswordHash!=HashPassword(creds.Password){
		log.Println("Incorrect password")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Declare the expiration time of the token
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &models.Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	JWT_KEY:=GetJWTKey()
	tokenString, err := token.SignedString(JWT_KEY)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		log.Println("Error while creating JWT:",JWT_KEY,err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Setting httpCookie
	http.SetCookie(w, &http.Cookie{
		Name:    "JWtoken",
		Value:   tokenString,
		Expires: expirationTime,
	})


}

func HandleRefresh(w http.ResponseWriter, r *http.Request){
	
}

func HandleLogout(w http.ResponseWriter, r *http.Request){
	
}
