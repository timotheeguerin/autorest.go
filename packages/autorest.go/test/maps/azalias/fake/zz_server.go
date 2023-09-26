//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"azalias"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Server is a fake server for instances of the azalias.Client type.
type Server struct {
	// Create is the fake for method Client.Create
	// HTTP status codes to indicate success: http.StatusCreated
	Create func(ctx context.Context, options *azalias.ClientCreateOptions) (resp azfake.Responder[azalias.ClientCreateResponse], errResp azfake.ErrorResponder)

	// GetScript is the fake for method Client.GetScript
	// HTTP status codes to indicate success: http.StatusOK
	GetScript func(ctx context.Context, headerCounts []int32, queryCounts []int64, headerTime time.Time, props azalias.GeoJSONObjectNamedCollection, explodedGroup azalias.ExplodedGroup, options *azalias.ClientGetScriptOptions) (resp azfake.Responder[azalias.ClientGetScriptResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method Client.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *azalias.ClientListOptions) (resp azfake.PagerResponder[azalias.ClientListResponse])

	// NewListWithSharedNextOnePager is the fake for method Client.NewListWithSharedNextOnePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListWithSharedNextOnePager func(options *azalias.ClientListWithSharedNextOneOptions) (resp azfake.PagerResponder[azalias.ClientListWithSharedNextOneResponse])

	// NewListWithSharedNextTwoPager is the fake for method Client.NewListWithSharedNextTwoPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListWithSharedNextTwoPager func(options *azalias.ClientListWithSharedNextTwoOptions) (resp azfake.PagerResponder[azalias.ClientListWithSharedNextTwoResponse])

	// PolicyAssignment is the fake for method Client.PolicyAssignment
	// HTTP status codes to indicate success: http.StatusOK
	PolicyAssignment func(ctx context.Context, things []azalias.Things, polymorphicParam azalias.GeoJSONObjectClassification, options *azalias.ClientPolicyAssignmentOptions) (resp azfake.Responder[azalias.ClientPolicyAssignmentResponse], errResp azfake.ErrorResponder)
}

// NewServerTransport creates a new instance of ServerTransport with the provided implementation.
// The returned ServerTransport instance is connected to an instance of azalias.Client via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerTransport(srv *Server) *ServerTransport {
	return &ServerTransport{
		srv:                           srv,
		newListPager:                  newTracker[azfake.PagerResponder[azalias.ClientListResponse]](),
		newListWithSharedNextOnePager: newTracker[azfake.PagerResponder[azalias.ClientListWithSharedNextOneResponse]](),
		newListWithSharedNextTwoPager: newTracker[azfake.PagerResponder[azalias.ClientListWithSharedNextTwoResponse]](),
	}
}

// ServerTransport connects instances of azalias.Client to instances of Server.
// Don't use this type directly, use NewServerTransport instead.
type ServerTransport struct {
	srv                           *Server
	newListPager                  *tracker[azfake.PagerResponder[azalias.ClientListResponse]]
	newListWithSharedNextOnePager *tracker[azfake.PagerResponder[azalias.ClientListWithSharedNextOneResponse]]
	newListWithSharedNextTwoPager *tracker[azfake.PagerResponder[azalias.ClientListWithSharedNextTwoResponse]]
}

