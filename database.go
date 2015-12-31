package main

import (
	"fmt"
	"log"

	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
)

type DB struct {
	conn *dbr.Connection
}

func NewDB(config *AppConfig) (*DB, error) {
	userName := config.Database.UserName
	dbName := config.Database.DatabaseName
	port := config.Database.Port
	sslMode := config.Database.SslMode
	server := config.Database.Server
	// "postgres://postgres@localhost:5432/uservoice_test?sslmode=disable"
	postgresDSN := fmt.Sprintf(
		"postgres://%s@%s:%s/%s?sslmode=%s", dbName, server, port, userName, sslMode)
	conn, err := dbr.Open("postgres", postgresDSN, nil)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
		return nil, err
	}
	db := DB{
		conn: conn,
	}
	return &db, nil
}
