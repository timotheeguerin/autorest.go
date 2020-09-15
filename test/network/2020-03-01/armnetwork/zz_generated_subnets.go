// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// SubnetsOperations contains the methods for the Subnets group.
type SubnetsOperations interface {
	// BeginCreateOrUpdate - Creates or updates a subnet in the specified virtual network.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet) (*SubnetPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (SubnetPoller, error)
	// BeginDelete - Deletes the specified subnet.
	BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified subnet by virtual network and resource group.
	Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetsGetOptions *SubnetsGetOptions) (*SubnetResponse, error)
	// List - Gets all subnets in a virtual network.
	List(resourceGroupName string, virtualNetworkName string) SubnetListResultPager
	// BeginPrepareNetworkPolicies - Prepares a subnet by applying network intent policies.
	BeginPrepareNetworkPolicies(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, prepareNetworkPoliciesRequestParameters PrepareNetworkPoliciesRequest) (*HTTPPollerResponse, error)
	// ResumePrepareNetworkPolicies - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePrepareNetworkPolicies(token string) (HTTPPoller, error)
	// BeginUnprepareNetworkPolicies - Unprepares a subnet by removing network intent policies.
	BeginUnprepareNetworkPolicies(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, unprepareNetworkPoliciesRequestParameters UnprepareNetworkPoliciesRequest) (*HTTPPollerResponse, error)
	// ResumeUnprepareNetworkPolicies - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeUnprepareNetworkPolicies(token string) (HTTPPoller, error)
}

// SubnetsClient implements the SubnetsOperations interface.
// Don't use this type directly, use NewSubnetsClient() instead.
type SubnetsClient struct {
	*Client
	subscriptionID string
}

// NewSubnetsClient creates a new instance of SubnetsClient with the specified values.
func NewSubnetsClient(c *Client, subscriptionID string) SubnetsOperations {
	return &SubnetsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *SubnetsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// CreateOrUpdate - Creates or updates a subnet in the specified virtual network.
func (client *SubnetsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet) (*SubnetPollerResponse, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, subnetParameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	result, err := client.CreateOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("SubnetsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &subnetPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*SubnetResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *SubnetsClient) ResumeCreateOrUpdate(token string) (SubnetPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("SubnetsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &subnetPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SubnetsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(subnetParameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SubnetsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*SubnetPollerResponse, error) {
	return &SubnetPollerResponse{RawResponse: resp.Response}, nil
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SubnetsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified subnet.
func (client *SubnetsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (*HTTPPollerResponse, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	result, err := client.DeleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("SubnetsClient.Delete", "location", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *SubnetsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("SubnetsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *SubnetsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleResponse handles the Delete response.
func (client *SubnetsClient) DeleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// DeleteHandleError handles the Delete error response.
func (client *SubnetsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified subnet by virtual network and resource group.
func (client *SubnetsClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetsGetOptions *SubnetsGetOptions) (*SubnetResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, subnetsGetOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *SubnetsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetsGetOptions *SubnetsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	if subnetsGetOptions != nil && subnetsGetOptions.Expand != nil {
		query.Set("$expand", *subnetsGetOptions.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *SubnetsClient) GetHandleResponse(resp *azcore.Response) (*SubnetResponse, error) {
	result := SubnetResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Subnet)
}

// GetHandleError handles the Get error response.
func (client *SubnetsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all subnets in a virtual network.
func (client *SubnetsClient) List(resourceGroupName string, virtualNetworkName string) SubnetListResultPager {
	return &subnetListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, virtualNetworkName)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *SubnetListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SubnetListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *SubnetsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
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
func (client *SubnetsClient) ListHandleResponse(resp *azcore.Response) (*SubnetListResultResponse, error) {
	result := SubnetListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SubnetListResult)
}

// ListHandleError handles the List error response.
func (client *SubnetsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PrepareNetworkPolicies - Prepares a subnet by applying network intent policies.
func (client *SubnetsClient) BeginPrepareNetworkPolicies(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, prepareNetworkPoliciesRequestParameters PrepareNetworkPoliciesRequest) (*HTTPPollerResponse, error) {
	req, err := client.PrepareNetworkPoliciesCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, prepareNetworkPoliciesRequestParameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.PrepareNetworkPoliciesHandleError(resp)
	}
	result, err := client.PrepareNetworkPoliciesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("SubnetsClient.PrepareNetworkPolicies", "location", resp, client.PrepareNetworkPoliciesHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *SubnetsClient) ResumePrepareNetworkPolicies(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("SubnetsClient.PrepareNetworkPolicies", token, client.PrepareNetworkPoliciesHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// PrepareNetworkPoliciesCreateRequest creates the PrepareNetworkPolicies request.
func (client *SubnetsClient) PrepareNetworkPoliciesCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, prepareNetworkPoliciesRequestParameters PrepareNetworkPoliciesRequest) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/PrepareNetworkPolicies"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(prepareNetworkPoliciesRequestParameters)
}

// PrepareNetworkPoliciesHandleResponse handles the PrepareNetworkPolicies response.
func (client *SubnetsClient) PrepareNetworkPoliciesHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// PrepareNetworkPoliciesHandleError handles the PrepareNetworkPolicies error response.
func (client *SubnetsClient) PrepareNetworkPoliciesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UnprepareNetworkPolicies - Unprepares a subnet by removing network intent policies.
func (client *SubnetsClient) BeginUnprepareNetworkPolicies(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, unprepareNetworkPoliciesRequestParameters UnprepareNetworkPoliciesRequest) (*HTTPPollerResponse, error) {
	req, err := client.UnprepareNetworkPoliciesCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, unprepareNetworkPoliciesRequestParameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.UnprepareNetworkPoliciesHandleError(resp)
	}
	result, err := client.UnprepareNetworkPoliciesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("SubnetsClient.UnprepareNetworkPolicies", "location", resp, client.UnprepareNetworkPoliciesHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *SubnetsClient) ResumeUnprepareNetworkPolicies(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("SubnetsClient.UnprepareNetworkPolicies", token, client.UnprepareNetworkPoliciesHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// UnprepareNetworkPoliciesCreateRequest creates the UnprepareNetworkPolicies request.
func (client *SubnetsClient) UnprepareNetworkPoliciesCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, unprepareNetworkPoliciesRequestParameters UnprepareNetworkPoliciesRequest) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/UnprepareNetworkPolicies"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(unprepareNetworkPoliciesRequestParameters)
}

// UnprepareNetworkPoliciesHandleResponse handles the UnprepareNetworkPolicies response.
func (client *SubnetsClient) UnprepareNetworkPoliciesHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// UnprepareNetworkPoliciesHandleError handles the UnprepareNetworkPolicies error response.
func (client *SubnetsClient) UnprepareNetworkPoliciesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
