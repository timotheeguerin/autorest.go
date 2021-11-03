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

type triggerRunClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newTriggerRunClient creates a new instance of triggerRunClient with the specified values.
func newTriggerRunClient(endpoint string, pl runtime.Pipeline) *triggerRunClient {
	client := &triggerRunClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// CancelTriggerInstance - Cancel single trigger instance by runId.
// If the operation fails it returns the *CloudError error type.
func (client *triggerRunClient) CancelTriggerInstance(ctx context.Context, triggerName string, runID string, options *TriggerRunCancelTriggerInstanceOptions) (TriggerRunCancelTriggerInstanceResponse, error) {
	req, err := client.cancelTriggerInstanceCreateRequest(ctx, triggerName, runID, options)
	if err != nil {
		return TriggerRunCancelTriggerInstanceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TriggerRunCancelTriggerInstanceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TriggerRunCancelTriggerInstanceResponse{}, client.cancelTriggerInstanceHandleError(resp)
	}
	return TriggerRunCancelTriggerInstanceResponse{RawResponse: resp}, nil
}

// cancelTriggerInstanceCreateRequest creates the CancelTriggerInstance request.
func (client *triggerRunClient) cancelTriggerInstanceCreateRequest(ctx context.Context, triggerName string, runID string, options *TriggerRunCancelTriggerInstanceOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/triggerRuns/{runId}/cancel"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// cancelTriggerInstanceHandleError handles the CancelTriggerInstance error response.
func (client *triggerRunClient) cancelTriggerInstanceHandleError(resp *http.Response) error {
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

// QueryTriggerRunsByWorkspace - Query trigger runs.
// If the operation fails it returns the *CloudError error type.
func (client *triggerRunClient) QueryTriggerRunsByWorkspace(ctx context.Context, filterParameters RunFilterParameters, options *TriggerRunQueryTriggerRunsByWorkspaceOptions) (TriggerRunQueryTriggerRunsByWorkspaceResponse, error) {
	req, err := client.queryTriggerRunsByWorkspaceCreateRequest(ctx, filterParameters, options)
	if err != nil {
		return TriggerRunQueryTriggerRunsByWorkspaceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TriggerRunQueryTriggerRunsByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TriggerRunQueryTriggerRunsByWorkspaceResponse{}, client.queryTriggerRunsByWorkspaceHandleError(resp)
	}
	return client.queryTriggerRunsByWorkspaceHandleResponse(resp)
}

// queryTriggerRunsByWorkspaceCreateRequest creates the QueryTriggerRunsByWorkspace request.
func (client *triggerRunClient) queryTriggerRunsByWorkspaceCreateRequest(ctx context.Context, filterParameters RunFilterParameters, options *TriggerRunQueryTriggerRunsByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/queryTriggerRuns"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, filterParameters)
}

// queryTriggerRunsByWorkspaceHandleResponse handles the QueryTriggerRunsByWorkspace response.
func (client *triggerRunClient) queryTriggerRunsByWorkspaceHandleResponse(resp *http.Response) (TriggerRunQueryTriggerRunsByWorkspaceResponse, error) {
	result := TriggerRunQueryTriggerRunsByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerRunsQueryResponse); err != nil {
		return TriggerRunQueryTriggerRunsByWorkspaceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// queryTriggerRunsByWorkspaceHandleError handles the QueryTriggerRunsByWorkspace error response.
func (client *triggerRunClient) queryTriggerRunsByWorkspaceHandleError(resp *http.Response) error {
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

// RerunTriggerInstance - Rerun single trigger instance by runId.
// If the operation fails it returns the *CloudError error type.
func (client *triggerRunClient) RerunTriggerInstance(ctx context.Context, triggerName string, runID string, options *TriggerRunRerunTriggerInstanceOptions) (TriggerRunRerunTriggerInstanceResponse, error) {
	req, err := client.rerunTriggerInstanceCreateRequest(ctx, triggerName, runID, options)
	if err != nil {
		return TriggerRunRerunTriggerInstanceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TriggerRunRerunTriggerInstanceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TriggerRunRerunTriggerInstanceResponse{}, client.rerunTriggerInstanceHandleError(resp)
	}
	return TriggerRunRerunTriggerInstanceResponse{RawResponse: resp}, nil
}

// rerunTriggerInstanceCreateRequest creates the RerunTriggerInstance request.
func (client *triggerRunClient) rerunTriggerInstanceCreateRequest(ctx context.Context, triggerName string, runID string, options *TriggerRunRerunTriggerInstanceOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/triggerRuns/{runId}/rerun"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// rerunTriggerInstanceHandleError handles the RerunTriggerInstance error response.
func (client *triggerRunClient) rerunTriggerInstanceHandleError(resp *http.Response) error {
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
