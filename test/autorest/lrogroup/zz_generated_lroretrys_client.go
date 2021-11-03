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

// LRORetrysClient contains the methods for the LRORetrys group.
// Don't use this type directly, use NewLRORetrysClient() instead.
type LRORetrysClient struct {
	pl runtime.Pipeline
}

// NewLRORetrysClient creates a new instance of LRORetrysClient with the specified values.
func NewLRORetrysClient(options *azcore.ClientOptions) *LRORetrysClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &LRORetrysClient{
		pl: runtime.NewPipeline(module, version, nil, nil, &cp),
	}
	return client
}

// BeginDelete202Retry200 - Long running delete request, service returns a 500, then a 202 to the initial request. Polls return this value until the last
// poll returns a ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) BeginDelete202Retry200(ctx context.Context, options *LRORetrysBeginDelete202Retry200Options) (LRORetrysDelete202Retry200PollerResponse, error) {
	resp, err := client.delete202Retry200(ctx, options)
	if err != nil {
		return LRORetrysDelete202Retry200PollerResponse{}, err
	}
	result := LRORetrysDelete202Retry200PollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LRORetrysClient.Delete202Retry200", "", resp, client.pl, client.delete202Retry200HandleError)
	if err != nil {
		return LRORetrysDelete202Retry200PollerResponse{}, err
	}
	result.Poller = &LRORetrysDelete202Retry200Poller{
		pt: pt,
	}
	return result, nil
}

// Delete202Retry200 - Long running delete request, service returns a 500, then a 202 to the initial request. Polls return this value until the last poll
// returns a ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) delete202Retry200(ctx context.Context, options *LRORetrysBeginDelete202Retry200Options) (*http.Response, error) {
	req, err := client.delete202Retry200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, client.delete202Retry200HandleError(resp)
	}
	return resp, nil
}

// delete202Retry200CreateRequest creates the Delete202Retry200 request.
func (client *LRORetrysClient) delete202Retry200CreateRequest(ctx context.Context, options *LRORetrysBeginDelete202Retry200Options) (*policy.Request, error) {
	urlPath := "/lro/retryerror/delete/202/retry/200"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// delete202Retry200HandleError handles the Delete202Retry200 error response.
func (client *LRORetrysClient) delete202Retry200HandleError(resp *http.Response) error {
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

// BeginDeleteAsyncRelativeRetrySucceeded - Long running delete request, service returns a 500, then a 202 to the initial request. Poll the endpoint indicated
// in the Azure-AsyncOperation header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) BeginDeleteAsyncRelativeRetrySucceeded(ctx context.Context, options *LRORetrysBeginDeleteAsyncRelativeRetrySucceededOptions) (LRORetrysDeleteAsyncRelativeRetrySucceededPollerResponse, error) {
	resp, err := client.deleteAsyncRelativeRetrySucceeded(ctx, options)
	if err != nil {
		return LRORetrysDeleteAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	result := LRORetrysDeleteAsyncRelativeRetrySucceededPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LRORetrysClient.DeleteAsyncRelativeRetrySucceeded", "", resp, client.pl, client.deleteAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return LRORetrysDeleteAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	result.Poller = &LRORetrysDeleteAsyncRelativeRetrySucceededPoller{
		pt: pt,
	}
	return result, nil
}

// DeleteAsyncRelativeRetrySucceeded - Long running delete request, service returns a 500, then a 202 to the initial request. Poll the endpoint indicated
// in the Azure-AsyncOperation header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) deleteAsyncRelativeRetrySucceeded(ctx context.Context, options *LRORetrysBeginDeleteAsyncRelativeRetrySucceededOptions) (*http.Response, error) {
	req, err := client.deleteAsyncRelativeRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, client.deleteAsyncRelativeRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// deleteAsyncRelativeRetrySucceededCreateRequest creates the DeleteAsyncRelativeRetrySucceeded request.
func (client *LRORetrysClient) deleteAsyncRelativeRetrySucceededCreateRequest(ctx context.Context, options *LRORetrysBeginDeleteAsyncRelativeRetrySucceededOptions) (*policy.Request, error) {
	urlPath := "/lro/retryerror/deleteasync/retry/succeeded"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteAsyncRelativeRetrySucceededHandleError handles the DeleteAsyncRelativeRetrySucceeded error response.
func (client *LRORetrysClient) deleteAsyncRelativeRetrySucceededHandleError(resp *http.Response) error {
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

// BeginDeleteProvisioning202Accepted200Succeeded - Long running delete request, service returns a 500, then a 202 to the initial request, with an entity
// that contains ProvisioningState=’Accepted’. Polls return this value until the last poll returns a
// ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) BeginDeleteProvisioning202Accepted200Succeeded(ctx context.Context, options *LRORetrysBeginDeleteProvisioning202Accepted200SucceededOptions) (LRORetrysDeleteProvisioning202Accepted200SucceededPollerResponse, error) {
	resp, err := client.deleteProvisioning202Accepted200Succeeded(ctx, options)
	if err != nil {
		return LRORetrysDeleteProvisioning202Accepted200SucceededPollerResponse{}, err
	}
	result := LRORetrysDeleteProvisioning202Accepted200SucceededPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LRORetrysClient.DeleteProvisioning202Accepted200Succeeded", "", resp, client.pl, client.deleteProvisioning202Accepted200SucceededHandleError)
	if err != nil {
		return LRORetrysDeleteProvisioning202Accepted200SucceededPollerResponse{}, err
	}
	result.Poller = &LRORetrysDeleteProvisioning202Accepted200SucceededPoller{
		pt: pt,
	}
	return result, nil
}

// DeleteProvisioning202Accepted200Succeeded - Long running delete request, service returns a 500, then a 202 to the initial request, with an entity that
// contains ProvisioningState=’Accepted’. Polls return this value until the last poll returns a
// ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) deleteProvisioning202Accepted200Succeeded(ctx context.Context, options *LRORetrysBeginDeleteProvisioning202Accepted200SucceededOptions) (*http.Response, error) {
	req, err := client.deleteProvisioning202Accepted200SucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.deleteProvisioning202Accepted200SucceededHandleError(resp)
	}
	return resp, nil
}

