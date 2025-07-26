package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/logeshwarann-dev/grpc-in-go/internal/api/handlers"
	"github.com/logeshwarann-dev/grpc-in-go/internal/repository/mongodb"
	pb "github.com/logeshwarann-dev/grpc-in-go/proto/rpcgen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {

	if err := mongodb.ConnectToMongo(); err != nil {
		log.Fatalf("Failed connecting to DB: %v", err.Error())
	}
	log.Println("DB Connection successful!")
}

type Server struct {
	pb.UnimplementedUserManagementServer
}

func (rpcServer *Server) AddUser(ctx context.Context, newUser *pb.NewUser) (*pb.UserResponse, error) {
	log.Println("Recieved Request for adding User: ", newUser)
	response, err := handlers.CreateUser(newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%v", err))
	}
	return response, nil
}

func (rpcServer *Server) DeleteUser(ctx context.Context, userId *pb.UserId) (*pb.ResponseMessage, error) {

	resp, err := handlers.DeleteUserUsingId(userId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%v", err))
	}
	return resp, nil
}

func (rpcServer *Server) GetUser(ctx context.Context, userIdReq *pb.UserId) (*pb.UserResponse, error) {
	response, err := handlers.GetUserbyId(userIdReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%v", err))
	}
	return &pb.UserResponse{
		Id:   userIdReq.Id,
		User: response,
	}, nil
}

func (rpcServer *Server) UpdateUser(ctx context.Context, updatedUser *pb.ModifiedUser) (*pb.UserResponse, error) {
	modifiedUser, err := handlers.UpdateUserById(updatedUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%v", err))
	}
	return modifiedUser, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("unable to listen on port: ", err.Error())
	}

	log.Println("TCP Port 8080 is assigned.")

	grpcServer := grpc.NewServer()

	pb.RegisterUserManagementServer(grpcServer, &Server{})
	log.Println("gRPC Server yet to start..")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalln("unable to start grpc server: ", err)
	}

	log.Println("gRPC Server has stopped!")
}
