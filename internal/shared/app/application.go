package app

import (
	"frisboo-bank/customers-service/internal/shared/configurations/customers"

	"frisboo-bank/pkg/application"
	"frisboo-bank/pkg/container"
	"frisboo-bank/pkg/environment"
	loggerContracts "frisboo-bank/pkg/logger/contracts"
)

type CustomersApplication struct {
	*customers.CustomersServiceConfigurator
}

func NewCustomerApplication(
	modules []container.Module,
	providers []container.Provider,
	decorators []container.Decorator,
	logger loggerContracts.Logger,
	environment environment.Environment,
) *CustomersApplication {
	app := application.NewApplication(
		modules,
		providers,
		decorators,
		logger,
		environment,
	)

	return &CustomersApplication{
		CustomersServiceConfigurator: customers.NewCustomersServiceConfigurator(app),
	}
}
