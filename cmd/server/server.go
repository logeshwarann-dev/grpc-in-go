package server

import (
	"context"

	pb "github.com/logeshwarann-dev/grpc-in-go/proto/rpcgen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedUserManagementServer
}

func (rpcServer *Server) AddUser(ctx context.Context, newUser *pb.NewUser) (*pb.UserResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}

func (rpcServer *Server) DeleteUser(context.Context, *pb.UserId) (*pb.ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func (rpcServer *Server) GetUser(context.Context, *pb.UserId) (*pb.UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func (rpcServer *Server) UpdateUser(context.Context, *pb.ModifiedUser) (*pb.UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
