package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID              int64
	Info            UserInfo
	Password        string
	PasswordConfirm string
	CreatedAt       time.Time
	UpdatedAt       sql.NullTime
}

type UserInfo struct {
	Email string
	Name  string
	Role  string
}
