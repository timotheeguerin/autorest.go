//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package arraygroup

import (
	"net/http"
	"time"
)

// ArrayClientGetArrayEmptyResponse contains the response from method ArrayClient.GetArrayEmpty.
type ArrayClientGetArrayEmptyResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// An empty array []
	StringArrayArray [][]*string
}

// ArrayClientGetArrayItemEmptyResponse contains the response from method ArrayClient.GetArrayItemEmpty.
type ArrayClientGetArrayItemEmptyResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// An array of array of strings [['1', '2', '3'], [], ['7', '8', '9']]
	StringArrayArray [][]*string
}

// ArrayClientGetArrayItemNullResponse contains the response from method ArrayClient.GetArrayItemNull.
type ArrayClientGetArrayItemNullResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// An array of array of strings [['1', '2', '3'], null, ['7', '8', '9']]
	StringArrayArray [][]*string
}

// ArrayClientGetArrayNullResponse contains the response from method ArrayClient.GetArrayNull.
type ArrayClientGetArrayNullResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// a null array
	StringArrayArray [][]*string
}

// ArrayClientGetArrayValidResponse contains the response from method ArrayClient.GetArrayValid.
type ArrayClientGetArrayValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// An array of array of strings [['1', '2', '3'], ['4', '5', '6'], ['7', '8', '9']]
	StringArrayArray [][]*string
}

