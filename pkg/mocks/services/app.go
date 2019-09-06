package services

import (
	"github.com/mjosc/rp/pkg/app"
	"github.com/mjosc/rp/pkg/mocks/services/bad"
	"github.com/mjosc/rp/pkg/mocks/services/goodbye"
	"github.com/mjosc/rp/pkg/mocks/services/hello"
	"go.uber.org/fx"
)

func Register(config *app.Config) fx.Option {
	return fx.Options(
		hello.Register(config),
		goodbye.Register(config),
		bad.Register(config),
	)
}
