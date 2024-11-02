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

// checks if the BriefPowerPanelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefPowerPanelRequest{}

// BriefPowerPanelRequest Adds support for custom fields and tags.
type BriefPowerPanelRequest struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefPowerPanelRequest BriefPowerPanelRequest

// NewBriefPowerPanelRequest instantiates a new BriefPowerPanelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefPowerPanelRequest(name string) *BriefPowerPanelRequest {
	this := BriefPowerPanelRequest{}
	this.Name = name
	return &this
}

// NewBriefPowerPanelRequestWithDefaults instantiates a new BriefPowerPanelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefPowerPanelRequestWithDefaults() *BriefPowerPanelRequest {
	this := BriefPowerPanelRequest{}
	return &this
}

// GetName returns the Name field value
func (o *BriefPowerPanelRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefPowerPanelRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefPowerPanelRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefPowerPanelRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefPowerPanelRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefPowerPanelRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefPowerPanelRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefPowerPanelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefPowerPanelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefPowerPanelRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBriefPowerPanelRequest := _BriefPowerPanelRequest{}

	err = json.Unmarshal(data, &varBriefPowerPanelRequest)

	if err != nil {
		return err
	}

	*o = BriefPowerPanelRequest(varBriefPowerPanelRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefPowerPanelRequest struct {
	value *BriefPowerPanelRequest
	isSet bool
}

func (v NullableBriefPowerPanelRequest) Get() *BriefPowerPanelRequest {
	return v.value
}

func (v *NullableBriefPowerPanelRequest) Set(val *BriefPowerPanelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefPowerPanelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefPowerPanelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefPowerPanelRequest(val *BriefPowerPanelRequest) *NullableBriefPowerPanelRequest {
	return &NullableBriefPowerPanelRequest{value: val, isSet: true}
}

func (v NullableBriefPowerPanelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefPowerPanelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


