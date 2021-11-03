//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package urlmultigroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// QueriesClient contains the methods for the Queries group.
// Don't use this type directly, use NewQueriesClient() instead.
type QueriesClient struct {
	pl runtime.Pipeline
}

// NewQueriesClient creates a new instance of QueriesClient with the specified values.
func NewQueriesClient(options *azcore.ClientOptions) *QueriesClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &QueriesClient{
		pl: runtime.NewPipeline(module, version, nil, nil, &cp),
	}
	return client
}

// ArrayStringMultiEmpty - Get an empty array [] of string using the multi-array format
// If the operation fails it returns the *Error error type.
func (client *QueriesClient) ArrayStringMultiEmpty(ctx context.Context, options *QueriesArrayStringMultiEmptyOptions) (QueriesArrayStringMultiEmptyResponse, error) {
	req, err := client.arrayStringMultiEmptyCreateRequest(ctx, options)
	if err != nil {
		return QueriesArrayStringMultiEmptyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueriesArrayStringMultiEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueriesArrayStringMultiEmptyResponse{}, client.arrayStringMultiEmptyHandleError(resp)
	}
	return QueriesArrayStringMultiEmptyResponse{RawResponse: resp}, nil
}

// arrayStringMultiEmptyCreateRequest creates the ArrayStringMultiEmpty request.
func (client *QueriesClient) arrayStringMultiEmptyCreateRequest(ctx context.Context, options *QueriesArrayStringMultiEmptyOptions) (*policy.Request, error) {
	urlPath := "/queries/array/multi/string/empty"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.ArrayQuery != nil {
		for _, qv := range options.ArrayQuery {
			reqQP.Add("arrayQuery", qv)
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// arrayStringMultiEmptyHandleError handles the ArrayStringMultiEmpty error response.
func (client *QueriesClient) arrayStringMultiEmptyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ArrayStringMultiNull - Get a null array of string using the multi-array format
// If the operation fails it returns the *Error error type.
func (client *QueriesClient) ArrayStringMultiNull(ctx context.Context, options *QueriesArrayStringMultiNullOptions) (QueriesArrayStringMultiNullResponse, error) {
	req, err := client.arrayStringMultiNullCreateRequest(ctx, options)
	if err != nil {
		return QueriesArrayStringMultiNullResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueriesArrayStringMultiNullResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueriesArrayStringMultiNullResponse{}, client.arrayStringMultiNullHandleError(resp)
	}
	return QueriesArrayStringMultiNullResponse{RawResponse: resp}, nil
}

// arrayStringMultiNullCreateRequest creates the ArrayStringMultiNull request.
func (client *QueriesClient) arrayStringMultiNullCreateRequest(ctx context.Context, options *QueriesArrayStringMultiNullOptions) (*policy.Request, error) {
	urlPath := "/queries/array/multi/string/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.ArrayQuery != nil {
		for _, qv := range options.ArrayQuery {
			reqQP.Add("arrayQuery", qv)
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// arrayStringMultiNullHandleError handles the ArrayStringMultiNull error response.
func (client *QueriesClient) arrayStringMultiNullHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ArrayStringMultiValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the mult-array format
// If the operation fails it returns the *Error error type.
func (client *QueriesClient) ArrayStringMultiValid(ctx context.Context, options *QueriesArrayStringMultiValidOptions) (QueriesArrayStringMultiValidResponse, error) {
	req, err := client.arrayStringMultiValidCreateRequest(ctx, options)
	if err != nil {
		return QueriesArrayStringMultiValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueriesArrayStringMultiValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueriesArrayStringMultiValidResponse{}, client.arrayStringMultiValidHandleError(resp)
	}
	return QueriesArrayStringMultiValidResponse{RawResponse: resp}, nil
}

// arrayStringMultiValidCreateRequest creates the ArrayStringMultiValid request.
func (client *QueriesClient) arrayStringMultiValidCreateRequest(ctx context.Context, options *QueriesArrayStringMultiValidOptions) (*policy.Request, error) {
	urlPath := "/queries/array/multi/string/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.ArrayQuery != nil {
		for _, qv := range options.ArrayQuery {
			reqQP.Add("arrayQuery", qv)
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// arrayStringMultiValidHandleError handles the ArrayStringMultiValid error response.
func (client *QueriesClient) arrayStringMultiValidHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
