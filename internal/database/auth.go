package database

import (
	"errors"

	"github.com/aswinbennyofficial/jwt-auth-golang/internal/utility"
)

// TODO
func GetPasswordHashFromDb(username string) (string,error){
	if(username=="aswinbenny"){
	return utility.HashPassword("password123")
	}
	return "",errors.New("User not found")
}

