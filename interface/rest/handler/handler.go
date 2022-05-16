package handler

import (
	"github.com/ymatzki/go-clean-architecture/usecase"
)

type handler struct {
	usecase usecase.Usecase
}

// get hello handler
func NewHandler(u usecase.Usecase) handler {
	return handler{
		u,
	}
}
