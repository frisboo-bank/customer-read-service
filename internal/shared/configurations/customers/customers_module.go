package customers

import (
	"frisboo-bank/customers-service/internal/customers"

	"frisboo-bank/pkg/container/dependencies/module"
)

var Module = module.NewModule("customers",
	customers.Module,
)
