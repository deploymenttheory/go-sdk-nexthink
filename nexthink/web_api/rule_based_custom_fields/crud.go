package rule_based_custom_fields

import (
	"context"
	"net/url"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

type RuleBasedCustomFieldsServiceInterface interface {
	Export(context.Context, string) (*ExportDocument, *interfaces.Response, error)
	List(context.Context) (*ListResponse, *interfaces.Response, error)
	Get(context.Context, string) (*CustomField, *interfaces.Response, error)
	Create(context.Context, *FieldInput) (*CustomField, *interfaces.Response, error)
	Update(context.Context, string, *FieldInput) (*CustomField, *interfaces.Response, error)
	Delete(context.Context, string, *DeleteRequest) (*interfaces.Response, error)
}

var _ RuleBasedCustomFieldsServiceInterface = (*Service)(nil)

// List retrieves the shared Custom Fields listing and keeps only RULE_BASED rows.
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	response, err := s.client.Get(ctx, EndpointList, nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	rows := result.Rows[:0]
	for _, row := range result.Rows {
		if row.Type == "RULE_BASED" {
			rows = append(rows, row)
		}
	}
	result.Rows = rows
	return &result, response, nil
}
func (s *Service) Get(ctx context.Context, id string) (*CustomField, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result CustomField
	response, err := s.client.Get(ctx, Endpoint+"/"+url.PathEscape(id), nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Create(ctx context.Context, request *FieldInput) (*CustomField, *interfaces.Response, error) {
	if err := ValidateInput(request, false); err != nil {
		return nil, nil, err
	}
	var result CustomField
	response, err := s.client.Post(ctx, Endpoint, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Update(ctx context.Context, id string, request *FieldInput) (*CustomField, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	if err := ValidateInput(request, true); err != nil {
		return nil, nil, err
	}
	var result CustomField
	response, err := s.client.Put(ctx, Endpoint+"/"+url.PathEscape(id), request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Delete uses the POST deletion contract observed in the UI; the response is not JSON.
func (s *Service) Delete(ctx context.Context, id string, request *DeleteRequest) (*interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, err
	}
	if err := ValidateDeleteRequest(request); err != nil {
		return nil, err
	}
	return s.client.Post(ctx, Endpoint+"/"+url.PathEscape(id), request, map[string]string{"Content-Type": "application/json"}, nil)
}

// Export retrieves the definition document without modifying field values.
func (s *Service) Export(ctx context.Context, id string) (*ExportDocument, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result ExportDocument
	resp, err := s.client.Get(ctx, Endpoint+"/export/"+url.PathEscape(id), nil, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
