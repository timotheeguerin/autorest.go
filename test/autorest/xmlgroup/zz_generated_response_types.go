//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package xmlgroup

import "net/http"

// XMLClientGetACLsResponse contains the response from method XMLClient.GetACLs.
type XMLClientGetACLsResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// a collection of signed identifiers
	SignedIdentifiers []*SignedIdentifier `xml:"SignedIdentifier"`
}

// XMLClientGetBytesResponse contains the response from method XMLClient.GetBytes.
type XMLClientGetBytesResponse struct {
	ModelWithByteProperty
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientGetComplexTypeRefNoMetaResponse contains the response from method XMLClient.GetComplexTypeRefNoMeta.
type XMLClientGetComplexTypeRefNoMetaResponse struct {
	RootWithRefAndNoMeta
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientGetComplexTypeRefWithMetaResponse contains the response from method XMLClient.GetComplexTypeRefWithMeta.
type XMLClientGetComplexTypeRefWithMetaResponse struct {
	RootWithRefAndMeta
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientGetEmptyChildElementResponse contains the response from method XMLClient.GetEmptyChildElement.
type XMLClientGetEmptyChildElementResponse struct {
	Banana
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientGetEmptyListResponse contains the response from method XMLClient.GetEmptyList.
type XMLClientGetEmptyListResponse struct {
	Slideshow
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientGetEmptyRootListResponse contains the response from method XMLClient.GetEmptyRootList.
type XMLClientGetEmptyRootListResponse struct {
	// Array of Banana
	Bananas []*Banana `xml:"banana"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientGetEmptyWrappedListsResponse contains the response from method XMLClient.GetEmptyWrappedLists.
type XMLClientGetEmptyWrappedListsResponse struct {
	AppleBarrel
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientGetHeadersResponse contains the response from method XMLClient.GetHeaders.
type XMLClientGetHeadersResponse struct {
	// CustomHeader contains the information returned from the Custom-Header header response.
	CustomHeader *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientGetRootListResponse contains the response from method XMLClient.GetRootList.
type XMLClientGetRootListResponse struct {
	// Array of Banana
	Bananas []*Banana `xml:"banana"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientGetRootListSingleItemResponse contains the response from method XMLClient.GetRootListSingleItem.
type XMLClientGetRootListSingleItemResponse struct {
	// Array of Banana
	Bananas []*Banana `xml:"banana"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientGetServicePropertiesResponse contains the response from method XMLClient.GetServiceProperties.
type XMLClientGetServicePropertiesResponse struct {
	StorageServiceProperties
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientGetSimpleResponse contains the response from method XMLClient.GetSimple.
type XMLClientGetSimpleResponse struct {
	Slideshow
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientGetURIResponse contains the response from method XMLClient.GetURI.
type XMLClientGetURIResponse struct {
	ModelWithURLProperty
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientGetWrappedListsResponse contains the response from method XMLClient.GetWrappedLists.
type XMLClientGetWrappedListsResponse struct {
	AppleBarrel
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientGetXMsTextResponse contains the response from method XMLClient.GetXMsText.
type XMLClientGetXMsTextResponse struct {
	ObjectWithXMsTextProperty
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientJSONInputResponse contains the response from method XMLClient.JSONInput.
type XMLClientJSONInputResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientJSONOutputResponse contains the response from method XMLClient.JSONOutput.
type XMLClientJSONOutputResponse struct {
	JSONOutput
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientListBlobsResponse contains the response from method XMLClient.ListBlobs.
type XMLClientListBlobsResponse struct {
	ListBlobsResponse
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientListContainersResponse contains the response from method XMLClient.ListContainers.
type XMLClientListContainersResponse struct {
	ListContainersResponse
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientPutACLsResponse contains the response from method XMLClient.PutACLs.
type XMLClientPutACLsResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientPutBinaryResponse contains the response from method XMLClient.PutBinary.
type XMLClientPutBinaryResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientPutComplexTypeRefNoMetaResponse contains the response from method XMLClient.PutComplexTypeRefNoMeta.
type XMLClientPutComplexTypeRefNoMetaResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientPutComplexTypeRefWithMetaResponse contains the response from method XMLClient.PutComplexTypeRefWithMeta.
type XMLClientPutComplexTypeRefWithMetaResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientPutEmptyChildElementResponse contains the response from method XMLClient.PutEmptyChildElement.
type XMLClientPutEmptyChildElementResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientPutEmptyListResponse contains the response from method XMLClient.PutEmptyList.
type XMLClientPutEmptyListResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientPutEmptyRootListResponse contains the response from method XMLClient.PutEmptyRootList.
type XMLClientPutEmptyRootListResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientPutEmptyWrappedListsResponse contains the response from method XMLClient.PutEmptyWrappedLists.
type XMLClientPutEmptyWrappedListsResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientPutRootListResponse contains the response from method XMLClient.PutRootList.
type XMLClientPutRootListResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientPutRootListSingleItemResponse contains the response from method XMLClient.PutRootListSingleItem.
type XMLClientPutRootListSingleItemResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientPutServicePropertiesResponse contains the response from method XMLClient.PutServiceProperties.
type XMLClientPutServicePropertiesResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientPutSimpleResponse contains the response from method XMLClient.PutSimple.
type XMLClientPutSimpleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientPutURIResponse contains the response from method XMLClient.PutURI.
type XMLClientPutURIResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLClientPutWrappedListsResponse contains the response from method XMLClient.PutWrappedLists.
type XMLClientPutWrappedListsResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
