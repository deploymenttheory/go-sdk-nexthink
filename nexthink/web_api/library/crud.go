package library

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"net/url"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// ListContents calls the corresponding Library UI operation.
func (s *Service) ListContents(ctx context.Context) (*ContentsResponse, *interfaces.Response, error) {
	var result ContentsResponse
	response, err := s.client.Get(ctx, Endpoint+"/builtincontent", nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// ListPacks calls the corresponding Library UI operation.
func (s *Service) ListPacks(ctx context.Context) (*PacksResponse, *interfaces.Response, error) {
	var result PacksResponse
	response, err := s.client.Get(ctx, Endpoint+"/packs", nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetContent calls the corresponding Library UI operation.
func (s *Service) GetContent(ctx context.Context, fileName string) (*Content, *interfaces.Response, error) {
	if err := validateReference("fileName", fileName); err != nil {
		return nil, nil, err
	}
	var result Content
	response, err := s.client.Get(ctx, Endpoint+"/builtincontent/"+url.PathEscape(fileName), nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetCustomContent calls the corresponding Library UI operation.
func (s *Service) GetCustomContent(ctx context.Context, contentID, resourceName string) (*Content, *interfaces.Response, error) {
	if err := validateReference("contentID", contentID); err != nil {
		return nil, nil, err
	}
	if err := validateReference("resourceName", resourceName); err != nil {
		return nil, nil, err
	}
	var result Content
	response, err := s.client.Get(ctx, Endpoint+"/builtincontent/contents/"+url.PathEscape(contentID)+"/resourceName/"+url.PathEscape(resourceName), nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetCreateCopyInfo calls the corresponding Library UI operation.
func (s *Service) GetCreateCopyInfo(ctx context.Context, contentType string) (*CreateCopyInfo, *interfaces.Response, error) {
	if err := validateReference("contentType", contentType); err != nil {
		return nil, nil, err
	}
	var result CreateCopyInfo
	response, err := s.client.Get(ctx, ContentAdminEndpoint+"/contents/"+url.PathEscape(contentType)+"/config", nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// InstallContent calls the corresponding Library UI operation.
func (s *Service) InstallContent(ctx context.Context, request *ContentInstallationRequest) (*OperationResult, *interfaces.Response, error) {
	if err := validateInstallContent(request); err != nil {
		return nil, nil, err
	}
	var result OperationResult
	response, err := s.client.Post(ctx, Endpoint+"/builtincontent/import", request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// InstallCustomContent calls the corresponding Library UI operation.
func (s *Service) InstallCustomContent(ctx context.Context, contentID, resourceName string) (*OperationResult, *interfaces.Response, error) {
	if err := validateReference("contentID", contentID); err != nil {
		return nil, nil, err
	}
	if err := validateReference("resourceName", resourceName); err != nil {
		return nil, nil, err
	}
	var result OperationResult
	response, err := s.client.Post(ctx, Endpoint+"/builtincontent/contents/"+url.PathEscape(contentID)+"/resourceName/"+url.PathEscape(resourceName)+"/install", nil, map[string]string{"Accept": "application/json", "nx-caller-service": CallerService}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// InstallPack calls the corresponding Library UI operation.
func (s *Service) InstallPack(ctx context.Context, request *Pack) (*Pack, *interfaces.Response, error) {
	if err := validateInstallPack(request); err != nil {
		return nil, nil, err
	}
	var result Pack
	response, err := s.client.Post(ctx, Endpoint+"/packs", request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// InstallCustomPack calls the corresponding Library UI operation.
func (s *Service) InstallCustomPack(ctx context.Context, packUUID string) (*OperationResult, *interfaces.Response, error) {
	if err := validateReference("packUUID", packUUID); err != nil {
		return nil, nil, err
	}
	var result OperationResult
	response, err := s.client.Post(ctx, Endpoint+"/packs/"+url.PathEscape(packUUID)+"/install", nil, map[string]string{"Accept": "application/json", "nx-caller-service": CallerService}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// UpdateContent calls the corresponding Library UI operation.
func (s *Service) UpdateContent(ctx context.Context, request *ContentUpdateRequest) (*OperationResult, *interfaces.Response, error) {
	if err := validateUpdateContent(request); err != nil {
		return nil, nil, err
	}
	var result OperationResult
	response, err := s.client.Put(ctx, Endpoint+"/builtincontent/update", request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// UpdatePack calls the corresponding Library UI operation.
func (s *Service) UpdatePack(ctx context.Context, packUUID string) (*OperationResult, *interfaces.Response, error) {
	if err := validateReference("packUUID", packUUID); err != nil {
		return nil, nil, err
	}
	var result OperationResult
	response, err := s.client.Post(ctx, Endpoint+"/packs/"+url.PathEscape(packUUID)+"/install/latest", nil, map[string]string{"Accept": "application/json", "nx-caller-service": CallerService}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// UpdateCustomContent calls the corresponding Library UI operation.
func (s *Service) UpdateCustomContent(ctx context.Context, contentID, resourceName string) (*OperationResult, *interfaces.Response, error) {
	if err := validateReference("contentID", contentID); err != nil {
		return nil, nil, err
	}
	if err := validateReference("resourceName", resourceName); err != nil {
		return nil, nil, err
	}
	var result OperationResult
	response, err := s.client.Post(ctx, Endpoint+"/builtincontent/contents/"+url.PathEscape(contentID)+"/resourceName/"+url.PathEscape(resourceName)+"/install/latest", nil, map[string]string{"Accept": "application/json", "nx-caller-service": CallerService}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetDependenciesStatus calls the corresponding Library UI operation.
func (s *Service) GetDependenciesStatus(ctx context.Context, request *DependenciesRequest) (*DependenciesResponse, *interfaces.Response, error) {
	if err := validateGetDependenciesStatus(request); err != nil {
		return nil, nil, err
	}
	var result DependenciesResponse
	response, err := s.client.Post(ctx, Endpoint+"/builtincontent/dependency-status", request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// InstallDependencies calls the corresponding Library UI operation.
func (s *Service) InstallDependencies(ctx context.Context, request *DependenciesRequest) (*OperationResult, *interfaces.Response, error) {
	if err := validateInstallDependencies(request); err != nil {
		return nil, nil, err
	}
	var result OperationResult
	response, err := s.client.Post(ctx, Endpoint+"/builtincontent/multi-import", request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetPack calls the corresponding Library UI operation.
func (s *Service) GetPack(ctx context.Context, fileName, packUUID string) (*Pack, *interfaces.Response, error) {
	if err := validateReference("fileName", fileName); err != nil {
		return nil, nil, err
	}
	query := map[string]string{}
	if packUUID != "" {
		query["packUuid"] = packUUID
	}
	var result Pack
	response, err := s.client.Get(ctx, Endpoint+"/packs/"+url.PathEscape(fileName), query, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// ImportCustomPack calls the corresponding Library UI operation.
func (s *Service) ImportCustomPack(ctx context.Context, libraryID string) (*OperationResult, *interfaces.Response, error) {
	if err := validateReference("libraryID", libraryID); err != nil {
		return nil, nil, err
	}
	var result OperationResult
	response, err := s.client.Post(ctx, Endpoint+"/packs/"+url.PathEscape(libraryID)+"/import", nil, map[string]string{"Accept": "application/json", "nx-caller-service": CallerService}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetPackInstallationStatus calls the corresponding Library UI operation.
func (s *Service) GetPackInstallationStatus(ctx context.Context, libraryID string) (*InstallationStatus, *interfaces.Response, error) {
	if err := validateReference("libraryID", libraryID); err != nil {
		return nil, nil, err
	}
	var result InstallationStatus
	response, err := s.client.Get(ctx, Endpoint+"/packs/"+url.PathEscape(libraryID)+"/status", nil, map[string]string{"Accept": "application/json", "nx-caller-service": CallerService}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetContentInstallationStatus calls the corresponding Library UI operation.
func (s *Service) GetContentInstallationStatus(ctx context.Context, contentID, resourceName string) (*InstallationStatus, *interfaces.Response, error) {
	if err := validateReference("contentID", contentID); err != nil {
		return nil, nil, err
	}
	if err := validateReference("resourceName", resourceName); err != nil {
		return nil, nil, err
	}
	var result InstallationStatus
	response, err := s.client.Get(ctx, Endpoint+"/builtincontent/contents/"+url.PathEscape(contentID)+"/resourceName/"+url.PathEscape(resourceName)+"/status", nil, map[string]string{"Accept": "application/json", "nx-caller-service": CallerService}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// DeleteCustomPack calls the corresponding Library UI operation.
func (s *Service) DeleteCustomPack(ctx context.Context, libraryID string) (*interfaces.Response, error) {
	if err := validateReference("libraryID", libraryID); err != nil {
		return nil, err
	}
	return s.client.Delete(ctx, Endpoint+"/packs/"+url.PathEscape(libraryID), nil, map[string]string{"Accept": "application/json", "nx-caller-service": CallerService}, nil)
}

// GetLocale calls the corresponding Library UI operation.
func (s *Service) GetLocale(ctx context.Context) (*LocaleResponse, *interfaces.Response, error) {
	var result LocaleResponse
	response, err := s.client.Get(ctx, Endpoint+"/locale", nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type LibraryServiceInterface interface {
	ListContents(ctx context.Context) (*ContentsResponse, *interfaces.Response, error)
	ListPacks(ctx context.Context) (*PacksResponse, *interfaces.Response, error)
	GetContent(ctx context.Context, fileName string) (*Content, *interfaces.Response, error)
	GetCustomContent(ctx context.Context, contentID, resourceName string) (*Content, *interfaces.Response, error)
	GetCreateCopyInfo(ctx context.Context, contentType string) (*CreateCopyInfo, *interfaces.Response, error)
	InstallContent(ctx context.Context, request *ContentInstallationRequest) (*OperationResult, *interfaces.Response, error)
	InstallCustomContent(ctx context.Context, contentID, resourceName string) (*OperationResult, *interfaces.Response, error)
	InstallPack(ctx context.Context, request *Pack) (*Pack, *interfaces.Response, error)
	InstallCustomPack(ctx context.Context, packUUID string) (*OperationResult, *interfaces.Response, error)
	UpdateContent(ctx context.Context, request *ContentUpdateRequest) (*OperationResult, *interfaces.Response, error)
	UpdatePack(ctx context.Context, packUUID string) (*OperationResult, *interfaces.Response, error)
	UpdateCustomContent(ctx context.Context, contentID, resourceName string) (*OperationResult, *interfaces.Response, error)
	GetDependenciesStatus(ctx context.Context, request *DependenciesRequest) (*DependenciesResponse, *interfaces.Response, error)
	InstallDependencies(ctx context.Context, request *DependenciesRequest) (*OperationResult, *interfaces.Response, error)
	GetPack(ctx context.Context, fileName, packUUID string) (*Pack, *interfaces.Response, error)
	ImportCustomPack(ctx context.Context, libraryID string) (*OperationResult, *interfaces.Response, error)
	GetPackInstallationStatus(ctx context.Context, libraryID string) (*InstallationStatus, *interfaces.Response, error)
	GetContentInstallationStatus(ctx context.Context, contentID, resourceName string) (*InstallationStatus, *interfaces.Response, error)
	DeleteCustomPack(ctx context.Context, libraryID string) (*interfaces.Response, error)
	GetLocale(ctx context.Context) (*LocaleResponse, *interfaces.Response, error)
}

var _ LibraryServiceInterface = (*Service)(nil)
