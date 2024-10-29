package application

// application package is mainly used for configuring the web server,
// booting the web server, linking handlers or any other resources like DB, cache to the web server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func NewApp() *App {
	app := &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	// pinging redis to check the connection
	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("FAILED TO CONNECT TO REDIS: %w", err)
	}

	// closing the redis connections during application shut down
	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("FAILED TO CLOSE REDIS", err)
		}
	}()

	fmt.Println("Starting server")

	// channel to listen to the below routine
	ch := make(chan error, 1)

	// go routine to listen and serve http requests
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("FAILED TO START SERVER: %w", err)
		}
		close(ch)
	}()

	// context routine
	ctx.Done()

	// handling both channels 
	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}

}
