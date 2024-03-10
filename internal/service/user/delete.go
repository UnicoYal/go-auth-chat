package user

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *serv) Delete(ctx context.Context, id int64) (*empty.Empty, error) {
	_, err := s.userRepository.Delete(ctx, id)

	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
