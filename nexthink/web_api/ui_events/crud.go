package ui_events

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// Poll reads one page. Pass an empty nextHref initially, then the returned _links.next.href. Absolute links must use the configured tenant origin.
// HTTP 422 is returned as an API error; the browser stops polling on that status.
func (s *Service) Poll(ctx context.Context, nextHref string) (*MessagesResponse, *interfaces.Response, error) {
	path, err := messagePath(nextHref)
	if err != nil {
		return nil, nil, err
	}
	var result MessagesResponse
	resp, err := s.client.Get(ctx, path, nil, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

type UIEventsServiceInterface interface {
	Poll(context.Context, string) (*MessagesResponse, *interfaces.Response, error)
}

var _ UIEventsServiceInterface = (*Service)(nil)
