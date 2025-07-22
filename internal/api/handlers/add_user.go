package handlers

import (
	"fmt"

	"github.com/logeshwarann-dev/grpc-in-go/internal/models"
	"github.com/logeshwarann-dev/grpc-in-go/internal/repository/mongodb"
	pb "github.com/logeshwarann-dev/grpc-in-go/proto/rpcgen"
)

func CreateUser(user *pb.NewUser) (*pb.UserResponse, error) {
	newUser := models.User{
		FistName: user.FirstName,
		LastName: user.LastName,
		Age:      int(user.Age),
		Email:    user.Email,
		PhNo:     user.PhNo,
	}
	userId, err := mongodb.InsertRecordInDB(newUser)
	if err != nil {
		return nil, fmt.Errorf("unable to Add new user: %v", err.Error())
	}
	userRes := pb.UserResponse{
		Id:   userId,
		User: user,
	}
	return &userRes, nil
}
