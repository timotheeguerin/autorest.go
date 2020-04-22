// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// ServiceOperations contains the methods for the Service group.
type ServiceOperations interface {
	// GetAccountInfo - Returns the sku name and account kind
	GetAccountInfo(ctx context.Context) (*ServiceGetAccountInfoResponse, error)
	// GetProperties - gets the properties of a storage account's Blob service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules.
	GetProperties(ctx context.Context, serviceGetPropertiesOptions *ServiceGetPropertiesOptions) (*StorageServicePropertiesResponse, error)
	// GetStatistics - Retrieves statistics related to replication for the Blob service. It is only available on the secondary location endpoint when read-access geo-redundant replication is enabled for the storage account.
	GetStatistics(ctx context.Context, serviceGetStatisticsOptions *ServiceGetStatisticsOptions) (*StorageServiceStatsResponse, error)
	// GetUserDelegationKey - Retrieves a user delegation key for the Blob service. This is only a valid operation when using bearer token authentication.
	GetUserDelegationKey(ctx context.Context, keyInfo KeyInfo, serviceGetUserDelegationKeyOptions *ServiceGetUserDelegationKeyOptions) (*UserDelegationKeyResponse, error)
	// ListContainersSegment - The List Containers Segment operation returns a list of the containers under the specified account
	ListContainersSegment(serviceListContainersSegmentOptions *ServiceListContainersSegmentOptions) (ListContainersSegmentResponsePager, error)
	// SetProperties - Sets properties for a storage account's Blob service endpoint, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules
	SetProperties(ctx context.Context, storageServiceProperties StorageServiceProperties, serviceSetPropertiesOptions *ServiceSetPropertiesOptions) (*ServiceSetPropertiesResponse, error)
	// SubmitBatch - The Batch operation allows multiple API calls to be embedded into a single HTTP request.
	SubmitBatch(ctx context.Context, contentLength int64, multipartContentType string, body azcore.ReadSeekCloser, serviceSubmitBatchOptions *ServiceSubmitBatchOptions) (*ServiceSubmitBatchResponse, error)
}

// serviceOperations implements the ServiceOperations interface.
type serviceOperations struct {
	*Client
}

// GetAccountInfo - Returns the sku name and account kind
func (client *serviceOperations) GetAccountInfo(ctx context.Context) (*ServiceGetAccountInfoResponse, error) {
	req, err := client.getAccountInfoCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getAccountInfoHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getAccountInfoCreateRequest creates the GetAccountInfo request.
func (client *serviceOperations) getAccountInfoCreateRequest() (*azcore.Request, error) {
	u := client.u
	query := u.Query()
	query.Set("restype", "account")
	query.Set("comp", "properties")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	req.Header.Set("x-ms-version", "2019-07-07")
	return req, nil
}

// getAccountInfoHandleResponse handles the GetAccountInfo response.
func (client *serviceOperations) getAccountInfoHandleResponse(resp *azcore.Response) (*ServiceGetAccountInfoResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newStorageError(resp)
	}
	result := ServiceGetAccountInfoResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestId = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestId = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-sku-name"); val != "" {
		result.SkuName = (*SkuName)(&val)
	}
	if val := resp.Header.Get("x-ms-account-kind"); val != "" {
		result.AccountKind = (*AccountKind)(&val)
	}
	return &result, nil
}

// GetProperties - gets the properties of a storage account's Blob service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules.
func (client *serviceOperations) GetProperties(ctx context.Context, serviceGetPropertiesOptions *ServiceGetPropertiesOptions) (*StorageServicePropertiesResponse, error) {
	req, err := client.getPropertiesCreateRequest(serviceGetPropertiesOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getPropertiesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getPropertiesCreateRequest creates the GetProperties request.
func (client *serviceOperations) getPropertiesCreateRequest(serviceGetPropertiesOptions *ServiceGetPropertiesOptions) (*azcore.Request, error) {
	u := client.u
	query := u.Query()
	query.Set("restype", "service")
	query.Set("comp", "properties")
	if serviceGetPropertiesOptions != nil && serviceGetPropertiesOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*serviceGetPropertiesOptions.Timeout), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	req.Header.Set("x-ms-version", "2019-07-07")
	if serviceGetPropertiesOptions != nil && serviceGetPropertiesOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *serviceGetPropertiesOptions.RequestId)
	}
	return req, nil
}

// getPropertiesHandleResponse handles the GetProperties response.
func (client *serviceOperations) getPropertiesHandleResponse(resp *azcore.Response) (*StorageServicePropertiesResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newStorageError(resp)
	}
	result := StorageServicePropertiesResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestId = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestId = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return &result, resp.UnmarshalAsXML(&result.StorageServiceProperties)
}

