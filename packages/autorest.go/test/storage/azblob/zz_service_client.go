//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azblob

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ServiceClient contains the methods for the Service group.
// Don't use this type directly, use a constructor function instead.
type ServiceClient struct {
	internal *azcore.Client
	endpoint string
	version  Enum2
}

// FilterBlobs - The Filter Blobs operation enables callers to list blobs across all containers whose tags match a given search
// expression. Filter blobs searches across all containers within a storage account but can
// be scoped within the expression to a single container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-02
//   - comp - comp
//   - options - ServiceClientFilterBlobsOptions contains the optional parameters for the ServiceClient.FilterBlobs method.
func (client *ServiceClient) FilterBlobs(ctx context.Context, comp Enum10, options *ServiceClientFilterBlobsOptions) (ServiceClientFilterBlobsResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "ServiceClient.FilterBlobs", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.filterBlobsCreateRequest(ctx, comp, options)
	if err != nil {
		return ServiceClientFilterBlobsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientFilterBlobsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServiceClientFilterBlobsResponse{}, err
	}
	resp, err := client.filterBlobsHandleResponse(httpResp)
	return resp, err
}

// filterBlobsCreateRequest creates the FilterBlobs request.
func (client *ServiceClient) filterBlobsCreateRequest(ctx context.Context, comp Enum10, options *ServiceClientFilterBlobsOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", string(comp))
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Where != nil {
		reqQP.Set("where", *options.Where)
	}
	if options != nil && options.Marker != nil {
		reqQP.Set("marker", *options.Marker)
	}
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	if options != nil && options.Include != nil {
		reqQP.Set("include", strings.Join(strings.Fields(strings.Trim(fmt.Sprint(options.Include), "[]")), ","))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{string(client.version)}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// filterBlobsHandleResponse handles the FilterBlobs response.
func (client *ServiceClient) filterBlobsHandleResponse(resp *http.Response) (ServiceClientFilterBlobsResponse, error) {
	result := ServiceClientFilterBlobsResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ServiceClientFilterBlobsResponse{}, err
		}
		result.Date = &date
	}
	if err := runtime.UnmarshalAsXML(resp, &result.FilterBlobSegment); err != nil {
		return ServiceClientFilterBlobsResponse{}, err
	}
	return result, nil
}

// GetAccountInfo - Returns the sku name and account kind
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-02
//   - restype - restype
//   - comp - comp
//   - options - ServiceClientGetAccountInfoOptions contains the optional parameters for the ServiceClient.GetAccountInfo method.
func (client *ServiceClient) GetAccountInfo(ctx context.Context, restype Enum8, comp Enum1, options *ServiceClientGetAccountInfoOptions) (ServiceClientGetAccountInfoResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "ServiceClient.GetAccountInfo", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAccountInfoCreateRequest(ctx, restype, comp, options)
	if err != nil {
		return ServiceClientGetAccountInfoResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientGetAccountInfoResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServiceClientGetAccountInfoResponse{}, err
	}
	resp, err := client.getAccountInfoHandleResponse(httpResp)
	return resp, err
}

// getAccountInfoCreateRequest creates the GetAccountInfo request.
func (client *ServiceClient) getAccountInfoCreateRequest(ctx context.Context, restype Enum8, comp Enum1, options *ServiceClientGetAccountInfoOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", string(restype))
	reqQP.Set("comp", string(comp))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{string(client.version)}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// getAccountInfoHandleResponse handles the GetAccountInfo response.
func (client *ServiceClient) getAccountInfoHandleResponse(resp *http.Response) (ServiceClientGetAccountInfoResponse, error) {
	result := ServiceClientGetAccountInfoResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ServiceClientGetAccountInfoResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-sku-name"); val != "" {
		result.SKUName = (*SKUName)(&val)
	}
	if val := resp.Header.Get("x-ms-account-kind"); val != "" {
		result.AccountKind = (*AccountKind)(&val)
	}
	if val := resp.Header.Get("x-ms-is-hns-enabled"); val != "" {
		isHierarchicalNamespaceEnabled, err := strconv.ParseBool(val)
		if err != nil {
			return ServiceClientGetAccountInfoResponse{}, err
		}
		result.IsHierarchicalNamespaceEnabled = &isHierarchicalNamespaceEnabled
	}
	return result, nil
}

