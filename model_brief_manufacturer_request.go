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

// checks if the BriefManufacturerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefManufacturerRequest{}

// BriefManufacturerRequest Adds support for custom fields and tags.
type BriefManufacturerRequest struct {
	Name                 string  `json:"name"`
	Slug                 string  `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefManufacturerRequest BriefManufacturerRequest

// NewBriefManufacturerRequest instantiates a new BriefManufacturerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefManufacturerRequest(name string, slug string) *BriefManufacturerRequest {
	this := BriefManufacturerRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewBriefManufacturerRequestWithDefaults instantiates a new BriefManufacturerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefManufacturerRequestWithDefaults() *BriefManufacturerRequest {
	this := BriefManufacturerRequest{}
	return &this
}

// GetName returns the Name field value
func (o *BriefManufacturerRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefManufacturerRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefManufacturerRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *BriefManufacturerRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BriefManufacturerRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BriefManufacturerRequest) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefManufacturerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefManufacturerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefManufacturerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefManufacturerRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefManufacturerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefManufacturerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefManufacturerRequest) UnmarshalJSON(data []byte) (err error) {
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
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBriefManufacturerRequest := _BriefManufacturerRequest{}

	err = json.Unmarshal(data, &varBriefManufacturerRequest)

	if err != nil {
		return err
	}

	*o = BriefManufacturerRequest(varBriefManufacturerRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefManufacturerRequest struct {
	value *BriefManufacturerRequest
	isSet bool
}

func (v NullableBriefManufacturerRequest) Get() *BriefManufacturerRequest {
	return v.value
}

func (v *NullableBriefManufacturerRequest) Set(val *BriefManufacturerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefManufacturerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefManufacturerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefManufacturerRequest(val *BriefManufacturerRequest) *NullableBriefManufacturerRequest {
	return &NullableBriefManufacturerRequest{value: val, isSet: true}
}

func (v NullableBriefManufacturerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefManufacturerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
