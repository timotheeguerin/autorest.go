// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package lrogroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// LroRetrysOperations contains the methods for the LroRetrys group.
type LroRetrysOperations interface {
	// BeginDelete202Retry200 - Long running delete request, service returns a 500, then a 202 to the initial request. Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
	BeginDelete202Retry200(ctx context.Context) (*HTTPPollerResponse, error)
	// ResumeDelete202Retry200 - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete202Retry200(token string) (HTTPPoller, error)
	// BeginDeleteAsyncRelativeRetrySucceeded - Long running delete request, service returns a 500, then a 202 to the initial request. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
	BeginDeleteAsyncRelativeRetrySucceeded(ctx context.Context) (*HTTPPollerResponse, error)
	// ResumeDeleteAsyncRelativeRetrySucceeded - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDeleteAsyncRelativeRetrySucceeded(token string) (HTTPPoller, error)
	// BeginDeleteProvisioning202Accepted200Succeeded - Long running delete request, service returns a 500, then a  202 to the initial request, with an entity that contains ProvisioningState=’Accepted’.  Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
	BeginDeleteProvisioning202Accepted200Succeeded(ctx context.Context) (*ProductPollerResponse, error)
	// ResumeDeleteProvisioning202Accepted200Succeeded - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDeleteProvisioning202Accepted200Succeeded(token string) (ProductPoller, error)
	// BeginPost202Retry200 - Long running post request, service returns a 500, then a 202 to the initial request, with 'Location' and 'Retry-After' headers, Polls return a 200 with a response body after success
	BeginPost202Retry200(ctx context.Context, lroRetrysPost202Retry200Options *LroRetrysPost202Retry200Options) (*HTTPPollerResponse, error)
	// ResumePost202Retry200 - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePost202Retry200(token string) (HTTPPoller, error)
	// BeginPostAsyncRelativeRetrySucceeded - Long running post request, service returns a 500, then a 202 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
	BeginPostAsyncRelativeRetrySucceeded(ctx context.Context, lroRetrysPostAsyncRelativeRetrySucceededOptions *LroRetrysPostAsyncRelativeRetrySucceededOptions) (*HTTPPollerResponse, error)
	// ResumePostAsyncRelativeRetrySucceeded - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePostAsyncRelativeRetrySucceeded(token string) (HTTPPoller, error)
	// BeginPut201CreatingSucceeded200 - Long running put request, service returns a 500, then a 201 to the initial request, with an entity that contains ProvisioningState=’Creating’.  Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
	BeginPut201CreatingSucceeded200(ctx context.Context, lroRetrysPut201CreatingSucceeded200Options *LroRetrysPut201CreatingSucceeded200Options) (*ProductPollerResponse, error)
	// ResumePut201CreatingSucceeded200 - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePut201CreatingSucceeded200(token string) (ProductPoller, error)
	// BeginPutAsyncRelativeRetrySucceeded - Long running put request, service returns a 500, then a 200 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
	BeginPutAsyncRelativeRetrySucceeded(ctx context.Context, lroRetrysPutAsyncRelativeRetrySucceededOptions *LroRetrysPutAsyncRelativeRetrySucceededOptions) (*ProductPollerResponse, error)
	// ResumePutAsyncRelativeRetrySucceeded - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePutAsyncRelativeRetrySucceeded(token string) (ProductPoller, error)
}

// LroRetrysClient implements the LroRetrysOperations interface.
// Don't use this type directly, use NewLroRetrysClient() instead.
type LroRetrysClient struct {
	*Client
}

// NewLroRetrysClient creates a new instance of LroRetrysClient with the specified values.
func NewLroRetrysClient(c *Client) LroRetrysOperations {
	return &LroRetrysClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *LroRetrysClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Delete202Retry200 - Long running delete request, service returns a 500, then a 202 to the initial request. Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client *LroRetrysClient) BeginDelete202Retry200(ctx context.Context) (*HTTPPollerResponse, error) {
	req, err := client.Delete202Retry200CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.Delete202Retry200HandleError(resp)
	}
	result, err := client.Delete202Retry200HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("LroRetrysClient.Delete202Retry200", "", resp, client.Delete202Retry200HandleError)
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

