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

// Generate a new JWT token
func generateToken(username string) (string, error) {
	// Setting expiration time to be 5 minutes from now
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &models.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declaring token with header and payload
	noSignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create a complete signed JWT
	signedToken, err := noSignedToken.SignedString(GetJWTKey())
	if err != nil {
		return "", err
	}

	return signedToken, nil
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

	// Compare the stored hashed password, with the hashed version of the password that was received
	if expectedPasswordHash!=HashPassword(creds.Password){
		log.Println("Incorrect password")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create a new JWT token
	signedToken, err := generateToken(creds.Username)
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
	claims, err := ParseAndValidateJWT(r)
	if err != nil {
		log.Println("ERROR WHILE PARSING/VALIDATING JWT: ", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// New token is not issued until 30s of expiration time
	if time.Until(claims.ExpiresAt.Time) > 240*time.Second {
		log.Println("NEW REFRESH ONLY BEFORE 4 mins OF EXPIRY")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Generate a new JWT token
	signedToken, err := generateToken(claims.Username)
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
