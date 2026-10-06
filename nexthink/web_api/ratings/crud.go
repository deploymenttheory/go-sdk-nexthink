package ratings

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	response, err := s.client.Get(ctx, EndpointList, nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Get(ctx context.Context, id string) (*Rating, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result Rating
	response, err := s.client.Get(ctx, Endpoint+"/"+url.PathEscape(id), nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Create(ctx context.Context, request *RatingInput) (*Rating, *interfaces.Response, error) {
	if err := ValidateInput(request); err != nil {
		return nil, nil, err
	}
	var result Rating
	response, err := s.client.Post(ctx, Endpoint, request, map[string]string{"Content-Type": "application/json", "Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Update sends the revision in both the query string and the JSON body, as the UI does.
func (s *Service) Update(ctx context.Context, request *UpdateRequest) (*Rating, *interfaces.Response, error) {
	if err := ValidateUpdateRequest(request); err != nil {
		return nil, nil, err
	}
	var result Rating
	response, err := s.client.Put(ctx, Endpoint+"/"+url.PathEscape(request.RatingID)+"?revision="+strconv.Itoa(request.Revision), request, map[string]string{"Content-Type": "application/json", "Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Delete sends an empty JSON object and the current revision. Its response is a JSON boolean.
func (s *Service) Delete(ctx context.Context, id string, revision int) (*DeleteResponse, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	if err := ValidateRevision(revision); err != nil {
		return nil, nil, err
	}
	var result DeleteResponse
	response, err := s.client.DeleteWithBody(ctx, Endpoint+"/"+url.PathEscape(id)+"?revision="+strconv.Itoa(revision), struct{}{}, map[string]string{"Content-Type": "application/json", "Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type RatingsServiceInterface interface {
	Export(ctx context.Context, id string) (*RatingInput, *interfaces.Response, error)
	GetRemoteActionLastValues(ctx context.Context, request *RemoteActionValuesRequest) ([]json.RawMessage, *interfaces.Response, error)
	ListRemoteActions(ctx context.Context) ([]TargetField, *interfaces.Response, error)
	ListFields(ctx context.Context) ([]TargetField, *interfaces.Response, error)
	List(context.Context) (*ListResponse, *interfaces.Response, error)
	Get(context.Context, string) (*Rating, *interfaces.Response, error)
	Create(context.Context, *RatingInput) (*Rating, *interfaces.Response, error)
	Update(context.Context, *UpdateRequest) (*Rating, *interfaces.Response, error)
	Delete(context.Context, string, int) (*DeleteResponse, *interfaces.Response, error)
}

var _ RatingsServiceInterface = (*Service)(nil)

func (s *Service) ListFields(ctx context.Context) ([]TargetField, *interfaces.Response, error) {
	var result []TargetField
	response, err := s.client.Get(ctx, EndpointMetadata+"/fields", nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

func (s *Service) ListRemoteActions(ctx context.Context) ([]TargetField, *interfaces.Response, error) {
	var result []TargetField
	response, err := s.client.Get(ctx, EndpointMetadata+"/remote-actions", nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

func (s *Service) GetRemoteActionLastValues(ctx context.Context, request *RemoteActionValuesRequest) ([]json.RawMessage, *interfaces.Response, error) {
	if request == nil || strings.TrimSpace(request.RemoteActionURI) == "" {
		return nil, nil, fmt.Errorf("remote action URI is required")
	}
	var result []json.RawMessage
	response, err := s.client.Post(ctx, EndpointMetadata+"/remote-action/last-values", request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

func (s *Service) Export(ctx context.Context, id string) (*RatingInput, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result RatingInput
	response, err := s.client.Get(ctx, Endpoint+"/export/"+url.PathEscape(id), nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
