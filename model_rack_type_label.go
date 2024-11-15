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

// RackTypeLabel the model 'RackTypeLabel'
type RackTypeLabel string

// List of Rack_type_label
const (
	RACKTYPELABEL__2_POST_FRAME                  RackTypeLabel = "2-post frame"
	RACKTYPELABEL__4_POST_FRAME                  RackTypeLabel = "4-post frame"
	RACKTYPELABEL__4_POST_CABINET                RackTypeLabel = "4-post cabinet"
	RACKTYPELABEL_WALL_MOUNTED_FRAME             RackTypeLabel = "Wall-mounted frame"
	RACKTYPELABEL_WALL_MOUNTED_FRAME__VERTICAL   RackTypeLabel = "Wall-mounted frame (vertical)"
	RACKTYPELABEL_WALL_MOUNTED_CABINET           RackTypeLabel = "Wall-mounted cabinet"
	RACKTYPELABEL_WALL_MOUNTED_CABINET__VERTICAL RackTypeLabel = "Wall-mounted cabinet (vertical)"
)

// All allowed values of RackTypeLabel enum
var AllowedRackTypeLabelEnumValues = []RackTypeLabel{
	"2-post frame",
	"4-post frame",
	"4-post cabinet",
	"Wall-mounted frame",
	"Wall-mounted frame (vertical)",
	"Wall-mounted cabinet",
	"Wall-mounted cabinet (vertical)",
}

func (v *RackTypeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RackTypeLabel(value)
	for _, existing := range AllowedRackTypeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RackTypeLabel", value)
}

// NewRackTypeLabelFromValue returns a pointer to a valid RackTypeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRackTypeLabelFromValue(v string) (*RackTypeLabel, error) {
	ev := RackTypeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RackTypeLabel: valid values are %v", v, AllowedRackTypeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RackTypeLabel) IsValid() bool {
	for _, existing := range AllowedRackTypeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Rack_type_label value
func (v RackTypeLabel) Ptr() *RackTypeLabel {
	return &v
}

type NullableRackTypeLabel struct {
	value *RackTypeLabel
	isSet bool
}

func (v NullableRackTypeLabel) Get() *RackTypeLabel {
	return v.value
}

func (v *NullableRackTypeLabel) Set(val *RackTypeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableRackTypeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableRackTypeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackTypeLabel(val *RackTypeLabel) *NullableRackTypeLabel {
	return &NullableRackTypeLabel{value: val, isSet: true}
}

func (v NullableRackTypeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackTypeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}