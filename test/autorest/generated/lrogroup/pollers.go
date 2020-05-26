// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package lrogroup

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"time"
)

// HTTPPoller provides polling facilities until the operation completes
type HTTPPoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse() *http.Response
	ResumeToken() (string, error)
}

type httpPoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
	// polling tracker
	pt pollingTracker
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *httpPoller) Done() bool {
	return p.pt.hasTerminated()
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *httpPoller) Poll(ctx context.Context) (*http.Response, error) {
	if lroPollDone(ctx, p.pipeline, p.pt) {
		return p.pt.latestResponse().Response, p.pt.pollingError()
	}
	return nil, p.pt.pollingError()
}

func (p *httpPoller) FinalResponse() *http.Response {
	return p.pt.latestResponse().Response
}

// ResumeToken generates the string token that can be used with the ResumeHTTPPoller method
// on the client to create a new poller from the data held in the current poller type
func (p *httpPoller) ResumeToken() (string, error) {
	if p.pt.hasTerminated() {
		return "", errors.New("cannot create a ResumeToken from a poller in a terminal state")
	}
	js, err := json.Marshal(p.pt)
	if err != nil {
		return "", err
	}
	return string(js), nil
}

func (p *httpPoller) pollUntilDone(ctx context.Context, frequency time.Duration) (*http.Response, error) {
	for {
		resp, err := p.Poll(ctx)
		if err != nil {
			return nil, err
		}
		if p.Done() {
			break
		}
		if delay := azcore.RetryAfter(resp); delay > 0 {
			time.Sleep(delay)
		} else {
			time.Sleep(frequency)
		}
	}
	return p.FinalResponse(), nil
}

// ProductPoller provides polling facilities until the operation completes
type ProductPoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*ProductResponse, error)
	ResumeToken() (string, error)
}

type productPoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
	// polling tracker
	pt pollingTracker
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *productPoller) Done() bool {
	return p.pt.hasTerminated()
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *productPoller) Poll(ctx context.Context) (*http.Response, error) {
	if lroPollDone(ctx, p.pipeline, p.pt) {
		return p.pt.latestResponse().Response, p.pt.pollingError()
	}
	return nil, p.pt.pollingError()
}

func (p *productPoller) FinalResponse(ctx context.Context) (*ProductResponse, error) {
	// checking if there was a FinalStateVia configuration to re-route the final GET
	// request to the value specified in the FinalStateVia property on the poller
	err := p.pt.setFinalState()
	if err != nil {
		return nil, err
	}
	if p.pt.finalGetURL() == "" {
		// we can end up in this situation if the async operation returns a 200
		// with no polling URLs.  in that case return the response which should
		// contain the JSON payload (only do this for successful terminal cases).
		if lr := p.pt.latestResponse(); lr != nil && p.pt.hasSucceeded() {
			result, err := p.handleResponse(lr)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, errors.New("missing URL for retrieving result")
	}
	u, err := url.Parse(p.pt.finalGetURL())
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	if err != nil {
		return nil, err
	}
	resp, err := p.pipeline.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	return p.handleResponse(resp)
}

// ResumeToken generates the string token that can be used with the ResumeProductPoller method
// on the client to create a new poller from the data held in the current poller type
func (p *productPoller) ResumeToken() (string, error) {
	if p.pt.hasTerminated() {
		return "", errors.New("cannot create a ResumeToken from a poller in a terminal state")
	}
	js, err := json.Marshal(p.pt)
	if err != nil {
		return "", err
	}
	return string(js), nil
}

func (p *productPoller) pollUntilDone(ctx context.Context, frequency time.Duration) (*ProductResponse, error) {
	for {
		resp, err := p.Poll(ctx)
		if err != nil {
			return nil, err
		}
		if p.Done() {
			break
		}
		if delay := azcore.RetryAfter(resp); delay > 0 {
			time.Sleep(delay)
		} else {
			time.Sleep(frequency)
		}
	}
	return p.FinalResponse(ctx)
}

func (p *productPoller) handleResponse(resp *azcore.Response) (*ProductResponse, error) {
	result := ProductResponse{RawResponse: resp.Response}
	if resp.HasStatusCode(http.StatusNoContent) {
		return &result, nil
	}
	if !resp.HasStatusCode(pollingCodes[:]...) {
		return nil, p.pt.handleError(resp)
	}
	return &result, resp.UnmarshalAsJSON(&result.Product)
}

// SkuPoller provides polling facilities until the operation completes
type SkuPoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*SkuResponse, error)
	ResumeToken() (string, error)
}

type skuPoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
	// polling tracker
	pt pollingTracker
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *skuPoller) Done() bool {
	return p.pt.hasTerminated()
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *skuPoller) Poll(ctx context.Context) (*http.Response, error) {
	if lroPollDone(ctx, p.pipeline, p.pt) {
		return p.pt.latestResponse().Response, p.pt.pollingError()
	}
	return nil, p.pt.pollingError()
}

