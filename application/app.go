package application

// application package is mainly used for configuring the web server,
// booting the web server, linking handlers or any other resources like DB, cache to the web server

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

func NewApp() *App {
	app := &App{
		router: loadRoutes(),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("FAILED TO START SERVER: %w", err)
	}

	return nil
}