// GetProperties - gets the properties of a storage account's Blob service, including properties for Storage Analytics and
// CORS (Cross-Origin Resource Sharing) rules.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-02
//   - restype - restype
//   - comp - comp
//   - options - ServiceClientGetPropertiesOptions contains the optional parameters for the ServiceClient.GetProperties method.
func (client *ServiceClient) GetProperties(ctx context.Context, restype Enum0, comp Enum1, options *ServiceClientGetPropertiesOptions) (ServiceClientGetPropertiesResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "ServiceClient.GetProperties", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getPropertiesCreateRequest(ctx, restype, comp, options)
	if err != nil {
		return ServiceClientGetPropertiesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientGetPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServiceClientGetPropertiesResponse{}, err
	}
	resp, err := client.getPropertiesHandleResponse(httpResp)
	return resp, err
}

// getPropertiesCreateRequest creates the GetProperties request.
func (client *ServiceClient) getPropertiesCreateRequest(ctx context.Context, restype Enum0, comp Enum1, options *ServiceClientGetPropertiesOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", string(restype))
	reqQP.Set("comp", string(comp))
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{string(client.version)}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// getPropertiesHandleResponse handles the GetProperties response.
func (client *ServiceClient) getPropertiesHandleResponse(resp *http.Response) (ServiceClientGetPropertiesResponse, error) {
	result := ServiceClientGetPropertiesResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if err := runtime.UnmarshalAsXML(resp, &result.StorageServiceProperties); err != nil {
		return ServiceClientGetPropertiesResponse{}, err
	}
	return result, nil
}

// GetStatistics - Retrieves statistics related to replication for the Blob service. It is only available on the secondary
// location endpoint when read-access geo-redundant replication is enabled for the storage account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-02
//   - restype - restype
//   - comp - comp
//   - options - ServiceClientGetStatisticsOptions contains the optional parameters for the ServiceClient.GetStatistics method.
func (client *ServiceClient) GetStatistics(ctx context.Context, restype Enum0, comp Enum3, options *ServiceClientGetStatisticsOptions) (ServiceClientGetStatisticsResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "ServiceClient.GetStatistics", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getStatisticsCreateRequest(ctx, restype, comp, options)
	if err != nil {
		return ServiceClientGetStatisticsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientGetStatisticsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServiceClientGetStatisticsResponse{}, err
	}
	resp, err := client.getStatisticsHandleResponse(httpResp)
	return resp, err
}

// getStatisticsCreateRequest creates the GetStatistics request.
func (client *ServiceClient) getStatisticsCreateRequest(ctx context.Context, restype Enum0, comp Enum3, options *ServiceClientGetStatisticsOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", string(restype))
	reqQP.Set("comp", string(comp))
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{string(client.version)}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// getStatisticsHandleResponse handles the GetStatistics response.
func (client *ServiceClient) getStatisticsHandleResponse(resp *http.Response) (ServiceClientGetStatisticsResponse, error) {
	result := ServiceClientGetStatisticsResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ServiceClientGetStatisticsResponse{}, err
		}
		result.Date = &date
	}
	if err := runtime.UnmarshalAsXML(resp, &result.StorageServiceStats); err != nil {
		return ServiceClientGetStatisticsResponse{}, err
	}
	return result, nil
}

