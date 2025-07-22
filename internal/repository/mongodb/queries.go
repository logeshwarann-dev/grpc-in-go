package mongodb

import (
	"context"

	"github.com/logeshwarann-dev/grpc-in-go/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOne(coll *mongo.Collection, document any) (*mongo.InsertOneResult, error) {
	return coll.InsertOne(context.TODO(), document)
}

func InsertMany(coll *mongo.Collection, documents []any) (*mongo.InsertManyResult, error) {
	return coll.InsertMany(context.TODO(), documents)
}

func FindMany(coll *mongo.Collection, filter interface{}) ([]models.User, error) {
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return []models.User{}, err
	}
	var docs []models.User
	if err := cursor.All(context.TODO(), &docs); err != nil {
		return []models.User{}, err
	}
	return docs, nil
}

func FindOne() {

}

func Update() {

}

func Delete() {

}

func GetCollection(db *mongo.Client, dbName string, collName string) *mongo.Collection {
	return db.Database(dbName).Collection(collName)
}
