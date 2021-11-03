//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ContainerServicesClient contains the methods for the ContainerServices group.
// Don't use this type directly, use NewContainerServicesClient() instead.
type ContainerServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewContainerServicesClient creates a new instance of ContainerServicesClient with the specified values.
func NewContainerServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ContainerServicesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &ContainerServicesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a container service with the specified configuration of orchestrator, masters, and agents.
// If the operation fails it returns a generic error.
func (client *ContainerServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, containerServiceName string, parameters ContainerService, options *ContainerServicesBeginCreateOrUpdateOptions) (ContainerServicesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, containerServiceName, parameters, options)
	if err != nil {
		return ContainerServicesCreateOrUpdatePollerResponse{}, err
	}
	result := ContainerServicesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ContainerServicesClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ContainerServicesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ContainerServicesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a container service with the specified configuration of orchestrator, masters, and agents.
// If the operation fails it returns a generic error.
func (client *ContainerServicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, containerServiceName string, parameters ContainerService, options *ContainerServicesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, containerServiceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ContainerServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, containerServiceName string, parameters ContainerService, options *ContainerServicesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerServiceName == "" {
		return nil, errors.New("parameter containerServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerServiceName}", url.PathEscape(containerServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-01-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ContainerServicesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Deletes the specified container service in the specified subscription and resource group. The operation does not delete other resources
// created as part of creating a container service, including
// storage accounts, VMs, and availability sets. All the other resources created with the container service are part of the same resource group and can
// be deleted individually.
// If the operation fails it returns a generic error.
func (client *ContainerServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesBeginDeleteOptions) (ContainerServicesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, containerServiceName, options)
	if err != nil {
		return ContainerServicesDeletePollerResponse{}, err
	}
	result := ContainerServicesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ContainerServicesClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ContainerServicesDeletePollerResponse{}, err
	}
	result.Poller = &ContainerServicesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified container service in the specified subscription and resource group. The operation does not delete other resources created
// as part of creating a container service, including
// storage accounts, VMs, and availability sets. All the other resources created with the container service are part of the same resource group and can
// be deleted individually.
// If the operation fails it returns a generic error.
func (client *ContainerServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, containerServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ContainerServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerServiceName == "" {
		return nil, errors.New("parameter containerServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerServiceName}", url.PathEscape(containerServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-01-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ContainerServicesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets the properties of the specified container service in the specified subscription and resource group. The operation returns the properties including
// state, orchestrator, number of masters and
// agents, and FQDNs of masters and agents.
// If the operation fails it returns a generic error.
func (client *ContainerServicesClient) Get(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesGetOptions) (ContainerServicesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, containerServiceName, options)
	if err != nil {
		return ContainerServicesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerServicesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainerServicesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ContainerServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerServiceName == "" {
		return nil, errors.New("parameter containerServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerServiceName}", url.PathEscape(containerServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-01-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContainerServicesClient) getHandleResponse(resp *http.Response) (ContainerServicesGetResponse, error) {
	result := ContainerServicesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerService); err != nil {
		return ContainerServicesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ContainerServicesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Gets a list of container services in the specified subscription. The operation returns properties of each container service including state, orchestrator,
// number of masters and agents, and FQDNs of
// masters and agents.
// If the operation fails it returns a generic error.
func (client *ContainerServicesClient) List(options *ContainerServicesListOptions) *ContainerServicesListPager {
	return &ContainerServicesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ContainerServicesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ContainerServiceListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ContainerServicesClient) listCreateRequest(ctx context.Context, options *ContainerServicesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/containerServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-01-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ContainerServicesClient) listHandleResponse(resp *http.Response) (ContainerServicesListResponse, error) {
	result := ContainerServicesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerServiceListResult); err != nil {
		return ContainerServicesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ContainerServicesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByResourceGroup - Gets a list of container services in the specified subscription and resource group. The operation returns properties of each container
// service including state, orchestrator, number of masters and
// agents, and FQDNs of masters and agents.
// If the operation fails it returns a generic error.
func (client *ContainerServicesClient) ListByResourceGroup(resourceGroupName string, options *ContainerServicesListByResourceGroupOptions) *ContainerServicesListByResourceGroupPager {
	return &ContainerServicesListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ContainerServicesListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ContainerServiceListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ContainerServicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ContainerServicesListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-01-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ContainerServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (ContainerServicesListByResourceGroupResponse, error) {
	result := ContainerServicesListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerServiceListResult); err != nil {
		return ContainerServicesListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ContainerServicesClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
