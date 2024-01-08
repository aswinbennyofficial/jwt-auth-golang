package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"context"
	"github.com/aswinbennyofficial/jwt-auth-golang/internal/database"
	"net/http"
	"github.com/aswinbennyofficial/jwt-auth-golang/internal/routes"
	
)



func main(){
	//load env variables
	err:=godotenv.Load(".env")
	if err != nil {
        log.Println("Error loading environment variables file")
		return
    }
	DB_URI:=os.Getenv("MONGODB_URI")
	DB_NAME:=os.Getenv("DB_NAME")
	DB_COLLECTION_NAME:=os.Getenv("DB_COLLECTION_NAME")


	// Creating a mongodb client using Db() function in db.go
	client:=database.DbConnect(DB_URI)
	
	// Create MongoDB collection obj
	coll:=client.Database(DB_NAME).Collection(DB_COLLECTION_NAME)
	log.Println("coll ", coll)
	
	
	// Invoking routes
	routes.Routes()



	// Defer disconnecting from the MongoDB client
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	http.ListenAndServe(":8080",nil)
}