//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import "net/http"

// AggregatedCostClientGetByManagementGroupResponse contains the response from method AggregatedCostClient.GetByManagementGroup.
type AggregatedCostClientGetByManagementGroupResponse struct {
	ManagementGroupAggregatedCostResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AggregatedCostClientGetForBillingPeriodByManagementGroupResponse contains the response from method AggregatedCostClient.GetForBillingPeriodByManagementGroup.
type AggregatedCostClientGetForBillingPeriodByManagementGroupResponse struct {
	ManagementGroupAggregatedCostResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BalancesClientGetByBillingAccountResponse contains the response from method BalancesClient.GetByBillingAccount.
type BalancesClientGetByBillingAccountResponse struct {
	Balance
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BalancesClientGetForBillingPeriodByBillingAccountResponse contains the response from method BalancesClient.GetForBillingPeriodByBillingAccount.
type BalancesClientGetForBillingPeriodByBillingAccountResponse struct {
	Balance
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BudgetsClientCreateOrUpdateResponse contains the response from method BudgetsClient.CreateOrUpdate.
type BudgetsClientCreateOrUpdateResponse struct {
	Budget
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BudgetsClientDeleteResponse contains the response from method BudgetsClient.Delete.
type BudgetsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BudgetsClientGetResponse contains the response from method BudgetsClient.Get.
type BudgetsClientGetResponse struct {
	Budget
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BudgetsClientListResponse contains the response from method BudgetsClient.List.
type BudgetsClientListResponse struct {
	BudgetsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ChargesClientListResponse contains the response from method ChargesClient.List.
type ChargesClientListResponse struct {
	ChargesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CreditsClientGetResponse contains the response from method CreditsClient.Get.
type CreditsClientGetResponse struct {
	CreditSummary
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EventsClientListResponse contains the response from method EventsClient.List.
type EventsClientListResponse struct {
	Events
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ForecastsClientListResponse contains the response from method ForecastsClient.List.
type ForecastsClientListResponse struct {
	ForecastsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// LotsClientListResponse contains the response from method LotsClient.List.
type LotsClientListResponse struct {
	Lots
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MarketplacesClientListResponse contains the response from method MarketplacesClient.List.
type MarketplacesClientListResponse struct {
	MarketplacesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PriceSheetClientGetByBillingPeriodResponse contains the response from method PriceSheetClient.GetByBillingPeriod.
type PriceSheetClientGetByBillingPeriodResponse struct {
	PriceSheetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PriceSheetClientGetResponse contains the response from method PriceSheetClient.Get.
type PriceSheetClientGetResponse struct {
	PriceSheetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ReservationRecommendationDetailsClientGetResponse contains the response from method ReservationRecommendationDetailsClient.Get.
type ReservationRecommendationDetailsClientGetResponse struct {
	ReservationRecommendationDetailsModel
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ReservationRecommendationsClientListResponse contains the response from method ReservationRecommendationsClient.List.
type ReservationRecommendationsClientListResponse struct {
	ReservationRecommendationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ReservationTransactionsClientListByBillingProfileResponse contains the response from method ReservationTransactionsClient.ListByBillingProfile.
type ReservationTransactionsClientListByBillingProfileResponse struct {
	ModernReservationTransactionsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ReservationTransactionsClientListResponse contains the response from method ReservationTransactionsClient.List.
type ReservationTransactionsClientListResponse struct {
	ReservationTransactionsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ReservationsDetailsClientListByReservationOrderAndReservationResponse contains the response from method ReservationsDetailsClient.ListByReservationOrderAndReservation.
type ReservationsDetailsClientListByReservationOrderAndReservationResponse struct {
	ReservationDetailsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ReservationsDetailsClientListByReservationOrderResponse contains the response from method ReservationsDetailsClient.ListByReservationOrder.
type ReservationsDetailsClientListByReservationOrderResponse struct {
	ReservationDetailsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ReservationsDetailsClientListResponse contains the response from method ReservationsDetailsClient.List.
type ReservationsDetailsClientListResponse struct {
	ReservationDetailsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ReservationsSummariesClientListByReservationOrderAndReservationResponse contains the response from method ReservationsSummariesClient.ListByReservationOrderAndReservation.
type ReservationsSummariesClientListByReservationOrderAndReservationResponse struct {
	ReservationSummariesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ReservationsSummariesClientListByReservationOrderResponse contains the response from method ReservationsSummariesClient.ListByReservationOrder.
type ReservationsSummariesClientListByReservationOrderResponse struct {
	ReservationSummariesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ReservationsSummariesClientListResponse contains the response from method ReservationsSummariesClient.List.
type ReservationsSummariesClientListResponse struct {
	ReservationSummariesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TagsClientGetResponse contains the response from method TagsClient.Get.
type TagsClientGetResponse struct {
	TagsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// UsageDetailsClientListResponse contains the response from method UsageDetailsClient.List.
type UsageDetailsClientListResponse struct {
	UsageDetailsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
