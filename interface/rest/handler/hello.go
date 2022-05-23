package handler

import (
	"io"
	"net/http"
)

// handle
func (h Handler) Hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, h.usecase.Hello())
}
