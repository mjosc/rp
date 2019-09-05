package handlers

import (
	"github.com/mjosc/rp/pkg/app"
	"go.uber.org/fx"
)

func Register(*app.Config) fx.Option {
	return fx.Options(
		fx.Provide(
			NewHello,
		),
	)
}
