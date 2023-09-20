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
	"generatortests/optionalgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"io"
	"net/http"
	"reflect"
	"strconv"
)

// ExplicitServer is a fake server for instances of the optionalgroup.ExplicitClient type.
type ExplicitServer struct {
	// PostOptionalArrayHeader is the fake for method ExplicitClient.PostOptionalArrayHeader
	// HTTP status codes to indicate success: http.StatusOK
	PostOptionalArrayHeader func(ctx context.Context, options *optionalgroup.ExplicitClientPostOptionalArrayHeaderOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostOptionalArrayHeaderResponse], errResp azfake.ErrorResponder)

	// PostOptionalArrayParameter is the fake for method ExplicitClient.PostOptionalArrayParameter
	// HTTP status codes to indicate success: http.StatusOK
	PostOptionalArrayParameter func(ctx context.Context, options *optionalgroup.ExplicitClientPostOptionalArrayParameterOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostOptionalArrayParameterResponse], errResp azfake.ErrorResponder)

	// PostOptionalArrayProperty is the fake for method ExplicitClient.PostOptionalArrayProperty
	// HTTP status codes to indicate success: http.StatusOK
	PostOptionalArrayProperty func(ctx context.Context, options *optionalgroup.ExplicitClientPostOptionalArrayPropertyOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostOptionalArrayPropertyResponse], errResp azfake.ErrorResponder)

	// PostOptionalClassParameter is the fake for method ExplicitClient.PostOptionalClassParameter
	// HTTP status codes to indicate success: http.StatusOK
	PostOptionalClassParameter func(ctx context.Context, options *optionalgroup.ExplicitClientPostOptionalClassParameterOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostOptionalClassParameterResponse], errResp azfake.ErrorResponder)

	// PostOptionalClassProperty is the fake for method ExplicitClient.PostOptionalClassProperty
	// HTTP status codes to indicate success: http.StatusOK
	PostOptionalClassProperty func(ctx context.Context, options *optionalgroup.ExplicitClientPostOptionalClassPropertyOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostOptionalClassPropertyResponse], errResp azfake.ErrorResponder)

	// PostOptionalIntegerHeader is the fake for method ExplicitClient.PostOptionalIntegerHeader
	// HTTP status codes to indicate success: http.StatusOK
	PostOptionalIntegerHeader func(ctx context.Context, options *optionalgroup.ExplicitClientPostOptionalIntegerHeaderOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostOptionalIntegerHeaderResponse], errResp azfake.ErrorResponder)

	// PostOptionalIntegerParameter is the fake for method ExplicitClient.PostOptionalIntegerParameter
	// HTTP status codes to indicate success: http.StatusOK
	PostOptionalIntegerParameter func(ctx context.Context, options *optionalgroup.ExplicitClientPostOptionalIntegerParameterOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostOptionalIntegerParameterResponse], errResp azfake.ErrorResponder)

	// PostOptionalIntegerProperty is the fake for method ExplicitClient.PostOptionalIntegerProperty
	// HTTP status codes to indicate success: http.StatusOK
	PostOptionalIntegerProperty func(ctx context.Context, options *optionalgroup.ExplicitClientPostOptionalIntegerPropertyOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostOptionalIntegerPropertyResponse], errResp azfake.ErrorResponder)

	// PostOptionalStringHeader is the fake for method ExplicitClient.PostOptionalStringHeader
	// HTTP status codes to indicate success: http.StatusOK
	PostOptionalStringHeader func(ctx context.Context, options *optionalgroup.ExplicitClientPostOptionalStringHeaderOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostOptionalStringHeaderResponse], errResp azfake.ErrorResponder)

	// PostOptionalStringParameter is the fake for method ExplicitClient.PostOptionalStringParameter
	// HTTP status codes to indicate success: http.StatusOK
	PostOptionalStringParameter func(ctx context.Context, options *optionalgroup.ExplicitClientPostOptionalStringParameterOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostOptionalStringParameterResponse], errResp azfake.ErrorResponder)

	// PostOptionalStringProperty is the fake for method ExplicitClient.PostOptionalStringProperty
	// HTTP status codes to indicate success: http.StatusOK
	PostOptionalStringProperty func(ctx context.Context, options *optionalgroup.ExplicitClientPostOptionalStringPropertyOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostOptionalStringPropertyResponse], errResp azfake.ErrorResponder)

	// PostRequiredArrayHeader is the fake for method ExplicitClient.PostRequiredArrayHeader
	// HTTP status codes to indicate success: http.StatusOK
	PostRequiredArrayHeader func(ctx context.Context, headerParameter []string, options *optionalgroup.ExplicitClientPostRequiredArrayHeaderOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostRequiredArrayHeaderResponse], errResp azfake.ErrorResponder)

	// PostRequiredArrayParameter is the fake for method ExplicitClient.PostRequiredArrayParameter
	// HTTP status codes to indicate success: http.StatusOK
	PostRequiredArrayParameter func(ctx context.Context, bodyParameter []*string, options *optionalgroup.ExplicitClientPostRequiredArrayParameterOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostRequiredArrayParameterResponse], errResp azfake.ErrorResponder)

	// PostRequiredArrayProperty is the fake for method ExplicitClient.PostRequiredArrayProperty
	// HTTP status codes to indicate success: http.StatusOK
	PostRequiredArrayProperty func(ctx context.Context, bodyParameter optionalgroup.ArrayWrapper, options *optionalgroup.ExplicitClientPostRequiredArrayPropertyOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostRequiredArrayPropertyResponse], errResp azfake.ErrorResponder)

	// PostRequiredClassParameter is the fake for method ExplicitClient.PostRequiredClassParameter
	// HTTP status codes to indicate success: http.StatusOK
	PostRequiredClassParameter func(ctx context.Context, bodyParameter optionalgroup.Product, options *optionalgroup.ExplicitClientPostRequiredClassParameterOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostRequiredClassParameterResponse], errResp azfake.ErrorResponder)

	// PostRequiredClassProperty is the fake for method ExplicitClient.PostRequiredClassProperty
	// HTTP status codes to indicate success: http.StatusOK
	PostRequiredClassProperty func(ctx context.Context, bodyParameter optionalgroup.ClassWrapper, options *optionalgroup.ExplicitClientPostRequiredClassPropertyOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostRequiredClassPropertyResponse], errResp azfake.ErrorResponder)

	// PostRequiredIntegerHeader is the fake for method ExplicitClient.PostRequiredIntegerHeader
	// HTTP status codes to indicate success: http.StatusOK
	PostRequiredIntegerHeader func(ctx context.Context, headerParameter int32, options *optionalgroup.ExplicitClientPostRequiredIntegerHeaderOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostRequiredIntegerHeaderResponse], errResp azfake.ErrorResponder)

	// PostRequiredIntegerParameter is the fake for method ExplicitClient.PostRequiredIntegerParameter
	// HTTP status codes to indicate success: http.StatusOK
	PostRequiredIntegerParameter func(ctx context.Context, bodyParameter int32, options *optionalgroup.ExplicitClientPostRequiredIntegerParameterOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostRequiredIntegerParameterResponse], errResp azfake.ErrorResponder)

	// PostRequiredIntegerProperty is the fake for method ExplicitClient.PostRequiredIntegerProperty
	// HTTP status codes to indicate success: http.StatusOK
	PostRequiredIntegerProperty func(ctx context.Context, bodyParameter optionalgroup.IntWrapper, options *optionalgroup.ExplicitClientPostRequiredIntegerPropertyOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostRequiredIntegerPropertyResponse], errResp azfake.ErrorResponder)

	// PostRequiredStringHeader is the fake for method ExplicitClient.PostRequiredStringHeader
	// HTTP status codes to indicate success: http.StatusOK
	PostRequiredStringHeader func(ctx context.Context, headerParameter string, options *optionalgroup.ExplicitClientPostRequiredStringHeaderOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostRequiredStringHeaderResponse], errResp azfake.ErrorResponder)

	// PostRequiredStringParameter is the fake for method ExplicitClient.PostRequiredStringParameter
	// HTTP status codes to indicate success: http.StatusOK
	PostRequiredStringParameter func(ctx context.Context, bodyParameter string, options *optionalgroup.ExplicitClientPostRequiredStringParameterOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostRequiredStringParameterResponse], errResp azfake.ErrorResponder)

	// PostRequiredStringProperty is the fake for method ExplicitClient.PostRequiredStringProperty
	// HTTP status codes to indicate success: http.StatusOK
	PostRequiredStringProperty func(ctx context.Context, bodyParameter optionalgroup.StringWrapper, options *optionalgroup.ExplicitClientPostRequiredStringPropertyOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostRequiredStringPropertyResponse], errResp azfake.ErrorResponder)

	// PutOptionalBinaryBody is the fake for method ExplicitClient.PutOptionalBinaryBody
	// HTTP status codes to indicate success: http.StatusOK
	PutOptionalBinaryBody func(ctx context.Context, options *optionalgroup.ExplicitClientPutOptionalBinaryBodyOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPutOptionalBinaryBodyResponse], errResp azfake.ErrorResponder)

	// PutRequiredBinaryBody is the fake for method ExplicitClient.PutRequiredBinaryBody
	// HTTP status codes to indicate success: http.StatusOK
	PutRequiredBinaryBody func(ctx context.Context, bodyParameter io.ReadSeekCloser, options *optionalgroup.ExplicitClientPutRequiredBinaryBodyOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPutRequiredBinaryBodyResponse], errResp azfake.ErrorResponder)
}

