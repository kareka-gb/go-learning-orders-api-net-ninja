package main

import (
	"context"
	"fmt"

	"githbub.com/kareka-gb/orders-api-net-ninja/application"
)

func main() {
	app := application.NewApp()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("FAILED TO START APP:", err)
	}
}
