package middleware

import (
	"context"
	"net/http"
)

func paginate() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "paginate", "temp")

			next.ServeHTTP(w, r.WithContext(ctx))

		})

	}
}
