package handler

import (
	"github.com/ymatzki/go-clean-architecture/usecase"
)

type Handler struct {
	usecase usecase.Usecase
}

// get hello handler
func NewHandler(u usecase.Usecase) Handler {
	return Handler{
		u,
	}
}
