/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the BriefCableRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefCableRequest{}

// BriefCableRequest Adds support for custom fields and tags.
type BriefCableRequest struct {
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefCableRequest BriefCableRequest

// NewBriefCableRequest instantiates a new BriefCableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefCableRequest() *BriefCableRequest {
	this := BriefCableRequest{}
	return &this
}

// NewBriefCableRequestWithDefaults instantiates a new BriefCableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefCableRequestWithDefaults() *BriefCableRequest {
	this := BriefCableRequest{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BriefCableRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefCableRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BriefCableRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BriefCableRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefCableRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefCableRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefCableRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefCableRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefCableRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefCableRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefCableRequest) UnmarshalJSON(data []byte) (err error) {
	varBriefCableRequest := _BriefCableRequest{}

	err = json.Unmarshal(data, &varBriefCableRequest)

	if err != nil {
		return err
	}

	*o = BriefCableRequest(varBriefCableRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefCableRequest struct {
	value *BriefCableRequest
	isSet bool
}

func (v NullableBriefCableRequest) Get() *BriefCableRequest {
	return v.value
}

func (v *NullableBriefCableRequest) Set(val *BriefCableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefCableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefCableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefCableRequest(val *BriefCableRequest) *NullableBriefCableRequest {
	return &NullableBriefCableRequest{value: val, isSet: true}
}

func (v NullableBriefCableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefCableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


