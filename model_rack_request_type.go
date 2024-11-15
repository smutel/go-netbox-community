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

// RackRequestType * `2-post-frame` - 2-post frame * `4-post-frame` - 4-post frame * `4-post-cabinet` - 4-post cabinet * `wall-frame` - Wall-mounted frame * `wall-frame-vertical` - Wall-mounted frame (vertical) * `wall-cabinet` - Wall-mounted cabinet * `wall-cabinet-vertical` - Wall-mounted cabinet (vertical)
type RackRequestType string

// List of RackRequest_type
const (
	RACKREQUESTTYPE__2_POST_FRAME         RackRequestType = "2-post-frame"
	RACKREQUESTTYPE__4_POST_FRAME         RackRequestType = "4-post-frame"
	RACKREQUESTTYPE__4_POST_CABINET       RackRequestType = "4-post-cabinet"
	RACKREQUESTTYPE_WALL_FRAME            RackRequestType = "wall-frame"
	RACKREQUESTTYPE_WALL_FRAME_VERTICAL   RackRequestType = "wall-frame-vertical"
	RACKREQUESTTYPE_WALL_CABINET          RackRequestType = "wall-cabinet"
	RACKREQUESTTYPE_WALL_CABINET_VERTICAL RackRequestType = "wall-cabinet-vertical"
	RACKREQUESTTYPE_EMPTY                 RackRequestType = ""
)

// All allowed values of RackRequestType enum
var AllowedRackRequestTypeEnumValues = []RackRequestType{
	"2-post-frame",
	"4-post-frame",
	"4-post-cabinet",
	"wall-frame",
	"wall-frame-vertical",
	"wall-cabinet",
	"wall-cabinet-vertical",
	"",
}

func (v *RackRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RackRequestType(value)
	for _, existing := range AllowedRackRequestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RackRequestType", value)
}

// NewRackRequestTypeFromValue returns a pointer to a valid RackRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRackRequestTypeFromValue(v string) (*RackRequestType, error) {
	ev := RackRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RackRequestType: valid values are %v", v, AllowedRackRequestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RackRequestType) IsValid() bool {
	for _, existing := range AllowedRackRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RackRequest_type value
func (v RackRequestType) Ptr() *RackRequestType {
	return &v
}

type NullableRackRequestType struct {
	value *RackRequestType
	isSet bool
}

func (v NullableRackRequestType) Get() *RackRequestType {
	return v.value
}

func (v *NullableRackRequestType) Set(val *RackRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullableRackRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullableRackRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackRequestType(val *RackRequestType) *NullableRackRequestType {
	return &NullableRackRequestType{value: val, isSet: true}
}

func (v NullableRackRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
