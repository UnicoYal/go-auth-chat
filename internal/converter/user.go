package converter

import (
	model "go-auth-chat/internal/model"
	desc "go-auth-chat/pkg/user/user_v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:              user.ID,
		UserInfo:        ToUserInfoFromService(&user.Info),
		CreatedAt:       timestamppb.New(user.CreatedAt),
		UpdatedAt:       updatedAt,
	}
}

func ToUserInfoFromService(info *model.UserInfo) *desc.UserInfo {
	var roleStringToEnum = map[string]desc.UserRoles{
		"admin": desc.UserRoles_admin,
		"user":  desc.UserRoles_user,
	}
	// Преобразование строки в значение перечисления UserRoles
	roleEnum := roleStringToEnum[info.Role]

	return &desc.UserInfo{
		Email: info.Email,
		Name:  info.Name,
		Role:  roleEnum,
		Password:        info.Password,
		PasswordConfirm: info.PasswordConfirm,
	}
}

func ToUserInfoFromDesc(info *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Email:           info.Email,
		Name:            info.Name,
		Role:            info.Role.String(),
		Password:        info.Password,
		PasswordConfirm: info.PasswordConfirm,
	}
}

func ToUserFromDesc(user *desc.User) *model.User {
	return &model.User{
		ID:   user.Id,
		Info: *ToUserInfoFromDesc(user.UserInfo),
	}
}
