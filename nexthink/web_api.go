package nexthink

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api"
)

// Do executes an observed operation from the catalog and retains the original JSON.
// Operations whose schemas remain undocumented expose raw JSON rather than
// inventing types or silently discarding fields.
func (c *WebAPIClient) Do(
	ctx context.Context,
	operationID string,
	request web_api.Request,
) (json.RawMessage, *interfaces.Response, error) {
	op, ok := web_api.OperationByID(operationID)
	if !ok {
		return nil, nil, fmt.Errorf("unknown web API operation %q", operationID)
	}
	path := op.Path
	for key, value := range request.PathParams {
		if value == "" || value == "." || value == ".." || strings.ContainsAny(value, "/\\") ||
			!strings.Contains(path, "{"+key+"}") {
			return nil, nil, fmt.Errorf("invalid path parameter %q", key)
		}
		path = strings.ReplaceAll(path, "{"+key+"}", url.PathEscape(value))
	}
	if strings.ContainsAny(path, "{}") {
		return nil, nil, fmt.Errorf("missing path parameters for %s", operationID)
	}
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	if request.ContentType != "" {
		headers["Content-Type"] = request.ContentType
	}
	var result json.RawMessage
	var resp *interfaces.Response
	var err error
	switch op.Method {
	case "GET":
		if request.Body != nil {
			return nil, nil, fmt.Errorf("GET operation does not accept a body")
		}
		resp, err = c.transport.Get(ctx, path, request.Query, headers, &result)
	case "POST":
		resp, err = c.transport.PostWithQuery(
			ctx,
			path,
			request.Query,
			request.Body,
			headers,
			&result,
		)
	case "PUT", "PATCH", "DELETE":
		if len(request.Query) > 0 {
			v := url.Values{}
			for k, x := range request.Query {
				v.Set(k, x)
			}
			path += "?" + v.Encode()
		}
		switch op.Method {
		case "PUT":
			resp, err = c.transport.Put(ctx, path, request.Body, headers, &result)
		case "PATCH":
			resp, err = c.transport.Patch(ctx, path, request.Body, headers, &result)
		case "DELETE":
			resp, err = c.transport.DeleteWithBody(ctx, path, request.Body, headers, &result)
		}
	default:
		return nil, nil, fmt.Errorf("unsupported operation method %q", op.Method)
	}
	return result, resp, err
}
