//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

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

// VirtualRouterPeeringsClient contains the methods for the VirtualRouterPeerings group.
// Don't use this type directly, use NewVirtualRouterPeeringsClient() instead.
type VirtualRouterPeeringsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualRouterPeeringsClient creates a new instance of VirtualRouterPeeringsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualRouterPeeringsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualRouterPeeringsClient, error) {
	cl, err := arm.NewClient(moduleName+".VirtualRouterPeeringsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualRouterPeeringsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the specified Virtual Router Peering.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - virtualRouterName - The name of the Virtual Router.
//   - peeringName - The name of the Virtual Router Peering.
//   - parameters - Parameters supplied to the create or update Virtual Router Peering operation.
//   - options - VirtualRouterPeeringsClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualRouterPeeringsClient.BeginCreateOrUpdate
//     method.
func (client *VirtualRouterPeeringsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, parameters VirtualRouterPeering, options *VirtualRouterPeeringsClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualRouterPeeringsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		var err error
		var endSpan func(error)
		ctx, endSpan = runtime.StartSpan(ctx, "VirtualRouterPeeringsClient.BeginCreateOrUpdate", client.internal.Tracer(), nil)
		defer func() { endSpan(err) }()
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualRouterName, peeringName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualRouterPeeringsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[VirtualRouterPeeringsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates the specified Virtual Router Peering.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VirtualRouterPeeringsClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, parameters VirtualRouterPeering, options *VirtualRouterPeeringsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualRouterName, peeringName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualRouterPeeringsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, parameters VirtualRouterPeering, options *VirtualRouterPeeringsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualRouterName == "" {
		return nil, errors.New("parameter virtualRouterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified peering from a Virtual Router.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - virtualRouterName - The name of the Virtual Router.
//   - peeringName - The name of the peering.
//   - options - VirtualRouterPeeringsClientBeginDeleteOptions contains the optional parameters for the VirtualRouterPeeringsClient.BeginDelete
//     method.
func (client *VirtualRouterPeeringsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *VirtualRouterPeeringsClientBeginDeleteOptions) (*runtime.Poller[VirtualRouterPeeringsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		var err error
		var endSpan func(error)
		ctx, endSpan = runtime.StartSpan(ctx, "VirtualRouterPeeringsClient.BeginDelete", client.internal.Tracer(), nil)
		defer func() { endSpan(err) }()
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualRouterName, peeringName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualRouterPeeringsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[VirtualRouterPeeringsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified peering from a Virtual Router.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VirtualRouterPeeringsClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *VirtualRouterPeeringsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualRouterName, peeringName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualRouterPeeringsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *VirtualRouterPeeringsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualRouterName == "" {
		return nil, errors.New("parameter virtualRouterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified Virtual Router Peering.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - virtualRouterName - The name of the Virtual Router.
//   - peeringName - The name of the Virtual Router Peering.
//   - options - VirtualRouterPeeringsClientGetOptions contains the optional parameters for the VirtualRouterPeeringsClient.Get
//     method.
func (client *VirtualRouterPeeringsClient) Get(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *VirtualRouterPeeringsClientGetOptions) (VirtualRouterPeeringsClientGetResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "VirtualRouterPeeringsClient.Get", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualRouterName, peeringName, options)
	if err != nil {
		return VirtualRouterPeeringsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualRouterPeeringsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualRouterPeeringsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VirtualRouterPeeringsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *VirtualRouterPeeringsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualRouterName == "" {
		return nil, errors.New("parameter virtualRouterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualRouterPeeringsClient) getHandleResponse(resp *http.Response) (VirtualRouterPeeringsClientGetResponse, error) {
	result := VirtualRouterPeeringsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualRouterPeering); err != nil {
		return VirtualRouterPeeringsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all Virtual Router Peerings in a Virtual Router resource.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - virtualRouterName - The name of the Virtual Router.
//   - options - VirtualRouterPeeringsClientListOptions contains the optional parameters for the VirtualRouterPeeringsClient.NewListPager
//     method.
func (client *VirtualRouterPeeringsClient) NewListPager(resourceGroupName string, virtualRouterName string, options *VirtualRouterPeeringsClientListOptions) *runtime.Pager[VirtualRouterPeeringsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualRouterPeeringsClientListResponse]{
		More: func(page VirtualRouterPeeringsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualRouterPeeringsClientListResponse) (VirtualRouterPeeringsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, virtualRouterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualRouterPeeringsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VirtualRouterPeeringsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualRouterPeeringsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *VirtualRouterPeeringsClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualRouterName string, options *VirtualRouterPeeringsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualRouterName == "" {
		return nil, errors.New("parameter virtualRouterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualRouterPeeringsClient) listHandleResponse(resp *http.Response) (VirtualRouterPeeringsClientListResponse, error) {
	result := VirtualRouterPeeringsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualRouterPeeringListResult); err != nil {
		return VirtualRouterPeeringsClientListResponse{}, err
	}
	return result, nil
}
