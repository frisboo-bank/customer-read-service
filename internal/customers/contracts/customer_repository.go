package contracts

import (
	"context"

	"frisboo-bank/customer-read-service/internal/customers/models/customers"
)

type CustomerRepository interface {
	GetAllCustomers(ctx context.Context) (map[string]*customers.Customer, error)
	GetCustomerByCin(ctx context.Context, cin string) (*customers.Customer, error)
}
