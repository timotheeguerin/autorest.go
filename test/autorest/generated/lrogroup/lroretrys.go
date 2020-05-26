// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package lrogroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// LroRetrysOperations contains the methods for the LroRetrys group.
type LroRetrysOperations interface {
	// BeginDelete202Retry200 - Long running delete request, service returns a 500, then a 202 to the initial request. Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
	BeginDelete202Retry200(ctx context.Context) (*HTTPResponse, error)
	// ResumeDelete202Retry200 - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete202Retry200(token string) (HTTPPoller, error)
	// BeginDeleteAsyncRelativeRetrySucceeded - Long running delete request, service returns a 500, then a 202 to the initial request. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
	BeginDeleteAsyncRelativeRetrySucceeded(ctx context.Context) (*HTTPResponse, error)
	// ResumeDeleteAsyncRelativeRetrySucceeded - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDeleteAsyncRelativeRetrySucceeded(token string) (HTTPPoller, error)
	// BeginDeleteProvisioning202Accepted200Succeeded - Long running delete request, service returns a 500, then a  202 to the initial request, with an entity that contains ProvisioningState=’Accepted’.  Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
	BeginDeleteProvisioning202Accepted200Succeeded(ctx context.Context) (*ProductResponse, error)
	// ResumeDeleteProvisioning202Accepted200Succeeded - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDeleteProvisioning202Accepted200Succeeded(token string) (ProductPoller, error)
	// BeginPost202Retry200 - Long running post request, service returns a 500, then a 202 to the initial request, with 'Location' and 'Retry-After' headers, Polls return a 200 with a response body after success
	BeginPost202Retry200(ctx context.Context, lroRetrysPost202Retry200Options *LroRetrysPost202Retry200Options) (*HTTPResponse, error)
	// ResumePost202Retry200 - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePost202Retry200(token string) (HTTPPoller, error)
	// BeginPostAsyncRelativeRetrySucceeded - Long running post request, service returns a 500, then a 202 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
	BeginPostAsyncRelativeRetrySucceeded(ctx context.Context, lroRetrysPostAsyncRelativeRetrySucceededOptions *LroRetrysPostAsyncRelativeRetrySucceededOptions) (*HTTPResponse, error)
	// ResumePostAsyncRelativeRetrySucceeded - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePostAsyncRelativeRetrySucceeded(token string) (HTTPPoller, error)
	// BeginPut201CreatingSucceeded200 - Long running put request, service returns a 500, then a 201 to the initial request, with an entity that contains ProvisioningState=’Creating’.  Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
	BeginPut201CreatingSucceeded200(ctx context.Context, lroRetrysPut201CreatingSucceeded200Options *LroRetrysPut201CreatingSucceeded200Options) (*ProductResponse, error)
	// ResumePut201CreatingSucceeded200 - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePut201CreatingSucceeded200(token string) (ProductPoller, error)
	// BeginPutAsyncRelativeRetrySucceeded - Long running put request, service returns a 500, then a 200 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
	BeginPutAsyncRelativeRetrySucceeded(ctx context.Context, lroRetrysPutAsyncRelativeRetrySucceededOptions *LroRetrysPutAsyncRelativeRetrySucceededOptions) (*ProductResponse, error)
	// ResumePutAsyncRelativeRetrySucceeded - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePutAsyncRelativeRetrySucceeded(token string) (ProductPoller, error)
}

// lroRetrysOperations implements the LroRetrysOperations interface.
type lroRetrysOperations struct {
	*Client
}

// Delete202Retry200 - Long running delete request, service returns a 500, then a 202 to the initial request. Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client *lroRetrysOperations) BeginDelete202Retry200(ctx context.Context) (*HTTPResponse, error) {
	req, err := client.delete202Retry200CreateRequest()
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.delete202Retry200HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("lroRetrysOperations.Delete202Retry200", "", resp, client.delete202Retry200HandleError)
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

