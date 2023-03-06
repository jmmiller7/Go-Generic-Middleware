package main

import (
	"generic-middleware/middleware"
	"net/http"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	logger.Info("starting api...")

	e := echo.New()
	e.Use(echo.WrapMiddleware(middleware.NewRequestLoggingMiddleware(logger)))
	e.GET("/", get)

	host := "localhost:5050"
	logger.Sugar().Infof("api running on %s\n", host)
	e.Logger.Fatal(e.Start(host))
}

func get(c echo.Context) error {
	return c.String(http.StatusOK, "everything is awesome on echo!")
}
