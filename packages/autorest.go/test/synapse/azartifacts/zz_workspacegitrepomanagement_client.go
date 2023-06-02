//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// WorkspaceGitRepoManagementClient contains the methods for the WorkspaceGitRepoManagement group.
// Don't use this type directly, use a constructor function instead.
type WorkspaceGitRepoManagementClient struct {
	internal *azcore.Client
	endpoint string
}

// GetGitHubAccessToken - Get the GitHub access token.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - options - WorkspaceGitRepoManagementClientGetGitHubAccessTokenOptions contains the optional parameters for the WorkspaceGitRepoManagementClient.GetGitHubAccessToken
//     method.
func (client *WorkspaceGitRepoManagementClient) GetGitHubAccessToken(ctx context.Context, gitHubAccessTokenRequest GitHubAccessTokenRequest, options *WorkspaceGitRepoManagementClientGetGitHubAccessTokenOptions) (WorkspaceGitRepoManagementClientGetGitHubAccessTokenResponse, error) {
	var err error
	req, err := client.getGitHubAccessTokenCreateRequest(ctx, gitHubAccessTokenRequest, options)
	if err != nil {
		return WorkspaceGitRepoManagementClientGetGitHubAccessTokenResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceGitRepoManagementClientGetGitHubAccessTokenResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceGitRepoManagementClientGetGitHubAccessTokenResponse{}, err
	}
	resp, err := client.getGitHubAccessTokenHandleResponse(httpResp)
	return resp, err
}

// getGitHubAccessTokenCreateRequest creates the GetGitHubAccessToken request.
func (client *WorkspaceGitRepoManagementClient) getGitHubAccessTokenCreateRequest(ctx context.Context, gitHubAccessTokenRequest GitHubAccessTokenRequest, options *WorkspaceGitRepoManagementClientGetGitHubAccessTokenOptions) (*policy.Request, error) {
	urlPath := "/getGitHubAccessToken"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.ClientRequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, gitHubAccessTokenRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// getGitHubAccessTokenHandleResponse handles the GetGitHubAccessToken response.
func (client *WorkspaceGitRepoManagementClient) getGitHubAccessTokenHandleResponse(resp *http.Response) (WorkspaceGitRepoManagementClientGetGitHubAccessTokenResponse, error) {
	result := WorkspaceGitRepoManagementClientGetGitHubAccessTokenResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitHubAccessTokenResponse); err != nil {
		return WorkspaceGitRepoManagementClientGetGitHubAccessTokenResponse{}, err
	}
	return result, nil
}