// Do implements the policy.Transporter interface for ServerTransport.
func (s *ServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "Client.Create":
		resp, err = s.dispatchCreate(req)
	case "Client.GetScript":
		resp, err = s.dispatchGetScript(req)
	case "Client.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	case "Client.NewListWithSharedNextOnePager":
		resp, err = s.dispatchNewListWithSharedNextOnePager(req)
	case "Client.NewListWithSharedNextTwoPager":
		resp, err = s.dispatchNewListWithSharedNextTwoPager(req)
	case "Client.PolicyAssignment":
		resp, err = s.dispatchPolicyAssignment(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if s.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	qp := req.URL.Query()
	creatorIDUnescaped, err := url.QueryUnescape(qp.Get("creator-id"))
	if err != nil {
		return nil, err
	}
	creatorIDParam, err := parseOptional(creatorIDUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	assignedIDParam, err := parseOptional(getHeaderValue(req.Header, "assigned-id"), func(v string) (float32, error) {
		p, parseErr := strconv.ParseFloat(v, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return float32(p), nil
	})
	if err != nil {
		return nil, err
	}
	groupByEscaped := qp["groupBy"]
	groupByUnescaped := make([]string, len(groupByEscaped))
	for i, v := range groupByEscaped {
		u, unescapeErr := url.QueryUnescape(v)
		if unescapeErr != nil {
			return nil, unescapeErr
		}
		groupByUnescaped[i] = u
	}
	groupByUnescapedElements := groupByUnescaped
	groupByParam := make([]azalias.SomethingCount, len(groupByUnescapedElements))
	for i := 0; i < len(groupByUnescapedElements); i++ {
		var parsedInt int64
		parsedInt, err = strconv.ParseInt(groupByUnescapedElements[i], 10, 32)
		if err != nil {
			return nil, err
		}
		groupByParam[i] = azalias.SomethingCount(parsedInt)
	}
	var options *azalias.ClientCreateOptions
	if creatorIDParam != nil || assignedIDParam != nil || len(groupByParam) > 0 {
		options = &azalias.ClientCreateOptions{
			CreatorID:  creatorIDParam,
			AssignedID: assignedIDParam,
			GroupBy:    groupByParam,
		}
	}
	respr, errRespr := s.srv.Create(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AliasesCreateResponse, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).AccessControlExposeHeaders; val != nil {
		resp.Header.Set("Access-Control-Expose-Headers", *val)
	}
	return resp, nil
}

func (s *ServerTransport) dispatchGetScript(req *http.Request) (*http.Response, error) {
	if s.srv.GetScript == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetScript not implemented")}
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[azalias.GeoJSONObjectNamedCollection](req)
	if err != nil {
		return nil, err
	}
	headerCountsHeader := getHeaderValue(req.Header, "headerCounts")
	headerCountsHeaderElements := splitHelper(headerCountsHeader, ",")
	headerCountsParam := make([]int32, len(headerCountsHeaderElements))
	for i := 0; i < len(headerCountsHeaderElements); i++ {
		var parsedInt int64
		parsedInt, err = strconv.ParseInt(headerCountsHeaderElements[i], 10, 32)
		if err != nil {
			return nil, err
		}
		headerCountsParam[i] = int32(parsedInt)
	}
	queryCountsUnescaped, err := url.QueryUnescape(qp.Get("queryCounts"))
	if err != nil {
		return nil, err
	}
	queryCountsUnescapedElements := splitHelper(queryCountsUnescaped, ",")
	queryCountsParam := make([]int64, len(queryCountsUnescapedElements))
	for i := 0; i < len(queryCountsUnescapedElements); i++ {
		var parsedInt int64
		parsedInt, err = strconv.ParseInt(queryCountsUnescapedElements[i], 10, 32)
		if err != nil {
			return nil, err
		}
		queryCountsParam[i] = int64(parsedInt)
	}
	explodedStuffEscaped := qp["explodedStuff"]
	explodedStuffUnescaped := make([]string, len(explodedStuffEscaped))
	for i, v := range explodedStuffEscaped {
		u, unescapeErr := url.QueryUnescape(v)
		if unescapeErr != nil {
			return nil, unescapeErr
		}
		explodedStuffUnescaped[i] = u
	}
	explodedStuffUnescapedElements := explodedStuffUnescaped
	explodedStuffParam := make([]int64, len(explodedStuffUnescapedElements))
	for i := 0; i < len(explodedStuffUnescapedElements); i++ {
		var parsedInt int64
		parsedInt, err = strconv.ParseInt(explodedStuffUnescapedElements[i], 10, 32)
		if err != nil {
			return nil, err
		}
		explodedStuffParam[i] = int64(parsedInt)
	}
	headerTimeParam, err := time.Parse("15:04:05.999999999Z07:00", getHeaderValue(req.Header, "headerTime"))
	if err != nil {
		return nil, err
	}
	explodedGroup := azalias.ExplodedGroup{
		ExplodedStuff: explodedStuffParam,
	}
	respr, errRespr := s.srv.GetScript(req.Context(), headerCountsParam, queryCountsParam, headerTimeParam, body, explodedGroup, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsText(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		qp := req.URL.Query()
		groupByEscaped := qp["groupBy"]
		groupByUnescaped := make([]string, len(groupByEscaped))
		for i, v := range groupByEscaped {
			u, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return nil, unescapeErr
			}
			groupByUnescaped[i] = u
		}
		groupByUnescapedElements := groupByUnescaped
		groupByParam := make([]azalias.LogMetricsGroupBy, len(groupByUnescapedElements))
		for i := 0; i < len(groupByUnescapedElements); i++ {
			groupByParam[i] = azalias.LogMetricsGroupBy(groupByUnescapedElements[i])
		}
		var options *azalias.ClientListOptions
		if len(groupByParam) > 0 {
			options = &azalias.ClientListOptions{
				GroupBy: groupByParam,
			}
		}
		resp := s.srv.NewListPager(options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *azalias.ClientListResponse, createLink func() string) {
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

func (s *ServerTransport) dispatchNewListWithSharedNextOnePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListWithSharedNextOnePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListWithSharedNextOnePager not implemented")}
	}
	newListWithSharedNextOnePager := s.newListWithSharedNextOnePager.get(req)
	if newListWithSharedNextOnePager == nil {
		resp := s.srv.NewListWithSharedNextOnePager(nil)
		newListWithSharedNextOnePager = &resp
		s.newListWithSharedNextOnePager.add(req, newListWithSharedNextOnePager)
		server.PagerResponderInjectNextLinks(newListWithSharedNextOnePager, req, func(page *azalias.ClientListWithSharedNextOneResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListWithSharedNextOnePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListWithSharedNextOnePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListWithSharedNextOnePager) {
		s.newListWithSharedNextOnePager.remove(req)
	}
	return resp, nil
}

func (s *ServerTransport) dispatchNewListWithSharedNextTwoPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListWithSharedNextTwoPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListWithSharedNextTwoPager not implemented")}
	}
	newListWithSharedNextTwoPager := s.newListWithSharedNextTwoPager.get(req)
	if newListWithSharedNextTwoPager == nil {
		resp := s.srv.NewListWithSharedNextTwoPager(nil)
		newListWithSharedNextTwoPager = &resp
		s.newListWithSharedNextTwoPager.add(req, newListWithSharedNextTwoPager)
		server.PagerResponderInjectNextLinks(newListWithSharedNextTwoPager, req, func(page *azalias.ClientListWithSharedNextTwoResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListWithSharedNextTwoPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListWithSharedNextTwoPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListWithSharedNextTwoPager) {
		s.newListWithSharedNextTwoPager.remove(req)
	}
	return resp, nil
}

func (s *ServerTransport) dispatchPolicyAssignment(req *http.Request) (*http.Response, error) {
	if s.srv.PolicyAssignment == nil {
		return nil, &nonRetriableError{errors.New("fake for method PolicyAssignment not implemented")}
	}
	qp := req.URL.Query()
	raw, err := readRequestBody(req)
	if err != nil {
		return nil, err
	}
	body, err := unmarshalGeoJSONObjectClassification(raw)
	if err != nil {
		return nil, err
	}
	thingsUnescaped, err := url.QueryUnescape(qp.Get("things"))
	if err != nil {
		return nil, err
	}
	thingsUnescapedElements := splitHelper(thingsUnescaped, ",")
	thingsParam := make([]azalias.Things, len(thingsUnescapedElements))
	for i := 0; i < len(thingsUnescapedElements); i++ {
		thingsParam[i] = azalias.Things(thingsUnescapedElements[i])
	}
	intervalUnescaped, err := url.QueryUnescape(qp.Get("interval"))
	if err != nil {
		return nil, err
	}
	intervalParam := getOptional(intervalUnescaped)
	uniqueUnescaped, err := url.QueryUnescape(qp.Get("unique"))
	if err != nil {
		return nil, err
	}
	uniqueParam := getOptional(uniqueUnescaped)
	var options *azalias.ClientPolicyAssignmentOptions
	if intervalParam != nil || uniqueParam != nil {
		options = &azalias.ClientPolicyAssignmentOptions{
			Interval: intervalParam,
			Unique:   uniqueParam,
		}
	}
	respr, errRespr := s.srv.PolicyAssignment(req.Context(), thingsParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyAssignmentProperties, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
