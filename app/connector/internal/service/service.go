package service

import (
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	pb "kratos-arod-im/api/connector/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService)

type ConnectorService struct {
	pb.UnimplementedConnectorServer

	naming naming_client.INamingClient

	// todo
}