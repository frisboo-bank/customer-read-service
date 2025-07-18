package repositories

type CustomerRepository interface {
	// GetCustomerByCin(ctx context.Context, cin cin.Cin) (*models.Customer, error)
	// CreateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	// UpdateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
}

type CustomerMongoRepository interface{}
