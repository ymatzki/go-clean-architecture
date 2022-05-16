package handler

import (
	"io"
	"net/http"
)

// handle
func (h handler) Hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, h.usecase.Hello())
}
