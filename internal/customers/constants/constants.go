package constants

import (
	httpserver "frisboo-bank/pkg/http/http_server"
	rpcserver "frisboo-bank/pkg/rpc/rpc_server"
)

const ServiceName = "customer-read-service"

// Dependency Injection keys
const (
	MainHTTPServer           = httpserver.HTTPServerProviderPrefix + "main"
	MainRPCServer            = rpcserver.RPCServerProviderPrefix + "main"
	HTTPServerCustomersGroup = "http-server-group:customer"
)
