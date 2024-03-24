package user

import (
	"go-auth-chat/internal/service"
	desc "go-auth-chat/pkg/user/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
