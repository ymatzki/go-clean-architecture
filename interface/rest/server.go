package rest

import (
	"log"
	"net/http"

	"github.com/ymatzki/go-clean-architecture/infra/local"
	"github.com/ymatzki/go-clean-architecture/interface/rest/handler"
	"github.com/ymatzki/go-clean-architecture/usecase"
)

func Execute() {
	// dependency injection
	// TODO: use DIContainer
	r := local.NewRepository()
	u := usecase.NewUseCase(r)
	h := handler.NewHandler(u)

	// execute
	http.HandleFunc("/", h.Hello)
	log.Printf("Start http server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
