package rest

import (
	"io"
	"log"
	"net/http"
)

func Execute() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})
	log.Printf("Start http server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
