package workspace_agents

import (
	"context"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"net/url"
	"strconv"
	"strings"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// ListSkills calls the Workspace UI GET contract.
func (s *Service) ListSkills(ctx context.Context, source string) ([]Skill, *interfaces.Response, error) {

	path := SkillsEndpoint
	if source != "" {
		path += "?" + url.Values{"source": {source}}.Encode()
	}
	var result []Skill
	response, err := s.client.Get(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// CheckSkillAvailability calls the Workspace UI GET contract.
func (s *Service) CheckSkillAvailability(ctx context.Context, source string) (*AvailabilityResponse, *interfaces.Response, error) {
	if strings.TrimSpace(source) == "" {
		return nil, nil, fmt.Errorf("source is required")
	}
	path := SkillsEndpoint + "?" + url.Values{"check-availability": {source}}.Encode()

	var result AvailabilityResponse
	response, err := s.client.Get(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetSkill calls the Workspace UI GET contract.
func (s *Service) GetSkill(ctx context.Context, id string) (*Skill, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	path := SkillsEndpoint + "/" + url.PathEscape(id)

	var result Skill
	response, err := s.client.Get(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// CreateSkill calls the Workspace UI POST contract.
func (s *Service) CreateSkill(ctx context.Context, request *SkillRequest) (*Skill, *interfaces.Response, error) {
	if err := validateSkill(request); err != nil {
		return nil, nil, err
	}
	if err := validateID(request.ID); err != nil {
		return nil, nil, err
	}
	path := SkillsEndpoint

	var result Skill
	response, err := s.client.Post(ctx, path, request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// UpdateSkill calls the Workspace UI PUT contract.
func (s *Service) UpdateSkill(ctx context.Context, id string, request *SkillRequest) (*Skill, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	if err := validateSkill(request); err != nil {
		return nil, nil, err
	}
	path := SkillsEndpoint + "/" + url.PathEscape(id)

	var result Skill
	response, err := s.client.Put(ctx, path, request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// DeleteSkill calls the Workspace UI DELETE contract.
func (s *Service) DeleteSkill(ctx context.Context, id string) (*interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}
	path := SkillsEndpoint + "/" + url.PathEscape(id)

	return s.client.Delete(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, nil)
}

// UploadSkillFile calls the Workspace UI POST contract.
func (s *Service) UploadSkillFile(ctx context.Context, id string, request *FileUploadRequest) (*Document, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	if request == nil || request.Filename == "" || request.MIMEType == "" || len(request.Data) == 0 {
		return nil, nil, fmt.Errorf("filename, mime_type and data are required")
	}
	path := SkillsEndpoint + "/skills/" + url.PathEscape(id) + "/knowledgebase/files"
	path += "?" + url.Values{"filename": {request.Filename}, "mime_type": {request.MIMEType}}.Encode()
	var result Document
	response, err := s.client.Post(ctx, path, request.Data, map[string]string{"Content-Type": "application/octet-stream"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// DeleteSkillFile calls the Workspace UI DELETE contract.
func (s *Service) DeleteSkillFile(ctx context.Context, id, fileID string) (*interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}
	if err := validateID(fileID); err != nil {
		return nil, err
	}
	path := SkillsEndpoint + "/skills/" + url.PathEscape(id) + "/knowledgebase/files" + "/" + url.PathEscape(fileID)

	return s.client.Delete(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, nil)
}

// StartSkillMultipartUpload calls the Workspace UI POST contract.
func (s *Service) StartSkillMultipartUpload(ctx context.Context, id string, request *StartMultipartRequest) (*MultipartUpload, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	if request == nil || request.Filename == "" || request.MIMEType == "" {
		return nil, nil, fmt.Errorf("filename and mime_type are required")
	}
	path := SkillsEndpoint + "/skills/" + url.PathEscape(id) + "/knowledgebase/files" + "/multipart"
	path += "?" + url.Values{"filename": {request.Filename}, "mime_type": {request.MIMEType}}.Encode()
	var result MultipartUpload
	response, err := s.client.Post(ctx, path, nil, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// UploadSkillPart calls the Workspace UI PUT contract.
func (s *Service) UploadSkillPart(ctx context.Context, id string, request *UploadPartRequest) (*UploadedPart, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	if request == nil || request.UploadID == "" || request.FileID == "" || request.PartNumber < 1 || len(request.Data) == 0 {
		return nil, nil, fmt.Errorf("uploadId, fileId, positive partNumber and data are required")
	}
	path := SkillsEndpoint + "/skills/" + url.PathEscape(id) + "/knowledgebase/files" + "/multipart"
	path += "?" + url.Values{"uploadId": {request.UploadID}, "fileId": {request.FileID}, "partNumber": {strconv.Itoa(request.PartNumber)}}.Encode()
	var result UploadedPart
	response, err := s.client.Put(ctx, path, request.Data, map[string]string{"Content-Type": "application/octet-stream"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// CompleteSkillMultipartUpload calls the Workspace UI POST contract.
func (s *Service) CompleteSkillMultipartUpload(ctx context.Context, id string, request *CompleteMultipartRequest) (*Document, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	if request == nil || request.UploadID == "" || request.FileID == "" || len(request.Parts) == 0 {
		return nil, nil, fmt.Errorf("uploadId, fileId and parts are required")
	}
	path := SkillsEndpoint + "/skills/" + url.PathEscape(id) + "/knowledgebase/files" + "/multipart/complete"
	path += "?" + url.Values{"uploadId": {request.UploadID}, "fileId": {request.FileID}}.Encode()
	var result Document
	response, err := s.client.Post(ctx, path, request.Parts, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
