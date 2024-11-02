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

// checks if the NestedProviderAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedProviderAccount{}

// NestedProviderAccount Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedProviderAccount struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	Name *string `json:"name,omitempty"`
	Account string `json:"account"`
	AdditionalProperties map[string]interface{}
}

type _NestedProviderAccount NestedProviderAccount

// NewNestedProviderAccount instantiates a new NestedProviderAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedProviderAccount(id int32, url string, display string, account string) *NestedProviderAccount {
	this := NestedProviderAccount{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Account = account
	return &this
}

// NewNestedProviderAccountWithDefaults instantiates a new NestedProviderAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedProviderAccountWithDefaults() *NestedProviderAccount {
	this := NestedProviderAccount{}
	return &this
}

// GetId returns the Id field value
func (o *NestedProviderAccount) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedProviderAccount) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedProviderAccount) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedProviderAccount) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedProviderAccount) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedProviderAccount) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedProviderAccount) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedProviderAccount) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedProviderAccount) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NestedProviderAccount) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedProviderAccount) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NestedProviderAccount) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NestedProviderAccount) SetName(v string) {
	o.Name = &v
}

// GetAccount returns the Account field value
func (o *NestedProviderAccount) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *NestedProviderAccount) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *NestedProviderAccount) SetAccount(v string) {
	o.Account = v
}

func (o NestedProviderAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedProviderAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["account"] = o.Account

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedProviderAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"account",
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

	varNestedProviderAccount := _NestedProviderAccount{}

	err = json.Unmarshal(data, &varNestedProviderAccount)

	if err != nil {
		return err
	}

	*o = NestedProviderAccount(varNestedProviderAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedProviderAccount struct {
	value *NestedProviderAccount
	isSet bool
}

func (v NullableNestedProviderAccount) Get() *NestedProviderAccount {
	return v.value
}

func (v *NullableNestedProviderAccount) Set(val *NestedProviderAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedProviderAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedProviderAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedProviderAccount(val *NestedProviderAccount) *NullableNestedProviderAccount {
	return &NullableNestedProviderAccount{value: val, isSet: true}
}

func (v NullableNestedProviderAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedProviderAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


