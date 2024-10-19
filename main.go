package main

import (
	"context"
	"log"

	"github.com/huannguyen2114/go-toy-project/api"
	db "github.com/huannguyen2114/go-toy-project/db/sqlc"
	"github.com/huannguyen2114/go-toy-project/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	defer conn.Close()

	store := db.NewStore(conn)
	server, err := api.NewServer(store, config)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
