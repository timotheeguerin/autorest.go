//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package accessgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// InternalOperationClient contains the methods for the _Specs_.Azure.ClientGenerator.Core.Access group.
// Don't use this type directly, use a constructor function instead.
type InternalOperationClient struct {
	internal *azcore.Client
}

func (client *InternalOperationClient) InternalDecoratorInInternal(ctx context.Context, name string, options *InternalOperationClientInternalDecoratorInInternalOptions) (InternalOperationClientInternalDecoratorInInternalResponse, error) {
	var err error
	req, err := client.internalDecoratorInInternalCreateRequest(ctx, name, options)
	if err != nil {
		return InternalOperationClientInternalDecoratorInInternalResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InternalOperationClientInternalDecoratorInInternalResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InternalOperationClientInternalDecoratorInInternalResponse{}, err
	}
	resp, err := client.internalDecoratorInInternalHandleResponse(httpResp)
	return resp, err
}

// internalDecoratorInInternalCreateRequest creates the InternalDecoratorInInternal request.
func (client *InternalOperationClient) internalDecoratorInInternalCreateRequest(ctx context.Context, name string, options *InternalOperationClientInternalDecoratorInInternalOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/access/internalOperation/internalDecoratorInInternal"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("name", name)
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// internalDecoratorInInternalHandleResponse handles the InternalDecoratorInInternal response.
func (client *InternalOperationClient) internalDecoratorInInternalHandleResponse(resp *http.Response) (InternalOperationClientInternalDecoratorInInternalResponse, error) {
	result := InternalOperationClientInternalDecoratorInInternalResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InternalDecoratorModelInInternal); err != nil {
		return InternalOperationClientInternalDecoratorInInternalResponse{}, err
	}
	return result, nil
}

func (client *InternalOperationClient) NoDecoratorInInternal(ctx context.Context, name string, options *InternalOperationClientNoDecoratorInInternalOptions) (InternalOperationClientNoDecoratorInInternalResponse, error) {
	var err error
	req, err := client.noDecoratorInInternalCreateRequest(ctx, name, options)
	if err != nil {
		return InternalOperationClientNoDecoratorInInternalResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InternalOperationClientNoDecoratorInInternalResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InternalOperationClientNoDecoratorInInternalResponse{}, err
	}
	resp, err := client.noDecoratorInInternalHandleResponse(httpResp)
	return resp, err
}

// noDecoratorInInternalCreateRequest creates the NoDecoratorInInternal request.
func (client *InternalOperationClient) noDecoratorInInternalCreateRequest(ctx context.Context, name string, options *InternalOperationClientNoDecoratorInInternalOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/access/internalOperation/noDecoratorInInternal"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("name", name)
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// noDecoratorInInternalHandleResponse handles the NoDecoratorInInternal response.
func (client *InternalOperationClient) noDecoratorInInternalHandleResponse(resp *http.Response) (InternalOperationClientNoDecoratorInInternalResponse, error) {
	result := InternalOperationClientNoDecoratorInInternalResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NoDecoratorModelInInternal); err != nil {
		return InternalOperationClientNoDecoratorInInternalResponse{}, err
	}
	return result, nil
}

func (client *InternalOperationClient) PublicDecoratorInInternal(ctx context.Context, name string, options *InternalOperationClientPublicDecoratorInInternalOptions) (InternalOperationClientPublicDecoratorInInternalResponse, error) {
	var err error
	req, err := client.publicDecoratorInInternalCreateRequest(ctx, name, options)
	if err != nil {
		return InternalOperationClientPublicDecoratorInInternalResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InternalOperationClientPublicDecoratorInInternalResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InternalOperationClientPublicDecoratorInInternalResponse{}, err
	}
	resp, err := client.publicDecoratorInInternalHandleResponse(httpResp)
	return resp, err
}

// publicDecoratorInInternalCreateRequest creates the PublicDecoratorInInternal request.
func (client *InternalOperationClient) publicDecoratorInInternalCreateRequest(ctx context.Context, name string, options *InternalOperationClientPublicDecoratorInInternalOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/access/internalOperation/publicDecoratorInInternal"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("name", name)
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// publicDecoratorInInternalHandleResponse handles the PublicDecoratorInInternal response.
func (client *InternalOperationClient) publicDecoratorInInternalHandleResponse(resp *http.Response) (InternalOperationClientPublicDecoratorInInternalResponse, error) {
	result := InternalOperationClientPublicDecoratorInInternalResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublicDecoratorModelInInternal); err != nil {
		return InternalOperationClientPublicDecoratorInInternalResponse{}, err
	}
	return result, nil
}
