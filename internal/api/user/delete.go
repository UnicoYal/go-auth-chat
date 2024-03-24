package user

import (
	"context"
	desc "go-auth-chat/pkg/user/user_v1"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) DeleteUser(ctx context.Context, req *desc.DeleteUserRequest) (*empty.Empty, error) {
	_, err := i.userService.Delete(ctx, req.GetId())
	if err != nil {
		return &emptypb.Empty{}, err
	}

	log.Printf("DeleteUser")

	return &emptypb.Empty{}, nil
}
