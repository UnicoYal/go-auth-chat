package user

import (
	"context"
	"go-auth-chat/internal/model"
)

func (s *serv) Create(ctx context.Context, userInfo *model.UserInfo) (int64, error) {
	id, err := s.userRepository.Create(ctx, userInfo)

	if err != nil {
		return 0, err
	}

	return id, nil
}
