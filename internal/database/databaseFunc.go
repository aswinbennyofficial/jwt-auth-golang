package database

import(
	"errors"
)

// TODO
func GetPasswordHashFromDb(username string) (string,error){
	if(username=="aswinbenny"){
	return "password123",nil
	}
	return "",errors.New("User not found")
}

