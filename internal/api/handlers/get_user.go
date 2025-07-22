package handlers

import (
	"github.com/logeshwarann-dev/grpc-in-go/internal/repository/mongodb"
	pb "github.com/logeshwarann-dev/grpc-in-go/proto/rpcgen"
)

func GetUserbyId(userIdReq *pb.UserId) (*pb.NewUser, error) {
	userDoc, err := mongodb.GetRecordFromDB(userIdReq.Id)
	if err != nil {
		return nil, err
	}
	return &pb.NewUser{
		FirstName: userDoc.FistName,
		LastName:  userDoc.LastName,
		Age:       uint32(userDoc.Age),
		Email:     userDoc.Email,
		PhNo:      userDoc.PhNo,
	}, nil
}
