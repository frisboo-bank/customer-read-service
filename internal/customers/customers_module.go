package customers

import (
	"os"

	"frisboo-bank/pkg/container/dependencies/module"
	"frisboo-bank/pkg/container/dependencies/provider"
	httpserverContracts "frisboo-bank/pkg/http/http_server/contracts"

	"github.com/davecgh/go-spew/spew"
)

var Module = module.NewModule(
	"customers",

	provider.Provide(func(httpServer httpserverContracts.HTTPServer) any {
		var g any
		httpServer.RouteBuilder().RegisterGroupFunc("/customers", func(group any) {
			g = group
		})

		spew.Dump(g)
		os.Exit(1)

		return g
	}, provider.ProvideWithOptions().With(provider.Name("customers-route-builder"))),
)
