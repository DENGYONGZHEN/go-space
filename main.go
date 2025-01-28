package main

import (
	"context"
	"log"
	"simple-bank/api"
	db "simple-bank/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

const serverAddress = "0.0.0.0:8080"

func main() {

	dbURL := "postgresql://deng:deng@192.168.193.158:5432/simple_bank?sslmode=disable"
	// dbURL := "postgresql://deng:deng@localhost:5432/simple_bank?sslmode=disable"
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	testStore := db.NewStore(pool)
	server := api.NewServer(&testStore)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
