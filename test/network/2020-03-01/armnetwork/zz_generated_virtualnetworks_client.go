//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualNetworksClient contains the methods for the VirtualNetworks group.
// Don't use this type directly, use NewVirtualNetworksClient() instead.
type VirtualNetworksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualNetworksClient creates a new instance of VirtualNetworksClient with the specified values.
func NewVirtualNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualNetworksClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &VirtualNetworksClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// CheckIPAddressAvailability - Checks whether a private IP address is available for use.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworksClient) CheckIPAddressAvailability(ctx context.Context, resourceGroupName string, virtualNetworkName string, ipAddress string, options *VirtualNetworksCheckIPAddressAvailabilityOptions) (VirtualNetworksCheckIPAddressAvailabilityResponse, error) {
	req, err := client.checkIPAddressAvailabilityCreateRequest(ctx, resourceGroupName, virtualNetworkName, ipAddress, options)
	if err != nil {
		return VirtualNetworksCheckIPAddressAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualNetworksCheckIPAddressAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworksCheckIPAddressAvailabilityResponse{}, client.checkIPAddressAvailabilityHandleError(resp)
	}
	return client.checkIPAddressAvailabilityHandleResponse(resp)
}

// checkIPAddressAvailabilityCreateRequest creates the CheckIPAddressAvailability request.
func (client *VirtualNetworksClient) checkIPAddressAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, ipAddress string, options *VirtualNetworksCheckIPAddressAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/CheckIPAddressAvailability"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("ipAddress", ipAddress)
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// checkIPAddressAvailabilityHandleResponse handles the CheckIPAddressAvailability response.
func (client *VirtualNetworksClient) checkIPAddressAvailabilityHandleResponse(resp *http.Response) (VirtualNetworksCheckIPAddressAvailabilityResponse, error) {
	result := VirtualNetworksCheckIPAddressAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPAddressAvailabilityResult); err != nil {
		return VirtualNetworksCheckIPAddressAvailabilityResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// checkIPAddressAvailabilityHandleError handles the CheckIPAddressAvailability error response.
func (client *VirtualNetworksClient) checkIPAddressAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginCreateOrUpdate - Creates or updates a virtual network in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork, options *VirtualNetworksBeginCreateOrUpdateOptions) (VirtualNetworksCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualNetworkName, parameters, options)
	if err != nil {
		return VirtualNetworksCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualNetworksCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualNetworksClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualNetworksCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualNetworksCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a virtual network in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworksClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork, options *VirtualNetworksBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualNetworkName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork, options *VirtualNetworksBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualNetworksClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the specified virtual network.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksBeginDeleteOptions) (VirtualNetworksDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, virtualNetworkName, options)
	if err != nil {
		return VirtualNetworksDeletePollerResponse{}, err
	}
	result := VirtualNetworksDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualNetworksClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return VirtualNetworksDeletePollerResponse{}, err
	}
	result.Poller = &VirtualNetworksDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified virtual network.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworksClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualNetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VirtualNetworksClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the specified virtual network by resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworksClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksGetOptions) (VirtualNetworksGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
	if err != nil {
		return VirtualNetworksGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualNetworksGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworksGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworksClient) getHandleResponse(resp *http.Response) (VirtualNetworksGetResponse, error) {
	result := VirtualNetworksGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetwork); err != nil {
		return VirtualNetworksGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualNetworksClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets all virtual networks in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworksClient) List(resourceGroupName string, options *VirtualNetworksListOptions) *VirtualNetworksListPager {
	return &VirtualNetworksListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp VirtualNetworksListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualNetworksClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualNetworksListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks"
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
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualNetworksClient) listHandleResponse(resp *http.Response) (VirtualNetworksListResponse, error) {
	result := VirtualNetworksListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkListResult); err != nil {
		return VirtualNetworksListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *VirtualNetworksClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListAll - Gets all virtual networks in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworksClient) ListAll(options *VirtualNetworksListAllOptions) *VirtualNetworksListAllPager {
	return &VirtualNetworksListAllPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp VirtualNetworksListAllResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *VirtualNetworksClient) listAllCreateRequest(ctx context.Context, options *VirtualNetworksListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *VirtualNetworksClient) listAllHandleResponse(resp *http.Response) (VirtualNetworksListAllResponse, error) {
	result := VirtualNetworksListAllResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkListResult); err != nil {
		return VirtualNetworksListAllResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listAllHandleError handles the ListAll error response.
func (client *VirtualNetworksClient) listAllHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListUsage - Lists usage stats.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworksClient) ListUsage(resourceGroupName string, virtualNetworkName string, options *VirtualNetworksListUsageOptions) *VirtualNetworksListUsagePager {
	return &VirtualNetworksListUsagePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listUsageCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
		},
		advancer: func(ctx context.Context, resp VirtualNetworksListUsageResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkListUsageResult.NextLink)
		},
	}
}

// listUsageCreateRequest creates the ListUsage request.
func (client *VirtualNetworksClient) listUsageCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksListUsageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/usages"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listUsageHandleResponse handles the ListUsage response.
func (client *VirtualNetworksClient) listUsageHandleResponse(resp *http.Response) (VirtualNetworksListUsageResponse, error) {
	result := VirtualNetworksListUsageResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkListUsageResult); err != nil {
		return VirtualNetworksListUsageResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listUsageHandleError handles the ListUsage error response.
func (client *VirtualNetworksClient) listUsageHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// UpdateTags - Updates a virtual network tags.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworksClient) UpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters TagsObject, options *VirtualNetworksUpdateTagsOptions) (VirtualNetworksUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, virtualNetworkName, parameters, options)
	if err != nil {
		return VirtualNetworksUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualNetworksUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworksUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VirtualNetworksClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters TagsObject, options *VirtualNetworksUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *VirtualNetworksClient) updateTagsHandleResponse(resp *http.Response) (VirtualNetworksUpdateTagsResponse, error) {
	result := VirtualNetworksUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetwork); err != nil {
		return VirtualNetworksUpdateTagsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *VirtualNetworksClient) updateTagsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
