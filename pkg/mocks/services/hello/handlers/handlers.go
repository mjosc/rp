package handlers

import (
	"github.com/mjosc/rp/pkg/mocks/services/hello/shared"
	"go.uber.org/fx"
)

type Handlers struct {
	fx.In
	Hello shared.Hello
}
