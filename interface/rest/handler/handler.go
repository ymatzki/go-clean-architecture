package handler

import (
	"net/http"

	"github.com/ymatzki/go-clean-architecture/usecase"
)

type handler interface {
	Hello(w http.ResponseWriter, req *http.Request)
}

type Handler struct {
	usecase usecase.Usecase
}

// get hello handler
func NewHandler(u usecase.Usecase) Handler {
	return Handler{
		u,
	}
}
