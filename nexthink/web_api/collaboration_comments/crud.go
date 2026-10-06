package collaboration_comments

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"net/url"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// ResolveIdentifier implements the RTC UI POST /identifier operation.
func (s *Service) ResolveIdentifier(ctx context.Context, request *IdentifierRequest) (*IdentifierResponse, *interfaces.Response, error) {
	if err := validateRequest(request); err != nil {
		return nil, nil, err
	}
	var result IdentifierResponse
	response, err := s.client.Post(ctx, Endpoint+"/identifier", request, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// ListComments implements the RTC UI GET /documents/{documentID}/comments operation.
func (s *Service) ListComments(ctx context.Context, documentID string) (*CommentsResponse, *interfaces.Response, error) {
	if err := validateReference("documentID", documentID); err != nil {
		return nil, nil, err
	}
	var result CommentsResponse
	response, err := s.client.Get(ctx, Endpoint+"/documents/"+url.PathEscape(documentID)+"/comments", nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// CreateComment implements the RTC UI POST /documents/{documentID}/comments operation.
func (s *Service) CreateComment(ctx context.Context, documentID string, request *CreateMessageRequest) (*interfaces.Response, error) {
	if err := validateReference("documentID", documentID); err != nil {
		return nil, err
	}
	if err := validateRequest(request); err != nil {
		return nil, err
	}
	return s.client.Post(ctx, Endpoint+"/documents/"+url.PathEscape(documentID)+"/comments", request, map[string]string{"Accept": "application/json"}, nil)
}

// CreateReply implements the RTC UI POST /documents/{documentID}/comments/{commentID}/replies operation.
func (s *Service) CreateReply(ctx context.Context, documentID string, commentID string, request *CreateMessageRequest) (*interfaces.Response, error) {
	if err := validateReference("documentID", documentID); err != nil {
		return nil, err
	}
	if err := validateReference("commentID", commentID); err != nil {
		return nil, err
	}
	if err := validateRequest(request); err != nil {
		return nil, err
	}
	return s.client.Post(ctx, Endpoint+"/documents/"+url.PathEscape(documentID)+"/comments/"+url.PathEscape(commentID)+"/replies", request, map[string]string{"Accept": "application/json"}, nil)
}

// ListUserMentions implements the RTC UI GET /user-mentions operation.
func (s *Service) ListUserMentions(ctx context.Context, search string) (*UserMentionsResponse, *interfaces.Response, error) {
	var result UserMentionsResponse
	response, err := s.client.Get(ctx, Endpoint+"/user-mentions", map[string]string{"u": search}, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// ArchiveComment implements the RTC UI PATCH /documents/{documentID}/comments/{commentID} operation.
func (s *Service) ArchiveComment(ctx context.Context, documentID string, commentID string) (*interfaces.Response, error) {
	if err := validateReference("documentID", documentID); err != nil {
		return nil, err
	}
	if err := validateReference("commentID", commentID); err != nil {
		return nil, err
	}
	return s.client.Patch(ctx, Endpoint+"/documents/"+url.PathEscape(documentID)+"/comments/"+url.PathEscape(commentID)+"", CommentOperation{Operation: "archive"}, map[string]string{"Accept": "application/json"}, nil)
}

// UnarchiveComment implements the RTC UI PATCH /documents/{documentID}/comments/{commentID} operation.
func (s *Service) UnarchiveComment(ctx context.Context, documentID string, commentID string) (*interfaces.Response, error) {
	if err := validateReference("documentID", documentID); err != nil {
		return nil, err
	}
	if err := validateReference("commentID", commentID); err != nil {
		return nil, err
	}
	return s.client.Patch(ctx, Endpoint+"/documents/"+url.PathEscape(documentID)+"/comments/"+url.PathEscape(commentID)+"", CommentOperation{Operation: "unarchive"}, map[string]string{"Accept": "application/json"}, nil)
}

// UpdateComment implements the RTC UI PATCH /documents/{documentID}/comments/{commentID} operation.
func (s *Service) UpdateComment(ctx context.Context, documentID string, commentID string, request *EditMessageRequest) (*interfaces.Response, error) {
	if err := validateReference("documentID", documentID); err != nil {
		return nil, err
	}
	if err := validateReference("commentID", commentID); err != nil {
		return nil, err
	}
	if err := validateRequest(request); err != nil {
		return nil, err
	}
	return s.client.Patch(ctx, Endpoint+"/documents/"+url.PathEscape(documentID)+"/comments/"+url.PathEscape(commentID)+"", CommentOperation{Operation: "editMessage", Data: request}, map[string]string{"Accept": "application/json"}, nil)
}

// UpdateReply implements the RTC UI PATCH /documents/{documentID}/comments/{commentID}/replies/{replyID} operation.
func (s *Service) UpdateReply(ctx context.Context, documentID string, commentID string, replyID string, request *EditMessageRequest) (*interfaces.Response, error) {
	if err := validateReference("documentID", documentID); err != nil {
		return nil, err
	}
	if err := validateReference("commentID", commentID); err != nil {
		return nil, err
	}
	if err := validateReference("replyID", replyID); err != nil {
		return nil, err
	}
	if err := validateRequest(request); err != nil {
		return nil, err
	}
	return s.client.Patch(ctx, Endpoint+"/documents/"+url.PathEscape(documentID)+"/comments/"+url.PathEscape(commentID)+"/replies/"+url.PathEscape(replyID)+"", request, map[string]string{"Accept": "application/json"}, nil)
}

// DeleteComment implements the RTC UI DELETE /documents/{documentID}/comments/{commentID} operation.
func (s *Service) DeleteComment(ctx context.Context, documentID string, commentID string) (*interfaces.Response, error) {
	if err := validateReference("documentID", documentID); err != nil {
		return nil, err
	}
	if err := validateReference("commentID", commentID); err != nil {
		return nil, err
	}
	return s.client.Delete(ctx, Endpoint+"/documents/"+url.PathEscape(documentID)+"/comments/"+url.PathEscape(commentID)+"", nil, map[string]string{"Accept": "application/json"}, nil)
}

// DeleteReply implements the RTC UI DELETE /documents/{documentID}/comments/{commentID}/replies/{replyID} operation.
func (s *Service) DeleteReply(ctx context.Context, documentID string, commentID string, replyID string) (*interfaces.Response, error) {
	if err := validateReference("documentID", documentID); err != nil {
		return nil, err
	}
	if err := validateReference("commentID", commentID); err != nil {
		return nil, err
	}
	if err := validateReference("replyID", replyID); err != nil {
		return nil, err
	}
	return s.client.Delete(ctx, Endpoint+"/documents/"+url.PathEscape(documentID)+"/comments/"+url.PathEscape(commentID)+"/replies/"+url.PathEscape(replyID)+"", nil, map[string]string{"Accept": "application/json"}, nil)
}
