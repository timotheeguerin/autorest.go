// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// AvailableServiceAliasesOperations contains the methods for the AvailableServiceAliases group.
type AvailableServiceAliasesOperations interface {
	// List - Gets all available service aliases for this subscription in this region.
	List(location string, options *AvailableServiceAliasesListOptions) AvailableServiceAliasesResultPager
	// ListByResourceGroup - Gets all available service aliases for this resource group in this region.
	ListByResourceGroup(resourceGroupName string, location string, options *AvailableServiceAliasesListByResourceGroupOptions) AvailableServiceAliasesResultPager
}

// AvailableServiceAliasesClient implements the AvailableServiceAliasesOperations interface.
// Don't use this type directly, use NewAvailableServiceAliasesClient() instead.
type AvailableServiceAliasesClient struct {
	*Client
	subscriptionID string
}

// NewAvailableServiceAliasesClient creates a new instance of AvailableServiceAliasesClient with the specified values.
func NewAvailableServiceAliasesClient(c *Client, subscriptionID string) AvailableServiceAliasesOperations {
	return &AvailableServiceAliasesClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *AvailableServiceAliasesClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// List - Gets all available service aliases for this subscription in this region.
func (client *AvailableServiceAliasesClient) List(location string, options *AvailableServiceAliasesListOptions) AvailableServiceAliasesResultPager {
	return &availableServiceAliasesResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, location, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *AvailableServiceAliasesResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AvailableServiceAliasesResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *AvailableServiceAliasesClient) ListCreateRequest(ctx context.Context, location string, options *AvailableServiceAliasesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availableServiceAliases"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *AvailableServiceAliasesClient) ListHandleResponse(resp *azcore.Response) (*AvailableServiceAliasesResultResponse, error) {
	result := AvailableServiceAliasesResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailableServiceAliasesResult)
}

// ListHandleError handles the List error response.
func (client *AvailableServiceAliasesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Gets all available service aliases for this resource group in this region.
func (client *AvailableServiceAliasesClient) ListByResourceGroup(resourceGroupName string, location string, options *AvailableServiceAliasesListByResourceGroupOptions) AvailableServiceAliasesResultPager {
	return &availableServiceAliasesResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByResourceGroupCreateRequest(ctx, resourceGroupName, location, options)
		},
		responder: client.ListByResourceGroupHandleResponse,
		errorer:   client.ListByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *AvailableServiceAliasesResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AvailableServiceAliasesResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AvailableServiceAliasesClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, location string, options *AvailableServiceAliasesListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availableServiceAliases"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AvailableServiceAliasesClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*AvailableServiceAliasesResultResponse, error) {
	result := AvailableServiceAliasesResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailableServiceAliasesResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *AvailableServiceAliasesClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
