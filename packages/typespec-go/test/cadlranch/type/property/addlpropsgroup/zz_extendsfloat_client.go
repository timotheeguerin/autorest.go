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

// ExtendsFloatClient contains the methods for the Type.Property.AdditionalProperties group.
// Don't use this type directly, use a constructor function instead.
type ExtendsFloatClient struct {
	internal *azcore.Client
}

// Get - Get call
func (client *ExtendsFloatClient) Get(ctx context.Context, options *ExtendsFloatClientGetOptions) (ExtendsFloatClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ExtendsFloatClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtendsFloatClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExtendsFloatClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ExtendsFloatClient) getCreateRequest(ctx context.Context, options *ExtendsFloatClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/extendsRecordFloat"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExtendsFloatClient) getHandleResponse(resp *http.Response) (ExtendsFloatClientGetResponse, error) {
	result := ExtendsFloatClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendsFloatAdditionalProperties); err != nil {
		return ExtendsFloatClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
//   - body - body
func (client *ExtendsFloatClient) Put(ctx context.Context, body ExtendsFloatAdditionalProperties, options *ExtendsFloatClientPutOptions) (ExtendsFloatClientPutResponse, error) {
	var err error
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return ExtendsFloatClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtendsFloatClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ExtendsFloatClientPutResponse{}, err
	}
	return ExtendsFloatClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ExtendsFloatClient) putCreateRequest(ctx context.Context, body ExtendsFloatAdditionalProperties, options *ExtendsFloatClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/extendsRecordFloat"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
