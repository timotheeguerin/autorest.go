//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CommunityGalleryImageVersionsClient contains the methods for the CommunityGalleryImageVersions group.
// Don't use this type directly, use NewCommunityGalleryImageVersionsClient() instead.
type CommunityGalleryImageVersionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCommunityGalleryImageVersionsClient creates a new instance of CommunityGalleryImageVersionsClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCommunityGalleryImageVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CommunityGalleryImageVersionsClient, error) {
	cl, err := arm.NewClient(moduleName+".CommunityGalleryImageVersionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CommunityGalleryImageVersionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a community gallery image version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - location - Resource location.
//   - publicGalleryName - The public name of the community gallery.
//   - galleryImageName - The name of the community gallery image definition.
//   - galleryImageVersionName - The name of the community gallery image version. Needs to follow semantic version name pattern:
//     The allowed characters are digit and period. Digits must be within the range of a 32-bit integer.
//     Format: ..
//   - options - CommunityGalleryImageVersionsClientGetOptions contains the optional parameters for the CommunityGalleryImageVersionsClient.Get
//     method.
func (client *CommunityGalleryImageVersionsClient) Get(ctx context.Context, location string, publicGalleryName string, galleryImageName string, galleryImageVersionName string, options *CommunityGalleryImageVersionsClientGetOptions) (CommunityGalleryImageVersionsClientGetResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "CommunityGalleryImageVersionsClient.Get", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, publicGalleryName, galleryImageName, galleryImageVersionName, options)
	if err != nil {
		return CommunityGalleryImageVersionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CommunityGalleryImageVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CommunityGalleryImageVersionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CommunityGalleryImageVersionsClient) getCreateRequest(ctx context.Context, location string, publicGalleryName string, galleryImageName string, galleryImageVersionName string, options *CommunityGalleryImageVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/communityGalleries/{publicGalleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publicGalleryName == "" {
		return nil, errors.New("parameter publicGalleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicGalleryName}", url.PathEscape(publicGalleryName))
	if galleryImageName == "" {
		return nil, errors.New("parameter galleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	if galleryImageVersionName == "" {
		return nil, errors.New("parameter galleryImageVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageVersionName}", url.PathEscape(galleryImageVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CommunityGalleryImageVersionsClient) getHandleResponse(resp *http.Response) (CommunityGalleryImageVersionsClientGetResponse, error) {
	result := CommunityGalleryImageVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommunityGalleryImageVersion); err != nil {
		return CommunityGalleryImageVersionsClientGetResponse{}, err
	}
	return result, nil
}
