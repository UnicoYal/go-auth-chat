package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	userAPI "go-auth-chat/internal/api/user"
	userRepository "go-auth-chat/internal/repository/user"
	userService "go-auth-chat/internal/service/user"
	desc "go-auth-chat/pkg/user/user_v1"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serverPort = 50050

func dbDSN() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	dbname := os.Getenv("PG_DATABASE_NAME")
	password := os.Getenv("PG_PASSWORD")

	return fmt.Sprintf("host=localhost port=%s user=%s dbname=%s password=%s sslmode=disable", port, user, dbname, password)
}

func main() {
	ctx := context.Background()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	// Создаем пул соединений с базой данных
	pool, err := pgxpool.Connect(ctx, dbDSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	defer pool.Close()

	userRepo := userRepository.NewRepository(pool)
	userSrv := userService.NewService(userRepo)

	s := grpc.NewServer()
	reflection.Register(s)

	desc.RegisterUserV1Server(s, userAPI.NewImplementation(userSrv))

	log.Printf("Server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
