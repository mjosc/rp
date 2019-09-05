package server

import (
	"context"
	"log"
	"net/http"

	"github.com/mjosc/rp/pkg/app"
	"github.com/mjosc/rp/pkg/handlers"
	"go.uber.org/fx"
)

func Register(*app.Config) fx.Option {
	return fx.Options(
		fx.Invoke(
			setup,
		),
	)
}

func setup(lc fx.Lifecycle, handlers handlers.Handlers) {

	mux := http.NewServeMux()
	mux.Handle("/hello/", handlers.HelloProxy)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					if err := server.ListenAndServe(); err != nil {
						if err == http.ErrServerClosed {
							log.Println("Server stopped")
						} else {
							log.Println(err, "Error shutting down server")
						}
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				if err := server.Shutdown(ctx); err != nil {
					log.Println(err, "Error shutting down server")
				}
				return nil
			},
		},
	)
}
