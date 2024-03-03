package main

import (
	"context"
	"database/sql"
	"log"
	"time"

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
		log.Fatal("Cannot connect to database: %v", err)
	}

	defer pool.Close()

	res, err := pool.Exec(ctx, "INSERT INTO users (email, name, role, password, password_confirm) VALUES($1, $2, $3, $4, $5)", gofakeit.Email(), gofakeit.Name(), "admin", "12345", "12345")

	if err != nil {
		log.Fatal("Error while inserting: %v", err)
	}

	log.Printf("Rows affected: %d", res.RowsAffected())

	rows, err := pool.Query(ctx, "SELECT * FROM users")

	if err != nil {
		log.Fatal("Error while selecting %v", err)
	}

	defer rows.Close()

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

	res, err = pool.Exec(ctx, "UPDATE users SET email=$1, updated_at=$2 WHERE id=$3", "yagmurov.igor@", time.Now(),id)

	if err != nil {
		log.Fatalf("Error while updating: %v", err)
	}

	log.Printf("Rows affected: %d", res.RowsAffected())

	row := pool.QueryRow(ctx, "SELECT * FROM users WHERE (id = $1)", id)

	err = row.Scan(&id, &email, &name, &role, &password, &password_confirm, &createdAt, &updatedAt)
	if err != nil {
		log.Fatalf("Cannot find row: %v", err)
	}

	log.Printf("id: %d, email: %s, name: %s, role: %s, created_at: %v, updated_at: %v\n", id, email, name, role, createdAt, updatedAt)
}
