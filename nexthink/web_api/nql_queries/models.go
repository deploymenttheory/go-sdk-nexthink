package nql_queries

import "encoding/json"

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
