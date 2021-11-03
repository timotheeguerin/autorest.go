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

// NetworkInterfaceIPConfigurationsClient contains the methods for the NetworkInterfaceIPConfigurations group.
// Don't use this type directly, use NewNetworkInterfaceIPConfigurationsClient() instead.
type NetworkInterfaceIPConfigurationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewNetworkInterfaceIPConfigurationsClient creates a new instance of NetworkInterfaceIPConfigurationsClient with the specified values.
func NewNetworkInterfaceIPConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *NetworkInterfaceIPConfigurationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &NetworkInterfaceIPConfigurationsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// Get - Gets the specified network interface ip configuration.
// If the operation fails it returns the *CloudError error type.
func (client *NetworkInterfaceIPConfigurationsClient) Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, ipConfigurationName string, options *NetworkInterfaceIPConfigurationsGetOptions) (NetworkInterfaceIPConfigurationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkInterfaceName, ipConfigurationName, options)
	if err != nil {
		return NetworkInterfaceIPConfigurationsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NetworkInterfaceIPConfigurationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetworkInterfaceIPConfigurationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NetworkInterfaceIPConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, ipConfigurationName string, options *NetworkInterfaceIPConfigurationsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkInterfaceName == "" {
		return nil, errors.New("parameter networkInterfaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	if ipConfigurationName == "" {
		return nil, errors.New("parameter ipConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipConfigurationName}", url.PathEscape(ipConfigurationName))
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

// getHandleResponse handles the Get response.
func (client *NetworkInterfaceIPConfigurationsClient) getHandleResponse(resp *http.Response) (NetworkInterfaceIPConfigurationsGetResponse, error) {
	result := NetworkInterfaceIPConfigurationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkInterfaceIPConfiguration); err != nil {
		return NetworkInterfaceIPConfigurationsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *NetworkInterfaceIPConfigurationsClient) getHandleError(resp *http.Response) error {
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

// List - Get all ip configurations in a network interface.
// If the operation fails it returns the *CloudError error type.
func (client *NetworkInterfaceIPConfigurationsClient) List(resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceIPConfigurationsListOptions) *NetworkInterfaceIPConfigurationsListPager {
	return &NetworkInterfaceIPConfigurationsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, networkInterfaceName, options)
		},
		advancer: func(ctx context.Context, resp NetworkInterfaceIPConfigurationsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.NetworkInterfaceIPConfigurationListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *NetworkInterfaceIPConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceIPConfigurationsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkInterfaceName == "" {
		return nil, errors.New("parameter networkInterfaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
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
func (client *NetworkInterfaceIPConfigurationsClient) listHandleResponse(resp *http.Response) (NetworkInterfaceIPConfigurationsListResponse, error) {
	result := NetworkInterfaceIPConfigurationsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkInterfaceIPConfigurationListResult); err != nil {
		return NetworkInterfaceIPConfigurationsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *NetworkInterfaceIPConfigurationsClient) listHandleError(resp *http.Response) error {
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
