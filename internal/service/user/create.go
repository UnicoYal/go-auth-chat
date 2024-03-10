package user

import (
	"context"
	"go-auth-chat/internal/service/model"
)

func (s *serv) Create(ctx context.Context, user *model.User) (int64, error) {
	id, err := s.userRepository.Create(ctx, user)

	if err != nil {
		return 0, err
	}

	return id, nil
}
