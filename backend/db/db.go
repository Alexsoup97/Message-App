package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	MessageRepo MessageRepo
	UserRepo    UserRepo
}

func InitalizeStorage() (*Storage, *pgxpool.Pool) {
	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println("An error has occured when connecting to the db")
		os.Exit(1)
	}

	return &Storage{
		MessageRepo: MessageRepo{db: dbPool},
		UserRepo:    UserRepo{db: dbPool},
	}, dbPool
}

func CreateUserTables(db *pgxpool.Pool) {
	_, err := db.Exec(context.Background(),
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

func CreateMessageTable(db *pgxpool.Pool) {
	_, err := db.Exec(context.Background(),
		`DROP TABLE IF EXISTS messages;
		CREATE TABLE messages(
		message_id varchar NOT NULL,
		conversation_id varchar NOT NULL,
		user_id varchar NOT NULL,
		message varchar NOT NULL,
		created_at timestamp,
		PRIMARY KEY(message_id)
	) `)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateConversationSettingTable(db *pgxpool.Pool) {
	_, err := db.Exec(context.Background(),
		`DROP TABLE IF EXISTS conversation_settings;
		CREATE TABLE conversation_settings(
		conversation_id varchar NOT NULL,
		conversation_name varchar,
		conversation_type int NOT NULL, 
		created_at Timestamp,
    last_message varchar,
		last_message_time Timestamp,
		PRIMARY KEY(conversation_id)
	) `)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateConversationUserTable(db *pgxpool.Pool) {

	_, err := db.Exec(context.Background(),
		`DROP TABLE IF EXISTS conversations;
		CREATE TABLE conversations (
	  conversation_id varchar NOT NULL,
		user_id varchar NOT NULL,
		permission_level int,
		PRIMARY KEY(conversation_id, user_id)
	) `)

	if err != nil {
		log.Fatal(err)
	}
}
