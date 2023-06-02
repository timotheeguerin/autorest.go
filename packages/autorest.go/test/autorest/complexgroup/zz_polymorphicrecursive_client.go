//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// PolymorphicrecursiveClient contains the methods for the Polymorphicrecursive group.
// Don't use this type directly, use a constructor function instead.
type PolymorphicrecursiveClient struct {
	internal *azcore.Client
}

// GetValid - Get complex types that are polymorphic and have recursive references
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - options - PolymorphicrecursiveClientGetValidOptions contains the optional parameters for the PolymorphicrecursiveClient.GetValid
//     method.
func (client *PolymorphicrecursiveClient) GetValid(ctx context.Context, options *PolymorphicrecursiveClientGetValidOptions) (PolymorphicrecursiveClientGetValidResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "PolymorphicrecursiveClient.GetValid", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return PolymorphicrecursiveClientGetValidResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PolymorphicrecursiveClientGetValidResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PolymorphicrecursiveClientGetValidResponse{}, err
	}
	resp, err := client.getValidHandleResponse(httpResp)
	return resp, err
}

// getValidCreateRequest creates the GetValid request.
func (client *PolymorphicrecursiveClient) getValidCreateRequest(ctx context.Context, options *PolymorphicrecursiveClientGetValidOptions) (*policy.Request, error) {
	urlPath := "/complex/polymorphicrecursive/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *PolymorphicrecursiveClient) getValidHandleResponse(resp *http.Response) (PolymorphicrecursiveClientGetValidResponse, error) {
	result := PolymorphicrecursiveClientGetValidResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return PolymorphicrecursiveClientGetValidResponse{}, err
	}
	return result, nil
}

// PutValid - Put complex types that are polymorphic and have recursive references
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - complexBody - Please put a salmon that looks like this: { "fishtype": "salmon", "species": "king", "length": 1, "age":
//     1, "location": "alaska", "iswild": true, "siblings": [ { "fishtype": "shark", "species":
//     "predator", "length": 20, "age": 6, "siblings": [ { "fishtype": "salmon", "species": "coho", "length": 2, "age": 2, "location":
//     "atlantic", "iswild": true, "siblings": [ { "fishtype": "shark",
//     "species": "predator", "length": 20, "age": 6 }, { "fishtype": "sawshark", "species": "dangerous", "length": 10, "age":
//     105 } ] }, { "fishtype": "sawshark", "species": "dangerous", "length": 10,
//     "age": 105 } ] }, { "fishtype": "sawshark", "species": "dangerous", "length": 10, "age": 105 } ] }
//   - options - PolymorphicrecursiveClientPutValidOptions contains the optional parameters for the PolymorphicrecursiveClient.PutValid
//     method.
func (client *PolymorphicrecursiveClient) PutValid(ctx context.Context, complexBody FishClassification, options *PolymorphicrecursiveClientPutValidOptions) (PolymorphicrecursiveClientPutValidResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "PolymorphicrecursiveClient.PutValid", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PolymorphicrecursiveClientPutValidResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PolymorphicrecursiveClientPutValidResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PolymorphicrecursiveClientPutValidResponse{}, err
	}
	return PolymorphicrecursiveClientPutValidResponse{}, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *PolymorphicrecursiveClient) putValidCreateRequest(ctx context.Context, complexBody FishClassification, options *PolymorphicrecursiveClientPutValidOptions) (*policy.Request, error) {
	urlPath := "/complex/polymorphicrecursive/valid"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, complexBody); err != nil {
		return nil, err
	}
	return req, nil
}
