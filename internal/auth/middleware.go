package auth

import (
	"context"
	"github.com/seagalputra/cashlog/internal/pkg/token"
	"net/http"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"userID"}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(rw, r)
				return
			}

			tokenStr := header
			userID, err := token.Parse(tokenStr)
			if err != nil {
				next.ServeHTTP(rw, r)
				return
			}
			ctx := context.WithValue(r.Context(), userCtxKey, userID)

			r = r.WithContext(ctx)
			next.ServeHTTP(rw, r)
		})
	}
}

func ForContext(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(string)
	return raw
}
