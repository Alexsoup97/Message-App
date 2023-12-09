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
		`DROP TABLE IF EXISTS MESSAGES;
		CREATE TABLE MESSAGES(
		MessageId varchar NOT NULL,
		ConversationId varchar NOT NULL,
		UserId varchar NOT NULL,
		message varchar NOT NULL,
		createdAt timestamp,
		PRIMARY KEY(messageId)
	) `)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateConversationSettingTable(db *pgxpool.Pool) {
	_, err := db.Exec(context.Background(),
		`DROP TABLE IF EXISTS CONVERSATION_SETTINGS;
		CREATE TABLE CONVERSATION_SETTINGS(
		ConversationId varchar NOT NULL,
		ConversationName varchar,
		ConversationType int NOT NULL, 
		CreatedAt Timestamp,
		lastMessage Timestamp,
		PRIMARY KEY(ConversationId)
	) `)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateConversationUserTable(db *pgxpool.Pool) {

	_, err := db.Exec(context.Background(),
		`DROP TABLE IF EXISTS CONVERSATIONS;
		CREATE TABLE CONVERSATIONS (
		ConversationId varchar NOT NULL,
		UserId varchar NOT NULL,
		PermissionLevel int,
		PRIMARY KEY(ConversationId, UserId)
	) `)

	if err != nil {
		log.Fatal(err)
	}
}
