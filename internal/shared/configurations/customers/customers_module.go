package customers

import (
	"frisboo-bank/customer-read-service/internal/shared/configurations/customers/infrastructure"
	"frisboo-bank/customer-read-service/internal/customers"
	applicationContracts "frisboo-bank/pkg/application/contracts"
	"frisboo-bank/pkg/container/dependencies/module"
	"frisboo-bank/pkg/validation"
)

func ModuleFunc(appBuilder applicationContracts.ApplicationBuilder) module.Module {
	validation.AssertNotNil("appBuilder", appBuilder)

	m := module.ModuleFunc(
		"customers",

		infrastructure.ModuleFunc(appBuilder),

		customers.ModuleFunc(appBuilder),
	)

	return m
}
