# gRPC in Go

A simple demonstration of gRPC implementation in Golang with MongoDB for user management operations. This project is designed for easy understanding and learning purposes.

## Overview

This project implements a User Management Service using gRPC and Protocol Buffers for efficient communication. It demonstrates how to build high-performance microservices with strict message structure and HTTP/2 protocol.

## Features

- **Add User**: Create new user records
- **Get User**: Fetch specific user details
- **Update User**: Modify existing user information
- **Delete User**: Remove user records
- CRUD operations with MongoDB integration
- Protocol Buffers for efficient serialization/deserialization
- HTTP/2 communication protocol

## Tech Stack

- **Language**: Go
- **Framework**: gRPC
- **Database**: MongoDB
- **Communication Protocol**: Protocol Buffers (Protobuf)
- **Transport Protocol**: HTTP/2

## Prerequisites

- Go 1.23
- Docker
- Protocol Buffers compiler (protoc)
- Git

## Installation

### 1. Install Protocol Buffers

**For Windows (PowerShell):**
```bash
winget install protobuf
```

### 2. Install Go Dependencies

```bash
# Install protoc-gen-go
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

# Install protoc-gen-go-grpc
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### 3. Clone and Setup Project

```bash
git clone https://github.com/logeshwarann-dev/grpc-in-go.git
cd grpc-in-go
go mod tidy
```

### 4. Generate Protocol Buffer Files

Navigate to the proto directory and generate the Go files:

```bash
cd proto
protoc --go_out=./rpcgen --go_opt=paths=source_relative --go-grpc_out=./rpcgen --go-grpc_opt=paths=source_relative ./user_mgmt.proto
```

### 5. Setup MongoDB

Ensure Docker is running on your system. Using Docker compose, MongoDB container will be started.
```bash
cd docker/db
docker-compose up -d
```

### 6. Run the Application

```bash
# Start the gRPC server
go run server/main.go

# In another terminal, run the client
go run client/main.go
```

## Protocol Buffers

Protocol Buffers provide several advantages over traditional JSON APIs:

- **Efficient Communication**: Optimized serialization and deserialization
- **Strict Message Structure**: Enforced data contracts
- **Compression**: Built-in compression for large data transfers
- **Binary Format**: Data converted to byte format for transmission
- **Language Agnostic**: Compatible with multiple programming languages
- **Performance**: Significantly faster than JSON-based APIs

## gRPC Benefits

- **High Performance**: Uses HTTP/2 for efficient communication
- **Streaming Support**: Supports client, server, and bidirectional streaming
- **Code Generation**: Automatic client and server code generation
- **Type Safety**: Strong typing with Protocol Buffers
- **Cross-Platform**: Works across different languages and platforms

## Project Structure

```
grpc-in-go/
│
├───cmd
│   ├───client                  # gRPC Client implementation
│   │       client.go
│   │
│   └───server                  # gRPC Server
│           server.go
│
├───docker                              # Docker compose
│   └───db
│           docker-compose.yml
│
├───internal
│   ├───api
│   │   └───handlers                    # Handlers
│   │           add_user.go
│   │           delete_user.go
│   │           get_user.go
│   │           update_user.go
│   │
│   ├───models
│   │       user.go                     # User data models
│   │
│   └───repository
│       └───mongodb
│               db_connect.go           # MongoDB connection
│               queries.go
│               query_exec.go
│               query_test.go
│
├───pkg
│   └───utils
│           string_utils.go
│
└───proto                               # Proto files
    │   user_mgmt.proto                 # Protocol buffer definition
    │
    └───rpcgen
            user_mgmt.pb.go             # Generated protobuf code
            user_mgmt_grpc.pb.go        # Generated gRPC code

```

## API Operations

### User Management Service

The service provides the following RPC methods:

#### AddUser
Creates a new user in the database.

**Request:**
```protobuf
message NewUser {
    string firstName = 1;
    string lastName = 2;
    uint32 age = 3;
    string email = 4;
    string phNo = 5;
}
```

**Response:**
```protobuf
message UserResponse {
    string id = 1;
    NewUser user = 2;
}
```

#### GetUser
Retrieves a specific user by ID.

**Request:**
```protobuf
message UserId {
    string id = 1;
}
```

**Response:**
```protobuf
message UserResponse {
    string id = 1;
    NewUser user = 2;
}
```

#### UpdateUser
Modifies an existing user's information.

**Request:**
```protobuf
message ModifiedUser {
    string id = 1;
    string firstName = 2;
    string lastName = 3;
    uint32 age = 4;
    string email = 5;
    string phNo = 6;
}
```

**Response:**
```protobuf
message UserResponse {
    string id = 1;
    NewUser user = 2;
}
```

#### DeleteUser
Removes a user from the database.

**Request:**
```protobuf
message UserId {
    string id = 1;
}
```

**Response:**
```protobuf
message ResponseMessage {
    string resp = 1;
}
```

## Testing

Test the gRPC services:

### Using gRPC Client
```bash
go run client/main.go
```

## Development Commands

### Regenerate Protocol Buffer Files
```bash
cd proto
protoc --go_out=./rpcgen --go_opt=paths=source_relative --go-grpc_out=./rpcgen --go-grpc_opt=paths=source_relative ./user_mgmt.proto
```

## Learning Resources

- [gRPC Official Documentation](https://grpc.io/docs/)
- [Protocol Buffers Guide](https://developers.google.com/protocol-buffers)
- [Go gRPC Tutorial](https://grpc.io/docs/languages/go/quickstart/)
- [Easy Go gRPC Guide](https://www.bradcypert.com/grpc-fundamentals-with-go/)

## License

This project is licensed under the MIT License.
