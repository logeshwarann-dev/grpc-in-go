package handlers

import (
	"fmt"

	"github.com/logeshwarann-dev/grpc-in-go/internal/repository/mongodb"
	pb "github.com/logeshwarann-dev/grpc-in-go/proto/rpcgen"
)

func DeleteUserUsingId(userId *pb.UserId) (*pb.ResponseMessage, error) {
	if err := mongodb.DeleteRecordInDB(userId.Id); err != nil {
		return nil, fmt.Errorf("unable to delete record: %v", err)
	}

	return &pb.ResponseMessage{
		Resp: "user deleted",
	}, nil
}
