//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package lrogroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// LROsCustomHeaderClient contains the methods for the LROsCustomHeader group.
// Don't use this type directly, use NewLROsCustomHeaderClient() instead.
type LROsCustomHeaderClient struct {
	pl runtime.Pipeline
}

// NewLROsCustomHeaderClient creates a new instance of LROsCustomHeaderClient with the specified values.
func NewLROsCustomHeaderClient(options *azcore.ClientOptions) *LROsCustomHeaderClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &LROsCustomHeaderClient{
		pl: runtime.NewPipeline(module, version, nil, nil, &cp),
	}
	return client
}

// BeginPost202Retry200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post request,
// service returns a 202 to the initial request, with 'Location' and
// 'Retry-After' headers, Polls return a 200 with a response body after success
// If the operation fails it returns the *CloudError error type.
func (client *LROsCustomHeaderClient) BeginPost202Retry200(ctx context.Context, options *LROsCustomHeaderBeginPost202Retry200Options) (LROsCustomHeaderPost202Retry200PollerResponse, error) {
	resp, err := client.post202Retry200(ctx, options)
	if err != nil {
		return LROsCustomHeaderPost202Retry200PollerResponse{}, err
	}
	result := LROsCustomHeaderPost202Retry200PollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LROsCustomHeaderClient.Post202Retry200", "", resp, client.pl, client.post202Retry200HandleError)
	if err != nil {
		return LROsCustomHeaderPost202Retry200PollerResponse{}, err
	}
	result.Poller = &LROsCustomHeaderPost202Retry200Poller{
		pt: pt,
	}
	return result, nil
}

// Post202Retry200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post request,
// service returns a 202 to the initial request, with 'Location' and
// 'Retry-After' headers, Polls return a 200 with a response body after success
// If the operation fails it returns the *CloudError error type.
func (client *LROsCustomHeaderClient) post202Retry200(ctx context.Context, options *LROsCustomHeaderBeginPost202Retry200Options) (*http.Response, error) {
	req, err := client.post202Retry200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, client.post202Retry200HandleError(resp)
	}
	return resp, nil
}

// post202Retry200CreateRequest creates the Post202Retry200 request.
func (client *LROsCustomHeaderClient) post202Retry200CreateRequest(ctx context.Context, options *LROsCustomHeaderBeginPost202Retry200Options) (*policy.Request, error) {
	urlPath := "/lro/customheader/post/202/retry/200"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Product != nil {
		return req, runtime.MarshalAsJSON(req, *options.Product)
	}
	return req, nil
}

// post202Retry200HandleError handles the Post202Retry200 error response.
func (client *LROsCustomHeaderClient) post202Retry200HandleError(resp *http.Response) error {
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

// BeginPostAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running
// post request, service returns a 202 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LROsCustomHeaderClient) BeginPostAsyncRetrySucceeded(ctx context.Context, options *LROsCustomHeaderBeginPostAsyncRetrySucceededOptions) (LROsCustomHeaderPostAsyncRetrySucceededPollerResponse, error) {
	resp, err := client.postAsyncRetrySucceeded(ctx, options)
	if err != nil {
		return LROsCustomHeaderPostAsyncRetrySucceededPollerResponse{}, err
	}
	result := LROsCustomHeaderPostAsyncRetrySucceededPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LROsCustomHeaderClient.PostAsyncRetrySucceeded", "", resp, client.pl, client.postAsyncRetrySucceededHandleError)
	if err != nil {
		return LROsCustomHeaderPostAsyncRetrySucceededPollerResponse{}, err
	}
	result.Poller = &LROsCustomHeaderPostAsyncRetrySucceededPoller{
		pt: pt,
	}
	return result, nil
}

// PostAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post
// request, service returns a 202 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LROsCustomHeaderClient) postAsyncRetrySucceeded(ctx context.Context, options *LROsCustomHeaderBeginPostAsyncRetrySucceededOptions) (*http.Response, error) {
	req, err := client.postAsyncRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, client.postAsyncRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// postAsyncRetrySucceededCreateRequest creates the PostAsyncRetrySucceeded request.
func (client *LROsCustomHeaderClient) postAsyncRetrySucceededCreateRequest(ctx context.Context, options *LROsCustomHeaderBeginPostAsyncRetrySucceededOptions) (*policy.Request, error) {
	urlPath := "/lro/customheader/postasync/retry/succeeded"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Product != nil {
		return req, runtime.MarshalAsJSON(req, *options.Product)
	}
	return req, nil
}