func (client *lroRetrysOperations) ResumeDelete202Retry200(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("lroRetrysOperations.Delete202Retry200", token, client.delete202Retry200HandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// delete202Retry200CreateRequest creates the Delete202Retry200 request.
func (client *lroRetrysOperations) delete202Retry200CreateRequest() (*azcore.Request, error) {
	urlPath := "/lro/retryerror/delete/202/retry/200"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// delete202Retry200HandleResponse handles the Delete202Retry200 response.
func (client *lroRetrysOperations) delete202Retry200HandleResponse(resp *azcore.Response) (*HTTPResponse, error) {
	if !resp.HasStatusCode(http.StatusAccepted, http.StatusNoContent) {
		return nil, client.delete202Retry200HandleError(resp)
	}
	return &HTTPResponse{RawResponse: resp.Response}, nil
}

// delete202Retry200HandleError handles the Delete202Retry200 error response.
func (client *lroRetrysOperations) delete202Retry200HandleError(resp *azcore.Response) error {
	err := CloudError{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// DeleteAsyncRelativeRetrySucceeded - Long running delete request, service returns a 500, then a 202 to the initial request. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client *lroRetrysOperations) BeginDeleteAsyncRelativeRetrySucceeded(ctx context.Context) (*HTTPResponse, error) {
	req, err := client.deleteAsyncRelativeRetrySucceededCreateRequest()
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteAsyncRelativeRetrySucceededHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("lroRetrysOperations.DeleteAsyncRelativeRetrySucceeded", "", resp, client.deleteAsyncRelativeRetrySucceededHandleError)
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

func (client *lroRetrysOperations) ResumeDeleteAsyncRelativeRetrySucceeded(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("lroRetrysOperations.DeleteAsyncRelativeRetrySucceeded", token, client.deleteAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteAsyncRelativeRetrySucceededCreateRequest creates the DeleteAsyncRelativeRetrySucceeded request.
func (client *lroRetrysOperations) deleteAsyncRelativeRetrySucceededCreateRequest() (*azcore.Request, error) {
	urlPath := "/lro/retryerror/deleteasync/retry/succeeded"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteAsyncRelativeRetrySucceededHandleResponse handles the DeleteAsyncRelativeRetrySucceeded response.
func (client *lroRetrysOperations) deleteAsyncRelativeRetrySucceededHandleResponse(resp *azcore.Response) (*HTTPResponse, error) {
	if !resp.HasStatusCode(http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteAsyncRelativeRetrySucceededHandleError(resp)
	}
	return &HTTPResponse{RawResponse: resp.Response}, nil
}

// deleteAsyncRelativeRetrySucceededHandleError handles the DeleteAsyncRelativeRetrySucceeded error response.
func (client *lroRetrysOperations) deleteAsyncRelativeRetrySucceededHandleError(resp *azcore.Response) error {
	err := CloudError{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// DeleteProvisioning202Accepted200Succeeded - Long running delete request, service returns a 500, then a  202 to the initial request, with an entity that contains ProvisioningState=’Accepted’.  Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client *lroRetrysOperations) BeginDeleteProvisioning202Accepted200Succeeded(ctx context.Context) (*ProductResponse, error) {
	req, err := client.deleteProvisioning202Accepted200SucceededCreateRequest()
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteProvisioning202Accepted200SucceededHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("lroRetrysOperations.DeleteProvisioning202Accepted200Succeeded", "", resp, client.deleteProvisioning202Accepted200SucceededHandleError)
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

func (client *lroRetrysOperations) ResumeDeleteProvisioning202Accepted200Succeeded(token string) (ProductPoller, error) {
	pt, err := resumePollingTracker("lroRetrysOperations.DeleteProvisioning202Accepted200Succeeded", token, client.deleteProvisioning202Accepted200SucceededHandleError)
	if err != nil {
		return nil, err
	}
	return &productPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteProvisioning202Accepted200SucceededCreateRequest creates the DeleteProvisioning202Accepted200Succeeded request.
func (client *lroRetrysOperations) deleteProvisioning202Accepted200SucceededCreateRequest() (*azcore.Request, error) {
	urlPath := "/lro/retryerror/delete/provisioning/202/accepted/200/succeeded"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteProvisioning202Accepted200SucceededHandleResponse handles the DeleteProvisioning202Accepted200Succeeded response.
func (client *lroRetrysOperations) deleteProvisioning202Accepted200SucceededHandleResponse(resp *azcore.Response) (*ProductResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteProvisioning202Accepted200SucceededHandleError(resp)
	}
	result := ProductResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Product)
}

// deleteProvisioning202Accepted200SucceededHandleError handles the DeleteProvisioning202Accepted200Succeeded error response.
func (client *lroRetrysOperations) deleteProvisioning202Accepted200SucceededHandleError(resp *azcore.Response) error {
	err := CloudError{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Post202Retry200 - Long running post request, service returns a 500, then a 202 to the initial request, with 'Location' and 'Retry-After' headers, Polls return a 200 with a response body after success
func (client *lroRetrysOperations) BeginPost202Retry200(ctx context.Context, lroRetrysPost202Retry200Options *LroRetrysPost202Retry200Options) (*HTTPResponse, error) {
	req, err := client.post202Retry200CreateRequest(lroRetrysPost202Retry200Options)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.post202Retry200HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("lroRetrysOperations.Post202Retry200", "", resp, client.post202Retry200HandleError)
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

func (client *lroRetrysOperations) ResumePost202Retry200(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("lroRetrysOperations.Post202Retry200", token, client.post202Retry200HandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// post202Retry200CreateRequest creates the Post202Retry200 request.
func (client *lroRetrysOperations) post202Retry200CreateRequest(lroRetrysPost202Retry200Options *LroRetrysPost202Retry200Options) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/post/202/retry/200"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	if lroRetrysPost202Retry200Options != nil {
		return req, req.MarshalAsJSON(lroRetrysPost202Retry200Options.Product)
	}
	return req, nil
}

// post202Retry200HandleResponse handles the Post202Retry200 response.
func (client *lroRetrysOperations) post202Retry200HandleResponse(resp *azcore.Response) (*HTTPResponse, error) {
	if !resp.HasStatusCode(http.StatusAccepted, http.StatusNoContent) {
		return nil, client.post202Retry200HandleError(resp)
	}
	return &HTTPResponse{RawResponse: resp.Response}, nil
}

// post202Retry200HandleError handles the Post202Retry200 error response.
func (client *lroRetrysOperations) post202Retry200HandleError(resp *azcore.Response) error {
	err := CloudError{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PostAsyncRelativeRetrySucceeded - Long running post request, service returns a 500, then a 202 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client *lroRetrysOperations) BeginPostAsyncRelativeRetrySucceeded(ctx context.Context, lroRetrysPostAsyncRelativeRetrySucceededOptions *LroRetrysPostAsyncRelativeRetrySucceededOptions) (*HTTPResponse, error) {
	req, err := client.postAsyncRelativeRetrySucceededCreateRequest(lroRetrysPostAsyncRelativeRetrySucceededOptions)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.postAsyncRelativeRetrySucceededHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("lroRetrysOperations.PostAsyncRelativeRetrySucceeded", "", resp, client.postAsyncRelativeRetrySucceededHandleError)
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

func (client *lroRetrysOperations) ResumePostAsyncRelativeRetrySucceeded(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("lroRetrysOperations.PostAsyncRelativeRetrySucceeded", token, client.postAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// postAsyncRelativeRetrySucceededCreateRequest creates the PostAsyncRelativeRetrySucceeded request.
func (client *lroRetrysOperations) postAsyncRelativeRetrySucceededCreateRequest(lroRetrysPostAsyncRelativeRetrySucceededOptions *LroRetrysPostAsyncRelativeRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/postasync/retry/succeeded"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	if lroRetrysPostAsyncRelativeRetrySucceededOptions != nil {
		return req, req.MarshalAsJSON(lroRetrysPostAsyncRelativeRetrySucceededOptions.Product)
	}
	return req, nil
}

// postAsyncRelativeRetrySucceededHandleResponse handles the PostAsyncRelativeRetrySucceeded response.
func (client *lroRetrysOperations) postAsyncRelativeRetrySucceededHandleResponse(resp *azcore.Response) (*HTTPResponse, error) {
	if !resp.HasStatusCode(http.StatusAccepted, http.StatusNoContent) {
		return nil, client.postAsyncRelativeRetrySucceededHandleError(resp)
	}
	return &HTTPResponse{RawResponse: resp.Response}, nil
}

// postAsyncRelativeRetrySucceededHandleError handles the PostAsyncRelativeRetrySucceeded error response.
func (client *lroRetrysOperations) postAsyncRelativeRetrySucceededHandleError(resp *azcore.Response) error {
	err := CloudError{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Put201CreatingSucceeded200 - Long running put request, service returns a 500, then a 201 to the initial request, with an entity that contains ProvisioningState=’Creating’.  Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client *lroRetrysOperations) BeginPut201CreatingSucceeded200(ctx context.Context, lroRetrysPut201CreatingSucceeded200Options *LroRetrysPut201CreatingSucceeded200Options) (*ProductResponse, error) {
	req, err := client.put201CreatingSucceeded200CreateRequest(lroRetrysPut201CreatingSucceeded200Options)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.put201CreatingSucceeded200HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("lroRetrysOperations.Put201CreatingSucceeded200", "", resp, client.put201CreatingSucceeded200HandleError)
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

func (client *lroRetrysOperations) ResumePut201CreatingSucceeded200(token string) (ProductPoller, error) {
	pt, err := resumePollingTracker("lroRetrysOperations.Put201CreatingSucceeded200", token, client.put201CreatingSucceeded200HandleError)
	if err != nil {
		return nil, err
	}
	return &productPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// put201CreatingSucceeded200CreateRequest creates the Put201CreatingSucceeded200 request.
func (client *lroRetrysOperations) put201CreatingSucceeded200CreateRequest(lroRetrysPut201CreatingSucceeded200Options *LroRetrysPut201CreatingSucceeded200Options) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/put/201/creating/succeeded/200"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	if lroRetrysPut201CreatingSucceeded200Options != nil {
		return req, req.MarshalAsJSON(lroRetrysPut201CreatingSucceeded200Options.Product)
	}
	return req, nil
}

// put201CreatingSucceeded200HandleResponse handles the Put201CreatingSucceeded200 response.
func (client *lroRetrysOperations) put201CreatingSucceeded200HandleResponse(resp *azcore.Response) (*ProductResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.put201CreatingSucceeded200HandleError(resp)
	}
	result := ProductResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Product)
}

// put201CreatingSucceeded200HandleError handles the Put201CreatingSucceeded200 error response.
func (client *lroRetrysOperations) put201CreatingSucceeded200HandleError(resp *azcore.Response) error {
	err := CloudError{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutAsyncRelativeRetrySucceeded - Long running put request, service returns a 500, then a 200 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client *lroRetrysOperations) BeginPutAsyncRelativeRetrySucceeded(ctx context.Context, lroRetrysPutAsyncRelativeRetrySucceededOptions *LroRetrysPutAsyncRelativeRetrySucceededOptions) (*ProductResponse, error) {
	req, err := client.putAsyncRelativeRetrySucceededCreateRequest(lroRetrysPutAsyncRelativeRetrySucceededOptions)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putAsyncRelativeRetrySucceededHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("lroRetrysOperations.PutAsyncRelativeRetrySucceeded", "", resp, client.putAsyncRelativeRetrySucceededHandleError)
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

func (client *lroRetrysOperations) ResumePutAsyncRelativeRetrySucceeded(token string) (ProductPoller, error) {
	pt, err := resumePollingTracker("lroRetrysOperations.PutAsyncRelativeRetrySucceeded", token, client.putAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	return &productPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// putAsyncRelativeRetrySucceededCreateRequest creates the PutAsyncRelativeRetrySucceeded request.
func (client *lroRetrysOperations) putAsyncRelativeRetrySucceededCreateRequest(lroRetrysPutAsyncRelativeRetrySucceededOptions *LroRetrysPutAsyncRelativeRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/putasync/retry/succeeded"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	if lroRetrysPutAsyncRelativeRetrySucceededOptions != nil {
		return req, req.MarshalAsJSON(lroRetrysPutAsyncRelativeRetrySucceededOptions.Product)
	}
	return req, nil
}

// putAsyncRelativeRetrySucceededHandleResponse handles the PutAsyncRelativeRetrySucceeded response.
func (client *lroRetrysOperations) putAsyncRelativeRetrySucceededHandleResponse(resp *azcore.Response) (*ProductResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.putAsyncRelativeRetrySucceededHandleError(resp)
	}
	result := ProductResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Product)
}

// putAsyncRelativeRetrySucceededHandleError handles the PutAsyncRelativeRetrySucceeded error response.
func (client *lroRetrysOperations) putAsyncRelativeRetrySucceededHandleError(resp *azcore.Response) error {
	err := CloudError{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
