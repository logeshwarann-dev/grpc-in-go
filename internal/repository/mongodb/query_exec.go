package mongodb

import (
	"fmt"
	"log"

	"github.com/logeshwarann-dev/grpc-in-go/internal/models"
)

const DatabaseName = "userdb"
const CollectionName = "users"

func InsertRecordInDB(newRecord models.User) (string, error) {
	db := GetMongoDB()
	coll := GetCollection(db, DatabaseName, CollectionName)
	res, err := InsertOne(coll, newRecord)
	if err != nil {
		log.Println("Insert record failed: ", err.Error())
		return "", fmt.Errorf("error while insert record in db: %v", err.Error())
	}
	log.Println("New Record Id: ", res.InsertedID)
	id := fmt.Sprintf("%v", res.InsertedID)
	return id, nil
}

func GetRecordFromDB(userId int) (models.User, error) {
	// db := GetMongoDB()

	return models.User{}, nil
}

func UpdateRecordInDB(userId int, updatedUser models.User) error {
	// db := GetMongoDB()
	return nil
}

func DeleteRecordInDB(userId int) error {
	// db := GetMongoDB()
	return nil
}
