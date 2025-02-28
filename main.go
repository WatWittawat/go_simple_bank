package main

import (
	"context"
	"log"

	"github.com/WatWittawat/go_simple_bank/api"
	db "github.com/WatWittawat/go_simple_bank/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource      = "postgresql://myuser:mypassword@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStrore(connPool)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
