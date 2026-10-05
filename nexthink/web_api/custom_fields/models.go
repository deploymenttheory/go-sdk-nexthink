package custom_fields

import "github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_administration"

// CreateRequest supports MANUAL and COMPUTED fields. Rule-based fields use a separate REST service.
type CreateRequest struct {
	Name            string `json:"name"`
	NQLID           string `json:"nqlId"`
	DataModelObject string `json:"dataModelObject"`
	FieldDataType   string `json:"fieldDataType"`
	Description     string `json:"description"`
	Type            string `json:"type"`
	NQLQuery        string `json:"nqlQuery"`
}
type UpdateRequest struct {
	DocUID          string `json:"docUid"`
	Name            string `json:"name"`
	NQLID           string `json:"nqlId"`
	DataModelObject string `json:"dataModelObject"`
	Description     string `json:"description"`
	Type            string `json:"type"`
	NQLQuery        string `json:"nqlQuery"`
	Revision        int    `json:"revision"`
}

// DeleteRequest uses Device/User/Binary/Package, while create and update use data-model URIs.
type DeleteRequest struct {
	DocUID          string `json:"docUid"`
	Name            string `json:"name"`
	NQLID           string `json:"nqlId"`
	DataModelObject string `json:"dataModelObject"`
	Type            string `json:"type"`
	Revision        int    `json:"revision"`
}
type NQL struct {
	NQLQuery string `json:"nqlQuery"`
}
type Definition struct {
	Name            string `json:"name"`
	NQLID           string `json:"nqlId"`
	DataModelObject string `json:"dataModelObject"`
	FieldDataType   string `json:"fieldDataType"`
	Description     string `json:"description"`
	Type            string `json:"type"`
	NQL             *NQL   `json:"nql"`
}
type CustomField struct {
	Definition
	DocUID   string `json:"docUid"`
	Revision int    `json:"revision"`
}
type UpdatedField struct {
	DocUID      string `json:"docUid"`
	Name        string `json:"name"`
	NQLID       string `json:"nqlId"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

// Create does not return docUid; use List to find the new field by nqlId.
type CreateResponse struct {
	CustomField *Definition `json:"createCustomField"`
}
type GetResponse struct {
	CustomField *CustomField `json:"customField"`
}
type UpdateResponse struct {
	CustomField *UpdatedField `json:"updateCustomField"`
}
type DeleteResponse struct {
	Deleted *bool `json:"deleteCustomField"`
}
type Summary struct {
	content_administration.Content
	DataModelObject string `json:"dataModelObject"`
	NQLID           string `json:"nqlId"`
	Type            string `json:"type"`
}

// List includes rule-based entries even though GraphQL methods manage only MANUAL/COMPUTED fields.
type ListResponse struct {
	User content_administration.ContentUser `json:"user"`
	Rows []Summary                          `json:"rows"`
}
