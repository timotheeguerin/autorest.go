//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armcompute"
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

// GalleryApplicationVersionsServer is a fake server for instances of the armcompute.GalleryApplicationVersionsClient type.
type GalleryApplicationVersionsServer struct {
	// BeginCreateOrUpdate is the fake for method GalleryApplicationVersionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion armcompute.GalleryApplicationVersion, options *armcompute.GalleryApplicationVersionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.GalleryApplicationVersionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method GalleryApplicationVersionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *armcompute.GalleryApplicationVersionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.GalleryApplicationVersionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method GalleryApplicationVersionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *armcompute.GalleryApplicationVersionsClientGetOptions) (resp azfake.Responder[armcompute.GalleryApplicationVersionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByGalleryApplicationPager is the fake for method GalleryApplicationVersionsClient.NewListByGalleryApplicationPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByGalleryApplicationPager func(resourceGroupName string, galleryName string, galleryApplicationName string, options *armcompute.GalleryApplicationVersionsClientListByGalleryApplicationOptions) (resp azfake.PagerResponder[armcompute.GalleryApplicationVersionsClientListByGalleryApplicationResponse])

	// BeginUpdate is the fake for method GalleryApplicationVersionsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginUpdate func(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion armcompute.GalleryApplicationVersionUpdate, options *armcompute.GalleryApplicationVersionsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.GalleryApplicationVersionsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewGalleryApplicationVersionsServerTransport creates a new instance of GalleryApplicationVersionsServerTransport with the provided implementation.
// The returned GalleryApplicationVersionsServerTransport instance is connected to an instance of armcompute.GalleryApplicationVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGalleryApplicationVersionsServerTransport(srv *GalleryApplicationVersionsServer) *GalleryApplicationVersionsServerTransport {
	return &GalleryApplicationVersionsServerTransport{
		srv:                              srv,
		beginCreateOrUpdate:              newTracker[azfake.PollerResponder[armcompute.GalleryApplicationVersionsClientCreateOrUpdateResponse]](),
		beginDelete:                      newTracker[azfake.PollerResponder[armcompute.GalleryApplicationVersionsClientDeleteResponse]](),
		newListByGalleryApplicationPager: newTracker[azfake.PagerResponder[armcompute.GalleryApplicationVersionsClientListByGalleryApplicationResponse]](),
		beginUpdate:                      newTracker[azfake.PollerResponder[armcompute.GalleryApplicationVersionsClientUpdateResponse]](),
	}
}

// GalleryApplicationVersionsServerTransport connects instances of armcompute.GalleryApplicationVersionsClient to instances of GalleryApplicationVersionsServer.
// Don't use this type directly, use NewGalleryApplicationVersionsServerTransport instead.
type GalleryApplicationVersionsServerTransport struct {
	srv                              *GalleryApplicationVersionsServer
	beginCreateOrUpdate              *tracker[azfake.PollerResponder[armcompute.GalleryApplicationVersionsClientCreateOrUpdateResponse]]
	beginDelete                      *tracker[azfake.PollerResponder[armcompute.GalleryApplicationVersionsClientDeleteResponse]]
	newListByGalleryApplicationPager *tracker[azfake.PagerResponder[armcompute.GalleryApplicationVersionsClientListByGalleryApplicationResponse]]
	beginUpdate                      *tracker[azfake.PollerResponder[armcompute.GalleryApplicationVersionsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for GalleryApplicationVersionsServerTransport.
func (g *GalleryApplicationVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GalleryApplicationVersionsClient.BeginCreateOrUpdate":
		resp, err = g.dispatchBeginCreateOrUpdate(req)
	case "GalleryApplicationVersionsClient.BeginDelete":
		resp, err = g.dispatchBeginDelete(req)
	case "GalleryApplicationVersionsClient.Get":
		resp, err = g.dispatchGet(req)
	case "GalleryApplicationVersionsClient.NewListByGalleryApplicationPager":
		resp, err = g.dispatchNewListByGalleryApplicationPager(req)
	case "GalleryApplicationVersionsClient.BeginUpdate":
		resp, err = g.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GalleryApplicationVersionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := g.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<galleryApplicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<galleryApplicationVersionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.GalleryApplicationVersion](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		galleryNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
		if err != nil {
			return nil, err
		}
		galleryApplicationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryApplicationName")])
		if err != nil {
			return nil, err
		}
		galleryApplicationVersionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryApplicationVersionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, galleryNameUnescaped, galleryApplicationNameUnescaped, galleryApplicationVersionNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		g.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		g.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		g.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (g *GalleryApplicationVersionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if g.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := g.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<galleryApplicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<galleryApplicationVersionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		galleryNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
		if err != nil {
			return nil, err
		}
		galleryApplicationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryApplicationName")])
		if err != nil {
			return nil, err
		}
		galleryApplicationVersionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryApplicationVersionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, galleryNameUnescaped, galleryApplicationNameUnescaped, galleryApplicationVersionNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		g.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		g.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		g.beginDelete.remove(req)
	}

	return resp, nil
}

func (g *GalleryApplicationVersionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<galleryApplicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<galleryApplicationVersionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	galleryNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
	if err != nil {
		return nil, err
	}
	galleryApplicationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryApplicationName")])
	if err != nil {
		return nil, err
	}
	galleryApplicationVersionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryApplicationVersionName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(armcompute.ReplicationStatusTypes(expandUnescaped))
	var options *armcompute.GalleryApplicationVersionsClientGetOptions
	if expandParam != nil {
		options = &armcompute.GalleryApplicationVersionsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := g.srv.Get(req.Context(), resourceGroupNameUnescaped, galleryNameUnescaped, galleryApplicationNameUnescaped, galleryApplicationVersionNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GalleryApplicationVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GalleryApplicationVersionsServerTransport) dispatchNewListByGalleryApplicationPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListByGalleryApplicationPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByGalleryApplicationPager not implemented")}
	}
	newListByGalleryApplicationPager := g.newListByGalleryApplicationPager.get(req)
	if newListByGalleryApplicationPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<galleryApplicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		galleryNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
		if err != nil {
			return nil, err
		}
		galleryApplicationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryApplicationName")])
		if err != nil {
			return nil, err
		}
		resp := g.srv.NewListByGalleryApplicationPager(resourceGroupNameUnescaped, galleryNameUnescaped, galleryApplicationNameUnescaped, nil)
		newListByGalleryApplicationPager = &resp
		g.newListByGalleryApplicationPager.add(req, newListByGalleryApplicationPager)
		server.PagerResponderInjectNextLinks(newListByGalleryApplicationPager, req, func(page *armcompute.GalleryApplicationVersionsClientListByGalleryApplicationResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByGalleryApplicationPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListByGalleryApplicationPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByGalleryApplicationPager) {
		g.newListByGalleryApplicationPager.remove(req)
	}
	return resp, nil
}

func (g *GalleryApplicationVersionsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := g.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<galleryApplicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<galleryApplicationVersionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.GalleryApplicationVersionUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		galleryNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
		if err != nil {
			return nil, err
		}
		galleryApplicationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryApplicationName")])
		if err != nil {
			return nil, err
		}
		galleryApplicationVersionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryApplicationVersionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginUpdate(req.Context(), resourceGroupNameUnescaped, galleryNameUnescaped, galleryApplicationNameUnescaped, galleryApplicationVersionNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		g.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		g.beginUpdate.remove(req)
	}

	return resp, nil
}
