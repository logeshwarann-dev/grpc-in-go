package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	FistName string
	LastName string
	Age      int
	Email    string
	PhNo     string
}

type MongoDBDocument struct {
	Id       primitive.ObjectID `bson:"_id, omitempty"`
	FistName string             `bson:"firstname"`
	LastName string             `bson:"lastname"`
	Age      int                `bson:"age"`
	Email    string             `bson:"email"`
	PhNo     string             `bson:"phno"`
}
