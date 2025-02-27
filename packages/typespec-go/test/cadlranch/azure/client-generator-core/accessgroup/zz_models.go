//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package accessgroup

// Used in internal operations, should be generated but not exported.
type AbstractModel struct {
	// REQUIRED
	Kind *string

	// REQUIRED
	Name *string
}

// GetAbstractModel implements the AbstractModelClassification interface for type AbstractModel.
func (a *AbstractModel) GetAbstractModel() *AbstractModel { return a }

// Used in internal operations, should be generated but not exported.
type BaseModel struct {
	// REQUIRED
	Name *string
}

// Used in internal operations, should be generated but not exported.
type InnerModel struct {
	// REQUIRED
	Name *string
}

// Used in an internal operation, should be generated but not exported.
type InternalDecoratorModelInInternal struct {
	// REQUIRED
	Name *string
}

// Used in an internal operation, should be generated but not exported.
type NoDecoratorModelInInternal struct {
	// REQUIRED
	Name *string
}

// Used in a public operation, should be generated and exported.
type NoDecoratorModelInPublic struct {
	// REQUIRED
	Name *string
}

// Used in internal operations, should be generated but not exported.
type OuterModel struct {
	// REQUIRED
	Inner *InnerModel

	// REQUIRED
	Name *string
}

// Used in an internal operation but with public decorator, should be generated and exported.
type PublicDecoratorModelInInternal struct {
	// REQUIRED
	Name *string
}

// Used in a public operation, should be generated and exported.
type PublicDecoratorModelInPublic struct {
	// REQUIRED
	Name *string
}

// Used in internal operations, should be generated but not exported.
type RealModel struct {
	// REQUIRED
	Kind *string

	// REQUIRED
	Name *string
}

// GetAbstractModel implements the AbstractModelClassification interface for type RealModel.
func (r *RealModel) GetAbstractModel() *AbstractModel {
	return &AbstractModel{
		Kind: r.Kind,
		Name: r.Name,
	}
}

// Used by both public and internal operation. It should be generated and exported.
type SharedModel struct {
	// REQUIRED
	Name *string
}
