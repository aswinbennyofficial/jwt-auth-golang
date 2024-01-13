package utility

import (
	"errors"
	"log"
	"math/rand"

	"net/http"
	"net/smtp"

	"time"

	"github.com/aswinbennyofficial/jwt-auth-golang/internal/config"
	"github.com/aswinbennyofficial/jwt-auth-golang/internal/database"
	"github.com/aswinbennyofficial/jwt-auth-golang/internal/models"

	"github.com/golang-jwt/jwt/v5"

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
		return config.LoadJWTSecret(), nil
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
	expirationTime := time.Now().Add(time.Duration(config.LoadJwtExpiresIn()) * time.Minute)

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
	signedToken, err := noSignedToken.SignedString(config.LoadJWTSecret())
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRandomString(length int) string{
	
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    seed := rand.NewSource(time.Now().UnixNano())
    random := rand.New(seed)

    result := make([]byte, length)
    for i := range result {
        result[i] = charset[random.Intn(len(charset))]
    }
    return string(result)
}

func SendMagicURLToUser(username string) error{
	// TODO

	// Send magic string to user's email
	magicString,err:=database.GetMagicString(username)
	if err!=nil{
		log.Println("Error while getting magic string from database: ",err)
		return err
	}

	// SMTP server Credentials from .env file
	SMTP_USERNAME := config.LoadSMTPUsername()
	SMTP_PASSWORD := config.LoadSMTPPassword()
	SMTP_HOST := config.LoadSMTPServer()
	FROM_EMAIL := config.LoadSMTPFromEmail()
	SMTP_PORT := config.LoadSMTPPort()
	REPLY_TO := config.LoadSMTPReply_to()
	WEBSITE_URL:=config.LoadWebsiteUrl()
	
	log.Println("SMTP CREDS init ",SMTP_USERNAME, " ", SMTP_PASSWORD," ",SMTP_HOST )
	
	// Setup authentication variable
	auth:=smtp.PlainAuth("",SMTP_USERNAME,SMTP_PASSWORD,SMTP_HOST)

	

	if REPLY_TO==""{
		REPLY_TO=FROM_EMAIL
	}

	// mail
	// TODO
	verifyLink:=WEBSITE_URL+"/verify?username="+username+"&magicString="+magicString
	subject:="Verify your email address"
	body := `
<html>
  <head>
    <style>
      body {
        font-family: 'Arial', sans-serif;
        background-color: #f4f4f4;
        text-align: center;
        margin: 30px;
      }
      h2 {
        color: #333;
      }
      a {
        display: inline-block;
        padding: 10px 20px;
        font-size: 16px;
        text-decoration: none;
        background-color: #4CAF50;
        color: #fff;
        border-radius: 5px;
      }
    </style>
  </head>
  <body>
    <h2>Hi, Verify your email on...</h2>
    <p>Click the button below to verify your email:</p>
    <a href="` + verifyLink + `">Verify Email</a>
  </body>
</html>
`
	
	

	var msg []byte
	//For basic text
	// msg = []byte(
	// 	"Reply-To: "+reply_to+"\r\n"+
	// 	"Subject: "+subject+"\r\n" +
	// 	"\r\n" +
	// 	body+"\r\n")

	//For rich html support
	msg = []byte(
		"From: "+FROM_EMAIL+"\r\n"+
		"Reply-To: " + REPLY_TO + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n" +
			"\r\n" +
			body + "\r\n")

	recieverEmail := []string{username} 
	
	// send the mail
	err = smtp.SendMail(SMTP_HOST+":"+SMTP_PORT, auth, FROM_EMAIL, recieverEmail, msg)

	// handling the errors
	if err != nil {
		log.Println(err)
		return err
	}


	log.Println("Successfully sent verification mail")

	return nil
}