//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package extenumsgroup

const host = "http://localhost:3000"

const (
	module  = "extenumsgroup"
	version = "v0.1.0"
)

// DaysOfWeekExtensibleEnum - Type of Pet
type DaysOfWeekExtensibleEnum string

const (
	DaysOfWeekExtensibleEnumFriday    DaysOfWeekExtensibleEnum = "Friday"
	DaysOfWeekExtensibleEnumMonday    DaysOfWeekExtensibleEnum = "Monday"
	DaysOfWeekExtensibleEnumSaturday  DaysOfWeekExtensibleEnum = "Saturday"
	DaysOfWeekExtensibleEnumSunday    DaysOfWeekExtensibleEnum = "Sunday"
	DaysOfWeekExtensibleEnumThursday  DaysOfWeekExtensibleEnum = "Thursday"
	DaysOfWeekExtensibleEnumTuesday   DaysOfWeekExtensibleEnum = "Tuesday"
	DaysOfWeekExtensibleEnumWednesday DaysOfWeekExtensibleEnum = "Wednesday"
)

// PossibleDaysOfWeekExtensibleEnumValues returns the possible values for the DaysOfWeekExtensibleEnum const type.
func PossibleDaysOfWeekExtensibleEnumValues() []DaysOfWeekExtensibleEnum {
	return []DaysOfWeekExtensibleEnum{
		DaysOfWeekExtensibleEnumFriday,
		DaysOfWeekExtensibleEnumMonday,
		DaysOfWeekExtensibleEnumSaturday,
		DaysOfWeekExtensibleEnumSunday,
		DaysOfWeekExtensibleEnumThursday,
		DaysOfWeekExtensibleEnumTuesday,
		DaysOfWeekExtensibleEnumWednesday,
	}
}

// ToPtr returns a *DaysOfWeekExtensibleEnum pointing to the current value.
func (c DaysOfWeekExtensibleEnum) ToPtr() *DaysOfWeekExtensibleEnum {
	return &c
}

type IntEnum string

const (
	// IntEnumOne - This is a really long comment to see what wrapping looks like. This comment is really long and it should wrap for readability. Please wrap.
	// This should wrap.
	IntEnumOne IntEnum = "1"
	// IntEnumThree - three
	IntEnumThree IntEnum = "3"
	// IntEnumTwo - two
	IntEnumTwo IntEnum = "2"
)

// PossibleIntEnumValues returns the possible values for the IntEnum const type.
func PossibleIntEnumValues() []IntEnum {
	return []IntEnum{
		IntEnumOne,
		IntEnumThree,
		IntEnumTwo,
	}
}

// ToPtr returns a *IntEnum pointing to the current value.
func (c IntEnum) ToPtr() *IntEnum {
	return &c
}
