//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azalias

import "net/http"

// clientCreateResponse contains the response from method client.Create.
type clientCreateResponse struct {
	AliasesCreateResponse
	// AccessControlExposeHeaders contains the information returned from the Access-Control-Expose-Headers header response.
	AccessControlExposeHeaders *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// clientGetScriptResponse contains the response from method client.GetScript.
type clientGetScriptResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	Value       *string
}

// clientListResponse contains the response from method client.List.
type clientListResponse struct {
	ListResponse
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// clientPolicyAssignmentResponse contains the response from method client.PolicyAssignment.
type clientPolicyAssignmentResponse struct {
	PolicyAssignmentProperties
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
