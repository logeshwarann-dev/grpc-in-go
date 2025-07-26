package mongodb

import (
	"fmt"
	"log"

	"github.com/logeshwarann-dev/grpc-in-go/internal/models"
	"github.com/logeshwarann-dev/grpc-in-go/pkg/utils"
)

const DatabaseName = "userdb"
const CollectionName = "users"

func InsertRecordInDB(newRecord models.User) (string, error) {
	db, err := GetMongoDB()
	if err != nil {
		log.Println("error in mongodb client: ", err)
		return "", err
	}
	coll := GetCollection(db, DatabaseName, CollectionName)
	res, err := InsertOne(coll, newRecord)
	if err != nil {
		log.Println("Insert record failed: ", err.Error())
		return "", fmt.Errorf("error while insert record in db: %v", err.Error())
	}
	id := ExtractObjectId(fmt.Sprintf("%v", res.InsertedID))
	log.Println("New Record has been added successfully\nObjectId: ", id)
	return id, nil
}

func GetRecordFromDB(userId string) (models.User, error) {
	db, err := GetMongoDB()
	if err != nil {
		log.Println("error in mongodb client: ", err)
		return models.User{}, err
	}
	coll := GetCollection(db, DatabaseName, CollectionName)
	bsonFilter, err := BuildFilterById(userId)
	if err != nil {
		log.Println("error in filter query: ", err.Error())
		return models.User{}, fmt.Errorf("error in filter query: %v", err.Error())
	}
	userRecord, err := FindOne(coll, bsonFilter)
	if err != nil {
		log.Println("error in finding document: ", err)
		return models.User{}, err
	}
	return userRecord, nil
}

func UpdateRecordInDB(userId string, updatedUser models.User) error {
	db, err := GetMongoDB()
	if err != nil {
		return fmt.Errorf("error in mongodb client: %v", err)
	}
	coll := GetCollection(db, DatabaseName, CollectionName)
	updateFilter, err := BuildFilterById(userId)
	if err != nil {
		return fmt.Errorf("error in filterById: %v", err)
	}
	replacementBson := BuildReplacementDoc(updatedUser)
	updatedRes, err := ReplaceDocument(coll, updateFilter, replacementBson)
	if err != nil {
		return fmt.Errorf("error in replacing document in db: %v", updatedRes)
	}
	log.Println("Total Docs Matched: ", updatedRes.MatchedCount, "\n Total Docs updated: ", updatedRes.ModifiedCount)
	return nil
}

func DeleteRecordInDB(userId string) error {
	db, err := GetMongoDB()
	if err != nil {
		return fmt.Errorf("error in mongo client: %v", err)
	}
	coll := GetCollection(db, DatabaseName, CollectionName)
	filter, err := BuildFilterById(userId)
	if err != nil {
		return fmt.Errorf("error in filterById: %v", err)
	}
	deleteRes, err := DeleteDocument(coll, filter)
	if err != nil {
		return fmt.Errorf("error in deleting document in db: %v", deleteRes)
	}
	log.Println("Total Documents Deleted: ", deleteRes.DeletedCount)
	return nil
}

func ExtractObjectId(objId string) string {
	return utils.Split(objId, "\"")[1]
}
