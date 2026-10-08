package product_configuration

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type ProductConfigurationServiceInterface interface {
	GetInstance(ctx context.Context, key string) (*InstanceResponse, *interfaces.Response, error)
	CreateInstance(ctx context.Context, request InstanceConfiguration) (*interfaces.Response, error)
	UpdateInstance(ctx context.Context, key string, request InstanceConfiguration) (*interfaces.Response, error)
}

var _ ProductConfigurationServiceInterface = (*Service)(nil)
