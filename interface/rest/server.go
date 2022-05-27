package rest

import (
	"fmt"
	"log"
	"net/http"
)

func Execute() {
	// dependency injection
	h := Initialize()

	// specify port
	p := 8080

	// execute
	http.HandleFunc("/", h.Get)
	log.Printf("About to listen on %d. Go to http://127.0.0.1:%d/", p, p)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", p), nil))
}
