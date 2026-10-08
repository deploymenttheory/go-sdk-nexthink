package workspace

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type WorkspaceServiceInterface interface {
	ListConversations(ctx context.Context, automationID string) ([]Conversation, *interfaces.Response, error)
	GetConversation(ctx context.Context, id string, options *PageOptions) (*Conversation, *interfaces.Response, error)
	GetSharedConversation(ctx context.Context, id string, options *PageOptions) (*Conversation, *interfaces.Response, error)
	UpdateConversation(ctx context.Context, id string, request *ConversationUpdate) (*Conversation, *interfaces.Response, error)
	DeleteConversation(ctx context.Context, id string) (*interfaces.Response, error)
	CancelConversation(ctx context.Context, id string) (*Document, *interfaces.Response, error)
	MarkConversationRead(ctx context.Context, id string) (*Conversation, *interfaces.Response, error)
	CreateConversationShare(ctx context.Context, id string) (*Conversation, *interfaces.Response, error)
	UploadConversationFile(ctx context.Context, id string, request *ConversationFileRequest) (*FileResponse, *interfaces.Response, error)
	DeleteConversationFile(ctx context.Context, id, fileID string) (*interfaces.Response, error)
	MCPProxy(ctx context.Context, request *MCPRequest) (*Document, *interfaces.Response, error)
	Chat(context.Context, *ChatRequest) (*ChatResponse, *interfaces.Response, error)
}

var _ WorkspaceServiceInterface = (*Service)(nil)
