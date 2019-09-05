package handlers

import (
	"net/http"

	"github.com/mjosc/rp/pkg/mocks/services/hello/shared"
)

func NewHello() shared.Hello {
	return &Hello{}
}

type Hello struct {
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello"))
}
