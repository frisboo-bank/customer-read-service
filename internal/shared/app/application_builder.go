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
		b.Modules(),
		b.Providers(),
		b.Decorators(),
		b.Container(),
		b.Logger(),
		b.Environment(),
	)
}
