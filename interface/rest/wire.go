//go:build wireinject
// +build wireinject

package rest

import (
	"github.com/google/wire"
	"github.com/ymatzki/go-clean-architecture/infra/local"
	"github.com/ymatzki/go-clean-architecture/interface/rest/handler"
	"github.com/ymatzki/go-clean-architecture/usecase"
)

func Initialize() handler.Handler {
	wire.Build(handler.NewHandler, usecase.NewUsecase, local.NewRepository)
	return handler.Handler{}
}
