package experimental

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type NQLQueriesService struct{ client *Client }

// SavedQuery is the read representation; the write API uses nql, not nqlQuery.
type SavedQuery struct {
	NQLAPIID           string                     `json:"nqlApiId"`
	Name               string                     `json:"name"`
	Description        string                     `json:"description"`
	NQLQuery           string                     `json:"nqlQuery"`
	ContentID          string                     `json:"contentId"`
	ParameterDataTypes map[string]json.RawMessage `json:"parameterDataTypes"`
}

type SaveQueryRequest struct {
	NQLAPIID    string `json:"nqlApiId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	NQL         string `json:"nql"`
	ContentID   string `json:"contentId"`
}

func (s *NQLQueriesService) Get(ctx context.Context, contentID string) (*SavedQuery, *interfaces.Response, error) {
	return decode[SavedQuery](s.client.Do(ctx, "nql_queries.get", Request{PathParams: map[string]string{"contentId": contentID}}))
}

func (s *NQLQueriesService) Create(ctx context.Context, req *SaveQueryRequest) (*SavedQuery, *interfaces.Response, error) {
	if req == nil || req.NQLAPIID == "" || req.Name == "" || req.NQL == "" {
		return nil, nil, fmt.Errorf("query ID, name and NQL are required")
	}
	return decode[SavedQuery](s.client.Do(ctx, "nql_queries.create", Request{Body: req}))
}

func (s *NQLQueriesService) Update(ctx context.Context, req *SaveQueryRequest) (*SavedQuery, *interfaces.Response, error) {
	if req == nil || req.ContentID == "" || req.NQLAPIID == "" || req.Name == "" || req.NQL == "" {
		return nil, nil, fmt.Errorf("content ID, query ID, name and NQL are required")
	}
	return decode[SavedQuery](s.client.Do(ctx, "nql_queries.update", Request{PathParams: map[string]string{"contentId": req.ContentID}, Body: req}))
}

func (s *NQLQueriesService) Delete(ctx context.Context, contentID string) (*interfaces.Response, error) {
	_, resp, err := s.client.Do(ctx, "nql_queries.delete", Request{PathParams: map[string]string{"contentId": contentID}})
	return resp, err
}
