package main

import (
	"context"
	"fmt"

	"github.com/S6-ICT-Sw/Backend/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("failed to start the app:", err)
	}
}
