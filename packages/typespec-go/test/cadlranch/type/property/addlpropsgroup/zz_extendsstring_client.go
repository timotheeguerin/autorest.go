//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package addlpropsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ExtendsStringClient contains the methods for the Type.Property.AdditionalProperties group.
// Don't use this type directly, use a constructor function instead.
type ExtendsStringClient struct {
	internal *azcore.Client
}

// Get - Get call
func (client *ExtendsStringClient) Get(ctx context.Context, options *ExtendsStringClientGetOptions) (ExtendsStringClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ExtendsStringClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtendsStringClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExtendsStringClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ExtendsStringClient) getCreateRequest(ctx context.Context, options *ExtendsStringClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/extendsRecordString"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExtendsStringClient) getHandleResponse(resp *http.Response) (ExtendsStringClientGetResponse, error) {
	result := ExtendsStringClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendsStringAdditionalProperties); err != nil {
		return ExtendsStringClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
//   - body - body
func (client *ExtendsStringClient) Put(ctx context.Context, body ExtendsStringAdditionalProperties, options *ExtendsStringClientPutOptions) (ExtendsStringClientPutResponse, error) {
	var err error
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return ExtendsStringClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtendsStringClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ExtendsStringClientPutResponse{}, err
	}
	return ExtendsStringClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ExtendsStringClient) putCreateRequest(ctx context.Context, body ExtendsStringAdditionalProperties, options *ExtendsStringClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/extendsRecordString"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
