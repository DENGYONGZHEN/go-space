package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	ctx := context.Background()
	db, err := pgx.Connect(ctx, "postgresql://deng:deng@192.168.193.158:5432/simple_bank?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(ctx)
	testQueries = New(db)
	os.Exit(m.Run())
}
