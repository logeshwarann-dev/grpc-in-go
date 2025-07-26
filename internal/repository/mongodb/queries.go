package mongodb

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/logeshwarann-dev/grpc-in-go/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func BuildReplacementDoc(updatedUser models.User) bson.D {
	return bson.D{
		{Key: "firstname", Value: updatedUser.FistName},
		{Key: "lastname", Value: updatedUser.LastName},
		{Key: "age", Value: updatedUser.Age},
		{Key: "email", Value: updatedUser.Email},
		{Key: "phno", Value: updatedUser.PhNo},
	}
}

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

func FindDocument(coll *mongo.Collection, filter interface{}) (models.User, error) {
	res := coll.FindOne(context.TODO(), filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return models.User{}, fmt.Errorf("invalid id. record not found")
		}
		return models.User{}, fmt.Errorf("error in finding document: %v", res.Err())
	}
	var userDoc models.MongoDBDocument
	if err := res.Decode(&userDoc); err != nil {
		return models.User{}, fmt.Errorf("error in decoding bson result: %v", err)
	}
	log.Println("User document in MongoDB: ", userDoc)
	return models.User{
		FistName: userDoc.FistName,
		LastName: userDoc.LastName,
		Age:      userDoc.Age,
		Email:    userDoc.Email,
		PhNo:     userDoc.PhNo,
	}, nil
}

func ReplaceDocument(coll *mongo.Collection, filter interface{}, replacement bson.D) (*mongo.UpdateResult, error) {
	return coll.ReplaceOne(context.TODO(), filter, replacement)
}

func DeleteDocument(coll *mongo.Collection, filter interface{}) (*mongo.DeleteResult, error) {
	return coll.DeleteOne(context.TODO(), filter)
}

func GetCollection(db *mongo.Client, dbName string, collName string) *mongo.Collection {
	return db.Database(dbName).Collection(collName)
}
