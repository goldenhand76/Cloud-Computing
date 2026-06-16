package main

import (
	"cloud-computing/users/api"
	db "cloud-computing/users/db/sqlc"
	"cloud-computing/users/util"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config :", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot Connect to db:", err)
	}

	store := db.NewStore(conn)
	runGinServer(config, store)
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
