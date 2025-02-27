//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package nodiscgroup

// The second level model in the normal multiple levels inheritance.
type Cat struct {
	// REQUIRED
	Age *int32

	// REQUIRED
	Name *string
}

// This is base model for not-discriminated normal multiple levels inheritance.
type Pet struct {
	// REQUIRED
	Name *string
}

// The third level model in the normal multiple levels inheritance.
type Siamese struct {
	// REQUIRED
	Age *int32

	// REQUIRED
	Name *string

	// REQUIRED
	Smart *bool
}
