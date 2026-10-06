package checklists

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

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
func (s *Service) Get(ctx context.Context, id string) (*Checklist, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result Checklist
	response, err := s.client.Get(ctx, Endpoint+"/"+url.PathEscape(id), nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Create(ctx context.Context, request *ChecklistInput, options *CreateOptions) (*Checklist, *interfaces.Response, error) {
	if err := ValidateInput(request); err != nil {
		return nil, nil, err
	}
	params := map[string]string{}
	if options != nil && options.LibraryUUID != "" {
		if err := ValidateID(options.LibraryUUID); err != nil {
			return nil, nil, err
		}
		params["libraryUUID"] = options.LibraryUUID
	}
	var result Checklist
	response, err := s.client.PostWithQuery(ctx, Endpoint, params, request, map[string]string{"Content-Type": "application/json", "Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Update(ctx context.Context, id string, revision int, request *ChecklistInput) (*Checklist, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	if err := ValidateRevision(revision); err != nil {
		return nil, nil, err
	}
	if err := ValidateInput(request); err != nil {
		return nil, nil, err
	}
	var result Checklist
	response, err := s.client.Put(ctx, Endpoint+"/"+url.PathEscape(id)+"?revisionNumber="+strconv.Itoa(revision), request, map[string]string{"Content-Type": "application/json", "Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Delete preserves the UI's nonempty DELETE body. The successful response is not JSON.
func (s *Service) Delete(ctx context.Context, id string, revision int) (*interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, err
	}
	if err := ValidateRevision(revision); err != nil {
		return nil, err
	}
	return s.client.DeleteWithBody(ctx, Endpoint+"/"+url.PathEscape(id)+"?revisionNumber="+strconv.Itoa(revision), map[string]string{"a": "fix"}, map[string]string{"Content-Type": "application/json"}, nil)
}

type ChecklistsServiceInterface interface {
	GetFromLibrary(context.Context, string) (*LibraryDocument, *interfaces.Response, error)
	Export(context.Context, string) (*ExportDocument, *interfaces.Response, error)
	Import(context.Context, *ExportDocument) (*interfaces.Response, error)
	ListGroupedFields(context.Context) ([]FieldGroup, *interfaces.Response, error)
	List(context.Context) (*ListResponse, *interfaces.Response, error)
	Get(context.Context, string) (*Checklist, *interfaces.Response, error)
	Create(context.Context, *ChecklistInput, *CreateOptions) (*Checklist, *interfaces.Response, error)
	Update(context.Context, string, int, *ChecklistInput) (*Checklist, *interfaces.Response, error)
	Delete(context.Context, string, int) (*interfaces.Response, error)
}

var _ ChecklistsServiceInterface = (*Service)(nil)

func (s *Service) Export(ctx context.Context, id string) (*ExportDocument, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result ExportDocument
	response, err := s.client.Get(ctx, EndpointAdmin+"/export/checklists/"+url.PathEscape(id), nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Import acknowledges the upload without returning an identifier. Use List to locate the imported checklist.
func (s *Service) Import(ctx context.Context, request *ExportDocument) (*interfaces.Response, error) {
	if request == nil || request.Label == "" || request.Type == "" || request.Version < 1 {
		return nil, fmt.Errorf("a versioned checklist export document is required")
	}
	return s.client.Post(ctx, EndpointAdmin+"/import/checklists", request, map[string]string{"Content-Type": "application/json"}, nil)
}
func (s *Service) ListGroupedFields(ctx context.Context) ([]FieldGroup, *interfaces.Response, error) {
	var result []FieldGroup
	response, err := s.client.Get(ctx, EndpointGroupedFields, nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// GetFromLibrary reads a built-in definition without installing it.
func (s *Service) GetFromLibrary(ctx context.Context, id string) (*LibraryDocument, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result LibraryDocument
	response, err := s.client.Get(ctx, EndpointLibrary+"/"+url.PathEscape(id), nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
