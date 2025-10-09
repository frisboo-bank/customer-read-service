package app

import (
	"frisboo-bank/pkg/application/builder"
	"frisboo-bank/pkg/application/contracts"
)

type CustomersApplicationBuilder struct {
	contracts.ApplicationBuilder
}

func NewCustomersApplicationBuilder() (*CustomersApplicationBuilder, error) {
	b, err := builder.NewApplicationBuilder()
	if err != nil {
		return nil, err
	}
	return &CustomersApplicationBuilder{b}, nil
}

func (b *CustomersApplicationBuilder) Build() contracts.Application {
	return NewCustomerApplication(
		b.Modules(),
		b.Providers(),
		b.Decorators(),
		b.Container(),
		b.Logger(),
		b.Environment())
}
