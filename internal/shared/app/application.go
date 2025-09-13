package app

import (
	"frisboo-bank/customers-service/internal/shared/configurations/customers"

	"frisboo-bank/pkg/application"
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
	app := application.NewApplication(
		modules,
		providers,
		decorators,
		container,
		logger,
		environment,
	)

	return &CustomersApplication{
		CustomersServiceConfigurator: customers.NewCustomersServiceConfigurator(app),
	}
}
