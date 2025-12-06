package db

import (
	"context"
	"log"
	"os"
	"simple-bank/util"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	pool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	testStore = NewStore(pool)
	os.Exit(m.Run())
}
