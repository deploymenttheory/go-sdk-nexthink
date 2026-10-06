package global_search

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/portalsession"
	"io"
	"strconv"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// Search buffers and decodes all category events. It retains partial results
// alongside malformed-stream or in-band provider errors. HTTP metadata remains
// available with transport errors.
func (s *Service) Search(ctx context.Context, r *SearchRequest) (*SearchResponse, *interfaces.Response, error) {
	if err := validateRequest(r); err != nil {
		return nil, nil, err
	}
	response, err := s.client.Post(ctx, Endpoint, r, map[string]string{"Accept": "application/x-json-stream", "Content-Type": "application/json"}, nil)
	if err != nil {
		return nil, response, err
	}
	if response == nil {
		return nil, nil, fmt.Errorf("global search: missing response")
	}
	result, err := decode(response.Body)
	return result, response, err
}
func decode(body []byte) (*SearchResponse, error) {
	result := &SearchResponse{Events: []SearchEvent{}}
	decoder := json.NewDecoder(bytes.NewReader(body))
	var problems []error
	for {
		var event *SearchEvent
		err := decoder.Decode(&event)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			problems = append(problems, fmt.Errorf("global search stream: %w", err))
			break
		}
		if event == nil {
			problems = append(problems, fmt.Errorf("global search stream: null event"))
			break
		}
		result.Events = append(result.Events, *event)
		if event.ErrorResponse != nil {
			problems = append(problems, event.ErrorResponse)
		}
	}
	return result, errors.Join(problems...)
}

// GetLegacyAuthToken obtains the token exposed by an authenticated legacy portal.
// It never discovers credentials or forwards the SDK browser bearer token.
func (s *Service) GetLegacyAuthToken(ctx context.Context, session *PortalSession) (*PortalResponse[PortalAuthToken], *interfaces.Response, error) {
	if err := validatePortalSession(session); err != nil {
		return nil, nil, err
	}
	result, response, err := portalCall[PortalAuthToken](ctx, s.client, session, map[string]string{"query": "getAuthToken"})
	if err == nil && result.Result.Token == "" {
		return result, response, fmt.Errorf("legacy portal returned an empty auth token")
	}
	return result, response, err
}

// SearchLegacyDashboards queries the legacy portal dashboard search. Limits are
// sent as provided; the UI adds one to each desired limit for its local hasMore check.
func (s *Service) SearchLegacyDashboards(ctx context.Context, session *PortalSession, r *LegacyDashboardSearchRequest) (*PortalResponse[LegacyDashboardSearchResults], *interfaces.Response, error) {
	if err := validatePortalSession(session); err != nil {
		return nil, nil, err
	}
	if err := validateLegacySearch(r); err != nil {
		return nil, nil, err
	}
	return portalCall[LegacyDashboardSearchResults](ctx, s.client, session, map[string]string{"query": "searchCustomDashboards", "search": r.Search, "maxPersonalResults": strconv.Itoa(r.MaxPersonalResults), "maxPublishedResults": strconv.Itoa(r.MaxPublishedResults), "maxRoleBasedResults": strconv.Itoa(r.MaxRoleBasedResults)})
}
func portalCall[T any](ctx context.Context, c interfaces.HTTPClient, session *PortalSession, form map[string]string) (*PortalResponse[T], *interfaces.Response, error) {
	var result PortalResponse[T]
	response, err := c.PostForm(portalsession.Context(ctx), portalsession.Endpoint, form, map[string]string{"Accept": "application/json", "x-auth-token": session.XAuthToken, "Cookie": session.Cookie}, &result)
	if err != nil {
		return nil, response, err
	}
	if result.ResultStatus == nil {
		return &result, response, fmt.Errorf("legacy portal returned no resultStatus")
	}
	if result.ResultStatus.Code != 0 {
		return &result, response, result.ResultStatus
	}
	if result.Result == nil {
		return &result, response, fmt.Errorf("legacy portal returned no result")
	}
	return &result, response, nil
}
