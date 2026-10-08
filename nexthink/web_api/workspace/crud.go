package workspace

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"net/url"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// ListConversations calls the Workspace UI GET contract.
func (s *Service) ListConversations(ctx context.Context, automationID string) ([]Conversation, *interfaces.Response, error) {

	path := Endpoint + "/conversations"
	if automationID != "" {
		path += "?" + url.Values{"automationId": {automationID}}.Encode()
	}
	var result []Conversation
	response, err := s.client.Get(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// GetConversation calls the Workspace UI GET contract.
func (s *Service) GetConversation(ctx context.Context, id string, options *PageOptions) (*Conversation, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	q, err := pageQuery(options)
	if err != nil {
		return nil, nil, err
	}
	path := Endpoint + "/conversations/" + url.PathEscape(id)
	path += q
	var result Conversation
	response, err := s.client.Get(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetSharedConversation calls the Workspace UI GET contract.
func (s *Service) GetSharedConversation(ctx context.Context, id string, options *PageOptions) (*Conversation, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	q, err := pageQuery(options)
	if err != nil {
		return nil, nil, err
	}
	path := Endpoint + "/share/" + url.PathEscape(id)
	path += q
	var result Conversation
	response, err := s.client.Get(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// UpdateConversation calls the Workspace UI PATCH contract.
func (s *Service) UpdateConversation(ctx context.Context, id string, request *ConversationUpdate) (*Conversation, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	path := Endpoint + "/conversations/" + url.PathEscape(id)

	var result Conversation
	response, err := s.client.Patch(ctx, path, request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// DeleteConversation calls the Workspace UI DELETE contract.
func (s *Service) DeleteConversation(ctx context.Context, id string) (*interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}
	path := Endpoint + "/conversations/" + url.PathEscape(id)

	return s.client.Delete(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, nil)
}

// CancelConversation calls the Workspace UI POST contract.
func (s *Service) CancelConversation(ctx context.Context, id string) (*Document, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	path := Endpoint + "/conversations/" + url.PathEscape(id) + "/cancel"

	var result Document
	response, err := s.client.Post(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// MarkConversationRead calls the Workspace UI PATCH contract.
func (s *Service) MarkConversationRead(ctx context.Context, id string) (*Conversation, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	path := Endpoint + "/conversations/" + url.PathEscape(id) + "/read"

	var result Conversation
	response, err := s.client.Patch(ctx, path, map[string]bool{"value": true}, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// CreateConversationShare calls the Workspace UI POST contract.
func (s *Service) CreateConversationShare(ctx context.Context, id string) (*Conversation, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	path := Endpoint + "/conversations/" + url.PathEscape(id) + "/share"

	var result Conversation
	response, err := s.client.Post(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// UploadConversationFile calls the Workspace UI POST contract.
func (s *Service) UploadConversationFile(ctx context.Context, id string, request *ConversationFileRequest) (*FileResponse, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	if request == nil || request.File == "" || request.Filename == "" || request.MIMEType == "" {
		return nil, nil, fmt.Errorf("file, filename and mime_type are required")
	}
	path := Endpoint + "/conversations/" + url.PathEscape(id) + "/files"

	var result FileResponse
	response, err := s.client.Post(ctx, path, request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// DeleteConversationFile calls the Workspace UI DELETE contract.
func (s *Service) DeleteConversationFile(ctx context.Context, id, fileID string) (*interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}
	if err := validateID(fileID); err != nil {
		return nil, err
	}
	path := Endpoint + "/conversations/" + url.PathEscape(id) + "/files/" + url.PathEscape(fileID)

	return s.client.Delete(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, nil)
}

// MCPProxy calls the Workspace UI POST contract.
func (s *Service) MCPProxy(ctx context.Context, request *MCPRequest) (*Document, *interfaces.Response, error) {
	if request == nil || request.Server == "" || (request.Method != "resources/read" && request.Method != "tools/call") || !json.Valid(request.Params) {
		return nil, nil, fmt.Errorf("server, resources/read or tools/call method and JSON params are required")
	}
	path := Endpoint + "/mcp-proxy"

	var result Document
	response, err := s.client.Post(ctx, path, request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
