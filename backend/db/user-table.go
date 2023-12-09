package db

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Username     sql.NullString
	PasswordHash sql.NullString
	Token        sql.NullString
}

type UserRepo struct {
	db *pgxpool.Pool
}

func (repo UserRepo) SaveUser(ctx context.Context, user User) error {
	_, err := repo.db.Exec(ctx, "INSERT INTO USERS (username, password) VALUES($1, $2)", user.Username, user.PasswordHash)
	return err
}

func (repo UserRepo) GetUserByToken(ctx context.Context, token string) (string, error) {
	var user string
	result := repo.db.QueryRow(ctx, "SELECT (username) FROM USERS WHERE logintoken=$1", token)

	if err := result.Scan(&user); err != nil {
		return "", err
	}

	return user, nil
}

func (repo UserRepo) GetUserByName(ctx context.Context, username string) (User, error) {
	var user User
	result := repo.db.QueryRow(ctx, "SELECT * FROM USERS WHERE username=$1", username)
	if err := result.Scan(&user.Username, &user.PasswordHash, &user.Token); err != nil {
		return user, err
	}
	return user, nil
}

func (repo UserRepo) UpdateUserToken(ctx context.Context, user User) error {
	user.Token.Valid = true
	_, err := repo.db.Exec(ctx, "UPDATE USERS SET logintoken=$1 WHERE username=$2", user.Token, user.Username)
	return err
}

func (repo UserRepo) DeleteUser() {

}
