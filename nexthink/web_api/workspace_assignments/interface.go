package workspace_assignments

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type WorkspaceAssignmentsServiceInterface interface {
	ListAssignments(ctx context.Context, options *AssignmentOptions) ([]Assignment, *interfaces.Response, error)
	GetAssignment(ctx context.Context, id string, options *AssignmentOptions) (*AssignmentResponse, *interfaces.Response, error)
	UpdateAssignment(ctx context.Context, id string, request *AssignmentUpdate) (*AssignmentResponse, *interfaces.Response, error)
	MarkAssignmentRead(ctx context.Context, id string) (*interfaces.Response, error)
	GetUnreadAssignmentCount(ctx context.Context, options *AssignmentOptions) (*UnreadCountResponse, *interfaces.Response, error)
	ListAssignees(ctx context.Context, id string) ([]Document, *interfaces.Response, error)
}

var _ WorkspaceAssignmentsServiceInterface = (*Service)(nil)
