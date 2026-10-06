package custom_fields

import "encoding/json"
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

type ValidationPatterns struct {
	RBCFLabelRegex                string `json:"RBCF_LABEL_REGEX"`
	NQLIDRegex                    string `json:"NQL_ID_REGEX"`
	NQLKeywordsRegex              string `json:"NQL_KEYWORDS_REGEX"`
	ForbiddenRBCFLabelPrefixRegex string `json:"FORBIDDEN_RBCF_LABEL_PREFIX_REGEX"`
	BuiltinContentNQLIDRegex      string `json:"BUILTIN_CONTENT_NQL_ID_REGEX"`
}
type ExportDocument struct {
	Name            string `json:"name"`
	NQLID           string `json:"nqlId"`
	Description     string `json:"description"`
	DataModelObject string `json:"dataModelObject"`
	FieldDataType   string `json:"fieldDataType"`
	Type            string `json:"type"`
	NQL             *NQL   `json:"nql,omitempty"`
}

// ContentFile contains serialized definition JSON, not base64; the UI sends metadata:null.
type ImportRequest struct {
	Metadata    json.RawMessage `json:"metadata"`
	ContentFile string          `json:"contentFile"`
}
type ImportedField struct {
	CustomField
	Category         *string         `json:"category"`
	ContentType      string          `json:"contentType"`
	LastModifiedDate string          `json:"lastModifiedDate"`
	DefaultTagEnum   int             `json:"defaultTagEnum"`
	TagConditions    json.RawMessage `json:"tagConditions"`
	Library          bool            `json:"library"`
}
