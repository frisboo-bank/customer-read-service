package app

import (
	"context"
	"fmt"
	"os"

	"frisboo-bank/customers-service/internal/shared/configurations/customers"
)

type Bootstrap struct{}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{}
}

func (b *Bootstrap) Run() {
	appBuilder := NewCustomersApplicationBuilder()
	appBuilder.ProvideModule(customers.Module)

	app := appBuilder.Build()

	err := app.ConfigureCustomers()
	if err != nil {
		fmt.Printf("bootstrap: ConfigureCustomers failed with error: %v\n", err)
		os.Exit(1)
	}

	app.MapCustomersEndpoints()

	fmt.Println("Starting customers_service application...")
	defer fmt.Println("Service customers_service stopped")

	err = app.Start(context.Background())
	if err != nil {
		fmt.Printf("bootstrap: app failed to start with error: %v\n", err)
		os.Exit(1)
	}
}
