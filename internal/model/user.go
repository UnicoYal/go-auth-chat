package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserInfo struct {
	Email           string
	Name            string
	Role            string
	Password        string
	PasswordConfirm string
}
