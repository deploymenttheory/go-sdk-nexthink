package rule_based_custom_fields

import "github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/custom_fields"

type FieldInput struct {
	Name            string              `json:"name"`
	NQLID           string              `json:"nqlId"`
	InventoryObject string              `json:"inventoryObject"`
	Type            string              `json:"type"`
	Description     string              `json:"description"`
	Revision        *int                `json:"revision,omitempty"`
	TagConditions   []TagConditionInput `json:"tagConditions"`
}
type TagConditionInput struct {
	Label      string `json:"label"`
	Statement  string `json:"statement"`
	RatingEnum int    `json:"ratingEnum"`
	TagEnum    *int   `json:"tagEnum,omitempty"`
}
type TagCondition struct {
	Label      string `json:"label"`
	Statement  string `json:"statement"`
	RatingEnum int    `json:"ratingEnum"`
	TagEnum    int    `json:"tagEnum"`
	Deleted    bool   `json:"deleted"`
}
type CustomField struct {
	DocUUID          string         `json:"docUuid"`
	Name             string         `json:"name"`
	NQLID            string         `json:"nqlId"`
	InventoryObject  string         `json:"inventoryObject"`
	Type             string         `json:"type"`
	Description      string         `json:"description"`
	Revision         int            `json:"revision"`
	TagConditions    []TagCondition `json:"tagConditions"`
	LastModifiedDate string         `json:"lastModifiedDate"`
	ContentType      string         `json:"contentType"`
	Library          bool           `json:"library"`
}

// Delete uses Device/User/Binary/Package, unlike the URI form used for create/update.
type DeleteRequest struct {
	Name            string `json:"name"`
	NQLID           string `json:"nqlId"`
	InventoryObject string `json:"inventoryObject"`
	Type            string `json:"type"`
	Revision        int    `json:"revision"`
}
type ListResponse = custom_fields.ListResponse
