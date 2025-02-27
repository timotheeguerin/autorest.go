//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package condreqgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ConditionalRequestClient contains the methods for the SpecialHeaders.ConditionalRequest group.
// Don't use this type directly, use a constructor function instead.
type ConditionalRequestClient struct {
	internal *azcore.Client
}

// PostIfMatch - Check when only If-Match in header is defined.
func (client *ConditionalRequestClient) PostIfMatch(ctx context.Context, options *ConditionalRequestClientPostIfMatchOptions) (ConditionalRequestClientPostIfMatchResponse, error) {
	var err error
	req, err := client.postIfMatchCreateRequest(ctx, options)
	if err != nil {
		return ConditionalRequestClientPostIfMatchResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConditionalRequestClientPostIfMatchResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConditionalRequestClientPostIfMatchResponse{}, err
	}
	return ConditionalRequestClientPostIfMatchResponse{}, nil
}

// postIfMatchCreateRequest creates the PostIfMatch request.
func (client *ConditionalRequestClient) postIfMatchCreateRequest(ctx context.Context, options *ConditionalRequestClientPostIfMatchOptions) (*policy.Request, error) {
	urlPath := "/special-headers/conditional-request/if-match"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	return req, nil
}

// PostIfNoneMatch - Check when only If-None-Match in header is defined.
func (client *ConditionalRequestClient) PostIfNoneMatch(ctx context.Context, options *ConditionalRequestClientPostIfNoneMatchOptions) (ConditionalRequestClientPostIfNoneMatchResponse, error) {
	var err error
	req, err := client.postIfNoneMatchCreateRequest(ctx, options)
	if err != nil {
		return ConditionalRequestClientPostIfNoneMatchResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConditionalRequestClientPostIfNoneMatchResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConditionalRequestClientPostIfNoneMatchResponse{}, err
	}
	return ConditionalRequestClientPostIfNoneMatchResponse{}, nil
}

// postIfNoneMatchCreateRequest creates the PostIfNoneMatch request.
func (client *ConditionalRequestClient) postIfNoneMatchCreateRequest(ctx context.Context, options *ConditionalRequestClientPostIfNoneMatchOptions) (*policy.Request, error) {
	urlPath := "/special-headers/conditional-request/if-none-match"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	return req, nil
}
