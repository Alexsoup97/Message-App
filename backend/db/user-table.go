package db

import (
	"context"
	"database/sql"
)

type User struct {
	Username     sql.NullString
	PasswordHash sql.NullString
	Token        sql.NullString
}


func (s Storage) SaveUser(ctx context.Context, user User) error {
	_, err := s.Db.Exec(ctx, "INSERT INTO USERS (username, password) VALUES($1, $2)", user.Username, user.PasswordHash)
	return err
}

func (s Storage) GetUserByToken(ctx context.Context, token string) (string, error) {
	var user string
	result := s.Db.QueryRow(ctx, "SELECT (LoginToken) FROM USERS WHERE logintoken=$1", token)

	if err := result.Scan(&user); err != nil {
		return "", err
	}

	return user, nil
}

func (s Storage) GetUserByName(ctx context.Context, username string) (User, error) {
	var user User
	result := s.Db.QueryRow(ctx, "SELECT * FROM USERS WHERE username=$1", username)
	if err := result.Scan(&user.Username, &user.PasswordHash, &user.Token); err != nil {
		return user, err
	}
	return user, nil
}

func (s Storage) UpdateUserToken(ctx context.Context, user User) error {
	user.Token.Valid = true
	_, err := s.Db.Exec(ctx, "UPDATE USERS SET logintoken=$1 WHERE username=$2", user.Token, user.Username)
	return err
}

func (s Storage) DeleteUser() {

}
