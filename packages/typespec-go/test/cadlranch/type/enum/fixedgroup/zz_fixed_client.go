//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fixedgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// FixedClient contains the methods for the Type.Enum.Fixed group.
// Don't use this type directly, use a constructor function instead.
type FixedClient struct {
	internal *azcore.Client
}

// GetKnownValue - getKnownValue
func (client *FixedClient) GetKnownValue(ctx context.Context, options *FixedClientGetKnownValueOptions) (FixedClientGetKnownValueResponse, error) {
	var err error
	req, err := client.getKnownValueCreateRequest(ctx, options)
	if err != nil {
		return FixedClientGetKnownValueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FixedClientGetKnownValueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FixedClientGetKnownValueResponse{}, err
	}
	resp, err := client.getKnownValueHandleResponse(httpResp)
	return resp, err
}

// getKnownValueCreateRequest creates the GetKnownValue request.
func (client *FixedClient) getKnownValueCreateRequest(ctx context.Context, options *FixedClientGetKnownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/fixed/string/known-value"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// getKnownValueHandleResponse handles the GetKnownValue response.
func (client *FixedClient) getKnownValueHandleResponse(resp *http.Response) (FixedClientGetKnownValueResponse, error) {
	result := FixedClientGetKnownValueResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return FixedClientGetKnownValueResponse{}, err
	}
	return result, nil
}

// PutKnownValue - putKnownValue
//   - body - _
func (client *FixedClient) PutKnownValue(ctx context.Context, body DaysOfWeekEnum, options *FixedClientPutKnownValueOptions) (FixedClientPutKnownValueResponse, error) {
	var err error
	req, err := client.putKnownValueCreateRequest(ctx, body, options)
	if err != nil {
		return FixedClientPutKnownValueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FixedClientPutKnownValueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return FixedClientPutKnownValueResponse{}, err
	}
	return FixedClientPutKnownValueResponse{}, nil
}

// putKnownValueCreateRequest creates the PutKnownValue request.
func (client *FixedClient) putKnownValueCreateRequest(ctx context.Context, body DaysOfWeekEnum, options *FixedClientPutKnownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/fixed/string/known-value"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// PutUnknownValue - putUnknownValue
//   - body - _
func (client *FixedClient) PutUnknownValue(ctx context.Context, body DaysOfWeekEnum, options *FixedClientPutUnknownValueOptions) (FixedClientPutUnknownValueResponse, error) {
	var err error
	req, err := client.putUnknownValueCreateRequest(ctx, body, options)
	if err != nil {
		return FixedClientPutUnknownValueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FixedClientPutUnknownValueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return FixedClientPutUnknownValueResponse{}, err
	}
	return FixedClientPutUnknownValueResponse{}, nil
}

// putUnknownValueCreateRequest creates the PutUnknownValue request.
func (client *FixedClient) putUnknownValueCreateRequest(ctx context.Context, body DaysOfWeekEnum, options *FixedClientPutUnknownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/fixed/string/unknown-value"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
