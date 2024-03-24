package user

import (
	"context"
	"go-auth-chat/internal/converter"
	desc "go-auth-chat/pkg/user/user_v1"
	"log"
)

func (i *Implementation) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	user, err := i.userService.Get(ctx, req.GetId())

	if err != nil {
		log.Print(err)
		return nil, err
	}

	log.Printf("GetUser")

	return &desc.GetUserResponse{
		UserInfo: converter.ToUserInfoFromService(&user.Info),
	}, nil
}
