//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package mediatypesgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"net/http"
	"strings"
)

// MediaTypesClient contains the methods for the MediaTypesClient group.
// Don't use this type directly, use NewMediaTypesClient() instead.
type MediaTypesClient struct {
	pl runtime.Pipeline
}

// NewMediaTypesClient creates a new instance of MediaTypesClient with the specified values.
func NewMediaTypesClient(options *azcore.ClientOptions) *MediaTypesClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &MediaTypesClient{
		pl: runtime.NewPipeline(module, version, nil, nil, &cp),
	}
	return client
}

// AnalyzeBody - Analyze body, that could be different media types.
// If the operation fails it returns a generic error.
func (client *MediaTypesClient) AnalyzeBody(ctx context.Context, contentType ContentType, options *MediaTypesClientAnalyzeBodyOptions) (MediaTypesClientAnalyzeBodyResponse, error) {
	req, err := client.analyzeBodyCreateRequest(ctx, contentType, options)
	if err != nil {
		return MediaTypesClientAnalyzeBodyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MediaTypesClientAnalyzeBodyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MediaTypesClientAnalyzeBodyResponse{}, client.analyzeBodyHandleError(resp)
	}
	return client.analyzeBodyHandleResponse(resp)
}

// analyzeBodyCreateRequest creates the AnalyzeBody request.
func (client *MediaTypesClient) analyzeBodyCreateRequest(ctx context.Context, contentType ContentType, options *MediaTypesClientAnalyzeBodyOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/analyze"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Content-Type", string(contentType))
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Input != nil {
		return req, req.SetBody(options.Input, string(contentType))
	}
	return req, nil
}

// analyzeBodyHandleResponse handles the AnalyzeBody response.
func (client *MediaTypesClient) analyzeBodyHandleResponse(resp *http.Response) (MediaTypesClientAnalyzeBodyResponse, error) {
	result := MediaTypesClientAnalyzeBodyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return MediaTypesClientAnalyzeBodyResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// analyzeBodyHandleError handles the AnalyzeBody error response.
func (client *MediaTypesClient) analyzeBodyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// AnalyzeBodyNoAcceptHeader - Analyze body, that could be different media types. Adds to AnalyzeBody by not having an accept type.
// If the operation fails it returns a generic error.
func (client *MediaTypesClient) AnalyzeBodyNoAcceptHeader(ctx context.Context, contentType ContentType, options *MediaTypesClientAnalyzeBodyNoAcceptHeaderOptions) (MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse, error) {
	req, err := client.analyzeBodyNoAcceptHeaderCreateRequest(ctx, contentType, options)
	if err != nil {
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse{}, client.analyzeBodyNoAcceptHeaderHandleError(resp)
	}
	return MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse{RawResponse: resp}, nil
}

// analyzeBodyNoAcceptHeaderCreateRequest creates the AnalyzeBodyNoAcceptHeader request.
func (client *MediaTypesClient) analyzeBodyNoAcceptHeaderCreateRequest(ctx context.Context, contentType ContentType, options *MediaTypesClientAnalyzeBodyNoAcceptHeaderOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/analyzeNoAccept"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Content-Type", string(contentType))
	if options != nil && options.Input != nil {
		return req, req.SetBody(options.Input, string(contentType))
	}
	return req, nil
}

// analyzeBodyNoAcceptHeaderHandleError handles the AnalyzeBodyNoAcceptHeader error response.
func (client *MediaTypesClient) analyzeBodyNoAcceptHeaderHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// AnalyzeBodyNoAcceptHeaderWithSourcePath - Analyze body, that could be different media types. Adds to AnalyzeBody by not having an accept type.
// If the operation fails it returns a generic error.
func (client *MediaTypesClient) AnalyzeBodyNoAcceptHeaderWithSourcePath(ctx context.Context, options *MediaTypesClientAnalyzeBodyNoAcceptHeaderWithSourcePathOptions) (MediaTypesClientAnalyzeBodyNoAcceptHeaderWithSourcePathResponse, error) {
	req, err := client.analyzeBodyNoAcceptHeaderWithSourcePathCreateRequest(ctx, options)
	if err != nil {
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderWithSourcePathResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderWithSourcePathResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderWithSourcePathResponse{}, client.analyzeBodyNoAcceptHeaderWithSourcePathHandleError(resp)
	}
	return MediaTypesClientAnalyzeBodyNoAcceptHeaderWithSourcePathResponse{RawResponse: resp}, nil
}

