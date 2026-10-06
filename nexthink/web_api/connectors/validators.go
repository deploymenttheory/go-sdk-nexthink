package connectors

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

var contentUUID = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

func ValidateID(id string) error { return validation.PathSegment(id) }
func ValidateInput(r *ConnectorInput) error {
	if r == nil {
		return fmt.Errorf("connector input is required")
	}
	if !contentUUID.MatchString(r.ContentID) {
		return fmt.Errorf("content_id must be a UUID")
	}
	for name, value := range map[string]string{"template name": r.Template.Name, "general name": r.General.Name, "nql_id": r.General.NQLID, "credential reference": r.Credentials.Reference, "credential type": r.Credentials.Type, "cron": r.Scheduling.Cron, "timezone": r.Scheduling.Timezone} {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("%s is required", name)
		}
	}
	if len(r.Streams) == 0 {
		return fmt.Errorf("at least one stream is required")
	}
	for _, stream := range r.Streams {
		if strings.TrimSpace(stream.Name) == "" {
			return fmt.Errorf("stream name is required")
		}
	}
	return nil
}
func ValidateDataModelObject(object string) error {
	if strings.TrimSpace(object) == "" {
		return fmt.Errorf("data model object is required")
	}
	return nil
}

func ValidateTest(r *TestRequest) error {
	if r == nil || strings.TrimSpace(r.TemplateID) == "" || strings.TrimSpace(r.Credentials.Reference) == "" {
		return fmt.Errorf("template ID and credential reference are required")
	}
	return nil
}
