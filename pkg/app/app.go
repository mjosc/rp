package app

import (
	"go.uber.org/fx"
)

// Config is used to pass any environment variables and/or other values to each package via an
// FXRegistrationFunc.
type Config struct {
}

// FXRegistrationFunc is the function signature of any function used to register dependencies or other
// FX options with the app.
type FXRegistrationFunc func(*Config) fx.Option

// New returns an instance of App, having passed the config to each package before running the
// underlying fx.App instance.
func New(config *Config, registrations ...FXRegistrationFunc) *App {
	options := make([]fx.Option, 0, len(registrations))
	for _, f := range registrations {
		options = append(options, f(config))
	}
	return &App{
		inner: fx.New(options...),
	}
}

// App wraps fx.App, providing the same core functionality.
type App struct {
	inner *fx.App
}

// Run is the simplest form of running an application via FX. We could use Start and Stop for additional
// customization but it's not needed in the scope of this application.
func (app *App) Run() {
	app.inner.Run()
}
