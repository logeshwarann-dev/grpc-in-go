package mongodb

import (
	"context"
	"fmt"

	"github.com/logeshwarann-dev/grpc-in-go/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func BuildFilterById(userId string) (bson.D, error) {
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, nil
	}
	return bson.D{{Key: "_id", Value: objId}}, nil
}

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

func FindOne(coll *mongo.Collection, filter interface{}) (models.User, error) {
	res := coll.FindOne(context.TODO(), filter)
	if res.Err() != nil {
		return models.User{}, fmt.Errorf("error in finding document: %v", res.Err())
	}
	var user models.User
	if err := res.Decode(&user); err != nil {
		return models.User{}, fmt.Errorf("error in decoding bson result: %v", err)
	}
	return user, nil
}

func Update() {

}

func Delete() {

}

func GetCollection(db *mongo.Client, dbName string, collName string) *mongo.Collection {
	return db.Database(dbName).Collection(collName)
}
