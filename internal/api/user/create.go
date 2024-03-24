package user

import (
	"context"
	"go-auth-chat/internal/converter"
	desc "go-auth-chat/pkg/user/user_v1"
	"log"
)

func (i *Implementation) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	log.Printf("CreateUser")

	id, err := i.userService.Create(ctx, converter.ToUserInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted note with id: %d", id)

	return &desc.CreateUserResponse{
		Id: id,
	}, nil
}
