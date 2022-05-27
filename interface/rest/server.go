package rest

import (
	"log"
	"net/http"
)

func Execute() {
	// dependency injection
	h := Initialize()

	// execute
	http.HandleFunc("/", h.Get)
	log.Printf("Start http server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
