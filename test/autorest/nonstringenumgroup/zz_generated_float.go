// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package nonstringenumgroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
)

// FloatOperations contains the methods for the Float group.
type FloatOperations interface {
	// Get - Get a float enum
	Get(ctx context.Context) (*FloatEnumResponse, error)
	// Put - Put a float enum
	Put(ctx context.Context, floatPutOptions *FloatPutOptions) (*StringResponse, error)
}

// FloatClient implements the FloatOperations interface.
// Don't use this type directly, use NewFloatClient() instead.
type FloatClient struct {
	*Client
}

// NewFloatClient creates a new instance of FloatClient with the specified values.
func NewFloatClient(c *Client) FloatOperations {
	return &FloatClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *FloatClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Get - Get a float enum
func (client *FloatClient) Get(ctx context.Context) (*FloatEnumResponse, error) {
	req, err := client.GetCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *FloatClient) GetCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/nonStringEnums/float/get"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *FloatClient) GetHandleResponse(resp *azcore.Response) (*FloatEnumResponse, error) {
	result := FloatEnumResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetHandleError handles the Get error response.
func (client *FloatClient) GetHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// Put - Put a float enum
func (client *FloatClient) Put(ctx context.Context, floatPutOptions *FloatPutOptions) (*StringResponse, error) {
	req, err := client.PutCreateRequest(ctx, floatPutOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutHandleError(resp)
	}
	result, err := client.PutHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutCreateRequest creates the Put request.
func (client *FloatClient) PutCreateRequest(ctx context.Context, floatPutOptions *FloatPutOptions) (*azcore.Request, error) {
	urlPath := "/nonStringEnums/float/put"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if floatPutOptions != nil {
		return req, req.MarshalAsJSON(floatPutOptions.Input)
	}
	return req, nil
}

// PutHandleResponse handles the Put response.
func (client *FloatClient) PutHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// PutHandleError handles the Put error response.
func (client *FloatClient) PutHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}
