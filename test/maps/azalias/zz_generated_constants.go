//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azalias

const (
	module  = "azalias"
	version = "v0.1.0"
)

// GeoJSONObjectType - Specifies the GeoJSON type. Must be one of the nine valid GeoJSON object types - Point, MultiPoint, LineString, MultiLineString,
// Polygon, MultiPolygon, GeometryCollection, Feature and
// FeatureCollection.
type GeoJSONObjectType string

const (
	// GeoJSONObjectTypeGeoJSONPoint - `GeoJSON Point` geometry.
	GeoJSONObjectTypeGeoJSONPoint GeoJSONObjectType = "Point"
	// GeoJSONObjectTypeGeoJSONMultiPoint - `GeoJSON MultiPoint` geometry.
	GeoJSONObjectTypeGeoJSONMultiPoint GeoJSONObjectType = "MultiPoint"
	// GeoJSONObjectTypeGeoJSONLineString - `GeoJSON LineString` geometry.
	GeoJSONObjectTypeGeoJSONLineString GeoJSONObjectType = "LineString"
	// GeoJSONObjectTypeGeoJSONMultiLineString - `GeoJSON MultiLineString` geometry.
	GeoJSONObjectTypeGeoJSONMultiLineString GeoJSONObjectType = "MultiLineString"
	// GeoJSONObjectTypeGeoJSONPolygon - `GeoJSON Polygon` geometry.
	GeoJSONObjectTypeGeoJSONPolygon GeoJSONObjectType = "Polygon"
	// GeoJSONObjectTypeGeoJSONMultiPolygon - `GeoJSON MultiPolygon` geometry.
	GeoJSONObjectTypeGeoJSONMultiPolygon GeoJSONObjectType = "MultiPolygon"
	// GeoJSONObjectTypeGeoJSONGeometryCollection - `GeoJSON GeometryCollection` geometry.
	GeoJSONObjectTypeGeoJSONGeometryCollection GeoJSONObjectType = "GeometryCollection"
	// GeoJSONObjectTypeGeoJSONFeature - `GeoJSON Feature` object.
	GeoJSONObjectTypeGeoJSONFeature GeoJSONObjectType = "Feature"
	// GeoJSONObjectTypeGeoJSONFeatureCollection - `GeoJSON FeatureCollection` object.
	GeoJSONObjectTypeGeoJSONFeatureCollection GeoJSONObjectType = "FeatureCollection"
)

// PossibleGeoJSONObjectTypeValues returns the possible values for the GeoJSONObjectType const type.
func PossibleGeoJSONObjectTypeValues() []GeoJSONObjectType {
	return []GeoJSONObjectType{
		GeoJSONObjectTypeGeoJSONPoint,
		GeoJSONObjectTypeGeoJSONMultiPoint,
		GeoJSONObjectTypeGeoJSONLineString,
		GeoJSONObjectTypeGeoJSONMultiLineString,
		GeoJSONObjectTypeGeoJSONPolygon,
		GeoJSONObjectTypeGeoJSONMultiPolygon,
		GeoJSONObjectTypeGeoJSONGeometryCollection,
		GeoJSONObjectTypeGeoJSONFeature,
		GeoJSONObjectTypeGeoJSONFeatureCollection,
	}
}

// ToPtr returns a *GeoJSONObjectType pointing to the current value.
func (c GeoJSONObjectType) ToPtr() *GeoJSONObjectType {
	return &c
}

type GeographicResourceLocation string

const (
	// GeographicResourceLocationEu - Used to access an Azure Maps Creator resource in Europe
	GeographicResourceLocationEu GeographicResourceLocation = "eu"
	// GeographicResourceLocationUs - Used to access an Azure Maps Creator resource in the United States
	GeographicResourceLocationUs GeographicResourceLocation = "us"
)

// PossibleGeographicResourceLocationValues returns the possible values for the GeographicResourceLocation const type.
func PossibleGeographicResourceLocationValues() []GeographicResourceLocation {
	return []GeographicResourceLocation{
		GeographicResourceLocationEu,
		GeographicResourceLocationUs,
	}
}

// ToPtr returns a *GeographicResourceLocation pointing to the current value.
func (c GeographicResourceLocation) ToPtr() *GeographicResourceLocation {
	return &c
}

// Geography - This parameter specifies where the Azure Maps Creator resource is located. Valid values are us and eu.
type Geography string

const (
	GeographyEu Geography = "eu"
	GeographyUs Geography = "us"
)

// PossibleGeographyValues returns the possible values for the Geography const type.
func PossibleGeographyValues() []Geography {
	return []Geography{
		GeographyEu,
		GeographyUs,
	}
}

// ToPtr returns a *Geography pointing to the current value.
func (c Geography) ToPtr() *Geography {
	return &c
}

type LogMetricsGroupBy string

const (
	LogMetricsGroupByCacheStatus    LogMetricsGroupBy = "cacheStatus"
	LogMetricsGroupByCountry        LogMetricsGroupBy = "country"
	LogMetricsGroupByCustomDomain   LogMetricsGroupBy = "customDomain"
	LogMetricsGroupByHTTPStatusCode LogMetricsGroupBy = "httpStatusCode"
	LogMetricsGroupByProtocol       LogMetricsGroupBy = "protocol"
)

// PossibleLogMetricsGroupByValues returns the possible values for the LogMetricsGroupBy const type.
func PossibleLogMetricsGroupByValues() []LogMetricsGroupBy {
	return []LogMetricsGroupBy{
		LogMetricsGroupByCacheStatus,
		LogMetricsGroupByCountry,
		LogMetricsGroupByCustomDomain,
		LogMetricsGroupByHTTPStatusCode,
		LogMetricsGroupByProtocol,
	}
}

// ToPtr returns a *LogMetricsGroupBy pointing to the current value.
func (c LogMetricsGroupBy) ToPtr() *LogMetricsGroupBy {
	return &c
}

type SomethingCount int32

const (
	SomethingCountTen    SomethingCount = 10
	SomethingCountTwenty SomethingCount = 20
	SomethingCountThirty SomethingCount = 30
	SomethingCountForty  SomethingCount = 40
)

// PossibleSomethingCountValues returns the possible values for the SomethingCount const type.
func PossibleSomethingCountValues() []SomethingCount {
	return []SomethingCount{
		SomethingCountTen,
		SomethingCountTwenty,
		SomethingCountThirty,
		SomethingCountForty,
	}
}

// ToPtr returns a *SomethingCount pointing to the current value.
func (c SomethingCount) ToPtr() *SomethingCount {
	return &c
}
