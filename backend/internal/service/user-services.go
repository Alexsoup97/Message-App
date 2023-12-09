package service

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/Alexsoup97/message-app/db"
	"golang.org/x/crypto/bcrypt"
)

var IncorrectPassword = errors.New("incorrect password")
var UserNotFound = errors.New("User not found")

type UserService struct {
	database *db.Storage
}

func CreateUserService(database *db.Storage) *UserService {
	return &UserService{
		database: database,
	}
}

func (user_serv UserService) CreateAccount(username string, password string) error {

	ctx := context.Background()

	hashedPass, err := hashPassword(password)
	if err != nil {
		return fmt.Errorf("Error hashing password")
	}

	user := &db.User{
		Username:     sql.NullString{String: username, Valid: true},
		PasswordHash: sql.NullString{String: hashedPass, Valid: true},
	}
	err = user_serv.database.UserRepo.SaveUser(ctx, *user)
	return err
}

func (user_serv UserService) Login(username string, password string) (string, error) {
	ctx := context.Background()
	user, err := user_serv.database.UserRepo.GetUserByName(ctx, username)

	if err != nil {
		return "", UserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.String), []byte(password))
	if err != nil {
		return "", IncorrectPassword
	}
	if err != nil {
		return "", err
	}

	if !user.Token.Valid {
		user.Token.String = generateRandomBytes(64)
		err = user_serv.database.UserRepo.UpdateUserToken(ctx, user)
		if err != nil {
			return "", err
		}
	}
	return user.Token.String, nil
}

func generateRandomBytes(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)

	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(b)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
