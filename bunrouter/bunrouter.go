package main

import (
	"errors"
	"generic-middleware/middleware"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
	"go.uber.org/zap"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatal(err)
	}
}

func realMain() error {
	logger, _ := zap.NewProduction()

	logger.Info("starting api...")

	router := bunrouter.New(
		bunrouter.Use(NewErrorLoggingMiddleware(logger)),
		bunrouter.Use(WrapMiddleware(middleware.NewRequestLoggingMiddleware(logger))),
	)
	router.GET("/", get)
	router.GET("/error", getWithError)

	host := "localhost:5050"
	logger.Sugar().Infof("api running on %s\n", host)
	return http.ListenAndServe(host, router)
}

func NewErrorLoggingMiddleware(logger *zap.Logger) func(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
		return func(w http.ResponseWriter, r bunrouter.Request) error {
			/*
				some panic handling here...
			*/

			err := next(w, r)
			if err != nil {
				logger.Error("an error occurred!", zap.Error(err))
				resp := http.Response{
					StatusCode: http.StatusInternalServerError,
				}
				resp.Write(w)
			}
			return nil
		}
	}
}

func WrapMiddleware(m func(http.Handler) http.Handler) bunrouter.MiddlewareFunc {
	return func(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
		return func(w http.ResponseWriter, r bunrouter.Request) (err error) {
			m(
				http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					err = next(w, r)
				})).ServeHTTP(w, r.Request)
			return
		}
	}
}

func get(w http.ResponseWriter, req bunrouter.Request) error {
	r := new(http.Response)
	r.StatusCode = http.StatusOK

	_, err := w.Write([]byte("everything is awesome on bunrouter!"))
	return err
}

func getWithError(w http.ResponseWriter, req bunrouter.Request) error {
	return errors.New("an error occurred")
}
