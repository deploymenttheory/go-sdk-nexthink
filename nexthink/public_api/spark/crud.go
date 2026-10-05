// Package spark hands off conversations to Nexthink Spark in Microsoft Teams.
// API reference: https://docs.nexthink.com/api/spark/handoff-api
package spark

import (
	"context"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type SparkServiceInterface interface {
	Handoff(
		ctx context.Context,
		upn, timezone string,
		req *HandoffRequest,
	) (*interfaces.Response, error)
}

var _ SparkServiceInterface = (*Service)(nil)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// Handoff sends a conversation to the specified user's Teams account.
// Successful requests return HTTP 204 with no JSON response body.
func (s *Service) Handoff(
	ctx context.Context,
	upn, timezone string,
	req *HandoffRequest,
) (*interfaces.Response, error) {
	if err := ValidateHandoffHeaders(upn, timezone); err != nil {
		return nil, err
	}
	if err := ValidateHandoffRequest(req); err != nil {
		return nil, err
	}
	headers := map[string]string{
		"Content-Type":        "application/json",
		"Accept":              "application/json",
		"User-Principal-Name": upn,
	}
	if timezone != "" {
		headers["Timezone"] = timezone
	}
	return s.client.Post(ctx, EndpointHandoff, req, headers, nil)
}
