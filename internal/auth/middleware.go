package auth

import "net/http"

func Middleware() func(http.Handler) http.Handler {
	return nil
}
