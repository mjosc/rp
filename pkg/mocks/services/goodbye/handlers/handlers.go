package handlers

import (
	"github.com/mjosc/rp/pkg/mocks/services/goodbye/shared"
	"go.uber.org/fx"
)

type Handlers struct {
	fx.In
	Goodbye shared.Goodbye
}
