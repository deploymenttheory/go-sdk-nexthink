package workspace_tasks

import (
	"context"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"net/url"
	"strconv"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// ListTasks calls the Workspace UI GET contract.
func (s *Service) ListTasks(ctx context.Context, options *TaskListOptions) (*TaskListResponse, *interfaces.Response, error) {
	if options != nil && (options.Page < 0 || options.PageSize < 0) {
		return nil, nil, fmt.Errorf("page and pageSize cannot be negative")
	}
	path := TasksEndpoint
	if options != nil {
		q := url.Values{"page": {strconv.Itoa(options.Page)}}
		if options.PageSize > 0 {
			q.Set("pageSize", strconv.Itoa(options.PageSize))
		}
		path += "?" + q.Encode()
	}
	var result TaskListResponse
	response, err := s.client.Get(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetTask calls the Workspace UI GET contract.
func (s *Service) GetTask(ctx context.Context, id string) (*Task, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	path := TasksEndpoint + "/" + url.PathEscape(id)

	var result Task
	response, err := s.client.Get(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// CreateTask calls the Workspace UI POST contract.
func (s *Service) CreateTask(ctx context.Context, request *TaskRequest) (*Task, *interfaces.Response, error) {
	if err := validateTask(request); err != nil {
		return nil, nil, err
	}
	path := TasksEndpoint

	var result Task
	response, err := s.client.Post(ctx, path, request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// UpdateTask calls the Workspace UI PUT contract.
func (s *Service) UpdateTask(ctx context.Context, id string, request *TaskRequest) (*Task, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	if err := validateTask(request); err != nil {
		return nil, nil, err
	}
	path := TasksEndpoint + "/" + url.PathEscape(id)

	var result Task
	response, err := s.client.Put(ctx, path, request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// DeleteTask calls the Workspace UI DELETE contract.
func (s *Service) DeleteTask(ctx context.Context, id string) (*interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}
	path := TasksEndpoint + "/" + url.PathEscape(id)

	return s.client.Delete(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, nil)
}

// ReconcileTaskAgentAccess calls the Workspace UI POST contract.
func (s *Service) ReconcileTaskAgentAccess(ctx context.Context) (*interfaces.Response, error) {

	path := TasksEndpoint + "/reconcile-it-agent-access"

	return s.client.Post(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, nil)
}
