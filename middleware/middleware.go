package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func UselessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func NewRequestLoggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqLogger := logger.With(
				zap.String("method", r.Method),
				zap.String("uri", r.RequestURI),
			)

			var start time.Time = time.Now()

			reqLogger.Info("Incoming Request Start")

			next.ServeHTTP(w, r)

			reqLogger.Info(
				"Incoming Request End",
				zap.Duration("time", time.Since(start)),
			)
		})
	}
}
