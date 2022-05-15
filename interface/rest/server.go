package rest

import (
	"log"
	"net/http"

	"github.com/ymatzki/go-clean-architecture/interface/rest/handler"
	"github.com/ymatzki/go-clean-architecture/usecase"
)

func Execute() {
	// dependency injection
	// TODO: use DIContainer
	helloUsecase := usecase.NewUseCase()
	helloHandler := handler.NewHandler(helloUsecase)

	// execute
	http.HandleFunc("/", helloHandler.Hello)
	log.Printf("Start http server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
