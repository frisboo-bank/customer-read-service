package customers

import (
	"frisboo-bank/customers-service/internal/shared/configurations/customers/infrastructure"
	"frisboo-bank/pkg/application/contracts"
	"frisboo-bank/pkg/container/dependencies/invoker"

	httpServerContacts "frisboo-bank/pkg/http/http_server/contracts"
)

type CustomersServiceConfigurator struct {
	contracts.Application
	infrastructureConfigurator *infrastructure.CustomersApplicationInfrastructure
}

func NewCustomersServiceConfigurator(app contracts.Application) *CustomersServiceConfigurator {
	infraConfigurator := infrastructure.NewCustomersApplicationInfrastructure(app)

	return &CustomersServiceConfigurator{
		Application:                app,
		infrastructureConfigurator: infraConfigurator,
	}
}

func (c *CustomersServiceConfigurator) ConfigureCustomers() error {
	if err := c.infrastructureConfigurator.ConfigureInfrastructure(); err != nil {
		return err
	}

	return nil
}

func (c *CustomersServiceConfigurator) MapCustomersEndpoints() {
	c.ResolveFunc(invoker.InvokerFunc(func(httpServers []httpServerContacts.HTTPServer) {
		// httpServer.RouteBuilder().RegisterRoutes(func(server any) {
		// 	server.(*gin.Engine).GET("", func(ctx *gin.Context) {
		// 		ctx.String(http.StatusOK, "%s is running", "customer service")
		// 	})
		// })
	}))
}
