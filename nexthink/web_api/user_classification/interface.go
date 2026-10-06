package user_classification

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type UserClassificationServiceInterface interface {
	List(context.Context) (*ListResponse, *interfaces.Response, error)
	Replace(ctx context.Context, request *ReplaceRequest) (*interfaces.Response, error)
}

var _ UserClassificationServiceInterface = (*Service)(nil)
