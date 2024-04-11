package main

import (
	"database/sql"
	"log"

	// _ "github.com/lib/pq" => have to use for connect db
	"github.com/Kiminoto1412/simplebank/api"
	db "github.com/Kiminoto1412/simplebank/db/sqlc"
	"github.com/Kiminoto1412/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".") // "." => app.env is the same location of this file
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connec the database")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