// ArrayClientGetBase64URLResponse contains the response from method ArrayClient.GetBase64URL.
type ArrayClientGetBase64URLResponse struct {
	// Get array value ['a string that gets encoded with base64url', 'test string' 'Lorem ipsum'] with the items base64url encoded
	ByteArrayArray [][]byte

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetBooleanInvalidNullResponse contains the response from method ArrayClient.GetBooleanInvalidNull.
type ArrayClientGetBooleanInvalidNullResponse struct {
	// The array value [true, null, false]
	BoolArray []*bool

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetBooleanInvalidStringResponse contains the response from method ArrayClient.GetBooleanInvalidString.
type ArrayClientGetBooleanInvalidStringResponse struct {
	// The array value [true, 'boolean', false]
	BoolArray []*bool

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetBooleanTfftResponse contains the response from method ArrayClient.GetBooleanTfft.
type ArrayClientGetBooleanTfftResponse struct {
	// The array value [true, false, false, true]
	BoolArray []*bool

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetByteInvalidNullResponse contains the response from method ArrayClient.GetByteInvalidNull.
type ArrayClientGetByteInvalidNullResponse struct {
	// The byte array value [hex(AB, AC, AD), null] with the first item base64 encoded
	ByteArrayArray [][]byte

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetByteValidResponse contains the response from method ArrayClient.GetByteValid.
type ArrayClientGetByteValidResponse struct {
	// The array value [hex(FF FF FF FA), hex(01 02 03), hex (25, 29, 43)] with each elementencoded in base 64
	ByteArrayArray [][]byte

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetComplexEmptyResponse contains the response from method ArrayClient.GetComplexEmpty.
type ArrayClientGetComplexEmptyResponse struct {
	// Empty array of complex type []
	ProductArray []*Product

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetComplexItemEmptyResponse contains the response from method ArrayClient.GetComplexItemEmpty.
type ArrayClientGetComplexItemEmptyResponse struct {
	// Array of complex type with empty item [{'integer': 1 'string': '2'}, {}, {'integer': 5, 'string': '6'}]
	ProductArray []*Product

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetComplexItemNullResponse contains the response from method ArrayClient.GetComplexItemNull.
type ArrayClientGetComplexItemNullResponse struct {
	// Array of complex type with null item [{'integer': 1 'string': '2'}, null, {'integer': 5, 'string': '6'}]
	ProductArray []*Product

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetComplexNullResponse contains the response from method ArrayClient.GetComplexNull.
type ArrayClientGetComplexNullResponse struct {
	// array of complex type with null value
	ProductArray []*Product

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetComplexValidResponse contains the response from method ArrayClient.GetComplexValid.
type ArrayClientGetComplexValidResponse struct {
	// array of complex type with [{'integer': 1 'string': '2'}, {'integer': 3, 'string': '4'}, {'integer': 5, 'string': '6'}]
	ProductArray []*Product

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDateInvalidCharsResponse contains the response from method ArrayClient.GetDateInvalidChars.
type ArrayClientGetDateInvalidCharsResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The array value ['2011-03-22', 'date']
	TimeArray []*time.Time
}

// ArrayClientGetDateInvalidNullResponse contains the response from method ArrayClient.GetDateInvalidNull.
type ArrayClientGetDateInvalidNullResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The array value ['2012-01-01', null, '1776-07-04']
	TimeArray []*time.Time
}

// ArrayClientGetDateTimeInvalidCharsResponse contains the response from method ArrayClient.GetDateTimeInvalidChars.
type ArrayClientGetDateTimeInvalidCharsResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The array value ['2000-12-01t00:00:01z', 'date-time']
	TimeArray []*time.Time
}

// ArrayClientGetDateTimeInvalidNullResponse contains the response from method ArrayClient.GetDateTimeInvalidNull.
type ArrayClientGetDateTimeInvalidNullResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The array value ['2000-12-01t00:00:01z', null]
	TimeArray []*time.Time
}

// ArrayClientGetDateTimeRFC1123ValidResponse contains the response from method ArrayClient.GetDateTimeRFC1123Valid.
type ArrayClientGetDateTimeRFC1123ValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The array value ['Fri, 01 Dec 2000 00:00:01 GMT', 'Wed, 02 Jan 1980 00:11:35 GMT', 'Wed, 12 Oct 1492 10:15:01 GMT']
	TimeArray []*time.Time
}

// ArrayClientGetDateTimeValidResponse contains the response from method ArrayClient.GetDateTimeValid.
type ArrayClientGetDateTimeValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The array value ['2000-12-01t00:00:01z', '1980-01-02T00:11:35+01:00', '1492-10-12T10:15:01-08:00']
	TimeArray []*time.Time
}

// ArrayClientGetDateValidResponse contains the response from method ArrayClient.GetDateValid.
type ArrayClientGetDateValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The array value ['2000-12-01', '1980-01-02', '1492-10-12']
	TimeArray []*time.Time
}

// ArrayClientGetDictionaryEmptyResponse contains the response from method ArrayClient.GetDictionaryEmpty.
type ArrayClientGetDictionaryEmptyResponse struct {
	// An array of Dictionaries of type <string, string> with value []
	MapOfStringArray []map[string]*string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDictionaryItemEmptyResponse contains the response from method ArrayClient.GetDictionaryItemEmpty.
type ArrayClientGetDictionaryItemEmptyResponse struct {
	// An array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, {}, {'7': 'seven',
	// '8': 'eight', '9': 'nine'}]
	MapOfStringArray []map[string]*string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDictionaryItemNullResponse contains the response from method ArrayClient.GetDictionaryItemNull.
type ArrayClientGetDictionaryItemNullResponse struct {
	// An array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, null, {'7': 'seven',
	// '8': 'eight', '9': 'nine'}]
	MapOfStringArray []map[string]*string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDictionaryNullResponse contains the response from method ArrayClient.GetDictionaryNull.
type ArrayClientGetDictionaryNullResponse struct {
	// An array of Dictionaries with value null
	MapOfStringArray []map[string]*string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDictionaryValidResponse contains the response from method ArrayClient.GetDictionaryValid.
type ArrayClientGetDictionaryValidResponse struct {
	// An array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, {'4': 'four', '5':
	// 'five', '6': 'six'}, {'7': 'seven', '8': 'eight', '9': 'nine'}]
	MapOfStringArray []map[string]*string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDoubleInvalidNullResponse contains the response from method ArrayClient.GetDoubleInvalidNull.
type ArrayClientGetDoubleInvalidNullResponse struct {
	// The array value [0.0, null, -1.2e20]
	Float64Array []*float64

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDoubleInvalidStringResponse contains the response from method ArrayClient.GetDoubleInvalidString.
type ArrayClientGetDoubleInvalidStringResponse struct {
	// The array value [1.0, 'number', 0.0]
	Float64Array []*float64

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDoubleValidResponse contains the response from method ArrayClient.GetDoubleValid.
type ArrayClientGetDoubleValidResponse struct {
	// The array value [0, -0.01, 1.2e20]
	Float64Array []*float64

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDurationValidResponse contains the response from method ArrayClient.GetDurationValid.
type ArrayClientGetDurationValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The array value ['P123DT22H14M12.011S', 'P5DT1H0M0S']
	StringArray []*string
}

// ArrayClientGetEmptyResponse contains the response from method ArrayClient.GetEmpty.
type ArrayClientGetEmptyResponse struct {
	// The empty array value []
	Int32Array []*int32

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetEnumValidResponse contains the response from method ArrayClient.GetEnumValid.
type ArrayClientGetEnumValidResponse struct {
	// The array value ['foo1', 'foo2', 'foo3']
	FooEnumArray []*FooEnum

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetFloatInvalidNullResponse contains the response from method ArrayClient.GetFloatInvalidNull.
type ArrayClientGetFloatInvalidNullResponse struct {
	// The array value [0.0, null, -1.2e20]
	Float32Array []*float32

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetFloatInvalidStringResponse contains the response from method ArrayClient.GetFloatInvalidString.
type ArrayClientGetFloatInvalidStringResponse struct {
	// The array value [1.0, 'number', 0.0]
	Float32Array []*float32

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetFloatValidResponse contains the response from method ArrayClient.GetFloatValid.
type ArrayClientGetFloatValidResponse struct {
	// The array value [0, -0.01, 1.2e20]
	Float32Array []*float32

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetIntInvalidNullResponse contains the response from method ArrayClient.GetIntInvalidNull.
type ArrayClientGetIntInvalidNullResponse struct {
	// The array value [1, null, 0]
	Int32Array []*int32

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetIntInvalidStringResponse contains the response from method ArrayClient.GetIntInvalidString.
type ArrayClientGetIntInvalidStringResponse struct {
	// The array value [1, 'integer', 0]
	Int32Array []*int32

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetIntegerValidResponse contains the response from method ArrayClient.GetIntegerValid.
type ArrayClientGetIntegerValidResponse struct {
	// The array value [1, -1, 3, 300]
	Int32Array []*int32

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetInvalidResponse contains the response from method ArrayClient.GetInvalid.
type ArrayClientGetInvalidResponse struct {
	// The invalid Array [1, 2, 3
	Int32Array []*int32

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetLongInvalidNullResponse contains the response from method ArrayClient.GetLongInvalidNull.
type ArrayClientGetLongInvalidNullResponse struct {
	// The array value [1, null, 0]
	Int64Array []*int64

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetLongInvalidStringResponse contains the response from method ArrayClient.GetLongInvalidString.
type ArrayClientGetLongInvalidStringResponse struct {
	// The array value [1, 'integer', 0]
	Int64Array []*int64

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetLongValidResponse contains the response from method ArrayClient.GetLongValid.
type ArrayClientGetLongValidResponse struct {
	// The array value [1, -1, 3, 300]
	Int64Array []*int64

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetNullResponse contains the response from method ArrayClient.GetNull.
type ArrayClientGetNullResponse struct {
	// The null Array value
	Int32Array []*int32

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetStringEnumValidResponse contains the response from method ArrayClient.GetStringEnumValid.
type ArrayClientGetStringEnumValidResponse struct {
	// The array value ['foo1', 'foo2', 'foo3']
	Enum0Array []*Enum0

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetStringValidResponse contains the response from method ArrayClient.GetStringValid.
type ArrayClientGetStringValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The array value ['foo1', 'foo2', 'foo3']
	StringArray []*string
}

// ArrayClientGetStringWithInvalidResponse contains the response from method ArrayClient.GetStringWithInvalid.
type ArrayClientGetStringWithInvalidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The array value ['foo', 123, 'foo2']
	StringArray []*string
}

// ArrayClientGetStringWithNullResponse contains the response from method ArrayClient.GetStringWithNull.
type ArrayClientGetStringWithNullResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The array value ['foo', null, 'foo2']
	StringArray []*string
}

// ArrayClientGetUUIDInvalidCharsResponse contains the response from method ArrayClient.GetUUIDInvalidChars.
type ArrayClientGetUUIDInvalidCharsResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The array value ['6dcc7237-45fe-45c4-8a6b-3a8a3f625652', 'foo']
	StringArray []*string
}

// ArrayClientGetUUIDValidResponse contains the response from method ArrayClient.GetUUIDValid.
type ArrayClientGetUUIDValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The array value ['6dcc7237-45fe-45c4-8a6b-3a8a3f625652', 'd1399005-30f7-40d6-8da6-dd7c89ad34db', 'f42f6aa1-a5bc-4ddf-907e-5f915de43205']
	StringArray []*string
}

// ArrayClientPutArrayValidResponse contains the response from method ArrayClient.PutArrayValid.
type ArrayClientPutArrayValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutBooleanTfftResponse contains the response from method ArrayClient.PutBooleanTfft.
type ArrayClientPutBooleanTfftResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutByteValidResponse contains the response from method ArrayClient.PutByteValid.
type ArrayClientPutByteValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutComplexValidResponse contains the response from method ArrayClient.PutComplexValid.
type ArrayClientPutComplexValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutDateTimeRFC1123ValidResponse contains the response from method ArrayClient.PutDateTimeRFC1123Valid.
type ArrayClientPutDateTimeRFC1123ValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutDateTimeValidResponse contains the response from method ArrayClient.PutDateTimeValid.
type ArrayClientPutDateTimeValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutDateValidResponse contains the response from method ArrayClient.PutDateValid.
type ArrayClientPutDateValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutDictionaryValidResponse contains the response from method ArrayClient.PutDictionaryValid.
type ArrayClientPutDictionaryValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutDoubleValidResponse contains the response from method ArrayClient.PutDoubleValid.
type ArrayClientPutDoubleValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutDurationValidResponse contains the response from method ArrayClient.PutDurationValid.
type ArrayClientPutDurationValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutEmptyResponse contains the response from method ArrayClient.PutEmpty.
type ArrayClientPutEmptyResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutEnumValidResponse contains the response from method ArrayClient.PutEnumValid.
type ArrayClientPutEnumValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutFloatValidResponse contains the response from method ArrayClient.PutFloatValid.
type ArrayClientPutFloatValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutIntegerValidResponse contains the response from method ArrayClient.PutIntegerValid.
type ArrayClientPutIntegerValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutLongValidResponse contains the response from method ArrayClient.PutLongValid.
type ArrayClientPutLongValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutStringEnumValidResponse contains the response from method ArrayClient.PutStringEnumValid.
type ArrayClientPutStringEnumValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutStringValidResponse contains the response from method ArrayClient.PutStringValid.
type ArrayClientPutStringValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutUUIDValidResponse contains the response from method ArrayClient.PutUUIDValid.
type ArrayClientPutUUIDValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