func (p *skuPoller) FinalResponse(ctx context.Context) (*SkuResponse, error) {
	// checking if there was a FinalStateVia configuration to re-route the final GET
	// request to the value specified in the FinalStateVia property on the poller
	err := p.pt.setFinalState()
	if err != nil {
		return nil, err
	}
	if p.pt.finalGetURL() == "" {
		// we can end up in this situation if the async operation returns a 200
		// with no polling URLs.  in that case return the response which should
		// contain the JSON payload (only do this for successful terminal cases).
		if lr := p.pt.latestResponse(); lr != nil && p.pt.hasSucceeded() {
			result, err := p.handleResponse(lr)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, errors.New("missing URL for retrieving result")
	}
	u, err := url.Parse(p.pt.finalGetURL())
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	if err != nil {
		return nil, err
	}
	resp, err := p.pipeline.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	return p.handleResponse(resp)
}

// ResumeToken generates the string token that can be used with the ResumeSkuPoller method
// on the client to create a new poller from the data held in the current poller type
func (p *skuPoller) ResumeToken() (string, error) {
	if p.pt.hasTerminated() {
		return "", errors.New("cannot create a ResumeToken from a poller in a terminal state")
	}
	js, err := json.Marshal(p.pt)
	if err != nil {
		return "", err
	}
	return string(js), nil
}

func (p *skuPoller) pollUntilDone(ctx context.Context, frequency time.Duration) (*SkuResponse, error) {
	for {
		resp, err := p.Poll(ctx)
		if err != nil {
			return nil, err
		}
		if p.Done() {
			break
		}
		if delay := azcore.RetryAfter(resp); delay > 0 {
			time.Sleep(delay)
		} else {
			time.Sleep(frequency)
		}
	}
	return p.FinalResponse(ctx)
}

func (p *skuPoller) handleResponse(resp *azcore.Response) (*SkuResponse, error) {
	result := SkuResponse{RawResponse: resp.Response}
	if resp.HasStatusCode(http.StatusNoContent) {
		return &result, nil
	}
	if !resp.HasStatusCode(pollingCodes[:]...) {
		return nil, p.pt.handleError(resp)
	}
	return &result, resp.UnmarshalAsJSON(&result.Sku)
}

// SubProductPoller provides polling facilities until the operation completes
type SubProductPoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*SubProductResponse, error)
	ResumeToken() (string, error)
}

type subProductPoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
	// polling tracker
	pt pollingTracker
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *subProductPoller) Done() bool {
	return p.pt.hasTerminated()
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *subProductPoller) Poll(ctx context.Context) (*http.Response, error) {
	if lroPollDone(ctx, p.pipeline, p.pt) {
		return p.pt.latestResponse().Response, p.pt.pollingError()
	}
	return nil, p.pt.pollingError()
}

func (p *subProductPoller) FinalResponse(ctx context.Context) (*SubProductResponse, error) {
	// checking if there was a FinalStateVia configuration to re-route the final GET
	// request to the value specified in the FinalStateVia property on the poller
	err := p.pt.setFinalState()
	if err != nil {
		return nil, err
	}
	if p.pt.finalGetURL() == "" {
		// we can end up in this situation if the async operation returns a 200
		// with no polling URLs.  in that case return the response which should
		// contain the JSON payload (only do this for successful terminal cases).
		if lr := p.pt.latestResponse(); lr != nil && p.pt.hasSucceeded() {
			result, err := p.handleResponse(lr)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, errors.New("missing URL for retrieving result")
	}
	u, err := url.Parse(p.pt.finalGetURL())
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	if err != nil {
		return nil, err
	}
	resp, err := p.pipeline.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	return p.handleResponse(resp)
}

// ResumeToken generates the string token that can be used with the ResumeSubProductPoller method
// on the client to create a new poller from the data held in the current poller type
func (p *subProductPoller) ResumeToken() (string, error) {
	if p.pt.hasTerminated() {
		return "", errors.New("cannot create a ResumeToken from a poller in a terminal state")
	}
	js, err := json.Marshal(p.pt)
	if err != nil {
		return "", err
	}
	return string(js), nil
}

func (p *subProductPoller) pollUntilDone(ctx context.Context, frequency time.Duration) (*SubProductResponse, error) {
	for {
		resp, err := p.Poll(ctx)
		if err != nil {
			return nil, err
		}
		if p.Done() {
			break
		}
		if delay := azcore.RetryAfter(resp); delay > 0 {
			time.Sleep(delay)
		} else {
			time.Sleep(frequency)
		}
	}
	return p.FinalResponse(ctx)
}

func (p *subProductPoller) handleResponse(resp *azcore.Response) (*SubProductResponse, error) {
	result := SubProductResponse{RawResponse: resp.Response}
	if resp.HasStatusCode(http.StatusNoContent) {
		return &result, nil
	}
	if !resp.HasStatusCode(pollingCodes[:]...) {
		return nil, p.pt.handleError(resp)
	}
	return &result, resp.UnmarshalAsJSON(&result.SubProduct)
}
