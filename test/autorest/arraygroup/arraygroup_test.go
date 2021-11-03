// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package arraygroup

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
)

func newArrayClient() *ArrayClient {
	return NewArrayClient(nil)
}

// GetArrayEmpty - Get an empty array []
func TestGetArrayEmpty(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetArrayEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StringArrayArray == nil {
		t.Fatal("unexpected nil array")
	}
	if r := cmp.Diff(len(resp.StringArrayArray), 0); r != "" {
		t.Fatal(r)
	}
}

// GetArrayItemEmpty - Get an array of array of strings [['1', '2', '3'], [], ['7', '8', '9']]
func TestGetArrayItemEmpty(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetArrayItemEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.StringArrayArray, [][]*string{
		to.StringPtrArray("1", "2", "3"),
		{},
		to.StringPtrArray("7", "8", "9"),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetArrayItemNull - Get an array of array of strings [['1', '2', '3'], null, ['7', '8', '9']]
func TestGetArrayItemNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetArrayItemNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.StringArrayArray, [][]*string{
		to.StringPtrArray("1", "2", "3"),
		nil,
		to.StringPtrArray("7", "8", "9"),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetArrayNull - Get a null array
func TestGetArrayNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetArrayNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StringArrayArray != nil {
		t.Fatal("expected nil array")
	}
}

// GetArrayValid - Get an array of array of strings [['1', '2', '3'], ['4', '5', '6'], ['7', '8', '9']]
func TestGetArrayValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetArrayValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.StringArrayArray, [][]*string{
		to.StringPtrArray("1", "2", "3"),
		to.StringPtrArray("4", "5", "6"),
		to.StringPtrArray("7", "8", "9"),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetBase64URL - Get array value ['a string that gets encoded with base64url', 'test string' 'Lorem ipsum'] with the items base64url encoded
func TestGetBase64URL(t *testing.T) {
	t.Skip("decoding fails")
	client := newArrayClient()
	resp, err := client.GetBase64URL(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.ByteArrayArray, [][]byte{
		{0},
		{0},
		{0},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetBooleanInvalidNull - Get boolean array value [true, null, false]
func TestGetBooleanInvalidNull(t *testing.T) {
	t.Skip("unmarshalling succeeds")
	client := newArrayClient()
	resp, err := client.GetBooleanInvalidNull(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetBooleanInvalidString - Get boolean array value [true, 'boolean', false]
func TestGetBooleanInvalidString(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetBooleanInvalidString(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetBooleanTfft - Get boolean array value [true, false, false, true]
func TestGetBooleanTfft(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetBooleanTfft(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.BoolArray, to.BoolPtrArray(true, false, false, true)); r != "" {
		t.Fatal(r)
	}
}

// GetByteInvalidNull - Get byte array value [hex(AB, AC, AD), null] with the first item base64 encoded
func TestGetByteInvalidNull(t *testing.T) {
	t.Skip("needs investigation")
	client := newArrayClient()
	resp, err := client.GetByteInvalidNull(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetByteValid - Get byte array value [hex(FF FF FF FA), hex(01 02 03), hex (25, 29, 43)] with each item encoded in base64
func TestGetByteValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetByteValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.ByteArrayArray, [][]byte{
		{255, 255, 255, 250},
		{1, 2, 3},
		{37, 41, 67},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetComplexEmpty - Get empty array of complex type []
func TestGetComplexEmpty(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetComplexEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.ProductArray == nil {
		t.Fatal("unexpected nil array")
	}
	if r := cmp.Diff(len(resp.ProductArray), 0); r != "" {
		t.Fatal(r)
	}
}

// GetComplexItemEmpty - Get array of complex type with empty item [{'integer': 1 'string': '2'}, {}, {'integer': 5, 'string': '6'}]
func TestGetComplexItemEmpty(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetComplexItemEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.ProductArray, []*Product{
		{Integer: to.Int32Ptr(1), String: to.StringPtr("2")},
		{},
		{Integer: to.Int32Ptr(5), String: to.StringPtr("6")},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetComplexItemNull - Get array of complex type with null item [{'integer': 1 'string': '2'}, null, {'integer': 5, 'string': '6'}]
func TestGetComplexItemNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetComplexItemNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.ProductArray, []*Product{
		{Integer: to.Int32Ptr(1), String: to.StringPtr("2")},
		nil,
		{Integer: to.Int32Ptr(5), String: to.StringPtr("6")},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetComplexNull - Get array of complex type null value
func TestGetComplexNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetComplexNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.ProductArray != nil {
		t.Fatal("expected nil array")
	}
}

// GetComplexValid - Get array of complex type with [{'integer': 1 'string': '2'}, {'integer': 3, 'string': '4'}, {'integer': 5, 'string': '6'}]
func TestGetComplexValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetComplexValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.ProductArray, []*Product{
		{Integer: to.Int32Ptr(1), String: to.StringPtr("2")},
		{Integer: to.Int32Ptr(3), String: to.StringPtr("4")},
		{Integer: to.Int32Ptr(5), String: to.StringPtr("6")},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateInvalidChars - Get date array value ['2011-03-22', 'date']
func TestGetDateInvalidChars(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDateInvalidChars(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetDateInvalidNull - Get date array value ['2012-01-01', null, '1776-07-04']
func TestGetDateInvalidNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDateInvalidNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	v1 := time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)
	v3 := time.Date(1776, 7, 4, 0, 0, 0, 0, time.UTC)
	if r := cmp.Diff(resp.TimeArray, []*time.Time{
		&v1,
		nil,
		&v3,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateTimeInvalidChars - Get date array value ['2000-12-01t00:00:01z', 'date-time']
func TestGetDateTimeInvalidChars(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDateTimeInvalidChars(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetDateTimeInvalidNull - Get date array value ['2000-12-01t00:00:01z', null]
func TestGetDateTimeInvalidNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDateTimeInvalidNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	v1, _ := time.Parse(time.RFC1123, "Fri, 01 Dec 2000 00:00:01 GMT")
	if r := cmp.Diff(resp.TimeArray, []*time.Time{
		&v1,
		nil,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateTimeRFC1123Valid - Get date-time array value ['Fri, 01 Dec 2000 00:00:01 GMT', 'Wed, 02 Jan 1980 00:11:35 GMT', 'Wed, 12 Oct 1492 10:15:01 GMT']
func TestGetDateTimeRFC1123Valid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDateTimeRFC1123Valid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	v1, _ := time.Parse(time.RFC1123, "Fri, 01 Dec 2000 00:00:01 GMT")
	v2, _ := time.Parse(time.RFC1123, "Wed, 02 Jan 1980 00:11:35 GMT")
	v3, _ := time.Parse(time.RFC1123, "Wed, 12 Oct 1492 10:15:01 GMT")
	if r := cmp.Diff(resp.TimeArray, []*time.Time{
		&v1,
		&v2,
		&v3,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateTimeValid - Get date-time array value ['2000-12-01t00:00:01z', '1980-01-02T00:11:35+01:00', '1492-10-12T10:15:01-08:00']
func TestGetDateTimeValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDateTimeValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	v1, _ := time.Parse(time.RFC3339, "2000-12-01T00:00:01Z")
	v2, _ := time.Parse(time.RFC3339, "1980-01-02T01:11:35+01:00")
	v3, _ := time.Parse(time.RFC3339, "1492-10-12T02:15:01-08:00")
	if r := cmp.Diff(resp.TimeArray, []*time.Time{
		&v1,
		&v2,
		&v3,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateValid - Get integer array value ['2000-12-01', '1980-01-02', '1492-10-12']
func TestGetDateValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDateValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.TimeArray, to.TimePtrArray(
		time.Date(2000, time.December, 01, 0, 0, 0, 0, time.UTC),
		time.Date(1980, time.January, 02, 0, 0, 0, 0, time.UTC),
		time.Date(1492, time.October, 12, 0, 0, 0, 0, time.UTC),
	)); r != "" {
		t.Fatal(r)
	}
}

// GetDictionaryEmpty - Get an array of Dictionaries of type <string, string> with value []
func TestGetDictionaryEmpty(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDictionaryEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.MapOfStringArray, []map[string]*string{}); r != "" {
		t.Fatal(r)
	}
}

// GetDictionaryItemEmpty - Get an array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, {}, {'7': 'seven', '8': 'eight', '9': 'nine'}]
func TestGetDictionaryItemEmpty(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDictionaryItemEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.MapOfStringArray, []map[string]*string{
		{
			"1": to.StringPtr("one"),
			"2": to.StringPtr("two"),
			"3": to.StringPtr("three"),
		},
		{},
		{
			"7": to.StringPtr("seven"),
			"8": to.StringPtr("eight"),
			"9": to.StringPtr("nine"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDictionaryItemNull - Get an array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, null, {'7': 'seven', '8': 'eight', '9': 'nine'}]
func TestGetDictionaryItemNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDictionaryItemNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.MapOfStringArray, []map[string]*string{
		{
			"1": to.StringPtr("one"),
			"2": to.StringPtr("two"),
			"3": to.StringPtr("three"),
		},
		nil,
		{
			"7": to.StringPtr("seven"),
			"8": to.StringPtr("eight"),
			"9": to.StringPtr("nine"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDictionaryNull - Get an array of Dictionaries with value null
func TestGetDictionaryNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDictionaryNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.MapOfStringArray != nil {
		t.Fatal("expected nil dictionary")
	}
}

// GetDictionaryValid - Get an array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, {'4': 'four', '5': 'five', '6': 'six'}, {'7': 'seven', '8': 'eight', '9': 'nine'}]
func TestGetDictionaryValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDictionaryValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.MapOfStringArray, []map[string]*string{
		{
			"1": to.StringPtr("one"),
			"2": to.StringPtr("two"),
			"3": to.StringPtr("three"),
		},
		{
			"4": to.StringPtr("four"),
			"5": to.StringPtr("five"),
			"6": to.StringPtr("six"),
		},
		{
			"7": to.StringPtr("seven"),
			"8": to.StringPtr("eight"),
			"9": to.StringPtr("nine"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDoubleInvalidNull - Get float array value [0.0, null, -1.2e20]
func TestGetDoubleInvalidNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDoubleInvalidNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Float64Array, []*float64{
		to.Float64Ptr(0),
		nil,
		to.Float64Ptr(-1.2e20),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDoubleInvalidString - Get boolean array value [1.0, 'number', 0.0]
func TestGetDoubleInvalidString(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDoubleInvalidString(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetDoubleValid - Get float array value [0, -0.01, 1.2e20]
func TestGetDoubleValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDoubleValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Float64Array, to.Float64PtrArray(0, -0.01, -1.2e20)); r != "" {
		t.Fatal(r)
	}
}

// GetDurationValid - Get duration array value ['P123DT22H14M12.011S', 'P5DT1H0M0S']
func TestGetDurationValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDurationValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.StringArray, to.StringPtrArray("P123DT22H14M12.011S", "P5DT1H0M0S")); r != "" {
		t.Fatal(r)
	}
}

// GetEmpty - Get empty array value []
func TestGetEmpty(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Int32Array == nil {
		t.Fatal("unexpected nil array")
	}
	if r := cmp.Diff(len(resp.Int32Array), 0); r != "" {
		t.Fatal(r)
	}
}

// GetEnumValid - Get enum array value ['foo1', 'foo2', 'foo3']
func TestGetEnumValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetEnumValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.FooEnumArray, []*FooEnum{
		FooEnumFoo1.ToPtr(), FooEnumFoo2.ToPtr(), FooEnumFoo3.ToPtr()}); r != "" {
		t.Fatal(r)
	}
}

// GetFloatInvalidNull - Get float array value [0.0, null, -1.2e20]
func TestGetFloatInvalidNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetFloatInvalidNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Float32Array, []*float32{
		to.Float32Ptr(0),
		nil,
		to.Float32Ptr(-1.2e20),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetFloatInvalidString - Get boolean array value [1.0, 'number', 0.0]
func TestGetFloatInvalidString(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetFloatInvalidString(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetFloatValid - Get float array value [0, -0.01, 1.2e20]
func TestGetFloatValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetFloatValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Float32Array, to.Float32PtrArray(0, -0.01, -1.2e20)); r != "" {
		t.Fatal(r)
	}
}

// GetIntInvalidNull - Get integer array value [1, null, 0]
func TestGetIntInvalidNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetIntInvalidNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Int32Array, []*int32{
		to.Int32Ptr(1),
		nil,
		to.Int32Ptr(0),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetIntInvalidString - Get integer array value [1, 'integer', 0]
func TestGetIntInvalidString(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetIntInvalidString(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetIntegerValid - Get integer array value [1, -1, 3, 300]
func TestGetIntegerValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetIntegerValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Int32Array, to.Int32PtrArray(1, -1, 3, 300)); r != "" {
		t.Fatal(r)
	}
}

// GetInvalid - Get invalid array [1, 2, 3
func TestGetInvalid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetInvalid(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetLongInvalidNull - Get long array value [1, null, 0]
func TestGetLongInvalidNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetLongInvalidNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Int64Array, []*int64{
		to.Int64Ptr(1),
		nil,
		to.Int64Ptr(0),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetLongInvalidString - Get long array value [1, 'integer', 0]
func TestGetLongInvalidString(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetLongInvalidString(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetLongValid - Get integer array value [1, -1, 3, 300]
func TestGetLongValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetLongValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Int64Array, to.Int64PtrArray(1, -1, 3, 300)); r != "" {
		t.Fatal(r)
	}
}

// GetNull - Get null array value
func TestGetNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Int32Array != nil {
		t.Fatal("expected nil array")
	}
}

// GetStringEnumValid - Get enum array value ['foo1', 'foo2', 'foo3']
func TestGetStringEnumValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetStringEnumValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Enum0Array, []*Enum0{
		Enum0Foo1.ToPtr(), Enum0Foo2.ToPtr(), Enum0Foo3.ToPtr()}); r != "" {
		t.Fatal(r)
	}
}

// GetStringValid - Get string array value ['foo1', 'foo2', 'foo3']
func TestGetStringValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetStringValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.StringArray, to.StringPtrArray("foo1", "foo2", "foo3")); r != "" {
		t.Fatal(r)
	}
}

// GetStringWithInvalid - Get string array value ['foo', 123, 'foo2']
func TestGetStringWithInvalid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetStringWithInvalid(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetStringWithNull - Get string array value ['foo', null, 'foo2']
func TestGetStringWithNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetStringWithNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.StringArray, []*string{to.StringPtr("foo"), nil, to.StringPtr("foo2")}); r != "" {
		t.Fatal(r)
	}
}

// GetUUIDInvalidChars - Get uuid array value ['6dcc7237-45fe-45c4-8a6b-3a8a3f625652', 'foo']
func TestGetUUIDInvalidChars(t *testing.T) {
	t.Skip("no strongly-typed UUID")
}

// GetUUIDValid - Get uuid array value ['6dcc7237-45fe-45c4-8a6b-3a8a3f625652', 'd1399005-30f7-40d6-8da6-dd7c89ad34db', 'f42f6aa1-a5bc-4ddf-907e-5f915de43205']
func TestGetUUIDValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetUUIDValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.StringArray, to.StringPtrArray("6dcc7237-45fe-45c4-8a6b-3a8a3f625652", "d1399005-30f7-40d6-8da6-dd7c89ad34db", "f42f6aa1-a5bc-4ddf-907e-5f915de43205")); r != "" {
		t.Fatal(r)
	}
}

// PutArrayValid - Put An array of array of strings [['1', '2', '3'], ['4', '5', '6'], ['7', '8', '9']]
func TestPutArrayValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutArrayValid(context.Background(), [][]*string{
		to.StringPtrArray("1", "2", "3"),
		to.StringPtrArray("4", "5", "6"),
		to.StringPtrArray("7", "8", "9"),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutBooleanTfft - Set array value empty [true, false, false, true]
func TestPutBooleanTfft(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutBooleanTfft(context.Background(), to.BoolPtrArray(true, false, false, true), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutByteValid - Put the array value [hex(FF FF FF FA), hex(01 02 03), hex (25, 29, 43)] with each elementencoded in base 64
func TestPutByteValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutByteValid(context.Background(), [][]byte{
		{0xFF, 0xFF, 0xFF, 0xFA},
		{0x01, 0x02, 0x03},
		{0x25, 0x29, 0x43},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutComplexValid - Put an array of complex type with values [{'integer': 1 'string': '2'}, {'integer': 3, 'string': '4'}, {'integer': 5, 'string': '6'}]
func TestPutComplexValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutComplexValid(context.Background(), []*Product{
		{Integer: to.Int32Ptr(1), String: to.StringPtr("2")},
		{Integer: to.Int32Ptr(3), String: to.StringPtr("4")},
		{Integer: to.Int32Ptr(5), String: to.StringPtr("6")},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutDateTimeRFC1123Valid - Set array value  ['Fri, 01 Dec 2000 00:00:01 GMT', 'Wed, 02 Jan 1980 00:11:35 GMT', 'Wed, 12 Oct 1492 10:15:01 GMT']
func TestPutDateTimeRFC1123Valid(t *testing.T) {
	client := newArrayClient()
	v1, _ := time.Parse(time.RFC1123, "Fri, 01 Dec 2000 00:00:01 GMT")
	v2, _ := time.Parse(time.RFC1123, "Wed, 02 Jan 1980 00:11:35 GMT")
	v3, _ := time.Parse(time.RFC1123, "Wed, 12 Oct 1492 10:15:01 GMT")
	resp, err := client.PutDateTimeRFC1123Valid(context.Background(), []*time.Time{&v1, &v2, &v3}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutDateTimeValid - Set array value  ['2000-12-01t00:00:01z', '1980-01-02T00:11:35+01:00', '1492-10-12T10:15:01-08:00']
func TestPutDateTimeValid(t *testing.T) {
	client := newArrayClient()
	v1, _ := time.Parse(time.RFC3339, "2000-12-01T00:00:01Z")
	v2, _ := time.Parse(time.RFC3339, "1980-01-02T00:11:35Z")
	v3, _ := time.Parse(time.RFC3339, "1492-10-12T10:15:01Z")
	resp, err := client.PutDateTimeValid(context.Background(), []*time.Time{&v1, &v2, &v3}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutDateValid - Set array value  ['2000-12-01', '1980-01-02', '1492-10-12']
func TestPutDateValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutDateValid(context.Background(), to.TimePtrArray(time.Date(2000, 12, 01, 0, 0, 0, 0, time.UTC),
		time.Date(1980, 01, 02, 0, 0, 0, 0, time.UTC),
		time.Date(1492, 10, 12, 0, 0, 0, 0, time.UTC)), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutDictionaryValid - Get an array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, {'4': 'four', '5': 'five', '6': 'six'}, {'7': 'seven', '8': 'eight', '9': 'nine'}]
func TestPutDictionaryValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutDictionaryValid(context.Background(), []map[string]*string{
		{
			"1": to.StringPtr("one"),
			"2": to.StringPtr("two"),
			"3": to.StringPtr("three"),
		},
		{
			"4": to.StringPtr("four"),
			"5": to.StringPtr("five"),
			"6": to.StringPtr("six"),
		},
		{
			"7": to.StringPtr("seven"),
			"8": to.StringPtr("eight"),
			"9": to.StringPtr("nine"),
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutDoubleValid - Set array value [0, -0.01, 1.2e20]
func TestPutDoubleValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutDoubleValid(context.Background(), to.Float64PtrArray(0, -0.01, -1.2e20), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutDurationValid - Set array value  ['P123DT22H14M12.011S', 'P5DT1H0M0S']
func TestPutDurationValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutDurationValid(context.Background(), to.StringPtrArray("P123DT22H14M12.011S", "P5DT1H"), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutEmpty - Set array value empty []
func TestPutEmpty(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutEmpty(context.Background(), []*string{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutEnumValid - Set array value ['foo1', 'foo2', 'foo3']
func TestPutEnumValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutEnumValid(context.Background(), []*FooEnum{
		FooEnumFoo1.ToPtr(), FooEnumFoo2.ToPtr(), FooEnumFoo3.ToPtr()}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutFloatValid - Set array value [0, -0.01, 1.2e20]
func TestPutFloatValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutFloatValid(context.Background(), to.Float32PtrArray(0, -0.01, -1.2e20), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutIntegerValid - Set array value empty [1, -1, 3, 300]
func TestPutIntegerValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutIntegerValid(context.Background(), to.Int32PtrArray(1, -1, 3, 300), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutLongValid - Set array value empty [1, -1, 3, 300]
func TestPutLongValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutLongValid(context.Background(), to.Int64PtrArray(1, -1, 3, 300), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutStringEnumValid - Set array value ['foo1', 'foo2', 'foo3']
func TestPutStringEnumValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutStringEnumValid(context.Background(), []*Enum1{
		Enum1Foo1.ToPtr(), Enum1Foo2.ToPtr(), Enum1Foo3.ToPtr()}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutStringValid - Set array value ['foo1', 'foo2', 'foo3']
func TestPutStringValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutStringValid(context.Background(), to.StringPtrArray("foo1", "foo2", "foo3"), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutUUIDValid - Set array value  ['6dcc7237-45fe-45c4-8a6b-3a8a3f625652', 'd1399005-30f7-40d6-8da6-dd7c89ad34db', 'f42f6aa1-a5bc-4ddf-907e-5f915de43205']
func TestPutUUIDValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutUUIDValid(context.Background(), to.StringPtrArray("6dcc7237-45fe-45c4-8a6b-3a8a3f625652", "d1399005-30f7-40d6-8da6-dd7c89ad34db", "f42f6aa1-a5bc-4ddf-907e-5f915de43205"), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
