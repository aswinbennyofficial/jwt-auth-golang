package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	//"context"
	//"github.com/aswinbennyofficial/jwt-auth-golang/internal/database"
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
	// DB_URI:=os.Getenv("MONGODB_URI")
	// DB_NAME:=os.Getenv("DB_NAME")
	// DB_COLLECTION_NAME:=os.Getenv("DB_COLLECTION_NAME")
	SERVER_PORT:=os.Getenv("PORT")
	


	// Creating a mongodb client using Db() function in db.go
	// client:=database.DbConnect(DB_URI)
	
	// // Create MongoDB collection obj
	// coll:=client.Database(DB_NAME).Collection(DB_COLLECTION_NAME)
		
	
	// Invoking routes
	routes.Routes()


	// Waste
	//log.Println("coll ", coll)
	


	// Defer disconnecting from the MongoDB client
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		log.Panic("Error while disconnecting MongoDB client: ",err)
	// 	}
	// }()

	if SERVER_PORT==""{
		SERVER_PORT="8080"
	}
	// Starting server
	log.Printf("Server starting in port %s....",SERVER_PORT)
	err=http.ListenAndServe(":"+SERVER_PORT,nil)
	if err!=nil{
		log.Panic("Error while starting server: ",err)
	}
}