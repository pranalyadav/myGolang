package controller

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/x/mongo/driver/mongocrypt/options"
)

var connectionString = os.Getenv("DB_CONNECTION")
const dbName = "netflix"
const colName = "watchlist"

// most important
var collection *mongo.Collection

// connect with mongoDB
func init() {
	godotenv.Load()
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	client, err :=mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance 
	fmt.Println("Collection instance is ready")
}