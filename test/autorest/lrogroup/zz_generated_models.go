//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package lrogroup

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// Implements the error and azcore.HTTPResponse interfaces.
type CloudError struct {
	raw     string
	Code    *int32  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// Error implements the error interface for type CloudError.
// The contents of the error text are not contractual and subject to change.
func (e CloudError) Error() string {
	return e.raw
}

// LRORetrysBeginDelete202Retry200Options contains the optional parameters for the LRORetrys.BeginDelete202Retry200 method.
type LRORetrysBeginDelete202Retry200Options struct {
	// placeholder for future optional parameters
}

// LRORetrysBeginDeleteAsyncRelativeRetrySucceededOptions contains the optional parameters for the LRORetrys.BeginDeleteAsyncRelativeRetrySucceeded method.
type LRORetrysBeginDeleteAsyncRelativeRetrySucceededOptions struct {
	// placeholder for future optional parameters
}

// LRORetrysBeginDeleteProvisioning202Accepted200SucceededOptions contains the optional parameters for the LRORetrys.BeginDeleteProvisioning202Accepted200Succeeded
// method.
type LRORetrysBeginDeleteProvisioning202Accepted200SucceededOptions struct {
	// placeholder for future optional parameters
}

// LRORetrysBeginPost202Retry200Options contains the optional parameters for the LRORetrys.BeginPost202Retry200 method.
type LRORetrysBeginPost202Retry200Options struct {
	// Product to put
	Product *Product
}

// LRORetrysBeginPostAsyncRelativeRetrySucceededOptions contains the optional parameters for the LRORetrys.BeginPostAsyncRelativeRetrySucceeded method.
type LRORetrysBeginPostAsyncRelativeRetrySucceededOptions struct {
	// Product to put
	Product *Product
}

// LRORetrysBeginPut201CreatingSucceeded200Options contains the optional parameters for the LRORetrys.BeginPut201CreatingSucceeded200 method.
type LRORetrysBeginPut201CreatingSucceeded200Options struct {
	// Product to put
	Product *Product
}

// LRORetrysBeginPutAsyncRelativeRetrySucceededOptions contains the optional parameters for the LRORetrys.BeginPutAsyncRelativeRetrySucceeded method.
type LRORetrysBeginPutAsyncRelativeRetrySucceededOptions struct {
	// Product to put
	Product *Product
}

// LROSADsBeginDelete202NonRetry400Options contains the optional parameters for the LROSADs.BeginDelete202NonRetry400 method.
type LROSADsBeginDelete202NonRetry400Options struct {
	// placeholder for future optional parameters
}

// LROSADsBeginDelete202RetryInvalidHeaderOptions contains the optional parameters for the LROSADs.BeginDelete202RetryInvalidHeader method.
type LROSADsBeginDelete202RetryInvalidHeaderOptions struct {
	// placeholder for future optional parameters
}

// LROSADsBeginDelete204SucceededOptions contains the optional parameters for the LROSADs.BeginDelete204Succeeded method.
type LROSADsBeginDelete204SucceededOptions struct {
	// placeholder for future optional parameters
}

// LROSADsBeginDeleteAsyncRelativeRetry400Options contains the optional parameters for the LROSADs.BeginDeleteAsyncRelativeRetry400 method.
type LROSADsBeginDeleteAsyncRelativeRetry400Options struct {
	// placeholder for future optional parameters
}

// LROSADsBeginDeleteAsyncRelativeRetryInvalidHeaderOptions contains the optional parameters for the LROSADs.BeginDeleteAsyncRelativeRetryInvalidHeader
// method.
type LROSADsBeginDeleteAsyncRelativeRetryInvalidHeaderOptions struct {
	// placeholder for future optional parameters
}

// LROSADsBeginDeleteAsyncRelativeRetryInvalidJSONPollingOptions contains the optional parameters for the LROSADs.BeginDeleteAsyncRelativeRetryInvalidJSONPolling
// method.
type LROSADsBeginDeleteAsyncRelativeRetryInvalidJSONPollingOptions struct {
	// placeholder for future optional parameters
}

// LROSADsBeginDeleteAsyncRelativeRetryNoStatusOptions contains the optional parameters for the LROSADs.BeginDeleteAsyncRelativeRetryNoStatus method.
type LROSADsBeginDeleteAsyncRelativeRetryNoStatusOptions struct {
	// placeholder for future optional parameters
}

// LROSADsBeginDeleteNonRetry400Options contains the optional parameters for the LROSADs.BeginDeleteNonRetry400 method.
type LROSADsBeginDeleteNonRetry400Options struct {
	// placeholder for future optional parameters
}

// LROSADsBeginPost202NoLocationOptions contains the optional parameters for the LROSADs.BeginPost202NoLocation method.
type LROSADsBeginPost202NoLocationOptions struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPost202NonRetry400Options contains the optional parameters for the LROSADs.BeginPost202NonRetry400 method.
type LROSADsBeginPost202NonRetry400Options struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPost202RetryInvalidHeaderOptions contains the optional parameters for the LROSADs.BeginPost202RetryInvalidHeader method.
type LROSADsBeginPost202RetryInvalidHeaderOptions struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPostAsyncRelativeRetry400Options contains the optional parameters for the LROSADs.BeginPostAsyncRelativeRetry400 method.
type LROSADsBeginPostAsyncRelativeRetry400Options struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPostAsyncRelativeRetryInvalidHeaderOptions contains the optional parameters for the LROSADs.BeginPostAsyncRelativeRetryInvalidHeader method.
type LROSADsBeginPostAsyncRelativeRetryInvalidHeaderOptions struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPostAsyncRelativeRetryInvalidJSONPollingOptions contains the optional parameters for the LROSADs.BeginPostAsyncRelativeRetryInvalidJSONPolling
// method.
type LROSADsBeginPostAsyncRelativeRetryInvalidJSONPollingOptions struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPostAsyncRelativeRetryNoPayloadOptions contains the optional parameters for the LROSADs.BeginPostAsyncRelativeRetryNoPayload method.
type LROSADsBeginPostAsyncRelativeRetryNoPayloadOptions struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPostNonRetry400Options contains the optional parameters for the LROSADs.BeginPostNonRetry400 method.
type LROSADsBeginPostNonRetry400Options struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPut200InvalidJSONOptions contains the optional parameters for the LROSADs.BeginPut200InvalidJSON method.
type LROSADsBeginPut200InvalidJSONOptions struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPutAsyncRelativeRetry400Options contains the optional parameters for the LROSADs.BeginPutAsyncRelativeRetry400 method.
type LROSADsBeginPutAsyncRelativeRetry400Options struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPutAsyncRelativeRetryInvalidHeaderOptions contains the optional parameters for the LROSADs.BeginPutAsyncRelativeRetryInvalidHeader method.
type LROSADsBeginPutAsyncRelativeRetryInvalidHeaderOptions struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPutAsyncRelativeRetryInvalidJSONPollingOptions contains the optional parameters for the LROSADs.BeginPutAsyncRelativeRetryInvalidJSONPolling
// method.
type LROSADsBeginPutAsyncRelativeRetryInvalidJSONPollingOptions struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPutAsyncRelativeRetryNoStatusOptions contains the optional parameters for the LROSADs.BeginPutAsyncRelativeRetryNoStatus method.
type LROSADsBeginPutAsyncRelativeRetryNoStatusOptions struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPutAsyncRelativeRetryNoStatusPayloadOptions contains the optional parameters for the LROSADs.BeginPutAsyncRelativeRetryNoStatusPayload method.
type LROSADsBeginPutAsyncRelativeRetryNoStatusPayloadOptions struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPutError201NoProvisioningStatePayloadOptions contains the optional parameters for the LROSADs.BeginPutError201NoProvisioningStatePayload
// method.
type LROSADsBeginPutError201NoProvisioningStatePayloadOptions struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPutNonRetry201Creating400InvalidJSONOptions contains the optional parameters for the LROSADs.BeginPutNonRetry201Creating400InvalidJSON method.
type LROSADsBeginPutNonRetry201Creating400InvalidJSONOptions struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPutNonRetry201Creating400Options contains the optional parameters for the LROSADs.BeginPutNonRetry201Creating400 method.
type LROSADsBeginPutNonRetry201Creating400Options struct {
	// Product to put
	Product *Product
}

// LROSADsBeginPutNonRetry400Options contains the optional parameters for the LROSADs.BeginPutNonRetry400 method.
type LROSADsBeginPutNonRetry400Options struct {
	// Product to put
	Product *Product
}

// LROsBeginDelete202NoRetry204Options contains the optional parameters for the LROs.BeginDelete202NoRetry204 method.
type LROsBeginDelete202NoRetry204Options struct {
	// placeholder for future optional parameters
}

// LROsBeginDelete202Retry200Options contains the optional parameters for the LROs.BeginDelete202Retry200 method.
type LROsBeginDelete202Retry200Options struct {
	// placeholder for future optional parameters
}

// LROsBeginDelete204SucceededOptions contains the optional parameters for the LROs.BeginDelete204Succeeded method.
type LROsBeginDelete204SucceededOptions struct {
	// placeholder for future optional parameters
}

// LROsBeginDeleteAsyncNoHeaderInRetryOptions contains the optional parameters for the LROs.BeginDeleteAsyncNoHeaderInRetry method.
type LROsBeginDeleteAsyncNoHeaderInRetryOptions struct {
	// placeholder for future optional parameters
}

// LROsBeginDeleteAsyncNoRetrySucceededOptions contains the optional parameters for the LROs.BeginDeleteAsyncNoRetrySucceeded method.
type LROsBeginDeleteAsyncNoRetrySucceededOptions struct {
	// placeholder for future optional parameters
}

// LROsBeginDeleteAsyncRetryFailedOptions contains the optional parameters for the LROs.BeginDeleteAsyncRetryFailed method.
type LROsBeginDeleteAsyncRetryFailedOptions struct {
	// placeholder for future optional parameters
}

// LROsBeginDeleteAsyncRetrySucceededOptions contains the optional parameters for the LROs.BeginDeleteAsyncRetrySucceeded method.
type LROsBeginDeleteAsyncRetrySucceededOptions struct {
	// placeholder for future optional parameters
}

// LROsBeginDeleteAsyncRetrycanceledOptions contains the optional parameters for the LROs.BeginDeleteAsyncRetrycanceled method.
type LROsBeginDeleteAsyncRetrycanceledOptions struct {
	// placeholder for future optional parameters
}

// LROsBeginDeleteNoHeaderInRetryOptions contains the optional parameters for the LROs.BeginDeleteNoHeaderInRetry method.
type LROsBeginDeleteNoHeaderInRetryOptions struct {
	// placeholder for future optional parameters
}

// LROsBeginDeleteProvisioning202Accepted200SucceededOptions contains the optional parameters for the LROs.BeginDeleteProvisioning202Accepted200Succeeded
// method.
type LROsBeginDeleteProvisioning202Accepted200SucceededOptions struct {
	// placeholder for future optional parameters
}

// LROsBeginDeleteProvisioning202DeletingFailed200Options contains the optional parameters for the LROs.BeginDeleteProvisioning202DeletingFailed200 method.
type LROsBeginDeleteProvisioning202DeletingFailed200Options struct {
	// placeholder for future optional parameters
}

// LROsBeginDeleteProvisioning202Deletingcanceled200Options contains the optional parameters for the LROs.BeginDeleteProvisioning202Deletingcanceled200
// method.
type LROsBeginDeleteProvisioning202Deletingcanceled200Options struct {
	// placeholder for future optional parameters
}

// LROsBeginPatch200SucceededIgnoreHeadersOptions contains the optional parameters for the LROs.BeginPatch200SucceededIgnoreHeaders method.
type LROsBeginPatch200SucceededIgnoreHeadersOptions struct {
	// Product to patch
	Product *Product
}

// LROsBeginPost200WithPayloadOptions contains the optional parameters for the LROs.BeginPost200WithPayload method.
type LROsBeginPost200WithPayloadOptions struct {
	// placeholder for future optional parameters
}

// LROsBeginPost202ListOptions contains the optional parameters for the LROs.BeginPost202List method.
type LROsBeginPost202ListOptions struct {
	// placeholder for future optional parameters
}

// LROsBeginPost202NoRetry204Options contains the optional parameters for the LROs.BeginPost202NoRetry204 method.
type LROsBeginPost202NoRetry204Options struct {
	// Product to put
	Product *Product
}

// LROsBeginPost202Retry200Options contains the optional parameters for the LROs.BeginPost202Retry200 method.
type LROsBeginPost202Retry200Options struct {
	// Product to put
	Product *Product
}

// LROsBeginPostAsyncNoRetrySucceededOptions contains the optional parameters for the LROs.BeginPostAsyncNoRetrySucceeded method.
type LROsBeginPostAsyncNoRetrySucceededOptions struct {
	// Product to put
	Product *Product
}

// LROsBeginPostAsyncRetryFailedOptions contains the optional parameters for the LROs.BeginPostAsyncRetryFailed method.
type LROsBeginPostAsyncRetryFailedOptions struct {
	// Product to put
	Product *Product
}

// LROsBeginPostAsyncRetrySucceededOptions contains the optional parameters for the LROs.BeginPostAsyncRetrySucceeded method.
type LROsBeginPostAsyncRetrySucceededOptions struct {
	// Product to put
	Product *Product
}

// LROsBeginPostAsyncRetrycanceledOptions contains the optional parameters for the LROs.BeginPostAsyncRetrycanceled method.
type LROsBeginPostAsyncRetrycanceledOptions struct {
	// Product to put
	Product *Product
}

// LROsBeginPostDoubleHeadersFinalAzureHeaderGetDefaultOptions contains the optional parameters for the LROs.BeginPostDoubleHeadersFinalAzureHeaderGetDefault
// method.
type LROsBeginPostDoubleHeadersFinalAzureHeaderGetDefaultOptions struct {
	// placeholder for future optional parameters
}

// LROsBeginPostDoubleHeadersFinalAzureHeaderGetOptions contains the optional parameters for the LROs.BeginPostDoubleHeadersFinalAzureHeaderGet method.
type LROsBeginPostDoubleHeadersFinalAzureHeaderGetOptions struct {
	// placeholder for future optional parameters
}

// LROsBeginPostDoubleHeadersFinalLocationGetOptions contains the optional parameters for the LROs.BeginPostDoubleHeadersFinalLocationGet method.
type LROsBeginPostDoubleHeadersFinalLocationGetOptions struct {
	// placeholder for future optional parameters
}

// LROsBeginPut200Acceptedcanceled200Options contains the optional parameters for the LROs.BeginPut200Acceptedcanceled200 method.
type LROsBeginPut200Acceptedcanceled200Options struct {
	// Product to put
	Product *Product
}

// LROsBeginPut200SucceededNoStateOptions contains the optional parameters for the LROs.BeginPut200SucceededNoState method.
type LROsBeginPut200SucceededNoStateOptions struct {
	// Product to put
	Product *Product
}

// LROsBeginPut200SucceededOptions contains the optional parameters for the LROs.BeginPut200Succeeded method.
type LROsBeginPut200SucceededOptions struct {
	// Product to put
	Product *Product
}

// LROsBeginPut200UpdatingSucceeded204Options contains the optional parameters for the LROs.BeginPut200UpdatingSucceeded204 method.
type LROsBeginPut200UpdatingSucceeded204Options struct {
	// Product to put
	Product *Product
}

// LROsBeginPut201CreatingFailed200Options contains the optional parameters for the LROs.BeginPut201CreatingFailed200 method.
type LROsBeginPut201CreatingFailed200Options struct {
	// Product to put
	Product *Product
}

// LROsBeginPut201CreatingSucceeded200Options contains the optional parameters for the LROs.BeginPut201CreatingSucceeded200 method.
type LROsBeginPut201CreatingSucceeded200Options struct {
	// Product to put
	Product *Product
}

// LROsBeginPut201SucceededOptions contains the optional parameters for the LROs.BeginPut201Succeeded method.
type LROsBeginPut201SucceededOptions struct {
	// Product to put
	Product *Product
}

// LROsBeginPut202Retry200Options contains the optional parameters for the LROs.BeginPut202Retry200 method.
type LROsBeginPut202Retry200Options struct {
	// Product to put
	Product *Product
}

// LROsBeginPutAsyncNoHeaderInRetryOptions contains the optional parameters for the LROs.BeginPutAsyncNoHeaderInRetry method.
type LROsBeginPutAsyncNoHeaderInRetryOptions struct {
	// Product to put
	Product *Product
}

// LROsBeginPutAsyncNoRetrySucceededOptions contains the optional parameters for the LROs.BeginPutAsyncNoRetrySucceeded method.
type LROsBeginPutAsyncNoRetrySucceededOptions struct {
	// Product to put
	Product *Product
}

// LROsBeginPutAsyncNoRetrycanceledOptions contains the optional parameters for the LROs.BeginPutAsyncNoRetrycanceled method.
type LROsBeginPutAsyncNoRetrycanceledOptions struct {
	// Product to put
	Product *Product
}

// LROsBeginPutAsyncNonResourceOptions contains the optional parameters for the LROs.BeginPutAsyncNonResource method.
type LROsBeginPutAsyncNonResourceOptions struct {
	// Sku to put
	SKU *SKU
}

// LROsBeginPutAsyncRetryFailedOptions contains the optional parameters for the LROs.BeginPutAsyncRetryFailed method.
type LROsBeginPutAsyncRetryFailedOptions struct {
	// Product to put
	Product *Product
}

// LROsBeginPutAsyncRetrySucceededOptions contains the optional parameters for the LROs.BeginPutAsyncRetrySucceeded method.
type LROsBeginPutAsyncRetrySucceededOptions struct {
	// Product to put
	Product *Product
}

// LROsBeginPutAsyncSubResourceOptions contains the optional parameters for the LROs.BeginPutAsyncSubResource method.
type LROsBeginPutAsyncSubResourceOptions struct {
	// Sub Product to put
	Product *SubProduct
}

// LROsBeginPutNoHeaderInRetryOptions contains the optional parameters for the LROs.BeginPutNoHeaderInRetry method.
type LROsBeginPutNoHeaderInRetryOptions struct {
	// Product to put
	Product *Product
}

// LROsBeginPutNonResourceOptions contains the optional parameters for the LROs.BeginPutNonResource method.
type LROsBeginPutNonResourceOptions struct {
	// sku to put
	SKU *SKU
}

// LROsBeginPutSubResourceOptions contains the optional parameters for the LROs.BeginPutSubResource method.
type LROsBeginPutSubResourceOptions struct {
	// Sub Product to put
	Product *SubProduct
}

// LROsCustomHeaderBeginPost202Retry200Options contains the optional parameters for the LROsCustomHeader.BeginPost202Retry200 method.
type LROsCustomHeaderBeginPost202Retry200Options struct {
	// Product to put
	Product *Product
}

// LROsCustomHeaderBeginPostAsyncRetrySucceededOptions contains the optional parameters for the LROsCustomHeader.BeginPostAsyncRetrySucceeded method.
type LROsCustomHeaderBeginPostAsyncRetrySucceededOptions struct {
	// Product to put
	Product *Product
}

// LROsCustomHeaderBeginPut201CreatingSucceeded200Options contains the optional parameters for the LROsCustomHeader.BeginPut201CreatingSucceeded200 method.
type LROsCustomHeaderBeginPut201CreatingSucceeded200Options struct {
	// Product to put
	Product *Product
}

// LROsCustomHeaderBeginPutAsyncRetrySucceededOptions contains the optional parameters for the LROsCustomHeader.BeginPutAsyncRetrySucceeded method.
type LROsCustomHeaderBeginPutAsyncRetrySucceededOptions struct {
	// Product to put
	Product *Product
}

type OperationResult struct {
	Error *OperationResultError `json:"error,omitempty"`

	// The status of the request
	Status *OperationResultStatus `json:"status,omitempty"`
}

type OperationResultError struct {
	// The error code for an operation failure
	Code *int32 `json:"code,omitempty"`

	// The detailed arror message
	Message *string `json:"message,omitempty"`
}

type Product struct {
	Resource
	Properties *ProductProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Product.
func (p Product) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	p.Resource.marshalInternal(objectMap)
	populate(objectMap, "properties", p.Properties)
	return json.Marshal(objectMap)
}

type ProductProperties struct {
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// READ-ONLY
	ProvisioningStateValues *ProductPropertiesProvisioningStateValues `json:"provisioningStateValues,omitempty" azure:"ro"`
}

type Resource struct {
	// Resource Location
	Location *string `json:"location,omitempty"`

	// Dictionary of
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Resource Id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource Name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource Type
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	r.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (r Resource) marshalInternal(objectMap map[string]interface{}) {
	populate(objectMap, "id", r.ID)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
}

type SKU struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type SubProduct struct {
	SubResource
	Properties *SubProductProperties `json:"properties,omitempty"`
}

type SubProductProperties struct {
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// READ-ONLY
	ProvisioningStateValues *SubProductPropertiesProvisioningStateValues `json:"provisioningStateValues,omitempty" azure:"ro"`
}

type SubResource struct {
	// READ-ONLY; Sub Resource Id
	ID *string `json:"id,omitempty" azure:"ro"`
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}
