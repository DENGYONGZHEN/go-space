package main

import (
	"context"
	"log"
	"simple-bank/api"
	db "simple-bank/db/sqlc"
	"simple-bank/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	pool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	testStore := db.NewStore(pool)
	server := api.NewServer(testStore)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
