package user

import (
	"context"
	modelServ "go-auth-chat/internal/model"
	"go-auth-chat/internal/repository"
	"go-auth-chat/internal/repository/user/converter"
	modelRepo "go-auth-chat/internal/repository/user/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	tableName             = "users"
	idColumn              = "id"
	emailColumn           = "email"
	nameColumn            = "name"
	roleColumn            = "role"
	passwordColumn        = "password"
	passwordConfirmColumn = "password_confirm"
	createdAtColumn       = "created_at"
	updatedAtColumn       = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, userInfo *modelServ.UserInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(emailColumn, nameColumn, roleColumn, passwordColumn, passwordConfirmColumn).
		Values(userInfo.Email, userInfo.Name, userInfo.Role, userInfo.Password, userInfo.PasswordConfirm).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()

	if err != nil {
		return 0, err
	}

	var userId int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&userId)

	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*modelServ.User, error) {
	builder := sq.Select(idColumn, emailColumn, nameColumn, roleColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"id": id}).
		Limit(1)

	query, args, err := builder.ToSql()

	if err != nil {
		return nil, err
	}

	var user = &modelRepo.User{}
	var email, name, role string
	err = r.db.QueryRow(ctx, query, args...).Scan(&user.ID, &email, &name, &role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	// Инициализация info перед использованием
	info := &modelRepo.UserInfo{
		Email: email,
		Name:  name,
		Role:  role,
	}

	// Присвоение info полю Info структуры user
	user.Info = *info

	return converter.ToUserFromRepo(user), nil
}

func (r *repo) Update(ctx context.Context, id int64, info *modelServ.UserInfo) (*empty.Empty, error) {
	builder := sq.Update(tableName).PlaceholderFormat(sq.Dollar).
		Set(emailColumn, info.Email).
		Set(nameColumn, info.Name).
		Set(roleColumn, info.Role).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return &emptypb.Empty{}, err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (r *repo) Delete(ctx context.Context, id int64) (*empty.Empty, error) {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()

	if err != nil {
		return &emptypb.Empty{}, err
	}

	_, err = r.db.Exec(ctx, query, args...)

	if err != nil {
		return &emptypb.Empty{}, err
	}

	return nil, nil
}
