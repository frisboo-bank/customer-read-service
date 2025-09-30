package customers

import (
	"fmt"
	"net/http"

	"frisboo-bank/customers-service/internal/shared/configurations/customers/infrastructure"
	"frisboo-bank/pkg/application/contracts"
	"frisboo-bank/pkg/container/dependencies/invoker"
	httpserver "frisboo-bank/pkg/http/http_server"
	httpserverContracts "frisboo-bank/pkg/http/http_server/contracts"

	"github.com/labstack/echo/v4"
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

func (c *CustomersServiceConfigurator) ConfigureCustomers() {
	c.infrastructureConfigurator.ConfigureInfrastructures()
}

type mapCustomersEndpointsParams struct {
	HTTPServer httpserverContracts.HTTPServer `name:"httpServerRef"`
}

func (c *CustomersServiceConfigurator) MapCustomersEndpoints() {
	c.ResolveFunc(invoker.InvokerFunc(func(params mapCustomersEndpointsParams) {
		srv := params.HTTPServer

		srv.RouteBuilder().RegisterRoutes(func(server any) {
			server.(*echo.Echo).GET("/", func(c echo.Context) error {
				return c.String(http.StatusOK, fmt.Sprintf("%s is running", "customer service"))
			})
		})
	},
		invoker.NamedDep("httpServerRef", fmt.Sprintf(httpserver.HTTPServerProvider, "customers-service")),
	))
}
