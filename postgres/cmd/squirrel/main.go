package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	dbDSN = "host=localhost port=54322 dbname=go_auth user=igor password=12345 sslmode=disable"
)

func main() {
	ctx := context.Background()

	pool, err := pgxpool.Connect(ctx, dbDSN)

	if err != nil {
		log.Fatalf("Error while connecting: %v", err)
	}

	builderInsert := sq.Insert("users").
		PlaceholderFormat(sq.Dollar).
		Columns("email", "name", "role", "password", "password_confirm").
		Values(gofakeit.Email(), gofakeit.Name(), "admin", "12345", "12345").
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()

	if err != nil {
		log.Fatalf("Cannot convert insert to sql: %v", err)
	}

	var userId int

	err = pool.QueryRow(ctx, query, args...).Scan(&userId)

	if err != nil {
		log.Fatalf("Error while exec insert: %v", err)
	}

	builderSelect := sq.Select("*").From("users").OrderBy("id DESC")

	query, args, err = builderSelect.ToSql()

	rows, err := pool.Query(ctx, query, args...)
	if err != nil {
		log.Fatalf("Error while selecting: %v", err)
	}

	var (
		id               int
		email            string
		name             string
		role             string
		password         string
		password_confirm string
		createdAt        time.Time
		updatedAt        sql.NullTime
	)

	for rows.Next() {
		err = rows.Scan(&id, &email, &name, &role, &password, &password_confirm, &createdAt, &updatedAt)

		if err != nil {
			log.Fatalf("Error while scanning %v", err)
		}

		log.Printf("id: %d, email: %s, name: %s, role: %s, created_at: %v, updated_at: %v\n", id, email, name, role, createdAt, updatedAt)
	}

	builderUpdate := sq.Update("users").
		PlaceholderFormat(sq.Dollar).
		Set("email", "yagmurov_igor@").
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": userId})

	query, args, err = builderUpdate.ToSql()

	res, err := pool.Exec(ctx, query, args...)

	if err != nil {
		log.Fatalf("Error while updating: %v", err)
	}

	log.Printf("Rows affected: %d", res.RowsAffected())

	builderGetOne := sq.Select("*").
		From("users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": userId})

	query, args, err = builderGetOne.ToSql()

	res_user := pool.QueryRow(ctx, query, args...)

	res_user.Scan(&id, &email, &name, &role, &password, &password_confirm, &createdAt, &updatedAt)
	log.Printf("id: %d, email: %s, name: %s, role: %s, created_at: %v, updated_at: %v\n", id, email, name, role, createdAt, updatedAt)
}
