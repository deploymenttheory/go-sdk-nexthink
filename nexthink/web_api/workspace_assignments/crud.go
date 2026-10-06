package workspace_assignments

import (
	"context"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"net/url"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// ListAssignments calls the Workspace UI GET contract.
func (s *Service) ListAssignments(ctx context.Context, options *AssignmentOptions) ([]Assignment, *interfaces.Response, error) {

	path := AssignmentsEndpoint
	path += assignmentQuery(options)
	var result []Assignment
	response, err := s.client.Get(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// GetAssignment calls the Workspace UI GET contract.
func (s *Service) GetAssignment(ctx context.Context, id string, options *AssignmentOptions) (*AssignmentResponse, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	path := AssignmentsEndpoint + "/" + url.PathEscape(id)
	path += assignmentQuery(options)
	var result AssignmentResponse
	response, err := s.client.Get(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	if result.ErrorCode == "assignment_not_found" || result.Item == nil {
		return nil, response, fmt.Errorf("workspace assignment %q not found", id)
	}
	return &result, response, nil
}

// UpdateAssignment calls the Workspace UI PATCH contract.
func (s *Service) UpdateAssignment(ctx context.Context, id string, request *AssignmentUpdate) (*AssignmentResponse, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	if err := validateAssignment(request); err != nil {
		return nil, nil, err
	}
	path := AssignmentsEndpoint + "/" + url.PathEscape(id)

	var result AssignmentResponse
	response, err := s.client.Patch(ctx, path, request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	if result.ErrorCode == "assignment_not_found" || result.Item == nil {
		return nil, response, fmt.Errorf("workspace assignment %q not found", id)
	}
	return &result, response, nil
}

// MarkAssignmentRead calls the Workspace UI POST contract.
func (s *Service) MarkAssignmentRead(ctx context.Context, id string) (*interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}
	path := AssignmentsEndpoint + "/" + url.PathEscape(id) + "/read"

	return s.client.Post(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, nil)
}

// GetUnreadAssignmentCount calls the Workspace UI GET contract.
func (s *Service) GetUnreadAssignmentCount(ctx context.Context, options *AssignmentOptions) (*UnreadCountResponse, *interfaces.Response, error) {

	path := AssignmentsEndpoint + "/summary/unread-count"
	path += assignmentQuery(options)
	var result UnreadCountResponse
	response, err := s.client.Get(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// ListAssignees calls the Workspace UI GET contract.
func (s *Service) ListAssignees(ctx context.Context, id string) ([]Document, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	path := AssigneesEndpoint + "?" + url.Values{"assignmentId": {id}}.Encode()

	var result []Document
	response, err := s.client.Get(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}
