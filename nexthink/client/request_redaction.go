package client

import (
	"errors"
	"net/url"
	"strings"
)

// requestLogPath excludes URL query values, which can contain browser service
// credentials. The actual request URL and query are never modified.
func requestLogPath(path string) string {
	u, err := url.Parse(path)
	if err != nil {
		if i := strings.IndexByte(path, '?'); i >= 0 {
			return path[:i]
		}
		return path
	}
	u.RawQuery = ""
	u.ForceQuery = false
	u.Fragment = ""
	u.User = nil
	return u.String()
}

type redactedRequestError struct {
	message string
	cause   error
}

func (e *redactedRequestError) Error() string { return e.message }
func (e *redactedRequestError) Unwrap() error { return e.cause }

// redactRequestError retains a sanitized url.Error and its underlying cause,
// allowing errors.As/Is without returning a credential-bearing URL.
func redactRequestError(err error) error {
	if err == nil {
		return nil
	}
	var original *url.Error
	if !errors.As(err, &original) {
		return err
	}
	clone := *original
	clone.URL = requestLogPath(original.URL)
	clone.Err = redactRequestError(original.Err)
	if err == original {
		return &clone
	}
	message := strings.ReplaceAll(err.Error(), original.Error(), clone.Error())
	return &redactedRequestError{message: message, cause: &clone}
}
