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

// VpnGatewaysOperations contains the methods for the VpnGateways group.
type VpnGatewaysOperations interface {
	// BeginCreateOrUpdate - Creates a virtual wan vpn gateway if it doesn't exist else updates the existing gateway.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters VpnGateway, options *VpnGatewaysCreateOrUpdateOptions) (*VpnGatewayPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VpnGatewayPoller, error)
	// BeginDelete - Deletes a virtual wan vpn gateway.
	BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, options *VpnGatewaysDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Retrieves the details of a virtual wan vpn gateway.
	Get(ctx context.Context, resourceGroupName string, gatewayName string, options *VpnGatewaysGetOptions) (*VpnGatewayResponse, error)
	// List - Lists all the VpnGateways in a subscription.
	List(options *VpnGatewaysListOptions) ListVpnGatewaysResultPager
	// ListByResourceGroup - Lists all the VpnGateways in a resource group.
	ListByResourceGroup(resourceGroupName string, options *VpnGatewaysListByResourceGroupOptions) ListVpnGatewaysResultPager
	// BeginReset - Resets the primary of the vpn gateway in the specified resource group.
	BeginReset(ctx context.Context, resourceGroupName string, gatewayName string, options *VpnGatewaysResetOptions) (*VpnGatewayPollerResponse, error)
	// ResumeReset - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeReset(token string) (VpnGatewayPoller, error)
	// UpdateTags - Updates virtual wan vpn gateway tags.
	UpdateTags(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters TagsObject, options *VpnGatewaysUpdateTagsOptions) (*VpnGatewayResponse, error)
}

// VpnGatewaysClient implements the VpnGatewaysOperations interface.
// Don't use this type directly, use NewVpnGatewaysClient() instead.
type VpnGatewaysClient struct {
	*Client
	subscriptionID string
}

// NewVpnGatewaysClient creates a new instance of VpnGatewaysClient with the specified values.
func NewVpnGatewaysClient(c *Client, subscriptionID string) VpnGatewaysOperations {
	return &VpnGatewaysClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *VpnGatewaysClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

func (client *VpnGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters VpnGateway, options *VpnGatewaysCreateOrUpdateOptions) (*VpnGatewayPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
	if err != nil {
		return nil, err
	}
	result := &VpnGatewayPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VpnGatewaysClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &vpnGatewayPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VpnGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VpnGatewaysClient) ResumeCreateOrUpdate(token string) (VpnGatewayPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VpnGatewaysClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &vpnGatewayPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates a virtual wan vpn gateway if it doesn't exist else updates the existing gateway.
func (client *VpnGatewaysClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters VpnGateway, options *VpnGatewaysCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VpnGatewaysClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters VpnGateway, options *VpnGatewaysCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(vpnGatewayParameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VpnGatewaysClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*VpnGatewayResponse, error) {
	result := VpnGatewayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnGateway)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VpnGatewaysClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *VpnGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, options *VpnGatewaysDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VpnGatewaysClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *VpnGatewaysClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VpnGatewaysClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Delete - Deletes a virtual wan vpn gateway.
func (client *VpnGatewaysClient) Delete(ctx context.Context, resourceGroupName string, gatewayName string, options *VpnGatewaysDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *VpnGatewaysClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VpnGatewaysDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
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

// DeleteHandleError handles the Delete error response.
func (client *VpnGatewaysClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves the details of a virtual wan vpn gateway.
func (client *VpnGatewaysClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, options *VpnGatewaysGetOptions) (*VpnGatewayResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, gatewayName, options)
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
func (client *VpnGatewaysClient) GetCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VpnGatewaysGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
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

// GetHandleResponse handles the Get response.
func (client *VpnGatewaysClient) GetHandleResponse(resp *azcore.Response) (*VpnGatewayResponse, error) {
	result := VpnGatewayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnGateway)
}

// GetHandleError handles the Get error response.
func (client *VpnGatewaysClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all the VpnGateways in a subscription.
func (client *VpnGatewaysClient) List(options *VpnGatewaysListOptions) ListVpnGatewaysResultPager {
	return &listVpnGatewaysResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *ListVpnGatewaysResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVpnGatewaysResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *VpnGatewaysClient) ListCreateRequest(ctx context.Context, options *VpnGatewaysListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnGateways"
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
func (client *VpnGatewaysClient) ListHandleResponse(resp *azcore.Response) (*ListVpnGatewaysResultResponse, error) {
	result := ListVpnGatewaysResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListVpnGatewaysResult)
}

// ListHandleError handles the List error response.
func (client *VpnGatewaysClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Lists all the VpnGateways in a resource group.
func (client *VpnGatewaysClient) ListByResourceGroup(resourceGroupName string, options *VpnGatewaysListByResourceGroupOptions) ListVpnGatewaysResultPager {
	return &listVpnGatewaysResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListByResourceGroupHandleResponse,
		errorer:   client.ListByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *ListVpnGatewaysResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVpnGatewaysResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VpnGatewaysClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VpnGatewaysListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *VpnGatewaysClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*ListVpnGatewaysResultResponse, error) {
	result := ListVpnGatewaysResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListVpnGatewaysResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *VpnGatewaysClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *VpnGatewaysClient) BeginReset(ctx context.Context, resourceGroupName string, gatewayName string, options *VpnGatewaysResetOptions) (*VpnGatewayPollerResponse, error) {
	resp, err := client.Reset(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return nil, err
	}
	result := &VpnGatewayPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VpnGatewaysClient.Reset", "location", resp, client.ResetHandleError)
	if err != nil {
		return nil, err
	}
	poller := &vpnGatewayPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VpnGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VpnGatewaysClient) ResumeReset(token string) (VpnGatewayPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VpnGatewaysClient.Reset", token, client.ResetHandleError)
	if err != nil {
		return nil, err
	}
	return &vpnGatewayPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Reset - Resets the primary of the vpn gateway in the specified resource group.
func (client *VpnGatewaysClient) Reset(ctx context.Context, resourceGroupName string, gatewayName string, options *VpnGatewaysResetOptions) (*azcore.Response, error) {
	req, err := client.ResetCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.ResetHandleError(resp)
	}
	return resp, nil
}

// ResetCreateRequest creates the Reset request.
func (client *VpnGatewaysClient) ResetCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VpnGatewaysResetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/reset"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ResetHandleResponse handles the Reset response.
func (client *VpnGatewaysClient) ResetHandleResponse(resp *azcore.Response) (*VpnGatewayResponse, error) {
	result := VpnGatewayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnGateway)
}

// ResetHandleError handles the Reset error response.
func (client *VpnGatewaysClient) ResetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates virtual wan vpn gateway tags.
func (client *VpnGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters TagsObject, options *VpnGatewaysUpdateTagsOptions) (*VpnGatewayResponse, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateTagsHandleError(resp)
	}
	result, err := client.UpdateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTagsCreateRequest creates the UpdateTags request.
func (client *VpnGatewaysClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters TagsObject, options *VpnGatewaysUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(vpnGatewayParameters)
}

// UpdateTagsHandleResponse handles the UpdateTags response.
func (client *VpnGatewaysClient) UpdateTagsHandleResponse(resp *azcore.Response) (*VpnGatewayResponse, error) {
	result := VpnGatewayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnGateway)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *VpnGatewaysClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
