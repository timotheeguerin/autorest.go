//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package morecustombaseurigroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PathsClient contains the methods for the Paths group.
// Don't use this type directly, use NewPathsClient() instead.
type PathsClient struct {
	dnsSuffix      string
	subscriptionID string
	pl             runtime.Pipeline
}

// PathsClientOptions contains the optional parameters for NewPathsClient.
type PathsClientOptions struct {
	azcore.ClientOptions
	DnsSuffix *string
}

// NewPathsClient creates a new instance of PathsClient with the specified values.
func NewPathsClient(subscriptionID string, options *PathsClientOptions) *PathsClient {
	cp := PathsClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &PathsClient{
		dnsSuffix:      "host",
		subscriptionID: subscriptionID,
		pl:             runtime.NewPipeline(module, version, nil, nil, &cp.ClientOptions),
	}
	if options.DnsSuffix != nil {
		client.dnsSuffix = *options.DnsSuffix
	}
	return client
}

// GetEmpty - Get a 200 to test a valid base uri
// If the operation fails it returns the *Error error type.
func (client *PathsClient) GetEmpty(ctx context.Context, vault string, secret string, keyName string, options *PathsGetEmptyOptions) (PathsGetEmptyResponse, error) {
	req, err := client.getEmptyCreateRequest(ctx, vault, secret, keyName, options)
	if err != nil {
		return PathsGetEmptyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PathsGetEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PathsGetEmptyResponse{}, client.getEmptyHandleError(resp)
	}
	return PathsGetEmptyResponse{RawResponse: resp}, nil
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *PathsClient) getEmptyCreateRequest(ctx context.Context, vault string, secret string, keyName string, options *PathsGetEmptyOptions) (*policy.Request, error) {
	host := "{vault}{secret}{dnsSuffix}"
	host = strings.ReplaceAll(host, "{dnsSuffix}", client.dnsSuffix)
	host = strings.ReplaceAll(host, "{vault}", vault)
	host = strings.ReplaceAll(host, "{secret}", secret)
	urlPath := "/customuri/{subscriptionId}/{keyName}"
	if keyName == "" {
		return nil, errors.New("parameter keyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyName}", url.PathEscape(keyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.KeyVersion != nil {
		reqQP.Set("keyVersion", *options.KeyVersion)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEmptyHandleError handles the GetEmpty error response.
func (client *PathsClient) getEmptyHandleError(resp *http.Response) error {
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
