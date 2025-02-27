//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurereportgroup

// AutoRestReportServiceForAzureClientGetReportOptions contains the optional parameters for the AutoRestReportServiceForAzureClient.GetReport
// method.
type AutoRestReportServiceForAzureClientGetReportOptions struct {
	// If specified, qualifies the generated report further (e.g. '2.7' vs '3.5' in for Python). The only effect is, that generators
	// that run all tests several times, can distinguish the generated reports.
	Qualifier *string
}