// deleteProvisioning202Accepted200SucceededCreateRequest creates the DeleteProvisioning202Accepted200Succeeded request.
func (client *LRORetrysClient) deleteProvisioning202Accepted200SucceededCreateRequest(ctx context.Context, options *LRORetrysBeginDeleteProvisioning202Accepted200SucceededOptions) (*policy.Request, error) {
	urlPath := "/lro/retryerror/delete/provisioning/202/accepted/200/succeeded"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteProvisioning202Accepted200SucceededHandleError handles the DeleteProvisioning202Accepted200Succeeded error response.
func (client *LRORetrysClient) deleteProvisioning202Accepted200SucceededHandleError(resp *http.Response) error {
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

// BeginPost202Retry200 - Long running post request, service returns a 500, then a 202 to the initial request, with 'Location' and 'Retry-After' headers,
// Polls return a 200 with a response body after success
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) BeginPost202Retry200(ctx context.Context, options *LRORetrysBeginPost202Retry200Options) (LRORetrysPost202Retry200PollerResponse, error) {
	resp, err := client.post202Retry200(ctx, options)
	if err != nil {
		return LRORetrysPost202Retry200PollerResponse{}, err
	}
	result := LRORetrysPost202Retry200PollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LRORetrysClient.Post202Retry200", "", resp, client.pl, client.post202Retry200HandleError)
	if err != nil {
		return LRORetrysPost202Retry200PollerResponse{}, err
	}
	result.Poller = &LRORetrysPost202Retry200Poller{
		pt: pt,
	}
	return result, nil
}

// Post202Retry200 - Long running post request, service returns a 500, then a 202 to the initial request, with 'Location' and 'Retry-After' headers, Polls
// return a 200 with a response body after success
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) post202Retry200(ctx context.Context, options *LRORetrysBeginPost202Retry200Options) (*http.Response, error) {
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
func (client *LRORetrysClient) post202Retry200CreateRequest(ctx context.Context, options *LRORetrysBeginPost202Retry200Options) (*policy.Request, error) {
	urlPath := "/lro/retryerror/post/202/retry/200"
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
func (client *LRORetrysClient) post202Retry200HandleError(resp *http.Response) error {
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

// BeginPostAsyncRelativeRetrySucceeded - Long running post request, service returns a 500, then a 202 to the initial request, with an entity that contains
// ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) BeginPostAsyncRelativeRetrySucceeded(ctx context.Context, options *LRORetrysBeginPostAsyncRelativeRetrySucceededOptions) (LRORetrysPostAsyncRelativeRetrySucceededPollerResponse, error) {
	resp, err := client.postAsyncRelativeRetrySucceeded(ctx, options)
	if err != nil {
		return LRORetrysPostAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	result := LRORetrysPostAsyncRelativeRetrySucceededPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LRORetrysClient.PostAsyncRelativeRetrySucceeded", "", resp, client.pl, client.postAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return LRORetrysPostAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	result.Poller = &LRORetrysPostAsyncRelativeRetrySucceededPoller{
		pt: pt,
	}
	return result, nil
}

// PostAsyncRelativeRetrySucceeded - Long running post request, service returns a 500, then a 202 to the initial request, with an entity that contains ProvisioningState=’Creating’.
// Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) postAsyncRelativeRetrySucceeded(ctx context.Context, options *LRORetrysBeginPostAsyncRelativeRetrySucceededOptions) (*http.Response, error) {
	req, err := client.postAsyncRelativeRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, client.postAsyncRelativeRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// postAsyncRelativeRetrySucceededCreateRequest creates the PostAsyncRelativeRetrySucceeded request.
func (client *LRORetrysClient) postAsyncRelativeRetrySucceededCreateRequest(ctx context.Context, options *LRORetrysBeginPostAsyncRelativeRetrySucceededOptions) (*policy.Request, error) {
	urlPath := "/lro/retryerror/postasync/retry/succeeded"
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

// postAsyncRelativeRetrySucceededHandleError handles the PostAsyncRelativeRetrySucceeded error response.
func (client *LRORetrysClient) postAsyncRelativeRetrySucceededHandleError(resp *http.Response) error {
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

// BeginPut201CreatingSucceeded200 - Long running put request, service returns a 500, then a 201 to the initial request, with an entity that contains ProvisioningState=’Creating’.
// Polls return this value until the last poll returns a
// ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) BeginPut201CreatingSucceeded200(ctx context.Context, options *LRORetrysBeginPut201CreatingSucceeded200Options) (LRORetrysPut201CreatingSucceeded200PollerResponse, error) {
	resp, err := client.put201CreatingSucceeded200(ctx, options)
	if err != nil {
		return LRORetrysPut201CreatingSucceeded200PollerResponse{}, err
	}
	result := LRORetrysPut201CreatingSucceeded200PollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LRORetrysClient.Put201CreatingSucceeded200", "", resp, client.pl, client.put201CreatingSucceeded200HandleError)
	if err != nil {
		return LRORetrysPut201CreatingSucceeded200PollerResponse{}, err
	}
	result.Poller = &LRORetrysPut201CreatingSucceeded200Poller{
		pt: pt,
	}
	return result, nil
}

// Put201CreatingSucceeded200 - Long running put request, service returns a 500, then a 201 to the initial request, with an entity that contains ProvisioningState=’Creating’.
// Polls return this value until the last poll returns a
// ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) put201CreatingSucceeded200(ctx context.Context, options *LRORetrysBeginPut201CreatingSucceeded200Options) (*http.Response, error) {
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
func (client *LRORetrysClient) put201CreatingSucceeded200CreateRequest(ctx context.Context, options *LRORetrysBeginPut201CreatingSucceeded200Options) (*policy.Request, error) {
	urlPath := "/lro/retryerror/put/201/creating/succeeded/200"
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
func (client *LRORetrysClient) put201CreatingSucceeded200HandleError(resp *http.Response) error {
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

// BeginPutAsyncRelativeRetrySucceeded - Long running put request, service returns a 500, then a 200 to the initial request, with an entity that contains
// ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) BeginPutAsyncRelativeRetrySucceeded(ctx context.Context, options *LRORetrysBeginPutAsyncRelativeRetrySucceededOptions) (LRORetrysPutAsyncRelativeRetrySucceededPollerResponse, error) {
	resp, err := client.putAsyncRelativeRetrySucceeded(ctx, options)
	if err != nil {
		return LRORetrysPutAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	result := LRORetrysPutAsyncRelativeRetrySucceededPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LRORetrysClient.PutAsyncRelativeRetrySucceeded", "", resp, client.pl, client.putAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return LRORetrysPutAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	result.Poller = &LRORetrysPutAsyncRelativeRetrySucceededPoller{
		pt: pt,
	}
	return result, nil
}

// PutAsyncRelativeRetrySucceeded - Long running put request, service returns a 500, then a 200 to the initial request, with an entity that contains ProvisioningState=’Creating’.
// Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) putAsyncRelativeRetrySucceeded(ctx context.Context, options *LRORetrysBeginPutAsyncRelativeRetrySucceededOptions) (*http.Response, error) {
	req, err := client.putAsyncRelativeRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, client.putAsyncRelativeRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// putAsyncRelativeRetrySucceededCreateRequest creates the PutAsyncRelativeRetrySucceeded request.
func (client *LRORetrysClient) putAsyncRelativeRetrySucceededCreateRequest(ctx context.Context, options *LRORetrysBeginPutAsyncRelativeRetrySucceededOptions) (*policy.Request, error) {
	urlPath := "/lro/retryerror/putasync/retry/succeeded"
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

// putAsyncRelativeRetrySucceededHandleError handles the PutAsyncRelativeRetrySucceeded error response.
func (client *LRORetrysClient) putAsyncRelativeRetrySucceededHandleError(resp *http.Response) error {
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
