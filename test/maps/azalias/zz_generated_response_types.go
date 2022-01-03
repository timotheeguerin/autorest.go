//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azalias

import "net/http"

// AliasCreateResponse contains the response from method Alias.Create.
type AliasCreateResponse struct {
	AliasCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AliasCreateResult contains the result from method Alias.Create.
type AliasCreateResult struct {
	AliasesCreateResponse
	// AccessControlExposeHeaders contains the information returned from the Access-Control-Expose-Headers header response.
	AccessControlExposeHeaders *string
}

// AliasGetScriptResponse contains the response from method Alias.GetScript.
type AliasGetScriptResponse struct {
	AliasGetScriptResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AliasGetScriptResult contains the result from method Alias.GetScript.
type AliasGetScriptResult struct {
	Value *string
}

// AliasListResponse contains the response from method Alias.List.
type AliasListResponse struct {
	AliasListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AliasListResult contains the result from method Alias.List.
type AliasListResult struct {
	ListResponse
}
