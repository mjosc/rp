package hello

import (
	"github.com/mjosc/rp/pkg/app"
	"github.com/mjosc/rp/pkg/mocks/services/hello/handlers"
	"github.com/mjosc/rp/pkg/mocks/services/hello/server"
	"go.uber.org/fx"
)

func Register(config *app.Config) fx.Option {
	return fx.Options(
		handlers.Register(config),
		server.Register(config),
	)
}