// NewExplicitServerTransport creates a new instance of ExplicitServerTransport with the provided implementation.
// The returned ExplicitServerTransport instance is connected to an instance of optionalgroup.ExplicitClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExplicitServerTransport(srv *ExplicitServer) *ExplicitServerTransport {
	return &ExplicitServerTransport{srv: srv}
}

// ExplicitServerTransport connects instances of optionalgroup.ExplicitClient to instances of ExplicitServer.
// Don't use this type directly, use NewExplicitServerTransport instead.
type ExplicitServerTransport struct {
	srv *ExplicitServer
}

// Do implements the policy.Transporter interface for ExplicitServerTransport.
func (e *ExplicitServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExplicitClient.PostOptionalArrayHeader":
		resp, err = e.dispatchPostOptionalArrayHeader(req)
	case "ExplicitClient.PostOptionalArrayParameter":
		resp, err = e.dispatchPostOptionalArrayParameter(req)
	case "ExplicitClient.PostOptionalArrayProperty":
		resp, err = e.dispatchPostOptionalArrayProperty(req)
	case "ExplicitClient.PostOptionalClassParameter":
		resp, err = e.dispatchPostOptionalClassParameter(req)
	case "ExplicitClient.PostOptionalClassProperty":
		resp, err = e.dispatchPostOptionalClassProperty(req)
	case "ExplicitClient.PostOptionalIntegerHeader":
		resp, err = e.dispatchPostOptionalIntegerHeader(req)
	case "ExplicitClient.PostOptionalIntegerParameter":
		resp, err = e.dispatchPostOptionalIntegerParameter(req)
	case "ExplicitClient.PostOptionalIntegerProperty":
		resp, err = e.dispatchPostOptionalIntegerProperty(req)
	case "ExplicitClient.PostOptionalStringHeader":
		resp, err = e.dispatchPostOptionalStringHeader(req)
	case "ExplicitClient.PostOptionalStringParameter":
		resp, err = e.dispatchPostOptionalStringParameter(req)
	case "ExplicitClient.PostOptionalStringProperty":
		resp, err = e.dispatchPostOptionalStringProperty(req)
	case "ExplicitClient.PostRequiredArrayHeader":
		resp, err = e.dispatchPostRequiredArrayHeader(req)
	case "ExplicitClient.PostRequiredArrayParameter":
		resp, err = e.dispatchPostRequiredArrayParameter(req)
	case "ExplicitClient.PostRequiredArrayProperty":
		resp, err = e.dispatchPostRequiredArrayProperty(req)
	case "ExplicitClient.PostRequiredClassParameter":
		resp, err = e.dispatchPostRequiredClassParameter(req)
	case "ExplicitClient.PostRequiredClassProperty":
		resp, err = e.dispatchPostRequiredClassProperty(req)
	case "ExplicitClient.PostRequiredIntegerHeader":
		resp, err = e.dispatchPostRequiredIntegerHeader(req)
	case "ExplicitClient.PostRequiredIntegerParameter":
		resp, err = e.dispatchPostRequiredIntegerParameter(req)
	case "ExplicitClient.PostRequiredIntegerProperty":
		resp, err = e.dispatchPostRequiredIntegerProperty(req)
	case "ExplicitClient.PostRequiredStringHeader":
		resp, err = e.dispatchPostRequiredStringHeader(req)
	case "ExplicitClient.PostRequiredStringParameter":
		resp, err = e.dispatchPostRequiredStringParameter(req)
	case "ExplicitClient.PostRequiredStringProperty":
		resp, err = e.dispatchPostRequiredStringProperty(req)
	case "ExplicitClient.PutOptionalBinaryBody":
		resp, err = e.dispatchPutOptionalBinaryBody(req)
	case "ExplicitClient.PutRequiredBinaryBody":
		resp, err = e.dispatchPutRequiredBinaryBody(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostOptionalArrayHeader(req *http.Request) (*http.Response, error) {
	if e.srv.PostOptionalArrayHeader == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostOptionalArrayHeader not implemented")}
	}
	headerParameterParam := splitHelper(getHeaderValue(req.Header, "headerParameter"), ",")
	var options *optionalgroup.ExplicitClientPostOptionalArrayHeaderOptions
	if len(headerParameterParam) > 0 {
		options = &optionalgroup.ExplicitClientPostOptionalArrayHeaderOptions{
			HeaderParameter: headerParameterParam,
		}
	}
	respr, errRespr := e.srv.PostOptionalArrayHeader(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostOptionalArrayParameter(req *http.Request) (*http.Response, error) {
	if e.srv.PostOptionalArrayParameter == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostOptionalArrayParameter not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[[]*string](req)
	if err != nil {
		return nil, err
	}
	var options *optionalgroup.ExplicitClientPostOptionalArrayParameterOptions
	if len(body) > 0 {
		options = &optionalgroup.ExplicitClientPostOptionalArrayParameterOptions{
			BodyParameter: body,
		}
	}
	respr, errRespr := e.srv.PostOptionalArrayParameter(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostOptionalArrayProperty(req *http.Request) (*http.Response, error) {
	if e.srv.PostOptionalArrayProperty == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostOptionalArrayProperty not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[optionalgroup.ArrayOptionalWrapper](req)
	if err != nil {
		return nil, err
	}
	var options *optionalgroup.ExplicitClientPostOptionalArrayPropertyOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &optionalgroup.ExplicitClientPostOptionalArrayPropertyOptions{
			BodyParameter: &body,
		}
	}
	respr, errRespr := e.srv.PostOptionalArrayProperty(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostOptionalClassParameter(req *http.Request) (*http.Response, error) {
	if e.srv.PostOptionalClassParameter == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostOptionalClassParameter not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[optionalgroup.Product](req)
	if err != nil {
		return nil, err
	}
	var options *optionalgroup.ExplicitClientPostOptionalClassParameterOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &optionalgroup.ExplicitClientPostOptionalClassParameterOptions{
			BodyParameter: &body,
		}
	}
	respr, errRespr := e.srv.PostOptionalClassParameter(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostOptionalClassProperty(req *http.Request) (*http.Response, error) {
	if e.srv.PostOptionalClassProperty == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostOptionalClassProperty not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[optionalgroup.ClassOptionalWrapper](req)
	if err != nil {
		return nil, err
	}
	var options *optionalgroup.ExplicitClientPostOptionalClassPropertyOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &optionalgroup.ExplicitClientPostOptionalClassPropertyOptions{
			BodyParameter: &body,
		}
	}
	respr, errRespr := e.srv.PostOptionalClassProperty(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostOptionalIntegerHeader(req *http.Request) (*http.Response, error) {
	if e.srv.PostOptionalIntegerHeader == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostOptionalIntegerHeader not implemented")}
	}
	headerParameterParam, err := parseOptional(getHeaderValue(req.Header, "headerParameter"), func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	var options *optionalgroup.ExplicitClientPostOptionalIntegerHeaderOptions
	if headerParameterParam != nil {
		options = &optionalgroup.ExplicitClientPostOptionalIntegerHeaderOptions{
			HeaderParameter: headerParameterParam,
		}
	}
	respr, errRespr := e.srv.PostOptionalIntegerHeader(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostOptionalIntegerParameter(req *http.Request) (*http.Response, error) {
	if e.srv.PostOptionalIntegerParameter == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostOptionalIntegerParameter not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[int32](req)
	if err != nil {
		return nil, err
	}
	var options *optionalgroup.ExplicitClientPostOptionalIntegerParameterOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &optionalgroup.ExplicitClientPostOptionalIntegerParameterOptions{
			BodyParameter: &body,
		}
	}
	respr, errRespr := e.srv.PostOptionalIntegerParameter(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostOptionalIntegerProperty(req *http.Request) (*http.Response, error) {
	if e.srv.PostOptionalIntegerProperty == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostOptionalIntegerProperty not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[optionalgroup.IntOptionalWrapper](req)
	if err != nil {
		return nil, err
	}
	var options *optionalgroup.ExplicitClientPostOptionalIntegerPropertyOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &optionalgroup.ExplicitClientPostOptionalIntegerPropertyOptions{
			BodyParameter: &body,
		}
	}
	respr, errRespr := e.srv.PostOptionalIntegerProperty(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostOptionalStringHeader(req *http.Request) (*http.Response, error) {
	if e.srv.PostOptionalStringHeader == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostOptionalStringHeader not implemented")}
	}
	bodyParameterParam := getOptional(getHeaderValue(req.Header, "bodyParameter"))
	var options *optionalgroup.ExplicitClientPostOptionalStringHeaderOptions
	if bodyParameterParam != nil {
		options = &optionalgroup.ExplicitClientPostOptionalStringHeaderOptions{
			BodyParameter: bodyParameterParam,
		}
	}
	respr, errRespr := e.srv.PostOptionalStringHeader(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostOptionalStringParameter(req *http.Request) (*http.Response, error) {
	if e.srv.PostOptionalStringParameter == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostOptionalStringParameter not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[string](req)
	if err != nil {
		return nil, err
	}
	var options *optionalgroup.ExplicitClientPostOptionalStringParameterOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &optionalgroup.ExplicitClientPostOptionalStringParameterOptions{
			BodyParameter: &body,
		}
	}
	respr, errRespr := e.srv.PostOptionalStringParameter(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostOptionalStringProperty(req *http.Request) (*http.Response, error) {
	if e.srv.PostOptionalStringProperty == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostOptionalStringProperty not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[optionalgroup.StringOptionalWrapper](req)
	if err != nil {
		return nil, err
	}
	var options *optionalgroup.ExplicitClientPostOptionalStringPropertyOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &optionalgroup.ExplicitClientPostOptionalStringPropertyOptions{
			BodyParameter: &body,
		}
	}
	respr, errRespr := e.srv.PostOptionalStringProperty(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostRequiredArrayHeader(req *http.Request) (*http.Response, error) {
	if e.srv.PostRequiredArrayHeader == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostRequiredArrayHeader not implemented")}
	}
	respr, errRespr := e.srv.PostRequiredArrayHeader(req.Context(), splitHelper(getHeaderValue(req.Header, "headerParameter"), ","), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostRequiredArrayParameter(req *http.Request) (*http.Response, error) {
	if e.srv.PostRequiredArrayParameter == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostRequiredArrayParameter not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[[]*string](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PostRequiredArrayParameter(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostRequiredArrayProperty(req *http.Request) (*http.Response, error) {
	if e.srv.PostRequiredArrayProperty == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostRequiredArrayProperty not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[optionalgroup.ArrayWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PostRequiredArrayProperty(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostRequiredClassParameter(req *http.Request) (*http.Response, error) {
	if e.srv.PostRequiredClassParameter == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostRequiredClassParameter not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[optionalgroup.Product](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PostRequiredClassParameter(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostRequiredClassProperty(req *http.Request) (*http.Response, error) {
	if e.srv.PostRequiredClassProperty == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostRequiredClassProperty not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[optionalgroup.ClassWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PostRequiredClassProperty(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostRequiredIntegerHeader(req *http.Request) (*http.Response, error) {
	if e.srv.PostRequiredIntegerHeader == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostRequiredIntegerHeader not implemented")}
	}
	headerParameterParam, err := parseWithCast(getHeaderValue(req.Header, "headerParameter"), func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PostRequiredIntegerHeader(req.Context(), int32(headerParameterParam), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostRequiredIntegerParameter(req *http.Request) (*http.Response, error) {
	if e.srv.PostRequiredIntegerParameter == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostRequiredIntegerParameter not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[int32](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PostRequiredIntegerParameter(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostRequiredIntegerProperty(req *http.Request) (*http.Response, error) {
	if e.srv.PostRequiredIntegerProperty == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostRequiredIntegerProperty not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[optionalgroup.IntWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PostRequiredIntegerProperty(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostRequiredStringHeader(req *http.Request) (*http.Response, error) {
	if e.srv.PostRequiredStringHeader == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostRequiredStringHeader not implemented")}
	}
	respr, errRespr := e.srv.PostRequiredStringHeader(req.Context(), getHeaderValue(req.Header, "headerParameter"), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostRequiredStringParameter(req *http.Request) (*http.Response, error) {
	if e.srv.PostRequiredStringParameter == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostRequiredStringParameter not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[string](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PostRequiredStringParameter(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPostRequiredStringProperty(req *http.Request) (*http.Response, error) {
	if e.srv.PostRequiredStringProperty == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostRequiredStringProperty not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[optionalgroup.StringWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PostRequiredStringProperty(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPutOptionalBinaryBody(req *http.Request) (*http.Response, error) {
	if e.srv.PutOptionalBinaryBody == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutOptionalBinaryBody not implemented")}
	}
	var options *optionalgroup.ExplicitClientPutOptionalBinaryBodyOptions
	if req.Body != nil {
		options = &optionalgroup.ExplicitClientPutOptionalBinaryBodyOptions{
			BodyParameter: req.Body.(io.ReadSeekCloser),
		}
	}
	respr, errRespr := e.srv.PutOptionalBinaryBody(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExplicitServerTransport) dispatchPutRequiredBinaryBody(req *http.Request) (*http.Response, error) {
	if e.srv.PutRequiredBinaryBody == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutRequiredBinaryBody not implemented")}
	}
	respr, errRespr := e.srv.PutRequiredBinaryBody(req.Context(), req.Body.(io.ReadSeekCloser), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
