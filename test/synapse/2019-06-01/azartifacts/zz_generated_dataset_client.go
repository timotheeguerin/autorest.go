//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

type datasetClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newDatasetClient creates a new instance of datasetClient with the specified values.
func newDatasetClient(endpoint string, pl runtime.Pipeline) *datasetClient {
	client := &datasetClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// BeginCreateOrUpdateDataset - Creates or updates a dataset.
// If the operation fails it returns the *CloudError error type.
func (client *datasetClient) BeginCreateOrUpdateDataset(ctx context.Context, datasetName string, dataset DatasetResource, options *DatasetBeginCreateOrUpdateDatasetOptions) (DatasetCreateOrUpdateDatasetPollerResponse, error) {
	resp, err := client.createOrUpdateDataset(ctx, datasetName, dataset, options)
	if err != nil {
		return DatasetCreateOrUpdateDatasetPollerResponse{}, err
	}
	result := DatasetCreateOrUpdateDatasetPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("datasetClient.CreateOrUpdateDataset", resp, client.pl, client.createOrUpdateDatasetHandleError)
	if err != nil {
		return DatasetCreateOrUpdateDatasetPollerResponse{}, err
	}
	result.Poller = &DatasetCreateOrUpdateDatasetPoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdateDataset - Creates or updates a dataset.
// If the operation fails it returns the *CloudError error type.
func (client *datasetClient) createOrUpdateDataset(ctx context.Context, datasetName string, dataset DatasetResource, options *DatasetBeginCreateOrUpdateDatasetOptions) (*http.Response, error) {
	req, err := client.createOrUpdateDatasetCreateRequest(ctx, datasetName, dataset, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateDatasetHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateDatasetCreateRequest creates the CreateOrUpdateDataset request.
func (client *datasetClient) createOrUpdateDatasetCreateRequest(ctx context.Context, datasetName string, dataset DatasetResource, options *DatasetBeginCreateOrUpdateDatasetOptions) (*policy.Request, error) {
	urlPath := "/datasets/{datasetName}"
	if datasetName == "" {
		return nil, errors.New("parameter datasetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dataset)
}

// createOrUpdateDatasetHandleError handles the CreateOrUpdateDataset error response.
func (client *datasetClient) createOrUpdateDatasetHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDeleteDataset - Deletes a dataset.
// If the operation fails it returns the *CloudError error type.
func (client *datasetClient) BeginDeleteDataset(ctx context.Context, datasetName string, options *DatasetBeginDeleteDatasetOptions) (DatasetDeleteDatasetPollerResponse, error) {
	resp, err := client.deleteDataset(ctx, datasetName, options)
	if err != nil {
		return DatasetDeleteDatasetPollerResponse{}, err
	}
	result := DatasetDeleteDatasetPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("datasetClient.DeleteDataset", resp, client.pl, client.deleteDatasetHandleError)
	if err != nil {
		return DatasetDeleteDatasetPollerResponse{}, err
	}
	result.Poller = &DatasetDeleteDatasetPoller{
		pt: pt,
	}
	return result, nil
}

// DeleteDataset - Deletes a dataset.
// If the operation fails it returns the *CloudError error type.
func (client *datasetClient) deleteDataset(ctx context.Context, datasetName string, options *DatasetBeginDeleteDatasetOptions) (*http.Response, error) {
	req, err := client.deleteDatasetCreateRequest(ctx, datasetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteDatasetHandleError(resp)
	}
	return resp, nil
}

// deleteDatasetCreateRequest creates the DeleteDataset request.
func (client *datasetClient) deleteDatasetCreateRequest(ctx context.Context, datasetName string, options *DatasetBeginDeleteDatasetOptions) (*policy.Request, error) {
	urlPath := "/datasets/{datasetName}"
	if datasetName == "" {
		return nil, errors.New("parameter datasetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteDatasetHandleError handles the DeleteDataset error response.
func (client *datasetClient) deleteDatasetHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetDataset - Gets a dataset.
// If the operation fails it returns the *CloudError error type.
func (client *datasetClient) GetDataset(ctx context.Context, datasetName string, options *DatasetGetDatasetOptions) (DatasetGetDatasetResponse, error) {
	req, err := client.getDatasetCreateRequest(ctx, datasetName, options)
	if err != nil {
		return DatasetGetDatasetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatasetGetDatasetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return DatasetGetDatasetResponse{}, client.getDatasetHandleError(resp)
	}
	return client.getDatasetHandleResponse(resp)
}

// getDatasetCreateRequest creates the GetDataset request.
func (client *datasetClient) getDatasetCreateRequest(ctx context.Context, datasetName string, options *DatasetGetDatasetOptions) (*policy.Request, error) {
	urlPath := "/datasets/{datasetName}"
	if datasetName == "" {
		return nil, errors.New("parameter datasetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getDatasetHandleResponse handles the GetDataset response.
func (client *datasetClient) getDatasetHandleResponse(resp *http.Response) (DatasetGetDatasetResponse, error) {
	result := DatasetGetDatasetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatasetResource); err != nil {
		return DatasetGetDatasetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getDatasetHandleError handles the GetDataset error response.
func (client *datasetClient) getDatasetHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetDatasetsByWorkspace - Lists datasets.
// If the operation fails it returns the *CloudError error type.
func (client *datasetClient) GetDatasetsByWorkspace(options *DatasetGetDatasetsByWorkspaceOptions) *DatasetGetDatasetsByWorkspacePager {
	return &DatasetGetDatasetsByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getDatasetsByWorkspaceCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DatasetGetDatasetsByWorkspaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DatasetListResponse.NextLink)
		},
	}
}

// getDatasetsByWorkspaceCreateRequest creates the GetDatasetsByWorkspace request.
func (client *datasetClient) getDatasetsByWorkspaceCreateRequest(ctx context.Context, options *DatasetGetDatasetsByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/datasets"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getDatasetsByWorkspaceHandleResponse handles the GetDatasetsByWorkspace response.
func (client *datasetClient) getDatasetsByWorkspaceHandleResponse(resp *http.Response) (DatasetGetDatasetsByWorkspaceResponse, error) {
	result := DatasetGetDatasetsByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatasetListResponse); err != nil {
		return DatasetGetDatasetsByWorkspaceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getDatasetsByWorkspaceHandleError handles the GetDatasetsByWorkspace error response.
func (client *datasetClient) getDatasetsByWorkspaceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginRenameDataset - Renames a dataset.
// If the operation fails it returns the *CloudError error type.
func (client *datasetClient) BeginRenameDataset(ctx context.Context, datasetName string, request ArtifactRenameRequest, options *DatasetBeginRenameDatasetOptions) (DatasetRenameDatasetPollerResponse, error) {
	resp, err := client.renameDataset(ctx, datasetName, request, options)
	if err != nil {
		return DatasetRenameDatasetPollerResponse{}, err
	}
	result := DatasetRenameDatasetPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("datasetClient.RenameDataset", resp, client.pl, client.renameDatasetHandleError)
	if err != nil {
		return DatasetRenameDatasetPollerResponse{}, err
	}
	result.Poller = &DatasetRenameDatasetPoller{
		pt: pt,
	}
	return result, nil
}

// RenameDataset - Renames a dataset.
// If the operation fails it returns the *CloudError error type.
func (client *datasetClient) renameDataset(ctx context.Context, datasetName string, request ArtifactRenameRequest, options *DatasetBeginRenameDatasetOptions) (*http.Response, error) {
	req, err := client.renameDatasetCreateRequest(ctx, datasetName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.renameDatasetHandleError(resp)
	}
	return resp, nil
}

// renameDatasetCreateRequest creates the RenameDataset request.
func (client *datasetClient) renameDatasetCreateRequest(ctx context.Context, datasetName string, request ArtifactRenameRequest, options *DatasetBeginRenameDatasetOptions) (*policy.Request, error) {
	urlPath := "/datasets/{datasetName}/rename"
	if datasetName == "" {
		return nil, errors.New("parameter datasetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// renameDatasetHandleError handles the RenameDataset error response.
func (client *datasetClient) renameDatasetHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
