package nql_editor

import (
	"context"
	"encoding/json"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type NQLEditorServiceInterface interface {
	Validate(
		context.Context,
		*ValidationRequest,
		string,
	) ([]Diagnostic, *interfaces.Response, error)
	Complete(
		context.Context,
		*PositionRequest,
		string,
	) ([]CompletionItem, *interfaces.Response, error)
	Hover(context.Context, *PositionRequest) (json.RawMessage, *interfaces.Response, error)
	Resolve(context.Context, CompletionItem) (CompletionItem, *interfaces.Response, error)
	GetHighlighting(context.Context) (json.RawMessage, *interfaces.Response, error)
}

var _ NQLEditorServiceInterface = (*Service)(nil)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// Validate returns diagnostics; an invalid NQL expression is not an HTTP error.
func (s *Service) Validate(
	ctx context.Context,
	req *ValidationRequest,
	nqlMode string,
) ([]Diagnostic, *interfaces.Response, error) {
	if err := ValidateValidationRequest(req); err != nil {
		return nil, nil, err
	}
	var result []Diagnostic
	resp, err := s.client.PostWithQuery(
		ctx,
		EndpointValidate,
		modeQuery(nqlMode),
		req,
		headers(),
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (s *Service) Complete(
	ctx context.Context,
	req *PositionRequest,
	nqlMode string,
) ([]CompletionItem, *interfaces.Response, error) {
	if err := ValidatePositionRequest(req); err != nil {
		return nil, nil, err
	}
	var result []CompletionItem
	resp, err := s.client.PostWithQuery(
		ctx,
		EndpointComplete,
		modeQuery(nqlMode),
		req,
		headers(),
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (s *Service) Hover(
	ctx context.Context,
	req *PositionRequest,
) (json.RawMessage, *interfaces.Response, error) {
	if err := ValidatePositionRequest(req); err != nil {
		return nil, nil, err
	}
	var result json.RawMessage
	resp, err := s.client.Post(ctx, EndpointHover, req, headers(), &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (s *Service) Resolve(
	ctx context.Context,
	item CompletionItem,
) (CompletionItem, *interfaces.Response, error) {
	if err := ValidateCompletionItem(item); err != nil {
		return nil, nil, err
	}
	var result CompletionItem
	resp, err := s.client.Post(ctx, EndpointResolve, item, headers(), &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (s *Service) GetHighlighting(
	ctx context.Context,
) (json.RawMessage, *interfaces.Response, error) {
	var result json.RawMessage
	resp, err := s.client.Get(ctx, EndpointHighlighting, nil, headers(), &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func modeQuery(mode string) map[string]string {
	if mode == "" {
		return nil
	}
	return map[string]string{"nqlMode": mode}
}

func headers() map[string]string {
	return map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
}