func (client *LroRetrysClient) ResumeDelete202Retry200(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LroRetrysClient.Delete202Retry200", token, client.Delete202Retry200HandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Delete202Retry200CreateRequest creates the Delete202Retry200 request.
func (client *LroRetrysClient) Delete202Retry200CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/delete/202/retry/200"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Delete202Retry200HandleResponse handles the Delete202Retry200 response.
func (client *LroRetrysClient) Delete202Retry200HandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// Delete202Retry200HandleError handles the Delete202Retry200 error response.
func (client *LroRetrysClient) Delete202Retry200HandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// DeleteAsyncRelativeRetrySucceeded - Long running delete request, service returns a 500, then a 202 to the initial request. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client *LroRetrysClient) BeginDeleteAsyncRelativeRetrySucceeded(ctx context.Context) (*HTTPPollerResponse, error) {
	req, err := client.DeleteAsyncRelativeRetrySucceededCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.DeleteAsyncRelativeRetrySucceededHandleError(resp)
	}
	result, err := client.DeleteAsyncRelativeRetrySucceededHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("LroRetrysClient.DeleteAsyncRelativeRetrySucceeded", "", resp, client.DeleteAsyncRelativeRetrySucceededHandleError)
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

func (client *LroRetrysClient) ResumeDeleteAsyncRelativeRetrySucceeded(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LroRetrysClient.DeleteAsyncRelativeRetrySucceeded", token, client.DeleteAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// DeleteAsyncRelativeRetrySucceededCreateRequest creates the DeleteAsyncRelativeRetrySucceeded request.
func (client *LroRetrysClient) DeleteAsyncRelativeRetrySucceededCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/deleteasync/retry/succeeded"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteAsyncRelativeRetrySucceededHandleResponse handles the DeleteAsyncRelativeRetrySucceeded response.
func (client *LroRetrysClient) DeleteAsyncRelativeRetrySucceededHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// DeleteAsyncRelativeRetrySucceededHandleError handles the DeleteAsyncRelativeRetrySucceeded error response.
func (client *LroRetrysClient) DeleteAsyncRelativeRetrySucceededHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// DeleteProvisioning202Accepted200Succeeded - Long running delete request, service returns a 500, then a  202 to the initial request, with an entity that contains ProvisioningState=’Accepted’.  Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client *LroRetrysClient) BeginDeleteProvisioning202Accepted200Succeeded(ctx context.Context) (*ProductPollerResponse, error) {
	req, err := client.DeleteProvisioning202Accepted200SucceededCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.DeleteProvisioning202Accepted200SucceededHandleError(resp)
	}
	result, err := client.DeleteProvisioning202Accepted200SucceededHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("LroRetrysClient.DeleteProvisioning202Accepted200Succeeded", "", resp, client.DeleteProvisioning202Accepted200SucceededHandleError)
	if err != nil {
		return nil, err
	}
	poller := &productPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ProductResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *LroRetrysClient) ResumeDeleteProvisioning202Accepted200Succeeded(token string) (ProductPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LroRetrysClient.DeleteProvisioning202Accepted200Succeeded", token, client.DeleteProvisioning202Accepted200SucceededHandleError)
	if err != nil {
		return nil, err
	}
	return &productPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// DeleteProvisioning202Accepted200SucceededCreateRequest creates the DeleteProvisioning202Accepted200Succeeded request.
func (client *LroRetrysClient) DeleteProvisioning202Accepted200SucceededCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/delete/provisioning/202/accepted/200/succeeded"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteProvisioning202Accepted200SucceededHandleResponse handles the DeleteProvisioning202Accepted200Succeeded response.
func (client *LroRetrysClient) DeleteProvisioning202Accepted200SucceededHandleResponse(resp *azcore.Response) (*ProductPollerResponse, error) {
	return &ProductPollerResponse{RawResponse: resp.Response}, nil
}

// DeleteProvisioning202Accepted200SucceededHandleError handles the DeleteProvisioning202Accepted200Succeeded error response.
func (client *LroRetrysClient) DeleteProvisioning202Accepted200SucceededHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Post202Retry200 - Long running post request, service returns a 500, then a 202 to the initial request, with 'Location' and 'Retry-After' headers, Polls return a 200 with a response body after success
func (client *LroRetrysClient) BeginPost202Retry200(ctx context.Context, lroRetrysPost202Retry200Options *LroRetrysPost202Retry200Options) (*HTTPPollerResponse, error) {
	req, err := client.Post202Retry200CreateRequest(ctx, lroRetrysPost202Retry200Options)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.Post202Retry200HandleError(resp)
	}
	result, err := client.Post202Retry200HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("LroRetrysClient.Post202Retry200", "", resp, client.Post202Retry200HandleError)
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

func (client *LroRetrysClient) ResumePost202Retry200(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LroRetrysClient.Post202Retry200", token, client.Post202Retry200HandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Post202Retry200CreateRequest creates the Post202Retry200 request.
func (client *LroRetrysClient) Post202Retry200CreateRequest(ctx context.Context, lroRetrysPost202Retry200Options *LroRetrysPost202Retry200Options) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/post/202/retry/200"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if lroRetrysPost202Retry200Options != nil {
		return req, req.MarshalAsJSON(lroRetrysPost202Retry200Options.Product)
	}
	return req, nil
}

