package converter

import (
	modelServ "go-auth-chat/internal/service/model"
	desc "go-auth-chat/pkg/user/user_v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromService(user *modelServ.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:              user.ID,
		UserInfo:        ToUserInfoFromService(&user.Info),
		Password:        user.Password,
		PasswordConfirm: user.PasswordConfirm,
		CreatedAt:       timestamppb.New(user.CreatedAt),
		UpdatedAt:       updatedAt,
	}
}

func ToUserInfoFromService(info *modelServ.UserInfo) *desc.UserInfo {
	var roleStringToEnum = map[string]desc.UserRoles{
		"admin": desc.UserRoles_admin,
		"user": desc.UserRoles_user,
	}
	// Преобразование строки в значение перечисления UserRoles
	roleEnum := roleStringToEnum[info.Role]

	return &desc.UserInfo{
		Email: info.Email,
		Name:  info.Name,
		Role:  roleEnum,
	}
}
