//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

import (
	"context"
	"errors"
	"fmt"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// StorageAccountCredentialsClient contains the methods for the StorageAccountCredentials group.
// Don't use this type directly, use NewStorageAccountCredentialsClient() instead.
type StorageAccountCredentialsClient struct {
	con            *connection
	subscriptionID string
}

// NewStorageAccountCredentialsClient creates a new instance of StorageAccountCredentialsClient with the specified values.
func NewStorageAccountCredentialsClient(con *connection, subscriptionID string) *StorageAccountCredentialsClient {
	return &StorageAccountCredentialsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates the storage account credential.
// If the operation fails it returns the *CloudError error type.
func (client *StorageAccountCredentialsClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, name string, resourceGroupName string, storageAccountCredential StorageAccountCredential, options *StorageAccountCredentialsBeginCreateOrUpdateOptions) (StorageAccountCredentialsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, deviceName, name, resourceGroupName, storageAccountCredential, options)
	if err != nil {
		return StorageAccountCredentialsCreateOrUpdatePollerResponse{}, err
	}
	result := StorageAccountCredentialsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("StorageAccountCredentialsClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return StorageAccountCredentialsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &StorageAccountCredentialsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the storage account credential.
// If the operation fails it returns the *CloudError error type.
func (client *StorageAccountCredentialsClient) createOrUpdate(ctx context.Context, deviceName string, name string, resourceGroupName string, storageAccountCredential StorageAccountCredential, options *StorageAccountCredentialsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, name, resourceGroupName, storageAccountCredential, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *StorageAccountCredentialsClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, storageAccountCredential StorageAccountCredential, options *StorageAccountCredentialsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, storageAccountCredential)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *StorageAccountCredentialsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes the storage account credential.
// If the operation fails it returns the *CloudError error type.
func (client *StorageAccountCredentialsClient) BeginDelete(ctx context.Context, deviceName string, name string, resourceGroupName string, options *StorageAccountCredentialsBeginDeleteOptions) (StorageAccountCredentialsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, deviceName, name, resourceGroupName, options)
	if err != nil {
		return StorageAccountCredentialsDeletePollerResponse{}, err
	}
	result := StorageAccountCredentialsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("StorageAccountCredentialsClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return StorageAccountCredentialsDeletePollerResponse{}, err
	}
	result.Poller = &StorageAccountCredentialsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the storage account credential.
// If the operation fails it returns the *CloudError error type.
func (client *StorageAccountCredentialsClient) deleteOperation(ctx context.Context, deviceName string, name string, resourceGroupName string, options *StorageAccountCredentialsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, deviceName, name, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StorageAccountCredentialsClient) deleteCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *StorageAccountCredentialsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *StorageAccountCredentialsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets the properties of the specified storage account credential.
// If the operation fails it returns the *CloudError error type.
func (client *StorageAccountCredentialsClient) Get(ctx context.Context, deviceName string, name string, resourceGroupName string, options *StorageAccountCredentialsGetOptions) (StorageAccountCredentialsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, name, resourceGroupName, options)
	if err != nil {
		return StorageAccountCredentialsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return StorageAccountCredentialsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageAccountCredentialsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *StorageAccountCredentialsClient) getCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *StorageAccountCredentialsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StorageAccountCredentialsClient) getHandleResponse(resp *http.Response) (StorageAccountCredentialsGetResponse, error) {
	result := StorageAccountCredentialsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageAccountCredential); err != nil {
		return StorageAccountCredentialsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *StorageAccountCredentialsClient) getHandleError(resp *http.Response) error {
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

// ListByDataBoxEdgeDevice - Gets all the storage account credentials in a Data Box Edge/Data Box Gateway device.
// If the operation fails it returns the *CloudError error type.
func (client *StorageAccountCredentialsClient) ListByDataBoxEdgeDevice(deviceName string, resourceGroupName string, options *StorageAccountCredentialsListByDataBoxEdgeDeviceOptions) *StorageAccountCredentialsListByDataBoxEdgeDevicePager {
	return &StorageAccountCredentialsListByDataBoxEdgeDevicePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDataBoxEdgeDeviceCreateRequest(ctx, deviceName, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp StorageAccountCredentialsListByDataBoxEdgeDeviceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.StorageAccountCredentialList.NextLink)
		},
	}
}

// listByDataBoxEdgeDeviceCreateRequest creates the ListByDataBoxEdgeDevice request.
func (client *StorageAccountCredentialsClient) listByDataBoxEdgeDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *StorageAccountCredentialsListByDataBoxEdgeDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDataBoxEdgeDeviceHandleResponse handles the ListByDataBoxEdgeDevice response.
func (client *StorageAccountCredentialsClient) listByDataBoxEdgeDeviceHandleResponse(resp *http.Response) (StorageAccountCredentialsListByDataBoxEdgeDeviceResponse, error) {
	result := StorageAccountCredentialsListByDataBoxEdgeDeviceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageAccountCredentialList); err != nil {
		return StorageAccountCredentialsListByDataBoxEdgeDeviceResponse{}, err
	}
	return result, nil
}

// listByDataBoxEdgeDeviceHandleError handles the ListByDataBoxEdgeDevice error response.
func (client *StorageAccountCredentialsClient) listByDataBoxEdgeDeviceHandleError(resp *http.Response) error {
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
