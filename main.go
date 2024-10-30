package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/kareka-gb/orders-api-net-ninja/application"
)

func main() {
	app := application.NewApp()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("FAILED TO START APP:", err)
	}
}
