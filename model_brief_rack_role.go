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

// checks if the BriefRackRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefRackRole{}

// BriefRackRole Adds support for custom fields and tags.
type BriefRackRole struct {
	Id                   int32   `json:"id"`
	Url                  string  `json:"url"`
	Display              string  `json:"display"`
	Name                 string  `json:"name"`
	Slug                 string  `json:"slug"`
	Description          *string `json:"description,omitempty"`
	RackCount            *int64  `json:"rack_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefRackRole BriefRackRole

// NewBriefRackRole instantiates a new BriefRackRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefRackRole(id int32, url string, display string, name string, slug string) *BriefRackRole {
	this := BriefRackRole{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	return &this
}

// NewBriefRackRoleWithDefaults instantiates a new BriefRackRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefRackRoleWithDefaults() *BriefRackRole {
	this := BriefRackRole{}
	return &this
}

// GetId returns the Id field value
func (o *BriefRackRole) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefRackRole) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefRackRole) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefRackRole) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefRackRole) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefRackRole) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefRackRole) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefRackRole) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefRackRole) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *BriefRackRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefRackRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefRackRole) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *BriefRackRole) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BriefRackRole) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BriefRackRole) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefRackRole) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefRackRole) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefRackRole) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefRackRole) SetDescription(v string) {
	o.Description = &v
}

// GetRackCount returns the RackCount field value if set, zero value otherwise.
func (o *BriefRackRole) GetRackCount() int64 {
	if o == nil || IsNil(o.RackCount) {
		var ret int64
		return ret
	}
	return *o.RackCount
}

// GetRackCountOk returns a tuple with the RackCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefRackRole) GetRackCountOk() (*int64, bool) {
	if o == nil || IsNil(o.RackCount) {
		return nil, false
	}
	return o.RackCount, true
}

// HasRackCount returns a boolean if a field has been set.
func (o *BriefRackRole) HasRackCount() bool {
	if o != nil && !IsNil(o.RackCount) {
		return true
	}

	return false
}

// SetRackCount gets a reference to the given int64 and assigns it to the RackCount field.
func (o *BriefRackRole) SetRackCount(v int64) {
	o.RackCount = &v
}

func (o BriefRackRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefRackRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.RackCount) {
		toSerialize["rack_count"] = o.RackCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefRackRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
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

	varBriefRackRole := _BriefRackRole{}

	err = json.Unmarshal(data, &varBriefRackRole)

	if err != nil {
		return err
	}

	*o = BriefRackRole(varBriefRackRole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		delete(additionalProperties, "rack_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefRackRole struct {
	value *BriefRackRole
	isSet bool
}

func (v NullableBriefRackRole) Get() *BriefRackRole {
	return v.value
}

func (v *NullableBriefRackRole) Set(val *BriefRackRole) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefRackRole) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefRackRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefRackRole(val *BriefRackRole) *NullableBriefRackRole {
	return &NullableBriefRackRole{value: val, isSet: true}
}

func (v NullableBriefRackRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefRackRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
