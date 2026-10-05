package remote_actions

import (
	"context"
	"encoding/base64"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
)

type Service struct {
	graphql *graphql.Service
	client  interfaces.HTTPClient
}

func NewService(
	c interfaces.HTTPClient,
) *Service {
	return &Service{graphql: graphql.NewService(c), client: c}
}

// Get uses the observed Nexthink management query.
func (s *Service) Get(
	ctx context.Context,
	uuid string,
) (*GetResponse, *interfaces.Response, error) {
	if err := ValidateUUID(uuid); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetResponse](
		ctx,
		s.graphql,
		operationID,
		graphql.GraphQLRequest{
			Query:         queryGet,
			OperationName: "GetRemoteAction",
			Variables:     map[string]any{"uuid": uuid},
		},
	)
}

// GetForView uses the observed Nexthink management query.
func (s *Service) GetForView(
	ctx context.Context,
	uuid string,
) (*ViewResponse, *interfaces.Response, error) {
	if err := ValidateUUID(uuid); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[ViewResponse](
		ctx,
		s.graphql,
		operationID,
		graphql.GraphQLRequest{
			Query:         queryGetForView,
			OperationName: "GetRemoteActionByUIdForView",
			Variables:     map[string]any{"uuid": uuid},
		},
	)
}

// GetContentVolume uses the observed Nexthink management query.
func (s *Service) GetContentVolume(
	ctx context.Context,
) (*VolumeResponse, *interfaces.Response, error) {
	return graphql.ExecuteData[VolumeResponse](
		ctx,
		s.graphql,
		operationID,
		graphql.GraphQLRequest{
			Query:         queryGetContentVolume,
			OperationName: "GetContentVolumeDetails",
		},
	)
}

// InspectBashScript inspects supplied bytes without executing the script.
// Supply a macOS tar.gz script archive. The SDK base64-encodes the bytes.
func (s *Service) InspectBashScript(
	ctx context.Context,
	script []byte,
) (*BashInspectionResponse, *interfaces.Response, error) {
	if err := ValidateScript(script); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[BashInspectionResponse](
		ctx,
		s.graphql,
		operationID,
		graphql.GraphQLRequest{
			Query:         queryInspectBashScript,
			OperationName: "GetInputAndOutputFromBashScript",
			Variables:     map[string]any{"script": base64.StdEncoding.EncodeToString(script)},
		},
	)
}

// InspectPowerShellScript inspects supplied bytes without executing the script.
// Supply UTF-8 PowerShell source including its BOM. The SDK base64-encodes the bytes.
func (s *Service) InspectPowerShellScript(
	ctx context.Context,
	script []byte,
) (*PowerShellInspectionResponse, *interfaces.Response, error) {
	if err := ValidateScript(script); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[PowerShellInspectionResponse](
		ctx,
		s.graphql,
		operationID,
		graphql.GraphQLRequest{
			Query:         queryInspectPowerShellScript,
			OperationName: "GetInputAndOutputFromPowershellScript",
			Variables:     map[string]any{"script": base64.StdEncoding.EncodeToString(script)},
		},
	)
}

// GetPowerShellSignature inspects supplied bytes without executing the script.
// Supply UTF-8 PowerShell source including its BOM. The SDK base64-encodes the bytes.
func (s *Service) GetPowerShellSignature(
	ctx context.Context,
	script []byte,
) (*SignatureResponse, *interfaces.Response, error) {
	if err := ValidateScript(script); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[SignatureResponse](
		ctx,
		s.graphql,
		operationID,
		graphql.GraphQLRequest{
			Query:         queryGetPowerShellSignature,
			OperationName: "GetPowershellScriptSignatureInfo",
			Variables:     map[string]any{"script": base64.StdEncoding.EncodeToString(script)},
		},
	)
}

type RemoteActionsServiceInterface interface {
	List(context.Context) (*ListResponse, *interfaces.Response, error)
	Delete(ctx context.Context, id string) (*DeleteResponse, *interfaces.Response, error)
	Update(
		ctx context.Context,
		request *RemoteActionInput,
	) (*UpdateResponse, *interfaces.Response, error)
	Create(
		ctx context.Context,
		request *RemoteActionInput,
	) (*CreateResponse, *interfaces.Response, error)
	Get(ctx context.Context, uuid string) (*GetResponse, *interfaces.Response, error)
	GetForView(ctx context.Context, uuid string) (*ViewResponse, *interfaces.Response, error)
	GetContentVolume(ctx context.Context) (*VolumeResponse, *interfaces.Response, error)
	InspectBashScript(
		ctx context.Context,
		script []byte,
	) (*BashInspectionResponse, *interfaces.Response, error)
	InspectPowerShellScript(
		ctx context.Context,
		script []byte,
	) (*PowerShellInspectionResponse, *interfaces.Response, error)
	GetPowerShellSignature(
		ctx context.Context,
		script []byte,
	) (*SignatureResponse, *interfaces.Response, error)
}

var _ RemoteActionsServiceInterface = (*Service)(nil)

// Create calls the UI management mutation. Inspect partial results even when
// err is non-nil; GraphQL can return both data and errors.
func (s *Service) Create(
	ctx context.Context,
	request *RemoteActionInput,
) (*CreateResponse, *interfaces.Response, error) {
	if err := ValidateCreateRequest(request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[CreateResponse](
		ctx,
		s.graphql,
		operationID,
		graphql.GraphQLRequest{
			Query:         queryCreate,
			OperationName: "CreateRemoteAction",
			Variables:     map[string]any{"remoteAction": request},
		},
	)
}

// Update calls the UI management mutation. Inspect partial results even when
// err is non-nil; GraphQL can return both data and errors.
func (s *Service) Update(
	ctx context.Context,
	request *RemoteActionInput,
) (*UpdateResponse, *interfaces.Response, error) {
	if err := ValidateUpdateRequest(request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[UpdateResponse](
		ctx,
		s.graphql,
		operationID,
		graphql.GraphQLRequest{
			Query:         queryUpdate,
			OperationName: "UpdateRemoteAction",
			Variables:     map[string]any{"remoteActionUpdateInput": request},
		},
	)
}

// Delete calls the UI management mutation. Inspect partial results even when
// err is non-nil; GraphQL can return both data and errors.
func (s *Service) Delete(
	ctx context.Context,
	id string,
) (*DeleteResponse, *interfaces.Response, error) {
	if err := ValidateNQLID(id); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[DeleteResponse](
		ctx,
		s.graphql,
		operationID,
		graphql.GraphQLRequest{
			Query:         queryDelete,
			OperationName: "DeleteRemoteAction",
			Variables:     map[string]any{"id": id},
		},
	)
}

// List uses the same content-administration listing as the Remote Actions UI.
// Each row supplies contentId for Get and nqlId for Update/Delete.
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	response, err := s.client.Get(
		ctx,
		EndpointList,
		nil,
		map[string]string{"Accept": "application/json"},
		&result,
	)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
