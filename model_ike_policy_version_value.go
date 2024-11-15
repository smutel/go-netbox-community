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

// IKEPolicyVersionValue * `1` - IKEv1 * `2` - IKEv2
type IKEPolicyVersionValue int32

// List of IKEPolicy_version_value
const (
	IKEPOLICYVERSIONVALUE__1 IKEPolicyVersionValue = 1
	IKEPOLICYVERSIONVALUE__2 IKEPolicyVersionValue = 2
)

// All allowed values of IKEPolicyVersionValue enum
var AllowedIKEPolicyVersionValueEnumValues = []IKEPolicyVersionValue{
	1,
	2,
}

func (v *IKEPolicyVersionValue) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IKEPolicyVersionValue(value)
	for _, existing := range AllowedIKEPolicyVersionValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IKEPolicyVersionValue", value)
}

// NewIKEPolicyVersionValueFromValue returns a pointer to a valid IKEPolicyVersionValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIKEPolicyVersionValueFromValue(v int32) (*IKEPolicyVersionValue, error) {
	ev := IKEPolicyVersionValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IKEPolicyVersionValue: valid values are %v", v, AllowedIKEPolicyVersionValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IKEPolicyVersionValue) IsValid() bool {
	for _, existing := range AllowedIKEPolicyVersionValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IKEPolicy_version_value value
func (v IKEPolicyVersionValue) Ptr() *IKEPolicyVersionValue {
	return &v
}

type NullableIKEPolicyVersionValue struct {
	value *IKEPolicyVersionValue
	isSet bool
}

func (v NullableIKEPolicyVersionValue) Get() *IKEPolicyVersionValue {
	return v.value
}

func (v *NullableIKEPolicyVersionValue) Set(val *IKEPolicyVersionValue) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEPolicyVersionValue) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEPolicyVersionValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEPolicyVersionValue(val *IKEPolicyVersionValue) *NullableIKEPolicyVersionValue {
	return &NullableIKEPolicyVersionValue{value: val, isSet: true}
}

func (v NullableIKEPolicyVersionValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEPolicyVersionValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
