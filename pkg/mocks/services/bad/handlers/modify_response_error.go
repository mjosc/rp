package handlers

import (
	"net/http"

	"github.com/mjosc/rp/pkg/mocks/services/bad/shared"
	common "github.com/mjosc/rp/pkg/shared"
)

func NewModifyResponseError() shared.ModifyResponseError {
	return &ModifyResponseError{}
}

type ModifyResponseError struct {
}

func (h *ModifyResponseError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(common.HeaderModifyResponseError, "error")
	w.WriteHeader(200)
}