// Post202Retry200HandleResponse handles the Post202Retry200 response.
func (client *LroRetrysClient) Post202Retry200HandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// Post202Retry200HandleError handles the Post202Retry200 error response.
func (client *LroRetrysClient) Post202Retry200HandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PostAsyncRelativeRetrySucceeded - Long running post request, service returns a 500, then a 202 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client *LroRetrysClient) BeginPostAsyncRelativeRetrySucceeded(ctx context.Context, lroRetrysPostAsyncRelativeRetrySucceededOptions *LroRetrysPostAsyncRelativeRetrySucceededOptions) (*HTTPPollerResponse, error) {
	req, err := client.PostAsyncRelativeRetrySucceededCreateRequest(ctx, lroRetrysPostAsyncRelativeRetrySucceededOptions)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.PostAsyncRelativeRetrySucceededHandleError(resp)
	}
	result, err := client.PostAsyncRelativeRetrySucceededHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("LroRetrysClient.PostAsyncRelativeRetrySucceeded", "", resp, client.PostAsyncRelativeRetrySucceededHandleError)
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

func (client *LroRetrysClient) ResumePostAsyncRelativeRetrySucceeded(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LroRetrysClient.PostAsyncRelativeRetrySucceeded", token, client.PostAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// PostAsyncRelativeRetrySucceededCreateRequest creates the PostAsyncRelativeRetrySucceeded request.
func (client *LroRetrysClient) PostAsyncRelativeRetrySucceededCreateRequest(ctx context.Context, lroRetrysPostAsyncRelativeRetrySucceededOptions *LroRetrysPostAsyncRelativeRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/postasync/retry/succeeded"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if lroRetrysPostAsyncRelativeRetrySucceededOptions != nil {
		return req, req.MarshalAsJSON(lroRetrysPostAsyncRelativeRetrySucceededOptions.Product)
	}
	return req, nil
}

