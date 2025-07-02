package infrastructure

import (
	"frisboo-bank/pkg/application"
	"frisboo-bank/pkg/application/contracts"
)

type CustomersApplicationInfrastructure struct {
	contracts.ApplicationInfrastructure
}

func NewCustomersApplicationInfrastructure(app contracts.Application) *CustomersApplicationInfrastructure {
	return &CustomersApplicationInfrastructure{application.NewApplicationInfrastructure(app)}
}

func (i *CustomersApplicationInfrastructure) configureMediator() error {
	return nil
}
