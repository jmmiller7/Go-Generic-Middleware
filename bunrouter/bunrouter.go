package main

import (
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

	router := bunrouter.New(bunrouter.Use(NewLoggerMiddleware(logger)))
	router.GET("/", Get)

	host := "localhost:5050"
	logger.Sugar().Infof("api running on %s\n", host)
	return http.ListenAndServe(host, router)
}

func NewLoggerMiddleware(logger *zap.Logger) bunrouter.MiddlewareFunc {
	return func(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
		// first, we need to convert "next" to a http.HandlerFunc to plug it into the middleware
		f := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// RISK: we will no longer see a returned error here
			next.ServeHTTP(w, r)
		})

		// then, we can init the new middleware piece
		h := middleware.NewRequestLoggingMiddleware(logger)

		// finally we can plug the middleware into bunrouter
		return bunrouter.HTTPHandler(h(f))
	}
}

func Get(w http.ResponseWriter, req bunrouter.Request) error {
	r := new(http.Response)
	r.StatusCode = http.StatusOK

	_, err := w.Write([]byte("everything is awesome on bunrouter!"))
	return err
}
