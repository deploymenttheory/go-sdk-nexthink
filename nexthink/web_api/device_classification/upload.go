package device_classification

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"mime"
	"mime/multipart"
	"net/textproto"
	"strings"
	"unicode/utf8"
)

var errRequestRequired = fmt.Errorf("request is required")

func validateUpload(request *RulesetUpload, create bool) error {
	if request == nil {
		return errRequestRequired
	}
	if strings.TrimSpace(request.Name) == "" || strings.ContainsAny(request.Name, "\r\n") {
		return fmt.Errorf("name is required and must not contain line breaks")
	}
	if create && len(request.CSV) == 0 {
		return fmt.Errorf("CSV file is required for create")
	}
	if request.CSV != nil {
		if request.Filename == "" {
			return fmt.Errorf("filename is required with CSV")
		}
		if len(request.CSV) > 5*1024*1024 {
			return fmt.Errorf("CSV exceeds the UI 5 MiB limit")
		}
		if !utf8.Valid(request.CSV) {
			return fmt.Errorf("CSV must be UTF-8")
		}
	}
	return nil
}
func (s *Service) upload(ctx context.Context, method, path string, request *RulesetUpload, create bool) (*interfaces.Response, error) {
	if err := validateUpload(request, create); err != nil {
		return nil, err
	}
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	if request.CSV != nil {
		h := textproto.MIMEHeader{}
		h.Set("Content-Disposition", mime.FormatMediaType("form-data", map[string]string{"name": "file", "filename": request.Filename}))
		h.Set("Content-Type", "text/csv")
		part, err := writer.CreatePart(h)
		if err != nil {
			return nil, err
		}
		if _, err = part.Write(bytes.TrimPrefix(request.CSV, []byte{0xef, 0xbb, 0xbf})); err != nil {
			return nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	headers := map[string]string{"Content-Type": writer.FormDataContentType(), "x-nxt-entities-name": request.Name, "x-nxt-entities-description": base64.StdEncoding.EncodeToString([]byte(request.Description))}
	if method == "POST" {
		return s.client.Post(ctx, path, buffer.Bytes(), headers, nil)
	}
	return s.client.Put(ctx, path, buffer.Bytes(), headers, nil)
}
