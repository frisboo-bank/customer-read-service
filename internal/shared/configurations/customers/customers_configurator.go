package customers

import (
	"fmt"
	"net/http"

	"frisboo-bank/customer-read-service/internal/shared/configurations/customers/infrastructure"
	"frisboo-bank/customer-read-service/internal/customers/constants"
	"frisboo-bank/pkg/application/contracts"
	"frisboo-bank/pkg/container/dependencies/invoker"

	httpserver "frisboo-bank/pkg/http/http_server"
	httpserverContracts "frisboo-bank/pkg/http/http_server/contracts"
)

type mapCustomersEndpointsParams struct {
	HTTPServer httpserverContracts.HTTPServer `name:"httpServerRef"`
}

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

func (c *CustomersServiceConfigurator) MapCustomersEndpoints() {
	c.ResolveFunc(invoker.InvokerFunc(func(params mapCustomersEndpointsParams) {
		srv := params.HTTPServer

		srv.RouteBuilder().Root().GET("/", http.HandlerFunc())

		srv.RouteBuilder().Root().GET("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s is running", constants.ServiceName)
		}))
	},
		invoker.NamedDep("httpServerRef", fmt.Sprintf(httpserver.HTTPServerProvider, "customer-read-service")),
	))
}
