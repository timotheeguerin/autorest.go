// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package mediatypesgroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"strings"
)

// MediaTypesClientOperations contains the methods for the MediaTypesClient group.
type MediaTypesClientOperations interface {
	// AnalyzeBody - Analyze body, that could be different media types.
	AnalyzeBody(ctx context.Context, contentType ContentType, input azcore.ReadSeekCloser) (*StringResponse, error)
	// AnalyzeBodyWithSourcePath - Analyze body, that could be different media types.
	AnalyzeBodyWithSourcePath(ctx context.Context, mediaTypesClientAnalyzeBodyWithSourcePathOptions *MediaTypesClientAnalyzeBodyWithSourcePathOptions) (*StringResponse, error)
	// ContentTypeWithEncoding - Pass in contentType 'text/plain; encoding=UTF-8' to pass test. Value for input does not matter
	ContentTypeWithEncoding(ctx context.Context, input string) (*StringResponse, error)
}

// MediaTypesClient implements the MediaTypesClientOperations interface.
// Don't use this type directly, use NewMediaTypesClient() instead.
type MediaTypesClient struct {
	*Client
}

// NewMediaTypesClient creates a new instance of MediaTypesClient with the specified values.
func NewMediaTypesClient(c *Client) MediaTypesClientOperations {
	return &MediaTypesClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *MediaTypesClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// AnalyzeBody - Analyze body, that could be different media types.
func (client *MediaTypesClient) AnalyzeBody(ctx context.Context, contentType ContentType, input azcore.ReadSeekCloser) (*StringResponse, error) {
	req, err := client.AnalyzeBodyCreateRequest(ctx, contentType, input)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.AnalyzeBodyHandleError(resp)
	}
	result, err := client.AnalyzeBodyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AnalyzeBodyCreateRequest creates the AnalyzeBody request.
func (client *MediaTypesClient) AnalyzeBodyCreateRequest(ctx context.Context, contentType ContentType, input azcore.ReadSeekCloser) (*azcore.Request, error) {
	urlPath := "/mediatypes/analyze"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", string(contentType))
	req.Header.Set("Accept", "application/json")
	return req, req.SetBody(input, string(contentType))
}

// AnalyzeBodyHandleResponse handles the AnalyzeBody response.
func (client *MediaTypesClient) AnalyzeBodyHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// AnalyzeBodyHandleError handles the AnalyzeBody error response.
func (client *MediaTypesClient) AnalyzeBodyHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// AnalyzeBodyWithSourcePath - Analyze body, that could be different media types.
func (client *MediaTypesClient) AnalyzeBodyWithSourcePath(ctx context.Context, mediaTypesClientAnalyzeBodyWithSourcePathOptions *MediaTypesClientAnalyzeBodyWithSourcePathOptions) (*StringResponse, error) {
	req, err := client.AnalyzeBodyWithSourcePathCreateRequest(ctx, mediaTypesClientAnalyzeBodyWithSourcePathOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.AnalyzeBodyWithSourcePathHandleError(resp)
	}
	result, err := client.AnalyzeBodyWithSourcePathHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AnalyzeBodyWithSourcePathCreateRequest creates the AnalyzeBodyWithSourcePath request.
func (client *MediaTypesClient) AnalyzeBodyWithSourcePathCreateRequest(ctx context.Context, mediaTypesClientAnalyzeBodyWithSourcePathOptions *MediaTypesClientAnalyzeBodyWithSourcePathOptions) (*azcore.Request, error) {
	urlPath := "/mediatypes/analyze"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if mediaTypesClientAnalyzeBodyWithSourcePathOptions != nil {
		return req, req.MarshalAsJSON(mediaTypesClientAnalyzeBodyWithSourcePathOptions.Input)
	}
	return req, nil
}

// AnalyzeBodyWithSourcePathHandleResponse handles the AnalyzeBodyWithSourcePath response.
func (client *MediaTypesClient) AnalyzeBodyWithSourcePathHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// AnalyzeBodyWithSourcePathHandleError handles the AnalyzeBodyWithSourcePath error response.
func (client *MediaTypesClient) AnalyzeBodyWithSourcePathHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// ContentTypeWithEncoding - Pass in contentType 'text/plain; encoding=UTF-8' to pass test. Value for input does not matter
func (client *MediaTypesClient) ContentTypeWithEncoding(ctx context.Context, input string) (*StringResponse, error) {
	req, err := client.ContentTypeWithEncodingCreateRequest(ctx, input)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ContentTypeWithEncodingHandleError(resp)
	}
	result, err := client.ContentTypeWithEncodingHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ContentTypeWithEncodingCreateRequest creates the ContentTypeWithEncoding request.
func (client *MediaTypesClient) ContentTypeWithEncodingCreateRequest(ctx context.Context, input string) (*azcore.Request, error) {
	urlPath := "/mediatypes/contentTypeWithEncoding"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	body := azcore.NopCloser(strings.NewReader(input))
	return req, req.SetBody(body, "text/plain; encoding=UTF-8")
}

// ContentTypeWithEncodingHandleResponse handles the ContentTypeWithEncoding response.
func (client *MediaTypesClient) ContentTypeWithEncodingHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// ContentTypeWithEncodingHandleError handles the ContentTypeWithEncoding error response.
func (client *MediaTypesClient) ContentTypeWithEncodingHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}
