package workspace_agents

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type WorkspaceAgentsServiceInterface interface {
	ListSkills(ctx context.Context, source string) ([]Skill, *interfaces.Response, error)
	CheckSkillAvailability(ctx context.Context, source string) (*AvailabilityResponse, *interfaces.Response, error)
	GetSkill(ctx context.Context, id string) (*Skill, *interfaces.Response, error)
	CreateSkill(ctx context.Context, request *SkillRequest) (*Skill, *interfaces.Response, error)
	UpdateSkill(ctx context.Context, id string, request *SkillRequest) (*Skill, *interfaces.Response, error)
	DeleteSkill(ctx context.Context, id string) (*interfaces.Response, error)
	UploadSkillFile(ctx context.Context, id string, request *FileUploadRequest) (*Document, *interfaces.Response, error)
	DeleteSkillFile(ctx context.Context, id, fileID string) (*interfaces.Response, error)
	StartSkillMultipartUpload(ctx context.Context, id string, request *StartMultipartRequest) (*MultipartUpload, *interfaces.Response, error)
	UploadSkillPart(ctx context.Context, id string, request *UploadPartRequest) (*UploadedPart, *interfaces.Response, error)
	CompleteSkillMultipartUpload(ctx context.Context, id string, request *CompleteMultipartRequest) (*Document, *interfaces.Response, error)
}

var _ WorkspaceAgentsServiceInterface = (*Service)(nil)
