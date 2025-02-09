package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type Database struct {
	db *pgxpool.Pool
}

func New(dsn string) *Database {
	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Printf("connect to db: %v", err)
	}
	return &Database{db: db}
}
