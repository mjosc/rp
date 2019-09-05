package handlers

import (
	"github.com/mjosc/rp/pkg/app"
	"github.com/mjosc/rp/pkg/shared"
	"go.uber.org/fx"
)

func Register(*app.Config) fx.Option {
	return fx.Options(
		fx.Provide(
			newHelloProxy,
			newGoodbyeProxy,
		),
	)
}

func newHelloProxy() (shared.HelloProxy, error) {
	return NewReverseProxy("http://localhost:8100")
}

func newGoodbyeProxy() (shared.GoodbyeProxy, error) {
	return NewReverseProxy("http://localhost:8200")
}
