package main

import (
	"log"
	"net/http"
)

func Logger() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Hello World!")
	}
}
