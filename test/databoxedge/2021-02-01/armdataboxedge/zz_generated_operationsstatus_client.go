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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// OperationsStatusClient contains the methods for the OperationsStatus group.
// Don't use this type directly, use NewOperationsStatusClient() instead.
type OperationsStatusClient struct {
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOperationsStatusClient creates a new instance of OperationsStatusClient with the specified values.
func NewOperationsStatusClient(subscriptionID string, options *azcore.ClientOptions) *OperationsStatusClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &OperationsStatusClient{
		subscriptionID: subscriptionID,
		pl:             runtime.NewPipeline(module, version, nil, nil, &cp),
	}
	return client
}

// Get - Gets the details of a specified job on a Data Box Edge/Data Box Gateway device.
// If the operation fails it returns the *CloudError error type.
func (client *OperationsStatusClient) Get(ctx context.Context, deviceName string, name string, resourceGroupName string, options *OperationsStatusGetOptions) (OperationsStatusGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, name, resourceGroupName, options)
	if err != nil {
		return OperationsStatusGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationsStatusGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationsStatusGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OperationsStatusClient) getCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *OperationsStatusGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/operationsStatus/{name}"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
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
func (client *OperationsStatusClient) getHandleResponse(resp *http.Response) (OperationsStatusGetResponse, error) {
	result := OperationsStatusGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Job); err != nil {
		return OperationsStatusGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *OperationsStatusClient) getHandleError(resp *http.Response) error {
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
