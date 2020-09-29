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

// ServiceEndpointPoliciesOperations contains the methods for the ServiceEndpointPolicies group.
type ServiceEndpointPoliciesOperations interface {
	// BeginCreateOrUpdate - Creates or updates a service Endpoint Policies.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters ServiceEndpointPolicy, options *ServiceEndpointPoliciesCreateOrUpdateOptions) (*ServiceEndpointPolicyPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (ServiceEndpointPolicyPoller, error)
	// BeginDelete - Deletes the specified service endpoint policy.
	BeginDelete(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified service Endpoint Policies in a specified resource group.
	Get(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesGetOptions) (*ServiceEndpointPolicyResponse, error)
	// List - Gets all the service endpoint policies in a subscription.
	List(options *ServiceEndpointPoliciesListOptions) ServiceEndpointPolicyListResultPager
	// ListByResourceGroup - Gets all service endpoint Policies in a resource group.
	ListByResourceGroup(resourceGroupName string, options *ServiceEndpointPoliciesListByResourceGroupOptions) ServiceEndpointPolicyListResultPager
	// UpdateTags - Updates tags of a service endpoint policy.
	UpdateTags(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters TagsObject, options *ServiceEndpointPoliciesUpdateTagsOptions) (*ServiceEndpointPolicyResponse, error)
}

// ServiceEndpointPoliciesClient implements the ServiceEndpointPoliciesOperations interface.
// Don't use this type directly, use NewServiceEndpointPoliciesClient() instead.
type ServiceEndpointPoliciesClient struct {
	*Client
	subscriptionID string
}

// NewServiceEndpointPoliciesClient creates a new instance of ServiceEndpointPoliciesClient with the specified values.
func NewServiceEndpointPoliciesClient(c *Client, subscriptionID string) ServiceEndpointPoliciesOperations {
	return &ServiceEndpointPoliciesClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *ServiceEndpointPoliciesClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

func (client *ServiceEndpointPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters ServiceEndpointPolicy, options *ServiceEndpointPoliciesCreateOrUpdateOptions) (*ServiceEndpointPolicyPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, serviceEndpointPolicyName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &ServiceEndpointPolicyPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ServiceEndpointPoliciesClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &serviceEndpointPolicyPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ServiceEndpointPolicyResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ServiceEndpointPoliciesClient) ResumeCreateOrUpdate(token string) (ServiceEndpointPolicyPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ServiceEndpointPoliciesClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &serviceEndpointPolicyPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a service Endpoint Policies.
func (client *ServiceEndpointPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters ServiceEndpointPolicy, options *ServiceEndpointPoliciesCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, parameters, options)
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
func (client *ServiceEndpointPoliciesClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters ServiceEndpointPolicy, options *ServiceEndpointPoliciesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ServiceEndpointPoliciesClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*ServiceEndpointPolicyResponse, error) {
	result := ServiceEndpointPolicyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ServiceEndpointPolicy)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ServiceEndpointPoliciesClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ServiceEndpointPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, serviceEndpointPolicyName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ServiceEndpointPoliciesClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *ServiceEndpointPoliciesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ServiceEndpointPoliciesClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified service endpoint policy.
func (client *ServiceEndpointPoliciesClient) Delete(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, options)
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
func (client *ServiceEndpointPoliciesClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
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

// DeleteHandleError handles the Delete error response.
func (client *ServiceEndpointPoliciesClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified service Endpoint Policies in a specified resource group.
func (client *ServiceEndpointPoliciesClient) Get(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesGetOptions) (*ServiceEndpointPolicyResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, options)
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
func (client *ServiceEndpointPoliciesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *ServiceEndpointPoliciesClient) GetHandleResponse(resp *azcore.Response) (*ServiceEndpointPolicyResponse, error) {
	result := ServiceEndpointPolicyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ServiceEndpointPolicy)
}

// GetHandleError handles the Get error response.
func (client *ServiceEndpointPoliciesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all the service endpoint policies in a subscription.
func (client *ServiceEndpointPoliciesClient) List(options *ServiceEndpointPoliciesListOptions) ServiceEndpointPolicyListResultPager {
	return &serviceEndpointPolicyListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *ServiceEndpointPolicyListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ServiceEndpointPolicyListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *ServiceEndpointPoliciesClient) ListCreateRequest(ctx context.Context, options *ServiceEndpointPoliciesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ServiceEndpointPolicies"
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
func (client *ServiceEndpointPoliciesClient) ListHandleResponse(resp *azcore.Response) (*ServiceEndpointPolicyListResultResponse, error) {
	result := ServiceEndpointPolicyListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ServiceEndpointPolicyListResult)
}

// ListHandleError handles the List error response.
func (client *ServiceEndpointPoliciesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Gets all service endpoint Policies in a resource group.
func (client *ServiceEndpointPoliciesClient) ListByResourceGroup(resourceGroupName string, options *ServiceEndpointPoliciesListByResourceGroupOptions) ServiceEndpointPolicyListResultPager {
	return &serviceEndpointPolicyListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListByResourceGroupHandleResponse,
		errorer:   client.ListByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *ServiceEndpointPolicyListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ServiceEndpointPolicyListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ServiceEndpointPoliciesClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ServiceEndpointPoliciesListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *ServiceEndpointPoliciesClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*ServiceEndpointPolicyListResultResponse, error) {
	result := ServiceEndpointPolicyListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ServiceEndpointPolicyListResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ServiceEndpointPoliciesClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates tags of a service endpoint policy.
func (client *ServiceEndpointPoliciesClient) UpdateTags(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters TagsObject, options *ServiceEndpointPoliciesUpdateTagsOptions) (*ServiceEndpointPolicyResponse, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, parameters, options)
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
func (client *ServiceEndpointPoliciesClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters TagsObject, options *ServiceEndpointPoliciesUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// UpdateTagsHandleResponse handles the UpdateTags response.
func (client *ServiceEndpointPoliciesClient) UpdateTagsHandleResponse(resp *azcore.Response) (*ServiceEndpointPolicyResponse, error) {
	result := ServiceEndpointPolicyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ServiceEndpointPolicy)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *ServiceEndpointPoliciesClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
