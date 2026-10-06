package observability

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"strconv"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// Submit sends a batch to the fixed Nexthink observability proxy; it never contacts a caller-supplied host.
func (s *Service) Submit(ctx context.Context, r *Submission) (*interfaces.Response, error) {
	if err := validateSubmission(r); err != nil {
		return nil, err
	}
	q := map[string]string{"ddsource": r.Source, "dd-api-key": r.ClientToken, "dd-evp-origin": r.Origin, "dd-evp-origin-version": r.OriginVersion, "dd-request-id": r.RequestID}
	for k, v := range map[string]string{"dd-evp-encoding": r.Encoding, "batch_time": r.BatchTime, "_dd.api": r.API, "_dd.retry_after": r.RetryAfter} {
		if v != "" {
			q[k] = v
		}
	}
	if r.RetryCount != nil {
		q["_dd.retry_count"] = strconv.Itoa(*r.RetryCount)
	}
	return s.client.PostWithQuery(ctx, Endpoint, q, r.Payload, map[string]string{"Content-Type": r.ContentType}, nil)
}

type ObservabilityServiceInterface interface {
	Submit(context.Context, *Submission) (*interfaces.Response, error)
}

var _ ObservabilityServiceInterface = (*Service)(nil)
