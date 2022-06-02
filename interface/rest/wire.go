//go:build wireinject
// +build wireinject

package rest

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/ymatzki/go-clean-architecture/infra/postgresql"
	"github.com/ymatzki/go-clean-architecture/interface/rest/handler"
	"github.com/ymatzki/go-clean-architecture/usecase"
)

func Initialize(db *sql.DB) handler.ServerInterface {
	wire.Build(handler.NewHandler, usecase.NewUsecase, postgresql.NewRepository)
	return handler.HelloHandler{}
}
