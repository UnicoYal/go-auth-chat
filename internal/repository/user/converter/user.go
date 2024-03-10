package converter

import (
	modelServ "go-auth-chat/internal/service/model"
	modelRepo "go-auth-chat/internal/repository/user/model"
)

func ToUserFromRepo(user *modelRepo.User) *modelServ.User {
	return &modelServ.User{
		ID:        user.ID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserInfoFromRepo(info modelRepo.UserInfo) modelServ.UserInfo {
	return modelServ.UserInfo{
		Email:   info.Email,
		Name: info.Name,
		Role: info.Role,
	}
}
