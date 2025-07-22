package handlers

import (
	"github.com/logeshwarann-dev/grpc-in-go/internal/models"
	"github.com/logeshwarann-dev/grpc-in-go/internal/repository/mongodb"
)

func CreateUser(newUser models.User) error {
	mongodb.InsertRecordInDB(newUser)
	return nil
}
