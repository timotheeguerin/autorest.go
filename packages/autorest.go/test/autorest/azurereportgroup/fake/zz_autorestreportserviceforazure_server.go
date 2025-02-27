//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/azurereportgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
)

// AutoRestReportServiceForAzureServer is a fake server for instances of the azurereportgroup.AutoRestReportServiceForAzureClient type.
type AutoRestReportServiceForAzureServer struct {
	// GetReport is the fake for method AutoRestReportServiceForAzureClient.GetReport
	// HTTP status codes to indicate success: http.StatusOK
	GetReport func(ctx context.Context, options *azurereportgroup.AutoRestReportServiceForAzureClientGetReportOptions) (resp azfake.Responder[azurereportgroup.AutoRestReportServiceForAzureClientGetReportResponse], errResp azfake.ErrorResponder)
}

// NewAutoRestReportServiceForAzureServerTransport creates a new instance of AutoRestReportServiceForAzureServerTransport with the provided implementation.
// The returned AutoRestReportServiceForAzureServerTransport instance is connected to an instance of azurereportgroup.AutoRestReportServiceForAzureClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAutoRestReportServiceForAzureServerTransport(srv *AutoRestReportServiceForAzureServer) *AutoRestReportServiceForAzureServerTransport {
	return &AutoRestReportServiceForAzureServerTransport{srv: srv}
}

// AutoRestReportServiceForAzureServerTransport connects instances of azurereportgroup.AutoRestReportServiceForAzureClient to instances of AutoRestReportServiceForAzureServer.
// Don't use this type directly, use NewAutoRestReportServiceForAzureServerTransport instead.
type AutoRestReportServiceForAzureServerTransport struct {
	srv *AutoRestReportServiceForAzureServer
}

// Do implements the policy.Transporter interface for AutoRestReportServiceForAzureServerTransport.
func (a *AutoRestReportServiceForAzureServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AutoRestReportServiceForAzureClient.GetReport":
		resp, err = a.dispatchGetReport(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AutoRestReportServiceForAzureServerTransport) dispatchGetReport(req *http.Request) (*http.Response, error) {
	if a.srv.GetReport == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetReport not implemented")}
	}
	qp := req.URL.Query()
	qualifierUnescaped, err := url.QueryUnescape(qp.Get("qualifier"))
	if err != nil {
		return nil, err
	}
	qualifierParam := getOptional(qualifierUnescaped)
	var options *azurereportgroup.AutoRestReportServiceForAzureClientGetReportOptions
	if qualifierParam != nil {
		options = &azurereportgroup.AutoRestReportServiceForAzureClientGetReportOptions{
			Qualifier: qualifierParam,
		}
	}
	respr, errRespr := a.srv.GetReport(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
