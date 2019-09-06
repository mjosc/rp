package handlers

import (
	"fmt"

	"github.com/mjosc/rp/pkg/app"
	"github.com/mjosc/rp/pkg/shared"
	"go.uber.org/fx"
)

var configuration *app.Config

func Register(config *app.Config) fx.Option {
	configuration = config
	return fx.Options(
		fx.Provide(
			newHelloProxy,
			newGoodbyeProxy,
			newBadProxy,
			newNoConnectionProxy,
		),
	)
}

func newHelloProxy() (shared.HelloProxy, error) {
	dst := fmt.Sprintf("http://localhost:%v", configuration.HelloServicePort)
	return NewReverseProxy(dst)
}

func newGoodbyeProxy() (shared.GoodbyeProxy, error) {
	dst := fmt.Sprintf("http://localhost:%v", configuration.GoodbyeServicePort)
	return NewReverseProxy(dst)
}

func newBadProxy() (shared.BadProxy, error) {
	dst := fmt.Sprintf("http://localhost:%v", configuration.BadServicePort)
	return NewReverseProxy(dst)
}

func newNoConnectionProxy() (shared.NoConnectionProxy, error) {
	dst := fmt.Sprintf("http://localhost:%v", configuration.NoConnectionServicePort)
	return NewReverseProxy(dst)
}
