package infrastructure

import (
	"frisboo-bank/pkg/application/builder"
	"frisboo-bank/pkg/application/contracts"
)

type CustomersApplicationInfrastructure struct {
	contracts.ApplicationInfrastructure
}

func NewCustomersApplicationInfrastructure(app contracts.Application) *CustomersApplicationInfrastructure {
	return &CustomersApplicationInfrastructure{builder.NewApplicationInfrastructure(app)}
}

func (i *CustomersApplicationInfrastructure) configureMediator() {
}
