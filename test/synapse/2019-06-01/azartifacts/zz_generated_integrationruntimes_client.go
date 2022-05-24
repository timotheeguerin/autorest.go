//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

type integrationRuntimesClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newIntegrationRuntimesClient creates a new instance of integrationRuntimesClient with the specified values.
// endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
// pl - the pipeline used for sending requests and handling responses.
func newIntegrationRuntimesClient(endpoint string, pl runtime.Pipeline) *integrationRuntimesClient {
	client := &integrationRuntimesClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// Get - Get Integration Runtime
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// integrationRuntimeName - The Integration Runtime name
// options - integrationRuntimesClientGetOptions contains the optional parameters for the integrationRuntimesClient.Get method.
func (client *integrationRuntimesClient) Get(ctx context.Context, integrationRuntimeName string, options *integrationRuntimesClientGetOptions) (IntegrationRuntimesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, integrationRuntimeName, options)
	if err != nil {
		return IntegrationRuntimesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationRuntimesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationRuntimesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *integrationRuntimesClient) getCreateRequest(ctx context.Context, integrationRuntimeName string, options *integrationRuntimesClientGetOptions) (*policy.Request, error) {
	urlPath := "/integrationRuntimes/{integrationRuntimeName}"
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *integrationRuntimesClient) getHandleResponse(resp *http.Response) (IntegrationRuntimesClientGetResponse, error) {
	result := IntegrationRuntimesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationRuntimeResource); err != nil {
		return IntegrationRuntimesClientGetResponse{}, err
	}
	return result, nil
}

// List - List Integration Runtimes
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// options - integrationRuntimesClientListOptions contains the optional parameters for the integrationRuntimesClient.List
// method.
func (client *integrationRuntimesClient) List(ctx context.Context, options *integrationRuntimesClientListOptions) (IntegrationRuntimesClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return IntegrationRuntimesClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationRuntimesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationRuntimesClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *integrationRuntimesClient) listCreateRequest(ctx context.Context, options *integrationRuntimesClientListOptions) (*policy.Request, error) {
	urlPath := "/integrationRuntimes"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *integrationRuntimesClient) listHandleResponse(resp *http.Response) (IntegrationRuntimesClientListResponse, error) {
	result := IntegrationRuntimesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationRuntimeListResponse); err != nil {
		return IntegrationRuntimesClientListResponse{}, err
	}
	return result, nil
}
