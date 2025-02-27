//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package coreusagegroup

// Usage override to roundtrip.
type InputModel struct {
	// REQUIRED
	Name *string
}

// Not used anywhere, but access is override to public so still need to be generated and exported with serialization.
type OrphanModel struct {
	// REQUIRED
	Name *string
}

// Usage override to roundtrip.
type OutputModel struct {
	// REQUIRED
	Name *string
}
