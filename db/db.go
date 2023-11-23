package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	Db *pgxpool.Pool
}

func InitalizeDb() *Storage {
	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println("An error has occured when connecting to the db")
		os.Exit(1)
	}

	return &Storage{
		Db: dbPool,
	}
}

func (s Storage) CreateTables() {
	_, err := s.Db.Exec(context.Background(),
		`DROP TABLE IF EXISTS users;
		CREATE TABLE USERS(
		Username varchar NOT NULL,
		Password varchar NOT NULL,
		LoginToken varchar,
		PRIMARY KEY(Username)
	) `)

	if err != nil {
		log.Fatal(err)
	}
}
