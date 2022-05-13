package rest

import (
	"io"
	"log"
	"net/http"

	"github.com/ymatzki/go-clean-architecture/interface/rest/handler"
)

func Execute() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, handler.Hello())
	})
	log.Printf("Start http server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
