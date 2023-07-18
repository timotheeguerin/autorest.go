//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armnetwork"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"net/url"
	"regexp"
)

// SubnetsServer is a fake server for instances of the armnetwork.SubnetsClient type.
type SubnetsServer struct {
	// BeginCreateOrUpdate is the fake for method SubnetsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters armnetwork.Subnet, options *armnetwork.SubnetsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.SubnetsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method SubnetsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *armnetwork.SubnetsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.SubnetsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SubnetsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *armnetwork.SubnetsClientGetOptions) (resp azfake.Responder[armnetwork.SubnetsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SubnetsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, virtualNetworkName string, options *armnetwork.SubnetsClientListOptions) (resp azfake.PagerResponder[armnetwork.SubnetsClientListResponse])

	// BeginPrepareNetworkPolicies is the fake for method SubnetsClient.BeginPrepareNetworkPolicies
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPrepareNetworkPolicies func(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, prepareNetworkPoliciesRequestParameters armnetwork.PrepareNetworkPoliciesRequest, options *armnetwork.SubnetsClientBeginPrepareNetworkPoliciesOptions) (resp azfake.PollerResponder[armnetwork.SubnetsClientPrepareNetworkPoliciesResponse], errResp azfake.ErrorResponder)

	// BeginUnprepareNetworkPolicies is the fake for method SubnetsClient.BeginUnprepareNetworkPolicies
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUnprepareNetworkPolicies func(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, unprepareNetworkPoliciesRequestParameters armnetwork.UnprepareNetworkPoliciesRequest, options *armnetwork.SubnetsClientBeginUnprepareNetworkPoliciesOptions) (resp azfake.PollerResponder[armnetwork.SubnetsClientUnprepareNetworkPoliciesResponse], errResp azfake.ErrorResponder)
}

// NewSubnetsServerTransport creates a new instance of SubnetsServerTransport with the provided implementation.
// The returned SubnetsServerTransport instance is connected to an instance of armnetwork.SubnetsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSubnetsServerTransport(srv *SubnetsServer) *SubnetsServerTransport {
	return &SubnetsServerTransport{
		srv:                           srv,
		beginCreateOrUpdate:           newTracker[azfake.PollerResponder[armnetwork.SubnetsClientCreateOrUpdateResponse]](),
		beginDelete:                   newTracker[azfake.PollerResponder[armnetwork.SubnetsClientDeleteResponse]](),
		newListPager:                  newTracker[azfake.PagerResponder[armnetwork.SubnetsClientListResponse]](),
		beginPrepareNetworkPolicies:   newTracker[azfake.PollerResponder[armnetwork.SubnetsClientPrepareNetworkPoliciesResponse]](),
		beginUnprepareNetworkPolicies: newTracker[azfake.PollerResponder[armnetwork.SubnetsClientUnprepareNetworkPoliciesResponse]](),
	}
}

// SubnetsServerTransport connects instances of armnetwork.SubnetsClient to instances of SubnetsServer.
// Don't use this type directly, use NewSubnetsServerTransport instead.
type SubnetsServerTransport struct {
	srv                           *SubnetsServer
	beginCreateOrUpdate           *tracker[azfake.PollerResponder[armnetwork.SubnetsClientCreateOrUpdateResponse]]
	beginDelete                   *tracker[azfake.PollerResponder[armnetwork.SubnetsClientDeleteResponse]]
	newListPager                  *tracker[azfake.PagerResponder[armnetwork.SubnetsClientListResponse]]
	beginPrepareNetworkPolicies   *tracker[azfake.PollerResponder[armnetwork.SubnetsClientPrepareNetworkPoliciesResponse]]
	beginUnprepareNetworkPolicies *tracker[azfake.PollerResponder[armnetwork.SubnetsClientUnprepareNetworkPoliciesResponse]]
}

// Do implements the policy.Transporter interface for SubnetsServerTransport.
func (s *SubnetsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SubnetsClient.BeginCreateOrUpdate":
		resp, err = s.dispatchBeginCreateOrUpdate(req)
	case "SubnetsClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "SubnetsClient.Get":
		resp, err = s.dispatchGet(req)
	case "SubnetsClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	case "SubnetsClient.BeginPrepareNetworkPolicies":
		resp, err = s.dispatchBeginPrepareNetworkPolicies(req)
	case "SubnetsClient.BeginUnprepareNetworkPolicies":
		resp, err = s.dispatchBeginUnprepareNetworkPolicies(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SubnetsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := s.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subnets/(?P<subnetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.Subnet](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
		if err != nil {
			return nil, err
		}
		subnetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("subnetName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, virtualNetworkNameUnescaped, subnetNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		s.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		s.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (s *SubnetsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subnets/(?P<subnetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
		if err != nil {
			return nil, err
		}
		subnetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("subnetName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, virtualNetworkNameUnescaped, subnetNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *SubnetsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subnets/(?P<subnetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
	if err != nil {
		return nil, err
	}
	subnetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("subnetName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armnetwork.SubnetsClientGetOptions
	if expandParam != nil {
		options = &armnetwork.SubnetsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameUnescaped, virtualNetworkNameUnescaped, subnetNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Subnet, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SubnetsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subnets`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListPager(resourceGroupNameUnescaped, virtualNetworkNameUnescaped, nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.SubnetsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}

func (s *SubnetsServerTransport) dispatchBeginPrepareNetworkPolicies(req *http.Request) (*http.Response, error) {
	if s.srv.BeginPrepareNetworkPolicies == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPrepareNetworkPolicies not implemented")}
	}
	beginPrepareNetworkPolicies := s.beginPrepareNetworkPolicies.get(req)
	if beginPrepareNetworkPolicies == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subnets/(?P<subnetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/PrepareNetworkPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.PrepareNetworkPoliciesRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
		if err != nil {
			return nil, err
		}
		subnetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("subnetName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginPrepareNetworkPolicies(req.Context(), resourceGroupNameUnescaped, virtualNetworkNameUnescaped, subnetNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPrepareNetworkPolicies = &respr
		s.beginPrepareNetworkPolicies.add(req, beginPrepareNetworkPolicies)
	}

	resp, err := server.PollerResponderNext(beginPrepareNetworkPolicies, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginPrepareNetworkPolicies.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPrepareNetworkPolicies) {
		s.beginPrepareNetworkPolicies.remove(req)
	}

	return resp, nil
}

func (s *SubnetsServerTransport) dispatchBeginUnprepareNetworkPolicies(req *http.Request) (*http.Response, error) {
	if s.srv.BeginUnprepareNetworkPolicies == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUnprepareNetworkPolicies not implemented")}
	}
	beginUnprepareNetworkPolicies := s.beginUnprepareNetworkPolicies.get(req)
	if beginUnprepareNetworkPolicies == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subnets/(?P<subnetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/UnprepareNetworkPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.UnprepareNetworkPoliciesRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
		if err != nil {
			return nil, err
		}
		subnetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("subnetName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginUnprepareNetworkPolicies(req.Context(), resourceGroupNameUnescaped, virtualNetworkNameUnescaped, subnetNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUnprepareNetworkPolicies = &respr
		s.beginUnprepareNetworkPolicies.add(req, beginUnprepareNetworkPolicies)
	}

	resp, err := server.PollerResponderNext(beginUnprepareNetworkPolicies, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginUnprepareNetworkPolicies.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUnprepareNetworkPolicies) {
		s.beginUnprepareNetworkPolicies.remove(req)
	}

	return resp, nil
}
