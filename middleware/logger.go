package middleware

import (
	"net/http"

	"github.com/syntaqx/buildkit"
)

func Logger(logger buildkit.Logger) buildkit.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Incoming request", "method", r.Method, "path", r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}
}
