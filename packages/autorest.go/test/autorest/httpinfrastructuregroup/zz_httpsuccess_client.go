//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package httpinfrastructuregroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// HTTPSuccessClient contains the methods for the HTTPSuccess group.
// Don't use this type directly, use a constructor function instead.
type HTTPSuccessClient struct {
	internal *azcore.Client
}

// Delete200 - Delete simple boolean value true returns 200
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientDelete200Options contains the optional parameters for the HTTPSuccessClient.Delete200 method.
func (client *HTTPSuccessClient) Delete200(ctx context.Context, options *HTTPSuccessClientDelete200Options) (HTTPSuccessClientDelete200Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Delete200", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.delete200CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientDelete200Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientDelete200Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientDelete200Response{}, err
	}
	return HTTPSuccessClientDelete200Response{}, nil
}

// delete200CreateRequest creates the Delete200 request.
func (client *HTTPSuccessClient) delete200CreateRequest(ctx context.Context, options *HTTPSuccessClientDelete200Options) (*policy.Request, error) {
	urlPath := "/http/success/200"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete202 - Delete true Boolean value in request returns 202 (accepted)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientDelete202Options contains the optional parameters for the HTTPSuccessClient.Delete202 method.
func (client *HTTPSuccessClient) Delete202(ctx context.Context, options *HTTPSuccessClientDelete202Options) (HTTPSuccessClientDelete202Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Delete202", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.delete202CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientDelete202Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientDelete202Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientDelete202Response{}, err
	}
	return HTTPSuccessClientDelete202Response{}, nil
}

// delete202CreateRequest creates the Delete202 request.
func (client *HTTPSuccessClient) delete202CreateRequest(ctx context.Context, options *HTTPSuccessClientDelete202Options) (*policy.Request, error) {
	urlPath := "/http/success/202"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete204 - Delete true Boolean value in request returns 204 (no content)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientDelete204Options contains the optional parameters for the HTTPSuccessClient.Delete204 method.
func (client *HTTPSuccessClient) Delete204(ctx context.Context, options *HTTPSuccessClientDelete204Options) (HTTPSuccessClientDelete204Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Delete204", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.delete204CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientDelete204Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientDelete204Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientDelete204Response{}, err
	}
	return HTTPSuccessClientDelete204Response{}, nil
}

// delete204CreateRequest creates the Delete204 request.
func (client *HTTPSuccessClient) delete204CreateRequest(ctx context.Context, options *HTTPSuccessClientDelete204Options) (*policy.Request, error) {
	urlPath := "/http/success/204"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Get200 - Get 200 success
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientGet200Options contains the optional parameters for the HTTPSuccessClient.Get200 method.
func (client *HTTPSuccessClient) Get200(ctx context.Context, options *HTTPSuccessClientGet200Options) (HTTPSuccessClientGet200Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Get200", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.get200CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientGet200Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientGet200Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientGet200Response{}, err
	}
	resp, err := client.get200HandleResponse(httpResp)
	return resp, err
}

