package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "go-auth-chat/pkg/user/user_v1"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const serverPort = 50050

type server struct {
	desc.UnimplementedUserV1Server
}

func (s *server) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	log.Printf("CreateUser")

	return &desc.CreateUserResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	log.Printf("GetUser")

	return &desc.GetUserResponse{
		UserInfo: &desc.UserInfo{
			Email: gofakeit.Email(),
			Name:  gofakeit.Name(),
			Role:  desc.UserRoles(gofakeit.Number(0, 1)),
		}}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *desc.UpdateUserRequest) (*empty.Empty, error) {
	log.Printf("UpdateUser")

	return &emptypb.Empty{}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *desc.DeleteUserRequest) (*empty.Empty, error) {
	log.Printf("DeleteUser")

	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	desc.RegisterUserV1Server(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