// PostAsyncRelativeRetrySucceededHandleResponse handles the PostAsyncRelativeRetrySucceeded response.
func (client *LroRetrysClient) PostAsyncRelativeRetrySucceededHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// PostAsyncRelativeRetrySucceededHandleError handles the PostAsyncRelativeRetrySucceeded error response.
func (client *LroRetrysClient) PostAsyncRelativeRetrySucceededHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Put201CreatingSucceeded200 - Long running put request, service returns a 500, then a 201 to the initial request, with an entity that contains ProvisioningState=’Creating’.  Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client *LroRetrysClient) BeginPut201CreatingSucceeded200(ctx context.Context, lroRetrysPut201CreatingSucceeded200Options *LroRetrysPut201CreatingSucceeded200Options) (*ProductPollerResponse, error) {
	req, err := client.Put201CreatingSucceeded200CreateRequest(ctx, lroRetrysPut201CreatingSucceeded200Options)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.Put201CreatingSucceeded200HandleError(resp)
	}
	result, err := client.Put201CreatingSucceeded200HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("LroRetrysClient.Put201CreatingSucceeded200", "", resp, client.Put201CreatingSucceeded200HandleError)
	if err != nil {
		return nil, err
	}
	poller := &productPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ProductResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *LroRetrysClient) ResumePut201CreatingSucceeded200(token string) (ProductPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LroRetrysClient.Put201CreatingSucceeded200", token, client.Put201CreatingSucceeded200HandleError)
	if err != nil {
		return nil, err
	}
	return &productPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Put201CreatingSucceeded200CreateRequest creates the Put201CreatingSucceeded200 request.
func (client *LroRetrysClient) Put201CreatingSucceeded200CreateRequest(ctx context.Context, lroRetrysPut201CreatingSucceeded200Options *LroRetrysPut201CreatingSucceeded200Options) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/put/201/creating/succeeded/200"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if lroRetrysPut201CreatingSucceeded200Options != nil {
		return req, req.MarshalAsJSON(lroRetrysPut201CreatingSucceeded200Options.Product)
	}
	return req, nil
}

// Put201CreatingSucceeded200HandleResponse handles the Put201CreatingSucceeded200 response.
func (client *LroRetrysClient) Put201CreatingSucceeded200HandleResponse(resp *azcore.Response) (*ProductPollerResponse, error) {
	return &ProductPollerResponse{RawResponse: resp.Response}, nil
}

// Put201CreatingSucceeded200HandleError handles the Put201CreatingSucceeded200 error response.
func (client *LroRetrysClient) Put201CreatingSucceeded200HandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutAsyncRelativeRetrySucceeded - Long running put request, service returns a 500, then a 200 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client *LroRetrysClient) BeginPutAsyncRelativeRetrySucceeded(ctx context.Context, lroRetrysPutAsyncRelativeRetrySucceededOptions *LroRetrysPutAsyncRelativeRetrySucceededOptions) (*ProductPollerResponse, error) {
	req, err := client.PutAsyncRelativeRetrySucceededCreateRequest(ctx, lroRetrysPutAsyncRelativeRetrySucceededOptions)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutAsyncRelativeRetrySucceededHandleError(resp)
	}
	result, err := client.PutAsyncRelativeRetrySucceededHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("LroRetrysClient.PutAsyncRelativeRetrySucceeded", "", resp, client.PutAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	poller := &productPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ProductResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *LroRetrysClient) ResumePutAsyncRelativeRetrySucceeded(token string) (ProductPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LroRetrysClient.PutAsyncRelativeRetrySucceeded", token, client.PutAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	return &productPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// PutAsyncRelativeRetrySucceededCreateRequest creates the PutAsyncRelativeRetrySucceeded request.
func (client *LroRetrysClient) PutAsyncRelativeRetrySucceededCreateRequest(ctx context.Context, lroRetrysPutAsyncRelativeRetrySucceededOptions *LroRetrysPutAsyncRelativeRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/putasync/retry/succeeded"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if lroRetrysPutAsyncRelativeRetrySucceededOptions != nil {
		return req, req.MarshalAsJSON(lroRetrysPutAsyncRelativeRetrySucceededOptions.Product)
	}
	return req, nil
}

// PutAsyncRelativeRetrySucceededHandleResponse handles the PutAsyncRelativeRetrySucceeded response.
func (client *LroRetrysClient) PutAsyncRelativeRetrySucceededHandleResponse(resp *azcore.Response) (*ProductPollerResponse, error) {
	return &ProductPollerResponse{RawResponse: resp.Response}, nil
}

// PutAsyncRelativeRetrySucceededHandleError handles the PutAsyncRelativeRetrySucceeded error response.
func (client *LroRetrysClient) PutAsyncRelativeRetrySucceededHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
