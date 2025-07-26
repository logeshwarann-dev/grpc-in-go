package mongodb

import (
	"fmt"
	"testing"
	"time"

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
		FistName: "Logan N",
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

func TestGetUserFromDB(t *testing.T) {
	TestConnectToMongo(t)
	userId := "688463fc2a44647baab22597"
	user, err := GetRecordFromDB(userId)
	if err != nil {
		t.Fatalf("error in fetching user from db: %v", err)
	}
	t.Log("User fetched successfully: ", user)
}

func TestDeleteUser(t *testing.T) {
	TestConnectToMongo(t)
	userId := "687fcf4fcd83c8a621c34382"
	if err := DeleteRecordInDB(userId); err != nil {
		t.Fatalf("error in deleting user in db: %v", err)
	}
	t.Log("User deleted successfully!")
}

func TestUpdateUser(t *testing.T) {
	TestConnectToMongo(t)
	userId := "688463fc2a44647baab22597"
	updateUser := models.User{
		FistName: "Logeshwaran N",
		LastName: "Gounder",
		Age:      25,
		Email:    "logeshwaran@gmail.com",
		PhNo:     "+911234567890",
	}
	err := UpdateRecordInDB(userId, updateUser)
	if err != nil {
		t.Fatalf("error in updating record in db: %v", err)
	}
	t.Log("Update operation successful!")
	time.Sleep(1 * time.Second)
	modifiedUser, err := GetRecordFromDB(userId)
	if err != nil {
		t.Logf("error in fetching record in db: %v", err)
	}
	t.Logf("Updated User: %v", modifiedUser)

}
func TestSplit(t *testing.T) {
	src := "ObjectId('687fa31f98d6042e08feb867')"
	fmt.Println(utils.Split(src, "'")[1])
}
