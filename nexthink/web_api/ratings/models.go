package ratings

import (
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_administration"
)

type RatingInput struct {
	TargetField               string `json:"targetField"`
	TargetFieldLabel          string `json:"targetFieldLabel"`
	DefaultValueInEnumeration int    `json:"defaultValueInEnumeration"`
	// Conditions maps enumeration values (1=poor, 2=average, 3=good) to complete NQL queries.
	Conditions map[string]string `json:"conditions"`
}
type UpdateRequest struct {
	RatingInput
	RatingID string `json:"ratingId"`
	Revision int    `json:"revision"`
}
type Rating struct {
	RatingInput
	RatingID             string `json:"ratingId"`
	Revision             int    `json:"revision"`
	DocumentLastModified string `json:"documentLastModified"`
	TenantUUID           string `json:"tenantUuid"`
	BuiltinFlag          bool   `json:"builtinFlag"`
	Title                string `json:"title"`
}
type Summary struct {
	content_administration.Content
	DocumentLastModified int64 `json:"documentLastModified"`
}
type ListResponse struct {
	User content_administration.ContentUser `json:"user"`
	Rows []Summary                          `json:"rows"`
}
type DeleteResponse bool
