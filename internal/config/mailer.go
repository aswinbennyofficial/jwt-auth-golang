package config

import (
	"os"
	"log"
	
)

func LoadSMTPUsername() string{
	SMTP_USERNAME:= os.Getenv("SMTP_USERNAME")
	if SMTP_USERNAME==""{
		log.Println("Error loading SMTP_USERNAME in config.GetSMTPUsername()")
		return ""
	}
	return SMTP_USERNAME
}

func LoadSMTPPassword() string{
	SMTP_PASSWORD:= os.Getenv("SMTP_PASSWORD")
	if SMTP_PASSWORD==""{
		log.Println("Error loading SMTP_PASSWORD in config.GetSMTPPassword()")
		return ""
	}
	return SMTP_PASSWORD
}


func LoadSMTPServer() string{
	SMTP_SERVER:= os.Getenv("SMTP_HOST")
	if SMTP_SERVER==""{
		log.Println("Error loading SMTP_SERVER in config.GetSMTPServer()")
		return ""
	}
	return SMTP_SERVER
}

func LoadSMTPPort() string{
	SMTP_PORT:= os.Getenv("SMTP_PORT")
	if SMTP_PORT==""{
		log.Println("Error loading SMTP_PORT in config.GetSMTPPort()")
		return "587"
	}
	return SMTP_PORT
}

func LoadSMTPReply_to() string{
	SMTP_REPLY_TO:= os.Getenv("REPLY_TO")
	if SMTP_REPLY_TO==""{
		log.Println("Error loading SMTP_REPLY_TO in config.GetSMTPReply_to()")
		return ""
	}
	return SMTP_REPLY_TO
}

func LoadSMTPFromEmail() string{
	SMTP_FROM_EMAIL:= os.Getenv("FROM_EMAIL")
	if SMTP_FROM_EMAIL==""{
		log.Println("Error loading SMTP_FROM_EMAIL in config.GetSMTPFromEmail()")
		return ""
	}
	return SMTP_FROM_EMAIL
}