// GetUserDelegationKey - Retrieves a user delegation key for the Blob service. This is only a valid operation when using
// bearer token authentication.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-02
//   - restype - restype
//   - comp - comp
//   - keyInfo - Key information
//   - options - ServiceClientGetUserDelegationKeyOptions contains the optional parameters for the ServiceClient.GetUserDelegationKey
//     method.
func (client *ServiceClient) GetUserDelegationKey(ctx context.Context, restype Enum0, comp Enum7, keyInfo KeyInfo, options *ServiceClientGetUserDelegationKeyOptions) (ServiceClientGetUserDelegationKeyResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "ServiceClient.GetUserDelegationKey", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getUserDelegationKeyCreateRequest(ctx, restype, comp, keyInfo, options)
	if err != nil {
		return ServiceClientGetUserDelegationKeyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientGetUserDelegationKeyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServiceClientGetUserDelegationKeyResponse{}, err
	}
	resp, err := client.getUserDelegationKeyHandleResponse(httpResp)
	return resp, err
}

// getUserDelegationKeyCreateRequest creates the GetUserDelegationKey request.
func (client *ServiceClient) getUserDelegationKeyCreateRequest(ctx context.Context, restype Enum0, comp Enum7, keyInfo KeyInfo, options *ServiceClientGetUserDelegationKeyOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPost, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", string(restype))
	reqQP.Set("comp", string(comp))
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{string(client.version)}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	if err := runtime.MarshalAsXML(req, keyInfo); err != nil {
		return nil, err
	}
	return req, nil
}

// getUserDelegationKeyHandleResponse handles the GetUserDelegationKey response.
func (client *ServiceClient) getUserDelegationKeyHandleResponse(resp *http.Response) (ServiceClientGetUserDelegationKeyResponse, error) {
	result := ServiceClientGetUserDelegationKeyResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ServiceClientGetUserDelegationKeyResponse{}, err
		}
		result.Date = &date
	}
	if err := runtime.UnmarshalAsXML(resp, &result.UserDelegationKey); err != nil {
		return ServiceClientGetUserDelegationKeyResponse{}, err
	}
	return result, nil
}