// GetStatistics - Retrieves statistics related to replication for the Blob service. It is only available on the secondary location endpoint when read-access geo-redundant replication is enabled for the storage account.
func (client *serviceOperations) GetStatistics(ctx context.Context, serviceGetStatisticsOptions *ServiceGetStatisticsOptions) (*StorageServiceStatsResponse, error) {
	req, err := client.getStatisticsCreateRequest(serviceGetStatisticsOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getStatisticsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getStatisticsCreateRequest creates the GetStatistics request.
func (client *serviceOperations) getStatisticsCreateRequest(serviceGetStatisticsOptions *ServiceGetStatisticsOptions) (*azcore.Request, error) {
	u := client.u
	query := u.Query()
	query.Set("restype", "service")
	query.Set("comp", "stats")
	if serviceGetStatisticsOptions != nil && serviceGetStatisticsOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*serviceGetStatisticsOptions.Timeout), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	req.Header.Set("x-ms-version", "2019-07-07")
	if serviceGetStatisticsOptions != nil && serviceGetStatisticsOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *serviceGetStatisticsOptions.RequestId)
	}
	return req, nil
}

// getStatisticsHandleResponse handles the GetStatistics response.
func (client *serviceOperations) getStatisticsHandleResponse(resp *azcore.Response) (*StorageServiceStatsResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newStorageError(resp)
	}
	result := StorageServiceStatsResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestId = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestId = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.Date = &date
	}
	return &result, resp.UnmarshalAsXML(&result.StorageServiceStats)
}

// GetUserDelegationKey - Retrieves a user delegation key for the Blob service. This is only a valid operation when using bearer token authentication.
func (client *serviceOperations) GetUserDelegationKey(ctx context.Context, keyInfo KeyInfo, serviceGetUserDelegationKeyOptions *ServiceGetUserDelegationKeyOptions) (*UserDelegationKeyResponse, error) {
	req, err := client.getUserDelegationKeyCreateRequest(keyInfo, serviceGetUserDelegationKeyOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getUserDelegationKeyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getUserDelegationKeyCreateRequest creates the GetUserDelegationKey request.
func (client *serviceOperations) getUserDelegationKeyCreateRequest(keyInfo KeyInfo, serviceGetUserDelegationKeyOptions *ServiceGetUserDelegationKeyOptions) (*azcore.Request, error) {
	u := client.u
	query := u.Query()
	query.Set("restype", "service")
	query.Set("comp", "userdelegationkey")
	if serviceGetUserDelegationKeyOptions != nil && serviceGetUserDelegationKeyOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*serviceGetUserDelegationKeyOptions.Timeout), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	req.Header.Set("x-ms-version", "2019-07-07")
	if serviceGetUserDelegationKeyOptions != nil && serviceGetUserDelegationKeyOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *serviceGetUserDelegationKeyOptions.RequestId)
	}
	return req, req.MarshalAsXML(keyInfo)
}

// getUserDelegationKeyHandleResponse handles the GetUserDelegationKey response.
func (client *serviceOperations) getUserDelegationKeyHandleResponse(resp *azcore.Response) (*UserDelegationKeyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newStorageError(resp)
	}
	result := UserDelegationKeyResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestId = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestId = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.Date = &date
	}
	return &result, resp.UnmarshalAsXML(&result.UserDelegationKey)
}

// ListContainersSegment - The List Containers Segment operation returns a list of the containers under the specified account
func (client *serviceOperations) ListContainersSegment(serviceListContainersSegmentOptions *ServiceListContainersSegmentOptions) (ListContainersSegmentResponsePager, error) {
	req, err := client.listContainersSegmentCreateRequest(serviceListContainersSegmentOptions)
	if err != nil {
		return nil, err
	}
	return &listContainersSegmentResponsePager{
		client:    client,
		request:   req,
		responder: client.listContainersSegmentHandleResponse,
		advancer: func(resp *ListContainersSegmentResponseResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.EnumerationResults.NextMarker)
			if err != nil {
				return nil, fmt.Errorf("invalid NextMarker: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextMarker %s", *resp.EnumerationResults.NextMarker)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listContainersSegmentCreateRequest creates the ListContainersSegment request.
func (client *serviceOperations) listContainersSegmentCreateRequest(serviceListContainersSegmentOptions *ServiceListContainersSegmentOptions) (*azcore.Request, error) {
	u := client.u
	query := u.Query()
	query.Set("comp", "list")
	if serviceListContainersSegmentOptions != nil && serviceListContainersSegmentOptions.Prefix != nil {
		query.Set("prefix", *serviceListContainersSegmentOptions.Prefix)
	}
	if serviceListContainersSegmentOptions != nil && serviceListContainersSegmentOptions.Marker != nil {
		query.Set("marker", *serviceListContainersSegmentOptions.Marker)
	}
	if serviceListContainersSegmentOptions != nil && serviceListContainersSegmentOptions.Maxresults != nil {
		query.Set("maxresults", strconv.FormatInt(int64(*serviceListContainersSegmentOptions.Maxresults), 10))
	}
	if serviceListContainersSegmentOptions != nil && serviceListContainersSegmentOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*serviceListContainersSegmentOptions.Timeout), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	req.Header.Set("x-ms-version", "2019-07-07")
	if serviceListContainersSegmentOptions != nil && serviceListContainersSegmentOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *serviceListContainersSegmentOptions.RequestId)
	}
	return req, nil
}

