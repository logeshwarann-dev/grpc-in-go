package handlers

import (
	"fmt"

	"github.com/logeshwarann-dev/grpc-in-go/internal/models"
	"github.com/logeshwarann-dev/grpc-in-go/internal/repository/mongodb"
	pb "github.com/logeshwarann-dev/grpc-in-go/proto/rpcgen"
)

func UpdateUserById(user *pb.ModifiedUser) (*pb.UserResponse, error) {
	userId := user.Id
	updatedUser := models.User{
		FistName: user.FirstName,
		LastName: user.LastName,
		Age:      int(user.Age),
		Email:    user.Email,
		PhNo:     user.PhNo,
	}
	if err := mongodb.UpdateRecordInDB(userId, updatedUser); err != nil {
		return nil, fmt.Errorf("unable to update record: %v", err)
	}

	modifiedUser, err := mongodb.GetRecordFromDB(userId)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch updated record: %v", err.Error())
	}

	return &pb.UserResponse{
		Id: userId,
		User: &pb.NewUser{
			FirstName: modifiedUser.Email,
			LastName:  modifiedUser.LastName,
			Age:       uint32(modifiedUser.Age),
			Email:     modifiedUser.Email,
			PhNo:      modifiedUser.PhNo,
		},
	}, nil

}
