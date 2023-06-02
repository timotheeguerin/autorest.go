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

// HTTPRetryClient contains the methods for the HTTPRetry group.
// Don't use this type directly, use a constructor function instead.
type HTTPRetryClient struct {
	internal *azcore.Client
}

// Delete503 - Return 503 status code, then 200 after retry
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRetryClientDelete503Options contains the optional parameters for the HTTPRetryClient.Delete503 method.
func (client *HTTPRetryClient) Delete503(ctx context.Context, options *HTTPRetryClientDelete503Options) (HTTPRetryClientDelete503Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPRetryClient.Delete503", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.delete503CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientDelete503Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryClientDelete503Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRetryClientDelete503Response{}, err
	}
	return HTTPRetryClientDelete503Response{}, nil
}

// delete503CreateRequest creates the Delete503 request.
func (client *HTTPRetryClient) delete503CreateRequest(ctx context.Context, options *HTTPRetryClientDelete503Options) (*policy.Request, error) {
	urlPath := "/http/retry/503"
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

// Get502 - Return 502 status code, then 200 after retry
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRetryClientGet502Options contains the optional parameters for the HTTPRetryClient.Get502 method.
func (client *HTTPRetryClient) Get502(ctx context.Context, options *HTTPRetryClientGet502Options) (HTTPRetryClientGet502Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPRetryClient.Get502", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.get502CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientGet502Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryClientGet502Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRetryClientGet502Response{}, err
	}
	return HTTPRetryClientGet502Response{}, nil
}

// get502CreateRequest creates the Get502 request.
func (client *HTTPRetryClient) get502CreateRequest(ctx context.Context, options *HTTPRetryClientGet502Options) (*policy.Request, error) {
	urlPath := "/http/retry/502"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Head408 - Return 408 status code, then 200 after retry
//
// Generated from API version 1.0.0
//   - options - HTTPRetryClientHead408Options contains the optional parameters for the HTTPRetryClient.Head408 method.
func (client *HTTPRetryClient) Head408(ctx context.Context, options *HTTPRetryClientHead408Options) (HTTPRetryClientHead408Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPRetryClient.Head408", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.head408CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientHead408Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryClientHead408Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRetryClientHead408Response{}, err
	}
	return HTTPRetryClientHead408Response{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// head408CreateRequest creates the Head408 request.
func (client *HTTPRetryClient) head408CreateRequest(ctx context.Context, options *HTTPRetryClientHead408Options) (*policy.Request, error) {
	urlPath := "/http/retry/408"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Options502 - Return 502 status code, then 200 after retry
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRetryClientOptions502Options contains the optional parameters for the HTTPRetryClient.Options502 method.
func (client *HTTPRetryClient) Options502(ctx context.Context, options *HTTPRetryClientOptions502Options) (HTTPRetryClientOptions502Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPRetryClient.Options502", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.options502CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientOptions502Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryClientOptions502Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRetryClientOptions502Response{}, err
	}
	resp, err := client.options502HandleResponse(httpResp)
	return resp, err
}

// options502CreateRequest creates the Options502 request.
func (client *HTTPRetryClient) options502CreateRequest(ctx context.Context, options *HTTPRetryClientOptions502Options) (*policy.Request, error) {
	urlPath := "/http/retry/502"
	req, err := runtime.NewRequest(ctx, http.MethodOptions, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// options502HandleResponse handles the Options502 response.
func (client *HTTPRetryClient) options502HandleResponse(resp *http.Response) (HTTPRetryClientOptions502Response, error) {
	result := HTTPRetryClientOptions502Response{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return HTTPRetryClientOptions502Response{}, err
	}
	return result, nil
}

// Patch500 - Return 500 status code, then 200 after retry
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRetryClientPatch500Options contains the optional parameters for the HTTPRetryClient.Patch500 method.
func (client *HTTPRetryClient) Patch500(ctx context.Context, options *HTTPRetryClientPatch500Options) (HTTPRetryClientPatch500Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPRetryClient.Patch500", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.patch500CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientPatch500Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryClientPatch500Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRetryClientPatch500Response{}, err
	}
	return HTTPRetryClientPatch500Response{}, nil
}

// patch500CreateRequest creates the Patch500 request.
func (client *HTTPRetryClient) patch500CreateRequest(ctx context.Context, options *HTTPRetryClientPatch500Options) (*policy.Request, error) {
	urlPath := "/http/retry/500"
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

// Patch504 - Return 504 status code, then 200 after retry
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRetryClientPatch504Options contains the optional parameters for the HTTPRetryClient.Patch504 method.
func (client *HTTPRetryClient) Patch504(ctx context.Context, options *HTTPRetryClientPatch504Options) (HTTPRetryClientPatch504Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPRetryClient.Patch504", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.patch504CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientPatch504Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryClientPatch504Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRetryClientPatch504Response{}, err
	}
	return HTTPRetryClientPatch504Response{}, nil
}

// patch504CreateRequest creates the Patch504 request.
func (client *HTTPRetryClient) patch504CreateRequest(ctx context.Context, options *HTTPRetryClientPatch504Options) (*policy.Request, error) {
	urlPath := "/http/retry/504"
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

// Post503 - Return 503 status code, then 200 after retry
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRetryClientPost503Options contains the optional parameters for the HTTPRetryClient.Post503 method.
func (client *HTTPRetryClient) Post503(ctx context.Context, options *HTTPRetryClientPost503Options) (HTTPRetryClientPost503Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPRetryClient.Post503", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.post503CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientPost503Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryClientPost503Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRetryClientPost503Response{}, err
	}
	return HTTPRetryClientPost503Response{}, nil
}

// post503CreateRequest creates the Post503 request.
func (client *HTTPRetryClient) post503CreateRequest(ctx context.Context, options *HTTPRetryClientPost503Options) (*policy.Request, error) {
	urlPath := "/http/retry/503"
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

// Put500 - Return 500 status code, then 200 after retry
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRetryClientPut500Options contains the optional parameters for the HTTPRetryClient.Put500 method.
func (client *HTTPRetryClient) Put500(ctx context.Context, options *HTTPRetryClientPut500Options) (HTTPRetryClientPut500Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPRetryClient.Put500", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.put500CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientPut500Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryClientPut500Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRetryClientPut500Response{}, err
	}
	return HTTPRetryClientPut500Response{}, nil
}

// put500CreateRequest creates the Put500 request.
func (client *HTTPRetryClient) put500CreateRequest(ctx context.Context, options *HTTPRetryClientPut500Options) (*policy.Request, error) {
	urlPath := "/http/retry/500"
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

// Put504 - Return 504 status code, then 200 after retry
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRetryClientPut504Options contains the optional parameters for the HTTPRetryClient.Put504 method.
func (client *HTTPRetryClient) Put504(ctx context.Context, options *HTTPRetryClientPut504Options) (HTTPRetryClientPut504Response, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "HTTPRetryClient.Put504", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.put504CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryClientPut504Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryClientPut504Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRetryClientPut504Response{}, err
	}
	return HTTPRetryClientPut504Response{}, nil
}

// put504CreateRequest creates the Put504 request.
func (client *HTTPRetryClient) put504CreateRequest(ctx context.Context, options *HTTPRetryClientPut504Options) (*policy.Request, error) {
	urlPath := "/http/retry/504"
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
