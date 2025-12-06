package main

import (
	"database/sql"
	"log"

	"go-api/cmd/api"
	"go-api/config"
	"go-api/db"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)
	server := api.NewAPIServer(":8000", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {

	err := db.Ping()

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 如果 Ping 成功，打印成功连接的信息
	log.Println("DB successfully connected!")
	log.Printf("Server listening on port: %s", config.Envs.Port)

}
