package content_sharing

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }
func (s *Service) GetActions(ctx context.Context, o *ActionsOptions) (*ActionsResponse, *interfaces.Response, error) {
	if err := validateActions(o); err != nil {
		return nil, nil, err
	}
	var result ActionsResponse
	resp, err := s.client.Get(ctx, Endpoint+"/v3/share/actions", map[string]string{"contentKey": o.ContentKey, "resourceName": o.ResourceName}, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
func (s *Service) GetProfiles(ctx context.Context, o *ProfilesOptions) (*ProfilesResponse, *interfaces.Response, error) {
	if err := validateProfiles(o); err != nil {
		return nil, nil, err
	}
	var result ProfilesResponse
	resp, err := s.client.Get(ctx, Endpoint+"/v3/share/profiles", map[string]string{"contentKey": o.ContentKey, "resourceName": o.ResourceName, "contentId": o.ContentID, "shared": strconv.FormatBool(o.Shared), "bcsName": o.BCSName}, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
func (s *Service) GetUser(ctx context.Context, id string) (*UserResponse, *interfaces.Response, error) {
	if err := validation.PathSegment(id); err != nil {
		return nil, nil, err
	}
	var result UserResponse
	resp, err := s.client.Get(ctx, Endpoint+"/v2/user/"+url.PathEscape(id), nil, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
func (s *Service) SetProfiles(ctx context.Context, contentKey string, request []ShareContent) (*interfaces.Response, error) {
	if err := required(contentKey); err != nil {
		return nil, err
	}
	if len(request) == 0 {
		return nil, fmt.Errorf("at least one content entry is required")
	}
	for _, entry := range request {
		if err := required(entry.ContentID, entry.ResourceName); err != nil {
			return nil, err
		}
		if entry.Profiles == nil {
			return nil, fmt.Errorf("profiles must be an array")
		}
		for _, grant := range entry.Profiles {
			if err := required(grant.RoleUUID); err != nil {
				return nil, fmt.Errorf("roleUuid must not be empty")
			}
			if grant.Actions == nil {
				return nil, fmt.Errorf("actions must be an array; use an explicit empty array to revoke a grant")
			}
		}
	}
	return s.client.PostWithQuery(ctx, Endpoint+"/v3/share/profiles", map[string]string{"contentKey": contentKey}, request, map[string]string{"Content-Type": "application/json"}, nil)
}
func (s *Service) GetLegacyActions(ctx context.Context, service string) (*LegacyActionsResponse, *interfaces.Response, error) {
	if err := required(service); err != nil {
		return nil, nil, err
	}
	var result LegacyActionsResponse
	resp, err := s.client.Get(ctx, EndpointLegacy+"/permissions", map[string]string{"service": service}, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetLegacyProfiles returns granted profiles when shared is true and ungranted
// profiles when false. A profile moves between these views as grants change.
func (s *Service) GetLegacyProfiles(ctx context.Context, o *LegacyOptions, shared bool) (*LegacyProfilesResponse, *interfaces.Response, error) {
	if err := validateLegacy(o); err != nil {
		return nil, nil, err
	}
	q := legacyQuery(o)
	q["shared"] = strconv.FormatBool(shared)
	var result LegacyProfilesResponse
	resp, err := s.client.Get(ctx, EndpointLegacy+"/profiles", q, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
func (s *Service) GetLegacyOwner(ctx context.Context, id string) (*UserResponse, *interfaces.Response, error) {
	if err := validation.PathSegment(id); err != nil {
		return nil, nil, err
	}
	var result UserResponse
	resp, err := s.client.Get(ctx, EndpointLegacy+"/content-owner/"+url.PathEscape(id), nil, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
func (s *Service) SetLegacyProfiles(ctx context.Context, o *LegacyOptions, request *LegacyUpdateRequest) (*LegacyProfilesResponse, *interfaces.Response, error) {
	if err := validateLegacy(o); err != nil {
		return nil, nil, err
	}
	if request == nil || request.Profiles == nil {
		return nil, nil, fmt.Errorf("request and profiles array required")
	}
	for _, grant := range request.Profiles {
		if grant.Actions == nil {
			return nil, nil, fmt.Errorf("actions must be an array; use an explicit empty array to revoke a grant")
		}
	}
	var result LegacyProfilesResponse
	resp, err := s.client.PostWithQuery(ctx, EndpointLegacy+"/profiles", legacyQuery(o), request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
func legacyQuery(o *LegacyOptions) map[string]string {
	return map[string]string{"service": o.Service, "contentId": o.ContentID, "bcsName": o.BCSName, "tag": o.Tag}
}

type ContentSharingServiceInterface interface {
	GetActions(context.Context, *ActionsOptions) (*ActionsResponse, *interfaces.Response, error)
	GetProfiles(context.Context, *ProfilesOptions) (*ProfilesResponse, *interfaces.Response, error)
	GetUser(context.Context, string) (*UserResponse, *interfaces.Response, error)
	SetProfiles(context.Context, string, []ShareContent) (*interfaces.Response, error)
	GetLegacyActions(context.Context, string) (*LegacyActionsResponse, *interfaces.Response, error)
	GetLegacyProfiles(context.Context, *LegacyOptions, bool) (*LegacyProfilesResponse, *interfaces.Response, error)
	GetLegacyOwner(context.Context, string) (*UserResponse, *interfaces.Response, error)
	SetLegacyProfiles(context.Context, *LegacyOptions, *LegacyUpdateRequest) (*LegacyProfilesResponse, *interfaces.Response, error)
}

var _ ContentSharingServiceInterface = (*Service)(nil)
