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

func (s Storage) CreateUserTables() {
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

func (s Storage) CreateMessageTable(){
	_, err := s.Db.Exec(context.Background(),
		`DROP TABLE IF EXISTS MESSAGES;
		CREATE TABLE MESSAGE(
		MessageId varchar NOT NULL
		ConversationId varchar NOT NULL,
		UserId varchar NOT NULL,
		message varchar NOT NULL
		PRIMARY KEY(messageId)
	) `)

	if err != nil {
		log.Fatal(err)
	}
}

func (s Storage) CreateConversationSettingTable(){	
	_, err := s.Db.Exec(context.Background(),
		`DROP TABLE IF EXISTS CONVERSATIONSETTING;
		CREATE TABLE CONVERSATIONSETTING(
		ConversationId varchar NOT NULL,
		ConversationType int NOT NULL
		PRIMARY KEY(ConversationId)
	) `)

	if err != nil {
		log.Fatal(err)
	}
}

}

func (s Storage) CreateConverstationTable(){

	_, err := s.Db.Exec(context.Background(),
		`DROP TABLE IF EXISTS CONVERSATIONS;
		CREATE TABLE CONVERSATIONS(
		ConversationId varchar NOT NULL,
		UserId varchar NOT NULL,
		PermissionLevel int 
		PRIMARY KEY(ConversationId, UserId)
	) `)

	if err != nil {
		log.Fatal(err)
	}
}