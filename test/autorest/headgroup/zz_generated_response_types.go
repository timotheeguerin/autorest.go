//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package headgroup

import "net/http"

// HTTPSuccessClientHead200Response contains the response from method HTTPSuccessClient.Head200.
type HTTPSuccessClientHead200Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Success indicates if the operation succeeded or failed.
	Success bool
}

// HTTPSuccessClientHead204Response contains the response from method HTTPSuccessClient.Head204.
type HTTPSuccessClientHead204Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Success indicates if the operation succeeded or failed.
	Success bool
}

// HTTPSuccessClientHead404Response contains the response from method HTTPSuccessClient.Head404.
type HTTPSuccessClientHead404Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Success indicates if the operation succeeded or failed.
	Success bool
}
