// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package rest

import (
	"github.com/ymatzki/go-clean-architecture/infra/local"
	"github.com/ymatzki/go-clean-architecture/interface/rest/handler"
	"github.com/ymatzki/go-clean-architecture/usecase"
)

// Injectors from wire.go:

func Initialize() handler.Handler {
	repository := local.NewRepository()
	usecaseUsecase := usecase.NewUsecase(repository)
	handlerHandler := handler.NewHandler(usecaseUsecase)
	return handlerHandler
}