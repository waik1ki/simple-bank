package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/waik1ki/simple-bank/api"
	db "github.com/waik1ki/simple-bank/db/sqlc"
	"github.com/waik1ki/simple-bank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connet to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
