/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// RackRequestOuterUnit * `mm` - Millimeters * `in` - Inches
type RackRequestOuterUnit string

// List of RackRequest_outer_unit
const (
	RACKREQUESTOUTERUNIT_MM RackRequestOuterUnit = "mm"
	RACKREQUESTOUTERUNIT_IN RackRequestOuterUnit = "in"
	RACKREQUESTOUTERUNIT_EMPTY RackRequestOuterUnit = ""
)

// All allowed values of RackRequestOuterUnit enum
var AllowedRackRequestOuterUnitEnumValues = []RackRequestOuterUnit{
	"mm",
	"in",
	"",
}

func (v *RackRequestOuterUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RackRequestOuterUnit(value)
	for _, existing := range AllowedRackRequestOuterUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RackRequestOuterUnit", value)
}

// NewRackRequestOuterUnitFromValue returns a pointer to a valid RackRequestOuterUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRackRequestOuterUnitFromValue(v string) (*RackRequestOuterUnit, error) {
	ev := RackRequestOuterUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RackRequestOuterUnit: valid values are %v", v, AllowedRackRequestOuterUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RackRequestOuterUnit) IsValid() bool {
	for _, existing := range AllowedRackRequestOuterUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RackRequest_outer_unit value
func (v RackRequestOuterUnit) Ptr() *RackRequestOuterUnit {
	return &v
}

type NullableRackRequestOuterUnit struct {
	value *RackRequestOuterUnit
	isSet bool
}

func (v NullableRackRequestOuterUnit) Get() *RackRequestOuterUnit {
	return v.value
}

func (v *NullableRackRequestOuterUnit) Set(val *RackRequestOuterUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableRackRequestOuterUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableRackRequestOuterUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackRequestOuterUnit(val *RackRequestOuterUnit) *NullableRackRequestOuterUnit {
	return &NullableRackRequestOuterUnit{value: val, isSet: true}
}

func (v NullableRackRequestOuterUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackRequestOuterUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

