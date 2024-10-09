package main

import (
	"context"
	"log"

	"github.com/huannguyen2114/go-toy-project/api"
	db "github.com/huannguyen2114/go-toy-project/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	defer conn.Close()

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
