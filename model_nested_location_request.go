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

// checks if the NestedLocationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedLocationRequest{}

// NestedLocationRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedLocationRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	AdditionalProperties map[string]interface{}
}

type _NestedLocationRequest NestedLocationRequest

// NewNestedLocationRequest instantiates a new NestedLocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedLocationRequest(name string, slug string) *NestedLocationRequest {
	this := NestedLocationRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedLocationRequestWithDefaults instantiates a new NestedLocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedLocationRequestWithDefaults() *NestedLocationRequest {
	this := NestedLocationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedLocationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedLocationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedLocationRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedLocationRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedLocationRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedLocationRequest) SetSlug(v string) {
	o.Slug = v
}

func (o NestedLocationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedLocationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedLocationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slug",
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

	varNestedLocationRequest := _NestedLocationRequest{}

	err = json.Unmarshal(data, &varNestedLocationRequest)

	if err != nil {
		return err
	}

	*o = NestedLocationRequest(varNestedLocationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedLocationRequest struct {
	value *NestedLocationRequest
	isSet bool
}

func (v NullableNestedLocationRequest) Get() *NestedLocationRequest {
	return v.value
}

func (v *NullableNestedLocationRequest) Set(val *NestedLocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedLocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedLocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedLocationRequest(val *NestedLocationRequest) *NullableNestedLocationRequest {
	return &NullableNestedLocationRequest{value: val, isSet: true}
}

func (v NullableNestedLocationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedLocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


