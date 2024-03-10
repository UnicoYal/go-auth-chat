package user

import (
	"context"
	"go-auth-chat/internal/service/model"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *serv) Update(ctx context.Context, id int64, info *model.UserInfo) (*empty.Empty, error) {
	_, err := s.userRepository.Update(ctx, id, info)

	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
