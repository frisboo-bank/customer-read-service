package app

import (
	"frisboo-bank/pkg/application"
	"frisboo-bank/pkg/application/contracts"
)

type CustomersApplicationBuilder struct {
	contracts.ApplicationBuilder
}

func NewCustomersApplicationBuilder() *CustomersApplicationBuilder {
	return &CustomersApplicationBuilder{application.NewApplicationBuilder()}
}

func (b *CustomersApplicationBuilder) Build() *CustomersApplication {
	return NewCustomerApplication(
		b.GetModules(),
		b.GetProviders(),
		b.GetDecorators(),
		b.Logger(),
		b.Environment(),
	)
}