// NewListContainersSegmentPager - The List Containers Segment operation returns a list of the containers under the specified
// account
//
// Generated from API version 2021-12-02
//   - comp - comp
//   - options - ServiceClientListContainersSegmentOptions contains the optional parameters for the ServiceClient.NewListContainersSegmentPager
//     method.
func (client *ServiceClient) NewListContainersSegmentPager(comp Enum5, options *ServiceClientListContainersSegmentOptions) *runtime.Pager[ServiceClientListContainersSegmentResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServiceClientListContainersSegmentResponse]{
		More: func(page ServiceClientListContainersSegmentResponse) bool {
			return page.NextMarker != nil && len(*page.NextMarker) > 0
		},
		Fetcher: func(ctx context.Context, page *ServiceClientListContainersSegmentResponse) (ServiceClientListContainersSegmentResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listContainersSegmentCreateRequest(ctx, comp, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextMarker)
			}
			if err != nil {
				return ServiceClientListContainersSegmentResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServiceClientListContainersSegmentResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServiceClientListContainersSegmentResponse{}, runtime.NewResponseError(resp)
			}
			return client.listContainersSegmentHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listContainersSegmentCreateRequest creates the ListContainersSegment request.
func (client *ServiceClient) listContainersSegmentCreateRequest(ctx context.Context, comp Enum5, options *ServiceClientListContainersSegmentOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", string(comp))
	if options != nil && options.Prefix != nil {
		reqQP.Set("prefix", *options.Prefix)
	}
	if options != nil && options.Marker != nil {
		reqQP.Set("marker", *options.Marker)
	}
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	if options != nil && options.Include != nil {
		reqQP.Set("include", strings.Join(strings.Fields(strings.Trim(fmt.Sprint(options.Include), "[]")), ","))
	}
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{string(client.version)}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// listContainersSegmentHandleResponse handles the ListContainersSegment response.
func (client *ServiceClient) listContainersSegmentHandleResponse(resp *http.Response) (ServiceClientListContainersSegmentResponse, error) {
	result := ServiceClientListContainersSegmentResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if err := runtime.UnmarshalAsXML(resp, &result.ListContainersSegmentResponse); err != nil {
		return ServiceClientListContainersSegmentResponse{}, err
	}
	return result, nil
}

// SetProperties - Sets properties for a storage account's Blob service endpoint, including properties for Storage Analytics
// and CORS (Cross-Origin Resource Sharing) rules
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-02
//   - restype - restype
//   - comp - comp
//   - storageServiceProperties - The StorageService properties.
//   - options - ServiceClientSetPropertiesOptions contains the optional parameters for the ServiceClient.SetProperties method.
func (client *ServiceClient) SetProperties(ctx context.Context, restype Enum0, comp Enum1, storageServiceProperties StorageServiceProperties, options *ServiceClientSetPropertiesOptions) (ServiceClientSetPropertiesResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "ServiceClient.SetProperties", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.setPropertiesCreateRequest(ctx, restype, comp, storageServiceProperties, options)
	if err != nil {
		return ServiceClientSetPropertiesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientSetPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return ServiceClientSetPropertiesResponse{}, err
	}
	resp, err := client.setPropertiesHandleResponse(httpResp)
	return resp, err
}

// setPropertiesCreateRequest creates the SetProperties request.
func (client *ServiceClient) setPropertiesCreateRequest(ctx context.Context, restype Enum0, comp Enum1, storageServiceProperties StorageServiceProperties, options *ServiceClientSetPropertiesOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", string(restype))
	reqQP.Set("comp", string(comp))
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{string(client.version)}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	if err := runtime.MarshalAsXML(req, storageServiceProperties); err != nil {
		return nil, err
	}
	return req, nil
}

// setPropertiesHandleResponse handles the SetProperties response.
func (client *ServiceClient) setPropertiesHandleResponse(resp *http.Response) (ServiceClientSetPropertiesResponse, error) {
	result := ServiceClientSetPropertiesResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return result, nil
}

// SubmitBatch - The Batch operation allows multiple API calls to be embedded into a single HTTP request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-02
//   - comp - comp
//   - contentLength - The length of the request.
//   - multipartContentType - Required. The value of this header must be multipart/mixed with a batch boundary. Example header
//     value: multipart/mixed; boundary=batch_
//   - body - Initial data
//   - options - ServiceClientSubmitBatchOptions contains the optional parameters for the ServiceClient.SubmitBatch method.
func (client *ServiceClient) SubmitBatch(ctx context.Context, comp Enum9, contentLength int64, multipartContentType string, body io.ReadSeekCloser, options *ServiceClientSubmitBatchOptions) (ServiceClientSubmitBatchResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "ServiceClient.SubmitBatch", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.submitBatchCreateRequest(ctx, comp, contentLength, multipartContentType, body, options)
	if err != nil {
		return ServiceClientSubmitBatchResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientSubmitBatchResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServiceClientSubmitBatchResponse{}, err
	}
	resp, err := client.submitBatchHandleResponse(httpResp)
	return resp, err
}

// submitBatchCreateRequest creates the SubmitBatch request.
func (client *ServiceClient) submitBatchCreateRequest(ctx context.Context, comp Enum9, contentLength int64, multipartContentType string, body io.ReadSeekCloser, options *ServiceClientSubmitBatchOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPost, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", string(comp))
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	runtime.SkipBodyDownload(req)
	req.Raw().Header["Content-Length"] = []string{strconv.FormatInt(contentLength, 10)}
	req.Raw().Header["Content-Type"] = []string{multipartContentType}
	req.Raw().Header["x-ms-version"] = []string{string(client.version)}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	if err := req.SetBody(body, "application/xml"); err != nil {
		return nil, err
	}
	return req, nil
}

// submitBatchHandleResponse handles the SubmitBatch response.
func (client *ServiceClient) submitBatchHandleResponse(resp *http.Response) (ServiceClientSubmitBatchResponse, error) {
	result := ServiceClientSubmitBatchResponse{Body: resp.Body}
	if val := resp.Header.Get("Content-Type"); val != "" {
		result.ContentType = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return result, nil
}
