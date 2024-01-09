package main

import (
	"context"
	"log"
	"net/http"
	"github.com/aswinbennyofficial/jwt-auth-golang/internal/config"
	"github.com/aswinbennyofficial/jwt-auth-golang/internal/database"
	"github.com/aswinbennyofficial/jwt-auth-golang/internal/routes"
	
)



func main(){

	//load env variables
	config.LoadEnv()

	DB_URI:=config.LoadMongoDBURI()
	DB_FOR_AUTH:=config.LoadMongoDBNameAuth()
	DB_COLLECTION_FOR_AUTH :=config.LoadMongoDBCollectionNameAuth()
	
	


	// Creating a mongodb client using Db() function in db.go
	client:=database.DbConnect(DB_URI)
	
	database.InitLoginCollection(client,DB_FOR_AUTH,DB_COLLECTION_FOR_AUTH)
	
	
	
	// Invoking routes
	routes.Routes()


	
	


	//Defer disconnecting from the MongoDB client
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Panic("Error while disconnecting MongoDB client: ",err)
		}
	}()

	// Starting server
	SERVER_PORT:=config.LoadServerPort()
	log.Printf("Server starting in port %s....",SERVER_PORT)
	err:=http.ListenAndServe(":"+SERVER_PORT,nil)
	if err!=nil{
		log.Panic("Error while starting server: ",err)
	}
}