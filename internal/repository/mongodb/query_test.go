package mongodb

import (
	"fmt"
	"testing"

	"github.com/logeshwarann-dev/grpc-in-go/internal/models"
	"github.com/logeshwarann-dev/grpc-in-go/pkg/utils"
)

func TestConnectToMongo(t *testing.T) {
	if err := ConnectToMongo(); err != nil {
		t.Fatalf("Failed connecting to DB: %v", err.Error())
	}
	t.Log("DB Connection successful!")
}

func TestAddUserInDB(t *testing.T) {
	TestConnectToMongo(t)
	newUser := models.User{
		FistName: "Logesh",
		LastName: "Gounder",
		Age:      24,
		Email:    "logesh@gmail.com",
		PhNo:     "+911234567890",
	}
	userId, err := InsertRecordInDB(newUser)
	if err != nil {
		t.Fatalf("Error while insert in record: %v", err.Error())
	}
	t.Logf("Record inserted. User Id: %v", userId)
}

func TestSplit(t *testing.T) {
	src := "ObjectId('687fa31f98d6042e08feb867')"
	fmt.Println(utils.Split(src, "'")[1])
}
