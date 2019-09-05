package main

import (
	"github.com/mjosc/rp/pkg/app"
	"github.com/mjosc/rp/pkg/handlers"
	"github.com/mjosc/rp/pkg/mocks/services"
	"github.com/mjosc/rp/pkg/server"
)

func main() {
	app := app.New(nil,
		server.Register,
		handlers.Register,
		services.Register,
	)
	app.Run()
}
