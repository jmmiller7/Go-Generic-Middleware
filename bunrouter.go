package main

import (
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatal(err)
	}
}

func realMain() error {
	log.Println("starting api...")

	router := bunrouter.New()
	router.GET("/", Get)

	host := "localhost:5050"
	log.Printf("api running on %s/n", host)
	return http.ListenAndServe(host, router)
}

func Get(w http.ResponseWriter, req bunrouter.Request) error {
	r := new(http.Response)
	r.StatusCode = http.StatusOK

	_, err := w.Write([]byte("everything is awesome"))
	return err
}
