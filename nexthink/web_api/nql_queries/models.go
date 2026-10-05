package nql_queries

import (
	"encoding/json"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_administration"
)

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

// ListResponse contains the saved-query content summaries; use ContentID with Get.
type ListResponse = content_administration.ListResponse
