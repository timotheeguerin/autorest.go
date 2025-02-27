//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WorkspaceConnectionsClient contains the methods for the WorkspaceConnections group.
// Don't use this type directly, use NewWorkspaceConnectionsClient() instead.
type WorkspaceConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceConnectionsClient creates a new instance of WorkspaceConnectionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Add a new workspace connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - connectionName - Friendly name of the workspace connection
//   - parameters - The object for creating or updating a new workspace connection
//   - options - WorkspaceConnectionsClientCreateOptions contains the optional parameters for the WorkspaceConnectionsClient.Create
//     method.
func (client *WorkspaceConnectionsClient) Create(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, parameters WorkspaceConnection, options *WorkspaceConnectionsClientCreateOptions) (WorkspaceConnectionsClientCreateResponse, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, workspaceName, connectionName, parameters, options)
	if err != nil {
		return WorkspaceConnectionsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceConnectionsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceConnectionsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *WorkspaceConnectionsClient) createCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, parameters WorkspaceConnection, options *WorkspaceConnectionsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/connections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *WorkspaceConnectionsClient) createHandleResponse(resp *http.Response) (WorkspaceConnectionsClientCreateResponse, error) {
	result := WorkspaceConnectionsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceConnection); err != nil {
		return WorkspaceConnectionsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a workspace connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - connectionName - Friendly name of the workspace connection
//   - options - WorkspaceConnectionsClientDeleteOptions contains the optional parameters for the WorkspaceConnectionsClient.Delete
//     method.
func (client *WorkspaceConnectionsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, options *WorkspaceConnectionsClientDeleteOptions) (WorkspaceConnectionsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, connectionName, options)
	if err != nil {
		return WorkspaceConnectionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceConnectionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceConnectionsClientDeleteResponse{}, err
	}
	return WorkspaceConnectionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, options *WorkspaceConnectionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/connections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the detail of a workspace connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - connectionName - Friendly name of the workspace connection
//   - options - WorkspaceConnectionsClientGetOptions contains the optional parameters for the WorkspaceConnectionsClient.Get
//     method.
func (client *WorkspaceConnectionsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, options *WorkspaceConnectionsClientGetOptions) (WorkspaceConnectionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, connectionName, options)
	if err != nil {
		return WorkspaceConnectionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceConnectionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkspaceConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, options *WorkspaceConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/connections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceConnectionsClient) getHandleResponse(resp *http.Response) (WorkspaceConnectionsClientGetResponse, error) {
	result := WorkspaceConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceConnection); err != nil {
		return WorkspaceConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all connections under a AML workspace.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - options - WorkspaceConnectionsClientListOptions contains the optional parameters for the WorkspaceConnectionsClient.NewListPager
//     method.
func (client *WorkspaceConnectionsClient) NewListPager(resourceGroupName string, workspaceName string, options *WorkspaceConnectionsClientListOptions) *runtime.Pager[WorkspaceConnectionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceConnectionsClientListResponse]{
		More: func(page WorkspaceConnectionsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *WorkspaceConnectionsClientListResponse) (WorkspaceConnectionsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			if err != nil {
				return WorkspaceConnectionsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return WorkspaceConnectionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspaceConnectionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WorkspaceConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/connections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	if options != nil && options.Target != nil {
		reqQP.Set("target", *options.Target)
	}
	if options != nil && options.Category != nil {
		reqQP.Set("category", *options.Category)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkspaceConnectionsClient) listHandleResponse(resp *http.Response) (WorkspaceConnectionsClientListResponse, error) {
	result := WorkspaceConnectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PaginatedWorkspaceConnectionsList); err != nil {
		return WorkspaceConnectionsClientListResponse{}, err
	}
	return result, nil
}
