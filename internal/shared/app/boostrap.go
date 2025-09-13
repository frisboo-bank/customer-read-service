package app

import (
	"context"
	"fmt"

	"frisboo-bank/customers-service/internal/shared/configurations/customers"
	"frisboo-bank/pkg/syserrors"
)

type Bootstrap struct{}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{}
}

func (b *Bootstrap) Run() error {
	appBuilder, err := NewCustomersApplicationBuilder()
	if err != nil {
		return syserrors.Wrap(err, "failed to initialize the application builder")
	}

	appBuilder.ProvideModule(customers.Module)

	app := appBuilder.Build()

	if err := app.ConfigureCustomers(); err != nil {
		return syserrors.Wrap(err, "failed to configure customers layer")
	}

	app.MapCustomersEndpoints()

	fmt.Println("Starting customers_service application...")
	defer fmt.Println("Service customers_service stopped")

	if err := app.Start(context.Background()); err != nil {
		return syserrors.Wrap(err, "failed to start app")
	}

	return nil
}
