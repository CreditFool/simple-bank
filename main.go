package main

import (
	"database/sql"
	"log"
	"os"
	"simple-bank/api"
	db "simple-bank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	dbSource, exists := os.LookupEnv("DATABASE")
	if !exists {
		log.Fatal("db address not configured at environment")
	}

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start the server", err)
	}
}
