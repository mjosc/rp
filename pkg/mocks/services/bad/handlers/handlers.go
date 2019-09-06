package handlers

import (
	"github.com/mjosc/rp/pkg/mocks/services/bad/shared"
	"go.uber.org/fx"
)

type Handlers struct {
	fx.In
	InternalServerError shared.InternalServerError
	BadRequestError     shared.BadRequestError
	ModifyResponseError shared.ModifyResponseError
}
