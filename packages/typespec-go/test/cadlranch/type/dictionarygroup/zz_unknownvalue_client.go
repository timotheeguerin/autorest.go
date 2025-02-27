//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package dictionarygroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// UnknownValueClient contains the methods for the Type.Dictionary group.
// Don't use this type directly, use a constructor function instead.
type UnknownValueClient struct {
	internal *azcore.Client
}

func (client *UnknownValueClient) Get(ctx context.Context, options *UnknownValueClientGetOptions) (UnknownValueClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return UnknownValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UnknownValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return UnknownValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *UnknownValueClient) getCreateRequest(ctx context.Context, options *UnknownValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/dictionary/unknown"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *UnknownValueClient) getHandleResponse(resp *http.Response) (UnknownValueClientGetResponse, error) {
	result := UnknownValueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return UnknownValueClientGetResponse{}, err
	}
	return result, nil
}

func (client *UnknownValueClient) Put(ctx context.Context, body map[string]any, options *UnknownValueClientPutOptions) (UnknownValueClientPutResponse, error) {
	var err error
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return UnknownValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UnknownValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return UnknownValueClientPutResponse{}, err
	}
	return UnknownValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *UnknownValueClient) putCreateRequest(ctx context.Context, body map[string]any, options *UnknownValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/dictionary/unknown"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
