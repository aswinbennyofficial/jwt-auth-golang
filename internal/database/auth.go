package database

import (
	"errors"

	"context"
	"log"

	"github.com/aswinbennyofficial/jwt-auth-golang/internal/models"
	//"github.com/aswinbennyofficial/jwt-auth-golang/internal/utility"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// has collection for login
var coll *mongo.Collection

func InitLoginCollection(client *mongo.Client, dbName, collName string) error {
	
	coll=client.Database(dbName).Collection(collName)
	return nil
}

func AddUserToDb(newuser models.NewUser) error {
	result, err := coll.InsertOne(context.TODO(), newuser)
	if err != nil {
		return(err)
	}
	log.Println(result.InsertedID)
	return nil
}

// TODO
func GetPasswordHashFromDb(username string) (string,error){
	isUserExist,err:=DoesUserExist(username)
	if err!=nil{
		log.Println("GetPasswordHashFromDb() :",err)
		return "",err
	}
	if isUserExist==false{
		return "",errors.New("User does not exist")
	}

	// Creating a filter
	filter := bson.D{{"username", username}}

	// Finding a single document
	var result models.NewUser

	err=coll.FindOne(context.TODO(),filter).Decode(&result)
	if err != nil {
		log.Println("GetPasswordHashFromDb() ",err)
		return "",err
	}

	return result.Password,nil
}

func DoesUserExist(username string) (bool,error) {
	// TODO
	opts := options.Count().SetHint("_id_")
	filter := bson.D{{"username", username}}
	// (ctx, filter,opts) 
	count, err := coll.CountDocuments(context.TODO(), filter, opts)
	if err != nil {
		log.Println(err)
		return true,err
	}
	if count == 0 {
		return false,nil
	}else{
		return true,nil
	}
}
