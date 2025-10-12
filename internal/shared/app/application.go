package app

import (
	"frisboo-bank/customer-read-service/internal/shared/configurations/customers"
	"frisboo-bank/pkg/application/app"
	containerContracts "frisboo-bank/pkg/container/contracts"
	"frisboo-bank/pkg/container/dependencies/decorator"
	"frisboo-bank/pkg/container/dependencies/module"
	"frisboo-bank/pkg/container/dependencies/provider"
	"frisboo-bank/pkg/environment"
	loggerContracts "frisboo-bank/pkg/logger/contracts"
)

type CustomersApplication struct {
	*customers.CustomersServiceConfigurator
}

func NewCustomerApplication(
	modules []module.Module,
	providers []provider.Provider,
	decorators []decorator.Decorator,
	container containerContracts.Container,
	logger loggerContracts.Logger,
	environment environment.Environment,
) *CustomersApplication {
	nApp := app.NewApplication(
		container,
		logger,
		environment,
		modules,
		providers,
		decorators,
	)

	return &CustomersApplication{
		CustomersServiceConfigurator: customers.NewCustomersServiceConfigurator(nApp),
	}
}
