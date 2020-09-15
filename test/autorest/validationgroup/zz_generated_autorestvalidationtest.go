// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package validationgroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// AutoRestValidationTestOperations contains the methods for the AutoRestValidationTest group.
type AutoRestValidationTestOperations interface {
	GetWithConstantInPath(ctx context.Context) (*http.Response, error)
	PostWithConstantInBody(ctx context.Context, autoRestValidationTestPostWithConstantInBodyOptions *AutoRestValidationTestPostWithConstantInBodyOptions) (*ProductResponse, error)
	// ValidationOfBody - Validates body parameters on the method. See swagger for details.
	ValidationOfBody(ctx context.Context, resourceGroupName string, id int32, autoRestValidationTestValidationOfBodyOptions *AutoRestValidationTestValidationOfBodyOptions) (*ProductResponse, error)
	// ValidationOfMethodParameters - Validates input parameters on the method. See swagger for details.
	ValidationOfMethodParameters(ctx context.Context, resourceGroupName string, id int32) (*ProductResponse, error)
}

// AutoRestValidationTestClient implements the AutoRestValidationTestOperations interface.
// Don't use this type directly, use NewAutoRestValidationTestClient() instead.
type AutoRestValidationTestClient struct {
	*Client
	subscriptionID string
}

// NewAutoRestValidationTestClient creates a new instance of AutoRestValidationTestClient with the specified values.
func NewAutoRestValidationTestClient(c *Client, subscriptionID string) AutoRestValidationTestOperations {
	return &AutoRestValidationTestClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *AutoRestValidationTestClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

func (client *AutoRestValidationTestClient) GetWithConstantInPath(ctx context.Context) (*http.Response, error) {
	req, err := client.GetWithConstantInPathCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetWithConstantInPathHandleError(resp)
	}
	return resp.Response, nil
}

// GetWithConstantInPathCreateRequest creates the GetWithConstantInPath request.
func (client *AutoRestValidationTestClient) GetWithConstantInPathCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/validation/constantsInPath/{constantParam}/value"
	urlPath = strings.ReplaceAll(urlPath, "{constantParam}", url.PathEscape("constant"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetWithConstantInPathHandleError handles the GetWithConstantInPath error response.
func (client *AutoRestValidationTestClient) GetWithConstantInPathHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

func (client *AutoRestValidationTestClient) PostWithConstantInBody(ctx context.Context, autoRestValidationTestPostWithConstantInBodyOptions *AutoRestValidationTestPostWithConstantInBodyOptions) (*ProductResponse, error) {
	req, err := client.PostWithConstantInBodyCreateRequest(ctx, autoRestValidationTestPostWithConstantInBodyOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostWithConstantInBodyHandleError(resp)
	}
	result, err := client.PostWithConstantInBodyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostWithConstantInBodyCreateRequest creates the PostWithConstantInBody request.
func (client *AutoRestValidationTestClient) PostWithConstantInBodyCreateRequest(ctx context.Context, autoRestValidationTestPostWithConstantInBodyOptions *AutoRestValidationTestPostWithConstantInBodyOptions) (*azcore.Request, error) {
	urlPath := "/validation/constantsInPath/{constantParam}/value"
	urlPath = strings.ReplaceAll(urlPath, "{constantParam}", url.PathEscape("constant"))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if autoRestValidationTestPostWithConstantInBodyOptions != nil {
		return req, req.MarshalAsJSON(autoRestValidationTestPostWithConstantInBodyOptions.Body)
	}
	return req, nil
}

// PostWithConstantInBodyHandleResponse handles the PostWithConstantInBody response.
func (client *AutoRestValidationTestClient) PostWithConstantInBodyHandleResponse(resp *azcore.Response) (*ProductResponse, error) {
	result := ProductResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Product)
}

// PostWithConstantInBodyHandleError handles the PostWithConstantInBody error response.
func (client *AutoRestValidationTestClient) PostWithConstantInBodyHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// ValidationOfBody - Validates body parameters on the method. See swagger for details.
func (client *AutoRestValidationTestClient) ValidationOfBody(ctx context.Context, resourceGroupName string, id int32, autoRestValidationTestValidationOfBodyOptions *AutoRestValidationTestValidationOfBodyOptions) (*ProductResponse, error) {
	req, err := client.ValidationOfBodyCreateRequest(ctx, resourceGroupName, id, autoRestValidationTestValidationOfBodyOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ValidationOfBodyHandleError(resp)
	}
	result, err := client.ValidationOfBodyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ValidationOfBodyCreateRequest creates the ValidationOfBody request.
func (client *AutoRestValidationTestClient) ValidationOfBodyCreateRequest(ctx context.Context, resourceGroupName string, id int32, autoRestValidationTestValidationOfBodyOptions *AutoRestValidationTestValidationOfBodyOptions) (*azcore.Request, error) {
	urlPath := "/fakepath/{subscriptionId}/{resourceGroupName}/{id}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(strconv.FormatInt(int64(id), 10)))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("apiVersion", "1.0.0")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	if autoRestValidationTestValidationOfBodyOptions != nil {
		return req, req.MarshalAsJSON(autoRestValidationTestValidationOfBodyOptions.Body)
	}
	return req, nil
}

// ValidationOfBodyHandleResponse handles the ValidationOfBody response.
func (client *AutoRestValidationTestClient) ValidationOfBodyHandleResponse(resp *azcore.Response) (*ProductResponse, error) {
	result := ProductResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Product)
}

// ValidationOfBodyHandleError handles the ValidationOfBody error response.
func (client *AutoRestValidationTestClient) ValidationOfBodyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ValidationOfMethodParameters - Validates input parameters on the method. See swagger for details.
func (client *AutoRestValidationTestClient) ValidationOfMethodParameters(ctx context.Context, resourceGroupName string, id int32) (*ProductResponse, error) {
	req, err := client.ValidationOfMethodParametersCreateRequest(ctx, resourceGroupName, id)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ValidationOfMethodParametersHandleError(resp)
	}
	result, err := client.ValidationOfMethodParametersHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ValidationOfMethodParametersCreateRequest creates the ValidationOfMethodParameters request.
func (client *AutoRestValidationTestClient) ValidationOfMethodParametersCreateRequest(ctx context.Context, resourceGroupName string, id int32) (*azcore.Request, error) {
	urlPath := "/fakepath/{subscriptionId}/{resourceGroupName}/{id}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(strconv.FormatInt(int64(id), 10)))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("apiVersion", "1.0.0")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ValidationOfMethodParametersHandleResponse handles the ValidationOfMethodParameters response.
func (client *AutoRestValidationTestClient) ValidationOfMethodParametersHandleResponse(resp *azcore.Response) (*ProductResponse, error) {
	result := ProductResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Product)
}

// ValidationOfMethodParametersHandleError handles the ValidationOfMethodParameters error response.
func (client *AutoRestValidationTestClient) ValidationOfMethodParametersHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
