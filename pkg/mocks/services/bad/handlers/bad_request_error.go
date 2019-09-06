package handlers

import (
	"net/http"

	"github.com/mjosc/rp/pkg/mocks/services/bad/shared"
)

func NewBadRequestError() shared.BadRequestError {
	return &BadRequestError{}
}

type BadRequestError struct {
}

func (h *BadRequestError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(400)
	w.Write([]byte("Bad Request"))
}
