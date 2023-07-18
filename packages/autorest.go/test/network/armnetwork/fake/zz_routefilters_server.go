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

// RouteFiltersServer is a fake server for instances of the armnetwork.RouteFiltersClient type.
type RouteFiltersServer struct {
	// BeginCreateOrUpdate is the fake for method RouteFiltersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters armnetwork.RouteFilter, options *armnetwork.RouteFiltersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.RouteFiltersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method RouteFiltersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, routeFilterName string, options *armnetwork.RouteFiltersClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.RouteFiltersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RouteFiltersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, routeFilterName string, options *armnetwork.RouteFiltersClientGetOptions) (resp azfake.Responder[armnetwork.RouteFiltersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RouteFiltersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armnetwork.RouteFiltersClientListOptions) (resp azfake.PagerResponder[armnetwork.RouteFiltersClientListResponse])

	// NewListByResourceGroupPager is the fake for method RouteFiltersClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armnetwork.RouteFiltersClientListByResourceGroupOptions) (resp azfake.PagerResponder[armnetwork.RouteFiltersClientListByResourceGroupResponse])

	// UpdateTags is the fake for method RouteFiltersClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, routeFilterName string, parameters armnetwork.TagsObject, options *armnetwork.RouteFiltersClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.RouteFiltersClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewRouteFiltersServerTransport creates a new instance of RouteFiltersServerTransport with the provided implementation.
// The returned RouteFiltersServerTransport instance is connected to an instance of armnetwork.RouteFiltersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRouteFiltersServerTransport(srv *RouteFiltersServer) *RouteFiltersServerTransport {
	return &RouteFiltersServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armnetwork.RouteFiltersClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armnetwork.RouteFiltersClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armnetwork.RouteFiltersClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armnetwork.RouteFiltersClientListByResourceGroupResponse]](),
	}
}

// RouteFiltersServerTransport connects instances of armnetwork.RouteFiltersClient to instances of RouteFiltersServer.
// Don't use this type directly, use NewRouteFiltersServerTransport instead.
type RouteFiltersServerTransport struct {
	srv                         *RouteFiltersServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armnetwork.RouteFiltersClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armnetwork.RouteFiltersClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armnetwork.RouteFiltersClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armnetwork.RouteFiltersClientListByResourceGroupResponse]]
}

// Do implements the policy.Transporter interface for RouteFiltersServerTransport.
func (r *RouteFiltersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RouteFiltersClient.BeginCreateOrUpdate":
		resp, err = r.dispatchBeginCreateOrUpdate(req)
	case "RouteFiltersClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "RouteFiltersClient.Get":
		resp, err = r.dispatchGet(req)
	case "RouteFiltersClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	case "RouteFiltersClient.NewListByResourceGroupPager":
		resp, err = r.dispatchNewListByResourceGroupPager(req)
	case "RouteFiltersClient.UpdateTags":
		resp, err = r.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RouteFiltersServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := r.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/routeFilters/(?P<routeFilterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.RouteFilter](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		routeFilterNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("routeFilterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, routeFilterNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		r.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		r.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		r.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (r *RouteFiltersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/routeFilters/(?P<routeFilterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		routeFilterNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("routeFilterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, routeFilterNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		r.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		r.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		r.beginDelete.remove(req)
	}

	return resp, nil
}

func (r *RouteFiltersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/routeFilters/(?P<routeFilterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	routeFilterNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("routeFilterName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armnetwork.RouteFiltersClientGetOptions
	if expandParam != nil {
		options = &armnetwork.RouteFiltersClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameUnescaped, routeFilterNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RouteFilter, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RouteFiltersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/routeFilters`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := r.srv.NewListPager(nil)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.RouteFiltersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}

func (r *RouteFiltersServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := r.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/routeFilters`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByResourceGroupPager(resourceGroupNameUnescaped, nil)
		newListByResourceGroupPager = &resp
		r.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armnetwork.RouteFiltersClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		r.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (r *RouteFiltersServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if r.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/routeFilters/(?P<routeFilterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	routeFilterNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("routeFilterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.UpdateTags(req.Context(), resourceGroupNameUnescaped, routeFilterNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RouteFilter, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
