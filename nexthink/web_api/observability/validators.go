package observability

import (
	"fmt"
	"mime"
	"strings"
)

func validateSubmission(r *Submission) error {
	if r == nil {
		return fmt.Errorf("submission is required")
	}
	for _, v := range []string{r.Source, r.ClientToken, r.Origin, r.OriginVersion, r.RequestID} {
		if strings.TrimSpace(v) == "" {
			return fmt.Errorf("source, client token, origin/version and request ID are required")
		}
	}
	if len(r.Payload) == 0 {
		return fmt.Errorf("event payload is required")
	}
	if _, _, err := mime.ParseMediaType(r.ContentType); err != nil {
		return fmt.Errorf("invalid content type: %w", err)
	}
	if r.RetryCount != nil && *r.RetryCount < 0 {
		return fmt.Errorf("retry count must not be negative")
	}
	return nil
}
