package utility

import(
	"log"
	"net/http"
	"os"
	"time"
	"errors"
	"github.com/aswinbennyofficial/jwt-auth-golang/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string,error) {
	// Hashing the password with the default cost of 10
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    return string(bytes), err
}




func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
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
func GenerateToken(username string) (string, error) {
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