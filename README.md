# grpc-in-go
This project is a simple demonstration of gRPC in Golang for easy understanding &amp; learning purposes


Protocol buffers(Protobuf):
 -> Efficient communication protocol between services
 -> Serialize & Deserialize the payload
 -> Strict Message structure
 -> Uses compression for sending large data 
 -> The data is converted into byte format for send/recieve.
 -> Better than JSON APIs
 -> Compatible with any programming language

gRPC:
 -> Remote Procedure Calls which use Protocol buffers for communication
 -> High-performance
 -> uses Http/2

Project Description:

User Management Service: 
Manages the user details in the MongoDB

Add User: Create new user

Get User: Fetch a specific user

Update User: Modify an existing user

Delete User: Remove an existing user

# Instal Protobuf compiler in Windows:

In Powershell, Execute the below command:

```
winget install protobuf
```

Then, Navigate to your project root directory and Execute the following commands:

```
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

```

# Protobuf compiler commands:

SIMPLE Command:
```
protoc --go_out=. .\<file-name>.proto
```

Go GRPC command:
Syntax: `protoc --go_out=<output dir> --go_opt=paths=source_relative --go-grpc_out=<output dir> --go-grpc_opt=paths=source_relative <proto main file>`
```
cd proto
protoc --go_out=./out --go_opt=paths=source_relative --go-grpc_out=./out --go-grpc_opt=paths=source_relative ./user_mgmt.proto     
```

