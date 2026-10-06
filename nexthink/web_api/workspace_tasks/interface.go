package workspace_tasks

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type WorkspaceTasksServiceInterface interface {
	ListTasks(ctx context.Context, options *TaskListOptions) (*TaskListResponse, *interfaces.Response, error)
	GetTask(ctx context.Context, id string) (*Task, *interfaces.Response, error)
	CreateTask(ctx context.Context, request *TaskRequest) (*Task, *interfaces.Response, error)
	UpdateTask(ctx context.Context, id string, request *TaskRequest) (*Task, *interfaces.Response, error)
	DeleteTask(ctx context.Context, id string) (*interfaces.Response, error)
	ReconcileTaskAgentAccess(ctx context.Context) (*interfaces.Response, error)
}

var _ WorkspaceTasksServiceInterface = (*Service)(nil)
