package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/Alexsoup97/message-app/db"
	"github.com/jackc/pgx/v5"
)

func AuthMiddleware(storage *db.Storage) func(http.Handler) http.Handler {
	return authenticate(storage)
}

func authenticate(storage *db.Storage) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("APIAuth")
			if err != nil {
				switch {

				case errors.Is(err, http.ErrNoCookie):
					log.Print(err)
					http.Error(w, "Please login", http.StatusUnauthorized)
				default:
					log.Println(err)
					http.Error(w, "server error", http.StatusInternalServerError)
				}
				return
			}

			user, err := storage.UserRepo.GetUserByToken(context.Background(), cookie.Value)
			if err != nil {
				switch {
				case err == pgx.ErrNoRows:
					http.Error(w, "Token is not valid", http.StatusUnauthorized)
				default:
					log.Print(err)
					http.Error(w, "server error", http.StatusInternalServerError)
				}
				return
			}
			ctx := context.WithValue(r.Context(), "User", user)
			next.ServeHTTP(w, r.WithContext(ctx))

		})

	}
}
