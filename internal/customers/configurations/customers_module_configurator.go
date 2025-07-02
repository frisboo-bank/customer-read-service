package configurations

import (
	appContracts "frisboo-bank/pkg/application/contracts"
)

type CustomersModuleConfigurator struct {
	appContracts.Application
}

func NewCustomersModuleConfigurator(app appContracts.Application) *CustomersModuleConfigurator {
	return &CustomersModuleConfigurator{
		Application: app,
	}
}

func (c *CustomersModuleConfigurator) ConfigureCustomersModule() {}
