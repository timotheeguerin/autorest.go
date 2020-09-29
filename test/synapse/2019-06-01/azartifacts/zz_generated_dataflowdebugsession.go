// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

type dataFlowDebugSessionClient struct {
	*client
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *dataFlowDebugSessionClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// AddDataFlow - Add a data flow into debug session.
func (client *dataFlowDebugSessionClient) AddDataFlow(ctx context.Context, request DataFlowDebugPackage, options *DataFlowDebugSessionAddDataFlowOptions) (*AddDataFlowToDebugSessionResponseResponse, error) {
	req, err := client.AddDataFlowCreateRequest(ctx, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.AddDataFlowHandleError(resp)
	}
	result, err := client.AddDataFlowHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AddDataFlowCreateRequest creates the AddDataFlow request.
func (client *dataFlowDebugSessionClient) AddDataFlowCreateRequest(ctx context.Context, request DataFlowDebugPackage, options *DataFlowDebugSessionAddDataFlowOptions) (*azcore.Request, error) {
	urlPath := "/addDataFlowToDebugSession"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// AddDataFlowHandleResponse handles the AddDataFlow response.
func (client *dataFlowDebugSessionClient) AddDataFlowHandleResponse(resp *azcore.Response) (*AddDataFlowToDebugSessionResponseResponse, error) {
	result := AddDataFlowToDebugSessionResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AddDataFlowToDebugSessionResponse)
}

// AddDataFlowHandleError handles the AddDataFlow error response.
func (client *dataFlowDebugSessionClient) AddDataFlowHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// CreateDataFlowDebugSession - Creates a data flow debug session.
func (client *dataFlowDebugSessionClient) CreateDataFlowDebugSession(ctx context.Context, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionCreateDataFlowDebugSessionOptions) (*azcore.Response, error) {
	req, err := client.CreateDataFlowDebugSessionCreateRequest(ctx, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.CreateDataFlowDebugSessionHandleError(resp)
	}
	return resp, nil
}

// CreateDataFlowDebugSessionCreateRequest creates the CreateDataFlowDebugSession request.
func (client *dataFlowDebugSessionClient) CreateDataFlowDebugSessionCreateRequest(ctx context.Context, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionCreateDataFlowDebugSessionOptions) (*azcore.Request, error) {
	urlPath := "/createDataFlowDebugSession"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// CreateDataFlowDebugSessionHandleResponse handles the CreateDataFlowDebugSession response.
func (client *dataFlowDebugSessionClient) CreateDataFlowDebugSessionHandleResponse(resp *azcore.Response) (*CreateDataFlowDebugSessionResponseResponse, error) {
	result := CreateDataFlowDebugSessionResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.CreateDataFlowDebugSessionResponse)
}

// CreateDataFlowDebugSessionHandleError handles the CreateDataFlowDebugSession error response.
func (client *dataFlowDebugSessionClient) CreateDataFlowDebugSessionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteDataFlowDebugSession - Deletes a data flow debug session.
func (client *dataFlowDebugSessionClient) DeleteDataFlowDebugSession(ctx context.Context, request DeleteDataFlowDebugSessionRequest, options *DataFlowDebugSessionDeleteDataFlowDebugSessionOptions) (*http.Response, error) {
	req, err := client.DeleteDataFlowDebugSessionCreateRequest(ctx, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.DeleteDataFlowDebugSessionHandleError(resp)
	}
	return resp.Response, nil
}

// DeleteDataFlowDebugSessionCreateRequest creates the DeleteDataFlowDebugSession request.
func (client *dataFlowDebugSessionClient) DeleteDataFlowDebugSessionCreateRequest(ctx context.Context, request DeleteDataFlowDebugSessionRequest, options *DataFlowDebugSessionDeleteDataFlowDebugSessionOptions) (*azcore.Request, error) {
	urlPath := "/deleteDataFlowDebugSession"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// DeleteDataFlowDebugSessionHandleError handles the DeleteDataFlowDebugSession error response.
func (client *dataFlowDebugSessionClient) DeleteDataFlowDebugSessionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ExecuteCommand - Execute a data flow debug command.
func (client *dataFlowDebugSessionClient) ExecuteCommand(ctx context.Context, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionExecuteCommandOptions) (*azcore.Response, error) {
	req, err := client.ExecuteCommandCreateRequest(ctx, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.ExecuteCommandHandleError(resp)
	}
	return resp, nil
}

// ExecuteCommandCreateRequest creates the ExecuteCommand request.
func (client *dataFlowDebugSessionClient) ExecuteCommandCreateRequest(ctx context.Context, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionExecuteCommandOptions) (*azcore.Request, error) {
	urlPath := "/executeDataFlowDebugCommand"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// ExecuteCommandHandleResponse handles the ExecuteCommand response.
func (client *dataFlowDebugSessionClient) ExecuteCommandHandleResponse(resp *azcore.Response) (*DataFlowDebugCommandResponseResponse, error) {
	result := DataFlowDebugCommandResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DataFlowDebugCommandResponse)
}

// ExecuteCommandHandleError handles the ExecuteCommand error response.
func (client *dataFlowDebugSessionClient) ExecuteCommandHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// QueryDataFlowDebugSessionsByWorkspace - Query all active data flow debug sessions.
func (client *dataFlowDebugSessionClient) QueryDataFlowDebugSessionsByWorkspace(options *DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspaceOptions) QueryDataFlowDebugSessionsResponsePager {
	return &queryDataFlowDebugSessionsResponsePager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.QueryDataFlowDebugSessionsByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.QueryDataFlowDebugSessionsByWorkspaceHandleResponse,
		errorer:   client.QueryDataFlowDebugSessionsByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp *QueryDataFlowDebugSessionsResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.QueryDataFlowDebugSessionsResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// QueryDataFlowDebugSessionsByWorkspaceCreateRequest creates the QueryDataFlowDebugSessionsByWorkspace request.
func (client *dataFlowDebugSessionClient) QueryDataFlowDebugSessionsByWorkspaceCreateRequest(ctx context.Context, options *DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/queryDataFlowDebugSessions"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// QueryDataFlowDebugSessionsByWorkspaceHandleResponse handles the QueryDataFlowDebugSessionsByWorkspace response.
func (client *dataFlowDebugSessionClient) QueryDataFlowDebugSessionsByWorkspaceHandleResponse(resp *azcore.Response) (*QueryDataFlowDebugSessionsResponseResponse, error) {
	result := QueryDataFlowDebugSessionsResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.QueryDataFlowDebugSessionsResponse)
}

// QueryDataFlowDebugSessionsByWorkspaceHandleError handles the QueryDataFlowDebugSessionsByWorkspace error response.
func (client *dataFlowDebugSessionClient) QueryDataFlowDebugSessionsByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
