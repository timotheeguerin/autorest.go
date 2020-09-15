// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package booleangroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// BoolOperations contains the methods for the Bool group.
type BoolOperations interface {
	// GetFalse - Get false Boolean value
	GetFalse(ctx context.Context) (*BoolResponse, error)
	// GetInvalid - Get invalid Boolean value
	GetInvalid(ctx context.Context) (*BoolResponse, error)
	// GetNull - Get null Boolean value
	GetNull(ctx context.Context) (*BoolResponse, error)
	// GetTrue - Get true Boolean value
	GetTrue(ctx context.Context) (*BoolResponse, error)
	// PutFalse - Set Boolean value false
	PutFalse(ctx context.Context) (*http.Response, error)
	// PutTrue - Set Boolean value true
	PutTrue(ctx context.Context) (*http.Response, error)
}

// BoolClient implements the BoolOperations interface.
// Don't use this type directly, use NewBoolClient() instead.
type BoolClient struct {
	*Client
}

// NewBoolClient creates a new instance of BoolClient with the specified values.
func NewBoolClient(c *Client) BoolOperations {
	return &BoolClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *BoolClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// GetFalse - Get false Boolean value
func (client *BoolClient) GetFalse(ctx context.Context) (*BoolResponse, error) {
	req, err := client.GetFalseCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetFalseHandleError(resp)
	}
	result, err := client.GetFalseHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFalseCreateRequest creates the GetFalse request.
func (client *BoolClient) GetFalseCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/bool/false"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetFalseHandleResponse handles the GetFalse response.
func (client *BoolClient) GetFalseHandleResponse(resp *azcore.Response) (*BoolResponse, error) {
	result := BoolResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetFalseHandleError handles the GetFalse error response.
func (client *BoolClient) GetFalseHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetInvalid - Get invalid Boolean value
func (client *BoolClient) GetInvalid(ctx context.Context) (*BoolResponse, error) {
	req, err := client.GetInvalidCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetInvalidHandleError(resp)
	}
	result, err := client.GetInvalidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInvalidCreateRequest creates the GetInvalid request.
func (client *BoolClient) GetInvalidCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/bool/invalid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetInvalidHandleResponse handles the GetInvalid response.
func (client *BoolClient) GetInvalidHandleResponse(resp *azcore.Response) (*BoolResponse, error) {
	result := BoolResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetInvalidHandleError handles the GetInvalid error response.
func (client *BoolClient) GetInvalidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetNull - Get null Boolean value
func (client *BoolClient) GetNull(ctx context.Context) (*BoolResponse, error) {
	req, err := client.GetNullCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetNullHandleError(resp)
	}
	result, err := client.GetNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNullCreateRequest creates the GetNull request.
func (client *BoolClient) GetNullCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/bool/null"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetNullHandleResponse handles the GetNull response.
func (client *BoolClient) GetNullHandleResponse(resp *azcore.Response) (*BoolResponse, error) {
	result := BoolResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetNullHandleError handles the GetNull error response.
func (client *BoolClient) GetNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetTrue - Get true Boolean value
func (client *BoolClient) GetTrue(ctx context.Context) (*BoolResponse, error) {
	req, err := client.GetTrueCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetTrueHandleError(resp)
	}
	result, err := client.GetTrueHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTrueCreateRequest creates the GetTrue request.
func (client *BoolClient) GetTrueCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/bool/true"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetTrueHandleResponse handles the GetTrue response.
func (client *BoolClient) GetTrueHandleResponse(resp *azcore.Response) (*BoolResponse, error) {
	result := BoolResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetTrueHandleError handles the GetTrue error response.
func (client *BoolClient) GetTrueHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutFalse - Set Boolean value false
func (client *BoolClient) PutFalse(ctx context.Context) (*http.Response, error) {
	req, err := client.PutFalseCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutFalseHandleError(resp)
	}
	return resp.Response, nil
}

// PutFalseCreateRequest creates the PutFalse request.
func (client *BoolClient) PutFalseCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/bool/false"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(false)
}

// PutFalseHandleError handles the PutFalse error response.
func (client *BoolClient) PutFalseHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutTrue - Set Boolean value true
func (client *BoolClient) PutTrue(ctx context.Context) (*http.Response, error) {
	req, err := client.PutTrueCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutTrueHandleError(resp)
	}
	return resp.Response, nil
}

// PutTrueCreateRequest creates the PutTrue request.
func (client *BoolClient) PutTrueCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/bool/true"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// PutTrueHandleError handles the PutTrue error response.
func (client *BoolClient) PutTrueHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
