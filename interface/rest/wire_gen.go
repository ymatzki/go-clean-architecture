// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package rest

import (
	"database/sql"
	"github.com/ymatzki/go-clean-architecture/infra/postgresql"
	"github.com/ymatzki/go-clean-architecture/interface/rest/handler"
	"github.com/ymatzki/go-clean-architecture/usecase"
)

// Injectors from wire.go:

func Initialize(db *sql.DB) handler.ServerInterface {
	repository := postgresql.NewRepository(db)
	usecaseUsecase := usecase.NewUsecase(repository)
	serverInterface := handler.NewHandler(usecaseUsecase)
	return serverInterface
}
