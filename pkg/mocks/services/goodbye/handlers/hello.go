package handlers

import (
	"net/http"

	"github.com/mjosc/rp/pkg/mocks/services/goodbye/shared"
)

func NewGoodbye() shared.Goodbye {
	return &Goodbye{}
}

type Goodbye struct {
}

func (h *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Goodbye"))
}
