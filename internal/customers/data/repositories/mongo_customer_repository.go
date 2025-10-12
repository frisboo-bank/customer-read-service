package repositories

import (
	"context"

	"frisboo-bank/customer-read-service/internal/customers/contracts"
	"frisboo-bank/customer-read-service/internal/customers/models/customers"
	loggerContracts "frisboo-bank/pkg/logger/contracts"
)

var _ contracts.CustomerRepository = (*mongoCustomerRepository)(nil)

type mongoCustomerRepository struct {
	logger loggerContracts.Logger
}

func NewMongoCustomerRepository(logger loggerContracts.Logger) *mongoCustomerRepository {
	return &mongoCustomerRepository{
		logger: logger,
	}
}

func (m *mongoCustomerRepository) GetAllCustomers(ctx context.Context) (map[string]*customers.Customer, error) {
	panic("unimplemented")
}

func (m *mongoCustomerRepository) GetCustomerByCin(ctx context.Context, cin string) (*customers.Customer, error) {
	panic("unimplemented")
}
