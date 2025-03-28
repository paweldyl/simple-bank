package main

import (
	"database/sql"
	"log"

	_ "github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
	"github.com/paweldyl/simplebank/api"
	db "github.com/paweldyl/simplebank/db/sqlc"
	"github.com/paweldyl/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// time.Sleep(5 * time.Second)
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
