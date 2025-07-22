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
	response, err := handlers.CreateUser(newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%v", err))
	}
	return response, nil
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

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("unable to listen on port: ", err.Error())
	}

	log.Println("TCP Port 8080 is assigned.")

	grpcServer := grpc.NewServer()

	pb.RegisterUserManagementServer(grpcServer, &Server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalln("unable to start grpc server: ", err)
	}

	log.Println("gRPC Server has started!")
}
