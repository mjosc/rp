package services

import (
	"github.com/mjosc/rp/pkg/app"
	"github.com/mjosc/rp/pkg/mocks/services/hello"
	"go.uber.org/fx"
)

func Register(config *app.Config) fx.Option {
	return fx.Options(
		hello.Register(config),
	)
}
