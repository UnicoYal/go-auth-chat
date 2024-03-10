package repository

import (
	"context"
	"go-auth-chat/internal/service/model"

	"github.com/golang/protobuf/ptypes/empty"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, id int64, info *model.UserInfo) (*empty.Empty, error)
	Delete(ctx context.Context, id int64) (*empty.Empty, error)
}
