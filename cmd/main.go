package main

import (
	"github.com/mjosc/rp/pkg/app"
	"github.com/mjosc/rp/pkg/handlers"
	mocks "github.com/mjosc/rp/pkg/mocks/services"
	"github.com/mjosc/rp/pkg/server"
)

func main() {
	config := &app.Config{
		ProxyServicePort:        8000,
		HelloServicePort:        8100,
		GoodbyeServicePort:      8200,
		BadServicePort:          8300,
		NoConnectionServicePort: 8400,
	}

	app := app.New(
		config,
		server.Register,
		handlers.Register,
		mocks.Register,
	)

	app.Run()
}
