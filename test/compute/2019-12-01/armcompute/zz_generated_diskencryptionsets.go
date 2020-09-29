// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// DiskEncryptionSetsOperations contains the methods for the DiskEncryptionSets group.
type DiskEncryptionSetsOperations interface {
	// BeginCreateOrUpdate - Creates or updates a disk encryption set
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSet, options *DiskEncryptionSetsCreateOrUpdateOptions) (*DiskEncryptionSetPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (DiskEncryptionSetPoller, error)
	// BeginDelete - Deletes a disk encryption set.
	BeginDelete(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets information about a disk encryption set.
	Get(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsGetOptions) (*DiskEncryptionSetResponse, error)
	// List - Lists all the disk encryption sets under a subscription.
	List(options *DiskEncryptionSetsListOptions) DiskEncryptionSetListPager
	// ListByResourceGroup - Lists all the disk encryption sets under a resource group.
	ListByResourceGroup(resourceGroupName string, options *DiskEncryptionSetsListByResourceGroupOptions) DiskEncryptionSetListPager
	// BeginUpdate - Updates (patches) a disk encryption set.
	BeginUpdate(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSetUpdate, options *DiskEncryptionSetsUpdateOptions) (*DiskEncryptionSetPollerResponse, error)
	// ResumeUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeUpdate(token string) (DiskEncryptionSetPoller, error)
}

// DiskEncryptionSetsClient implements the DiskEncryptionSetsOperations interface.
// Don't use this type directly, use NewDiskEncryptionSetsClient() instead.
type DiskEncryptionSetsClient struct {
	*Client
	subscriptionID string
}

// NewDiskEncryptionSetsClient creates a new instance of DiskEncryptionSetsClient with the specified values.
func NewDiskEncryptionSetsClient(c *Client, subscriptionID string) DiskEncryptionSetsOperations {
	return &DiskEncryptionSetsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *DiskEncryptionSetsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

func (client *DiskEncryptionSetsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSet, options *DiskEncryptionSetsCreateOrUpdateOptions) (*DiskEncryptionSetPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet, options)
	if err != nil {
		return nil, err
	}
	result := &DiskEncryptionSetPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("DiskEncryptionSetsClient.CreateOrUpdate", "", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &diskEncryptionSetPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*DiskEncryptionSetResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *DiskEncryptionSetsClient) ResumeCreateOrUpdate(token string) (DiskEncryptionSetPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("DiskEncryptionSetsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &diskEncryptionSetPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a disk encryption set
func (client *DiskEncryptionSetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSet, options *DiskEncryptionSetsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet, options)
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
func (client *DiskEncryptionSetsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSet, options *DiskEncryptionSetsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-11-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(diskEncryptionSet)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DiskEncryptionSetsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*DiskEncryptionSetResponse, error) {
	result := DiskEncryptionSetResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DiskEncryptionSet)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DiskEncryptionSetsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *DiskEncryptionSetsClient) BeginDelete(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, diskEncryptionSetName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("DiskEncryptionSetsClient.Delete", "", resp, client.DeleteHandleError)
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

func (client *DiskEncryptionSetsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("DiskEncryptionSetsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Delete - Deletes a disk encryption set.
func (client *DiskEncryptionSetsClient) Delete(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, options)
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
func (client *DiskEncryptionSetsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-11-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *DiskEncryptionSetsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets information about a disk encryption set.
func (client *DiskEncryptionSetsClient) Get(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsGetOptions) (*DiskEncryptionSetResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, options)
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
func (client *DiskEncryptionSetsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-11-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *DiskEncryptionSetsClient) GetHandleResponse(resp *azcore.Response) (*DiskEncryptionSetResponse, error) {
	result := DiskEncryptionSetResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DiskEncryptionSet)
}

// GetHandleError handles the Get error response.
func (client *DiskEncryptionSetsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all the disk encryption sets under a subscription.
func (client *DiskEncryptionSetsClient) List(options *DiskEncryptionSetsListOptions) DiskEncryptionSetListPager {
	return &diskEncryptionSetListPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *DiskEncryptionSetListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DiskEncryptionSetList.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *DiskEncryptionSetsClient) ListCreateRequest(ctx context.Context, options *DiskEncryptionSetsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/diskEncryptionSets"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-11-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *DiskEncryptionSetsClient) ListHandleResponse(resp *azcore.Response) (*DiskEncryptionSetListResponse, error) {
	result := DiskEncryptionSetListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DiskEncryptionSetList)
}

// ListHandleError handles the List error response.
func (client *DiskEncryptionSetsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Lists all the disk encryption sets under a resource group.
func (client *DiskEncryptionSetsClient) ListByResourceGroup(resourceGroupName string, options *DiskEncryptionSetsListByResourceGroupOptions) DiskEncryptionSetListPager {
	return &diskEncryptionSetListPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListByResourceGroupHandleResponse,
		errorer:   client.ListByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *DiskEncryptionSetListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DiskEncryptionSetList.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DiskEncryptionSetsClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DiskEncryptionSetsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-11-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DiskEncryptionSetsClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*DiskEncryptionSetListResponse, error) {
	result := DiskEncryptionSetListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DiskEncryptionSetList)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *DiskEncryptionSetsClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *DiskEncryptionSetsClient) BeginUpdate(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSetUpdate, options *DiskEncryptionSetsUpdateOptions) (*DiskEncryptionSetPollerResponse, error) {
	resp, err := client.Update(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet, options)
	if err != nil {
		return nil, err
	}
	result := &DiskEncryptionSetPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("DiskEncryptionSetsClient.Update", "", resp, client.UpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &diskEncryptionSetPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*DiskEncryptionSetResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *DiskEncryptionSetsClient) ResumeUpdate(token string) (DiskEncryptionSetPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("DiskEncryptionSetsClient.Update", token, client.UpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &diskEncryptionSetPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Update - Updates (patches) a disk encryption set.
func (client *DiskEncryptionSetsClient) Update(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSetUpdate, options *DiskEncryptionSetsUpdateOptions) (*azcore.Response, error) {
	req, err := client.UpdateCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.UpdateHandleError(resp)
	}
	return resp, nil
}

// UpdateCreateRequest creates the Update request.
func (client *DiskEncryptionSetsClient) UpdateCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSetUpdate, options *DiskEncryptionSetsUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-11-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(diskEncryptionSet)
}

// UpdateHandleResponse handles the Update response.
func (client *DiskEncryptionSetsClient) UpdateHandleResponse(resp *azcore.Response) (*DiskEncryptionSetResponse, error) {
	result := DiskEncryptionSetResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DiskEncryptionSet)
}

// UpdateHandleError handles the Update error response.
func (client *DiskEncryptionSetsClient) UpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