// listContainersSegmentHandleResponse handles the ListContainersSegment response.
func (client *serviceOperations) listContainersSegmentHandleResponse(resp *azcore.Response) (*ListContainersSegmentResponseResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newStorageError(resp)
	}
	result := ListContainersSegmentResponseResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestId = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestId = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return &result, resp.UnmarshalAsXML(&result.EnumerationResults)
}

// SetProperties - Sets properties for a storage account's Blob service endpoint, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules
func (client *serviceOperations) SetProperties(ctx context.Context, storageServiceProperties StorageServiceProperties, serviceSetPropertiesOptions *ServiceSetPropertiesOptions) (*ServiceSetPropertiesResponse, error) {
	req, err := client.setPropertiesCreateRequest(storageServiceProperties, serviceSetPropertiesOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.setPropertiesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// setPropertiesCreateRequest creates the SetProperties request.
func (client *serviceOperations) setPropertiesCreateRequest(storageServiceProperties StorageServiceProperties, serviceSetPropertiesOptions *ServiceSetPropertiesOptions) (*azcore.Request, error) {
	u := client.u
	query := u.Query()
	query.Set("restype", "service")
	query.Set("comp", "properties")
	if serviceSetPropertiesOptions != nil && serviceSetPropertiesOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*serviceSetPropertiesOptions.Timeout), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	req.Header.Set("x-ms-version", "2019-07-07")
	if serviceSetPropertiesOptions != nil && serviceSetPropertiesOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *serviceSetPropertiesOptions.RequestId)
	}
	return req, req.MarshalAsXML(storageServiceProperties)
}

// setPropertiesHandleResponse handles the SetProperties response.
func (client *serviceOperations) setPropertiesHandleResponse(resp *azcore.Response) (*ServiceSetPropertiesResponse, error) {
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, newStorageError(resp)
	}
	result := ServiceSetPropertiesResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestId = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestId = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return &result, nil
}

// SubmitBatch - The Batch operation allows multiple API calls to be embedded into a single HTTP request.
func (client *serviceOperations) SubmitBatch(ctx context.Context, contentLength int64, multipartContentType string, body azcore.ReadSeekCloser, serviceSubmitBatchOptions *ServiceSubmitBatchOptions) (*ServiceSubmitBatchResponse, error) {
	req, err := client.submitBatchCreateRequest(contentLength, multipartContentType, body, serviceSubmitBatchOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.submitBatchHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// submitBatchCreateRequest creates the SubmitBatch request.
func (client *serviceOperations) submitBatchCreateRequest(contentLength int64, multipartContentType string, body azcore.ReadSeekCloser, serviceSubmitBatchOptions *ServiceSubmitBatchOptions) (*azcore.Request, error) {
	u := client.u
	query := u.Query()
	query.Set("comp", "batch")
	if serviceSubmitBatchOptions != nil && serviceSubmitBatchOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*serviceSubmitBatchOptions.Timeout), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	req.SkipBodyDownload()
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	req.Header.Set("Content-Type", multipartContentType)
	req.Header.Set("x-ms-version", "2019-07-07")
	if serviceSubmitBatchOptions != nil && serviceSubmitBatchOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *serviceSubmitBatchOptions.RequestId)
	}
	return req, req.MarshalAsXML(body)
}

// submitBatchHandleResponse handles the SubmitBatch response.
func (client *serviceOperations) submitBatchHandleResponse(resp *azcore.Response) (*ServiceSubmitBatchResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newStorageError(resp)
	}
	result := ServiceSubmitBatchResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Content-Type"); val != "" {
		result.ContentType = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestId = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return &result, nil
}