// analyzeBodyNoAcceptHeaderWithSourcePathCreateRequest creates the AnalyzeBodyNoAcceptHeaderWithSourcePath request.
func (client *MediaTypesClient) analyzeBodyNoAcceptHeaderWithSourcePathCreateRequest(ctx context.Context, options *MediaTypesClientAnalyzeBodyNoAcceptHeaderWithSourcePathOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/analyzeNoAccept"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.Input != nil {
		return req, runtime.MarshalAsJSON(req, *options.Input)
	}
	return req, nil
}

// analyzeBodyNoAcceptHeaderWithSourcePathHandleError handles the AnalyzeBodyNoAcceptHeaderWithSourcePath error response.
func (client *MediaTypesClient) analyzeBodyNoAcceptHeaderWithSourcePathHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// AnalyzeBodyWithSourcePath - Analyze body, that could be different media types.
// If the operation fails it returns a generic error.
func (client *MediaTypesClient) AnalyzeBodyWithSourcePath(ctx context.Context, options *MediaTypesClientAnalyzeBodyWithSourcePathOptions) (MediaTypesClientAnalyzeBodyWithSourcePathResponse, error) {
	req, err := client.analyzeBodyWithSourcePathCreateRequest(ctx, options)
	if err != nil {
		return MediaTypesClientAnalyzeBodyWithSourcePathResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MediaTypesClientAnalyzeBodyWithSourcePathResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MediaTypesClientAnalyzeBodyWithSourcePathResponse{}, client.analyzeBodyWithSourcePathHandleError(resp)
	}
	return client.analyzeBodyWithSourcePathHandleResponse(resp)
}

// analyzeBodyWithSourcePathCreateRequest creates the AnalyzeBodyWithSourcePath request.
func (client *MediaTypesClient) analyzeBodyWithSourcePathCreateRequest(ctx context.Context, options *MediaTypesClientAnalyzeBodyWithSourcePathOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/analyze"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Input != nil {
		return req, runtime.MarshalAsJSON(req, *options.Input)
	}
	return req, nil
}

// analyzeBodyWithSourcePathHandleResponse handles the AnalyzeBodyWithSourcePath response.
func (client *MediaTypesClient) analyzeBodyWithSourcePathHandleResponse(resp *http.Response) (MediaTypesClientAnalyzeBodyWithSourcePathResponse, error) {
	result := MediaTypesClientAnalyzeBodyWithSourcePathResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return MediaTypesClientAnalyzeBodyWithSourcePathResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// analyzeBodyWithSourcePathHandleError handles the AnalyzeBodyWithSourcePath error response.
func (client *MediaTypesClient) analyzeBodyWithSourcePathHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ContentTypeWithEncoding - Pass in contentType 'text/plain; encoding=UTF-8' to pass test. Value for input does not matter
// If the operation fails it returns a generic error.
func (client *MediaTypesClient) ContentTypeWithEncoding(ctx context.Context, options *MediaTypesClientContentTypeWithEncodingOptions) (MediaTypesClientContentTypeWithEncodingResponse, error) {
	req, err := client.contentTypeWithEncodingCreateRequest(ctx, options)
	if err != nil {
		return MediaTypesClientContentTypeWithEncodingResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MediaTypesClientContentTypeWithEncodingResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MediaTypesClientContentTypeWithEncodingResponse{}, client.contentTypeWithEncodingHandleError(resp)
	}
	return client.contentTypeWithEncodingHandleResponse(resp)
}

// contentTypeWithEncodingCreateRequest creates the ContentTypeWithEncoding request.
func (client *MediaTypesClient) contentTypeWithEncodingCreateRequest(ctx context.Context, options *MediaTypesClientContentTypeWithEncodingOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/contentTypeWithEncoding"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Input != nil {
		body := streaming.NopCloser(strings.NewReader(*options.Input))
		return req, req.SetBody(body, "text/plain; encoding=UTF-8")
	}
	return req, nil
}

// contentTypeWithEncodingHandleResponse handles the ContentTypeWithEncoding response.
func (client *MediaTypesClient) contentTypeWithEncodingHandleResponse(resp *http.Response) (MediaTypesClientContentTypeWithEncodingResponse, error) {
	result := MediaTypesClientContentTypeWithEncodingResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return MediaTypesClientContentTypeWithEncodingResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// contentTypeWithEncodingHandleError handles the ContentTypeWithEncoding error response.
func (client *MediaTypesClient) contentTypeWithEncodingHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