// get200CreateRequest creates the Get200 request.
func (client *HTTPSuccessClient) get200CreateRequest(ctx context.Context, options *HTTPSuccessClientGet200Options) (*policy.Request, error) {
	urlPath := "/http/success/200"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// get200HandleResponse handles the Get200 response.
func (client *HTTPSuccessClient) get200HandleResponse(resp *http.Response) (HTTPSuccessClientGet200Response, error) {
	result := HTTPSuccessClientGet200Response{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return HTTPSuccessClientGet200Response{}, err
	}
	return result, nil
}

// Head200 - Return 200 status code if successful
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientHead200Options contains the optional parameters for the HTTPSuccessClient.Head200 method.
func (client *HTTPSuccessClient) Head200(ctx context.Context, options *HTTPSuccessClientHead200Options) (HTTPSuccessClientHead200Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Head200", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.head200CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientHead200Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientHead200Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientHead200Response{}, err
	}
	return HTTPSuccessClientHead200Response{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// head200CreateRequest creates the Head200 request.
func (client *HTTPSuccessClient) head200CreateRequest(ctx context.Context, options *HTTPSuccessClientHead200Options) (*policy.Request, error) {
	urlPath := "/http/success/200"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Head204 - Return 204 status code if successful
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientHead204Options contains the optional parameters for the HTTPSuccessClient.Head204 method.
func (client *HTTPSuccessClient) Head204(ctx context.Context, options *HTTPSuccessClientHead204Options) (HTTPSuccessClientHead204Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Head204", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.head204CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientHead204Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientHead204Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientHead204Response{}, err
	}
	return HTTPSuccessClientHead204Response{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// head204CreateRequest creates the Head204 request.
func (client *HTTPSuccessClient) head204CreateRequest(ctx context.Context, options *HTTPSuccessClientHead204Options) (*policy.Request, error) {
	urlPath := "/http/success/204"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Head404 - Return 404 status code
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientHead404Options contains the optional parameters for the HTTPSuccessClient.Head404 method.
func (client *HTTPSuccessClient) Head404(ctx context.Context, options *HTTPSuccessClientHead404Options) (HTTPSuccessClientHead404Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Head404", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.head404CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientHead404Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientHead404Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientHead404Response{}, err
	}
	return HTTPSuccessClientHead404Response{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// head404CreateRequest creates the Head404 request.
func (client *HTTPSuccessClient) head404CreateRequest(ctx context.Context, options *HTTPSuccessClientHead404Options) (*policy.Request, error) {
	urlPath := "/http/success/404"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Options200 - Options 200 success
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientOptions200Options contains the optional parameters for the HTTPSuccessClient.Options200 method.
func (client *HTTPSuccessClient) Options200(ctx context.Context, options *HTTPSuccessClientOptions200Options) (HTTPSuccessClientOptions200Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Options200", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.options200CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientOptions200Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientOptions200Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientOptions200Response{}, err
	}
	resp, err := client.options200HandleResponse(httpResp)
	return resp, err
}

// options200CreateRequest creates the Options200 request.
func (client *HTTPSuccessClient) options200CreateRequest(ctx context.Context, options *HTTPSuccessClientOptions200Options) (*policy.Request, error) {
	urlPath := "/http/success/200"
	req, err := runtime.NewRequest(ctx, http.MethodOptions, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// options200HandleResponse handles the Options200 response.
func (client *HTTPSuccessClient) options200HandleResponse(resp *http.Response) (HTTPSuccessClientOptions200Response, error) {
	result := HTTPSuccessClientOptions200Response{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return HTTPSuccessClientOptions200Response{}, err
	}
	return result, nil
}

// Patch200 - Patch true Boolean value in request returning 200
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientPatch200Options contains the optional parameters for the HTTPSuccessClient.Patch200 method.
func (client *HTTPSuccessClient) Patch200(ctx context.Context, options *HTTPSuccessClientPatch200Options) (HTTPSuccessClientPatch200Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Patch200", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.patch200CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPatch200Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientPatch200Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientPatch200Response{}, err
	}
	return HTTPSuccessClientPatch200Response{}, nil
}

// patch200CreateRequest creates the Patch200 request.
func (client *HTTPSuccessClient) patch200CreateRequest(ctx context.Context, options *HTTPSuccessClientPatch200Options) (*policy.Request, error) {
	urlPath := "/http/success/200"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Patch202 - Patch true Boolean value in request returns 202
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientPatch202Options contains the optional parameters for the HTTPSuccessClient.Patch202 method.
func (client *HTTPSuccessClient) Patch202(ctx context.Context, options *HTTPSuccessClientPatch202Options) (HTTPSuccessClientPatch202Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Patch202", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.patch202CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPatch202Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientPatch202Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientPatch202Response{}, err
	}
	return HTTPSuccessClientPatch202Response{}, nil
}

// patch202CreateRequest creates the Patch202 request.
func (client *HTTPSuccessClient) patch202CreateRequest(ctx context.Context, options *HTTPSuccessClientPatch202Options) (*policy.Request, error) {
	urlPath := "/http/success/202"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Patch204 - Patch true Boolean value in request returns 204 (no content)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientPatch204Options contains the optional parameters for the HTTPSuccessClient.Patch204 method.
func (client *HTTPSuccessClient) Patch204(ctx context.Context, options *HTTPSuccessClientPatch204Options) (HTTPSuccessClientPatch204Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Patch204", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.patch204CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPatch204Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientPatch204Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientPatch204Response{}, err
	}
	return HTTPSuccessClientPatch204Response{}, nil
}

// patch204CreateRequest creates the Patch204 request.
func (client *HTTPSuccessClient) patch204CreateRequest(ctx context.Context, options *HTTPSuccessClientPatch204Options) (*policy.Request, error) {
	urlPath := "/http/success/204"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Post200 - Post bollean value true in request that returns a 200
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientPost200Options contains the optional parameters for the HTTPSuccessClient.Post200 method.
func (client *HTTPSuccessClient) Post200(ctx context.Context, options *HTTPSuccessClientPost200Options) (HTTPSuccessClientPost200Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Post200", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.post200CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPost200Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientPost200Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientPost200Response{}, err
	}
	return HTTPSuccessClientPost200Response{}, nil
}

// post200CreateRequest creates the Post200 request.
func (client *HTTPSuccessClient) post200CreateRequest(ctx context.Context, options *HTTPSuccessClientPost200Options) (*policy.Request, error) {
	urlPath := "/http/success/200"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Post201 - Post true Boolean value in request returns 201 (Created)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientPost201Options contains the optional parameters for the HTTPSuccessClient.Post201 method.
func (client *HTTPSuccessClient) Post201(ctx context.Context, options *HTTPSuccessClientPost201Options) (HTTPSuccessClientPost201Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Post201", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.post201CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPost201Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientPost201Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientPost201Response{}, err
	}
	return HTTPSuccessClientPost201Response{}, nil
}

// post201CreateRequest creates the Post201 request.
func (client *HTTPSuccessClient) post201CreateRequest(ctx context.Context, options *HTTPSuccessClientPost201Options) (*policy.Request, error) {
	urlPath := "/http/success/201"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Post202 - Post true Boolean value in request returns 202 (Accepted)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientPost202Options contains the optional parameters for the HTTPSuccessClient.Post202 method.
func (client *HTTPSuccessClient) Post202(ctx context.Context, options *HTTPSuccessClientPost202Options) (HTTPSuccessClientPost202Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Post202", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.post202CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPost202Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientPost202Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientPost202Response{}, err
	}
	return HTTPSuccessClientPost202Response{}, nil
}

// post202CreateRequest creates the Post202 request.
func (client *HTTPSuccessClient) post202CreateRequest(ctx context.Context, options *HTTPSuccessClientPost202Options) (*policy.Request, error) {
	urlPath := "/http/success/202"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Post204 - Post true Boolean value in request returns 204 (no content)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientPost204Options contains the optional parameters for the HTTPSuccessClient.Post204 method.
func (client *HTTPSuccessClient) Post204(ctx context.Context, options *HTTPSuccessClientPost204Options) (HTTPSuccessClientPost204Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Post204", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.post204CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPost204Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientPost204Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientPost204Response{}, err
	}
	return HTTPSuccessClientPost204Response{}, nil
}

// post204CreateRequest creates the Post204 request.
func (client *HTTPSuccessClient) post204CreateRequest(ctx context.Context, options *HTTPSuccessClientPost204Options) (*policy.Request, error) {
	urlPath := "/http/success/204"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Put200 - Put boolean value true returning 200 success
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientPut200Options contains the optional parameters for the HTTPSuccessClient.Put200 method.
func (client *HTTPSuccessClient) Put200(ctx context.Context, options *HTTPSuccessClientPut200Options) (HTTPSuccessClientPut200Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Put200", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.put200CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPut200Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientPut200Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientPut200Response{}, err
	}
	return HTTPSuccessClientPut200Response{}, nil
}

// put200CreateRequest creates the Put200 request.
func (client *HTTPSuccessClient) put200CreateRequest(ctx context.Context, options *HTTPSuccessClientPut200Options) (*policy.Request, error) {
	urlPath := "/http/success/200"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Put201 - Put true Boolean value in request returns 201
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientPut201Options contains the optional parameters for the HTTPSuccessClient.Put201 method.
func (client *HTTPSuccessClient) Put201(ctx context.Context, options *HTTPSuccessClientPut201Options) (HTTPSuccessClientPut201Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Put201", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.put201CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPut201Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientPut201Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientPut201Response{}, err
	}
	return HTTPSuccessClientPut201Response{}, nil
}

// put201CreateRequest creates the Put201 request.
func (client *HTTPSuccessClient) put201CreateRequest(ctx context.Context, options *HTTPSuccessClientPut201Options) (*policy.Request, error) {
	urlPath := "/http/success/201"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Put202 - Put true Boolean value in request returns 202 (Accepted)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientPut202Options contains the optional parameters for the HTTPSuccessClient.Put202 method.
func (client *HTTPSuccessClient) Put202(ctx context.Context, options *HTTPSuccessClientPut202Options) (HTTPSuccessClientPut202Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Put202", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.put202CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPut202Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientPut202Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientPut202Response{}, err
	}
	return HTTPSuccessClientPut202Response{}, nil
}

// put202CreateRequest creates the Put202 request.
func (client *HTTPSuccessClient) put202CreateRequest(ctx context.Context, options *HTTPSuccessClientPut202Options) (*policy.Request, error) {
	urlPath := "/http/success/202"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Put204 - Put true Boolean value in request returns 204 (no content)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientPut204Options contains the optional parameters for the HTTPSuccessClient.Put204 method.
func (client *HTTPSuccessClient) Put204(ctx context.Context, options *HTTPSuccessClientPut204Options) (HTTPSuccessClientPut204Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPSuccessClient.Put204", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.put204CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientPut204Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientPut204Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientPut204Response{}, err
	}
	return HTTPSuccessClientPut204Response{}, nil
}

// put204CreateRequest creates the Put204 request.
func (client *HTTPSuccessClient) put204CreateRequest(ctx context.Context, options *HTTPSuccessClientPut204Options) (*policy.Request, error) {
	urlPath := "/http/success/204"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}
