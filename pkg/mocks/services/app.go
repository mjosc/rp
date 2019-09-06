package services

import (
	"github.com/mjosc/rp/pkg/app"
	"github.com/mjosc/rp/pkg/mocks/services/bad"
	"github.com/mjosc/rp/pkg/mocks/services/goodbye"
	"github.com/mjosc/rp/pkg/mocks/services/hello"
	"go.uber.org/fx"
)

// Register registers any FX options required in this package. It must be provided to app.New in order
// to be used in the application. To avoid polluting app.New with each register function pertaining to
// the various mock services, register those services here prior to calling this function in app.New.
func Register(config *app.Config) fx.Option {
	return fx.Options(
		hello.Register(config),
		goodbye.Register(config),
		bad.Register(config),
	)
}
