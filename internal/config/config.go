package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


func LoadEnv(){
	err:=godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading environment variables file")
		return
	}
}

func LoadWebsiteUrl() string{
	websiteUrl:=os.Getenv("WEBSITE_URL")
	return websiteUrl
}

