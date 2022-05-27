//go:generate oapi-codegen -config config.yaml ../../../schema/openapi/api.yaml

package handler

import (
	"io"
	"net/http"

	"github.com/ymatzki/go-clean-architecture/usecase"
)

type HelloHandler struct {
	usecase usecase.Usecase
}

func NewHandler(u usecase.Usecase) ServerInterface {
	return HelloHandler{
		u,
	}
}

func (h HelloHandler) Get(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, h.usecase.Hello())
}
