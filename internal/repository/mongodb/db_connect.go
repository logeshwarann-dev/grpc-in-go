package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const mongoDSN = "mongodb://localhost:27017"

var MongoDBClient *mongo.Client

func ConnectToMongo() error {
	clientOptions := options.Client().ApplyURI(mongoDSN)
	db, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("DB connection failed: ", err.Error())
		return err
	}
	if err := db.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Println("DB Ping check failed: ", err.Error())
		return err
	}
	MongoDBClient = db
	log.Println("MongoDB connection successful!")
	return nil
}

func GetMongoDB() *mongo.Client {
	return MongoDBClient
}