// postAsyncRetrySucceededHandleError handles the PostAsyncRetrySucceeded error response.
func (client *LROsCustomHeaderClient) postAsyncRetrySucceededHandleError(resp *http.Response) error {
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

// BeginPut201CreatingSucceeded200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running
// put request, service returns a 201 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns the *CloudError error type.
func (client *LROsCustomHeaderClient) BeginPut201CreatingSucceeded200(ctx context.Context, options *LROsCustomHeaderBeginPut201CreatingSucceeded200Options) (LROsCustomHeaderPut201CreatingSucceeded200PollerResponse, error) {
	resp, err := client.put201CreatingSucceeded200(ctx, options)
	if err != nil {
		return LROsCustomHeaderPut201CreatingSucceeded200PollerResponse{}, err
	}
	result := LROsCustomHeaderPut201CreatingSucceeded200PollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LROsCustomHeaderClient.Put201CreatingSucceeded200", "", resp, client.pl, client.put201CreatingSucceeded200HandleError)
	if err != nil {
		return LROsCustomHeaderPut201CreatingSucceeded200PollerResponse{}, err
	}
	result.Poller = &LROsCustomHeaderPut201CreatingSucceeded200Poller{
		pt: pt,
	}
	return result, nil
}

// Put201CreatingSucceeded200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running
// put request, service returns a 201 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns the *CloudError error type.
func (client *LROsCustomHeaderClient) put201CreatingSucceeded200(ctx context.Context, options *LROsCustomHeaderBeginPut201CreatingSucceeded200Options) (*http.Response, error) {
	req, err := client.put201CreatingSucceeded200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.put201CreatingSucceeded200HandleError(resp)
	}
	return resp, nil
}

// put201CreatingSucceeded200CreateRequest creates the Put201CreatingSucceeded200 request.
func (client *LROsCustomHeaderClient) put201CreatingSucceeded200CreateRequest(ctx context.Context, options *LROsCustomHeaderBeginPut201CreatingSucceeded200Options) (*policy.Request, error) {
	urlPath := "/lro/customheader/put/201/creating/succeeded/200"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Product != nil {
		return req, runtime.MarshalAsJSON(req, *options.Product)
	}
	return req, nil
}

// put201CreatingSucceeded200HandleError handles the Put201CreatingSucceeded200 error response.
func (client *LROsCustomHeaderClient) put201CreatingSucceeded200HandleError(resp *http.Response) error {
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

// BeginPutAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running
// put request, service returns a 200 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LROsCustomHeaderClient) BeginPutAsyncRetrySucceeded(ctx context.Context, options *LROsCustomHeaderBeginPutAsyncRetrySucceededOptions) (LROsCustomHeaderPutAsyncRetrySucceededPollerResponse, error) {
	resp, err := client.putAsyncRetrySucceeded(ctx, options)
	if err != nil {
		return LROsCustomHeaderPutAsyncRetrySucceededPollerResponse{}, err
	}
	result := LROsCustomHeaderPutAsyncRetrySucceededPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LROsCustomHeaderClient.PutAsyncRetrySucceeded", "", resp, client.pl, client.putAsyncRetrySucceededHandleError)
	if err != nil {
		return LROsCustomHeaderPutAsyncRetrySucceededPollerResponse{}, err
	}
	result.Poller = &LROsCustomHeaderPutAsyncRetrySucceededPoller{
		pt: pt,
	}
	return result, nil
}

// PutAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running put
// request, service returns a 200 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LROsCustomHeaderClient) putAsyncRetrySucceeded(ctx context.Context, options *LROsCustomHeaderBeginPutAsyncRetrySucceededOptions) (*http.Response, error) {
	req, err := client.putAsyncRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, client.putAsyncRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// putAsyncRetrySucceededCreateRequest creates the PutAsyncRetrySucceeded request.
func (client *LROsCustomHeaderClient) putAsyncRetrySucceededCreateRequest(ctx context.Context, options *LROsCustomHeaderBeginPutAsyncRetrySucceededOptions) (*policy.Request, error) {
	urlPath := "/lro/customheader/putasync/retry/succeeded"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Product != nil {
		return req, runtime.MarshalAsJSON(req, *options.Product)
	}
	return req, nil
}

// putAsyncRetrySucceededHandleError handles the PutAsyncRetrySucceeded error response.
func (client *LROsCustomHeaderClient) putAsyncRetrySucceededHandleError(resp *http.Response) error {
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
