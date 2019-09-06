package handlers

import (
	"net/http"

	"github.com/mjosc/rp/pkg/mocks/services/bad/shared"
)

func NewInternalServerError() shared.InternalServerError {
	return &InternalServerError{}
}

type InternalServerError struct {
}

func (h *InternalServerError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	w.Write([]byte("Internal Server Error"))
}
