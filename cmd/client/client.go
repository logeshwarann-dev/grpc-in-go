package main

import (
	"context"
	"log"
	"time"

	pb "github.com/logeshwarann-dev/grpc-in-go/proto/rpcgen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	ServerHost = "localhost"
	ServerPort = "8080"
)

func main() {
	serverAddr := ServerHost + ":" + ServerPort
	grpcConn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to gRPC server at %v: %v", serverAddr, err)
	}
	defer grpcConn.Close()
	log.Println("Connected to Server..")
	client := pb.NewUserManagementClient(grpcConn)

	//-------CREATE NEW USER-----
	newUser := &pb.NewUser{
		FirstName: "Logan N",
		LastName:  "Gounder",
		Age:       24,
		Email:     "logesh@gmail.com",
		PhNo:      "+911234567890",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	resp, err := client.AddUser(ctx, newUser)
	if err != nil {
		log.Fatalf("error in adding user: %v", err)
	}
	log.Println("Response from AddUser: ", resp)

	//---------GET USER--------
	userId := &pb.UserId{
		Id: resp.Id,
	}
	fetchResp, err := client.GetUser(ctx, userId)
	if err != nil {
		log.Fatalf("error in fetching userId: %v : %v", userId, err)
	}
	log.Println("Response from GetUser(): ", fetchResp)
}
