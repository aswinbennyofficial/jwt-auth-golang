package config

import (
	"os"
	"log"
)

func LoadJWTSecret() []byte{
	JWT_SECRET:= os.Getenv("JWT_SECRET")
	if JWT_SECRET==""{
		log.Println("Error loading JWT_SECRET in config.LoadJWTSecret()")
		return []byte("")
	}
	return []byte(JWT_SECRET)
}
