package handlers

import (
	"github.com/mjosc/rp/pkg/shared"
	"go.uber.org/fx"
)

type Handlers struct {
	fx.In
	HelloProxy        shared.HelloProxy
	GoodbyeProxy      shared.GoodbyeProxy
	BadProxy          shared.BadProxy
	NoConnectionProxy shared.NoConnectionProxy
}
