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

// checks if the PatchedBookmarkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBookmarkRequest{}

// PatchedBookmarkRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedBookmarkRequest struct {
	ObjectType           *string           `json:"object_type,omitempty"`
	ObjectId             *int64            `json:"object_id,omitempty"`
	User                 *BriefUserRequest `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBookmarkRequest PatchedBookmarkRequest

// NewPatchedBookmarkRequest instantiates a new PatchedBookmarkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBookmarkRequest() *PatchedBookmarkRequest {
	this := PatchedBookmarkRequest{}
	return &this
}

// NewPatchedBookmarkRequestWithDefaults instantiates a new PatchedBookmarkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBookmarkRequestWithDefaults() *PatchedBookmarkRequest {
	this := PatchedBookmarkRequest{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *PatchedBookmarkRequest) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBookmarkRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *PatchedBookmarkRequest) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *PatchedBookmarkRequest) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *PatchedBookmarkRequest) GetObjectId() int64 {
	if o == nil || IsNil(o.ObjectId) {
		var ret int64
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBookmarkRequest) GetObjectIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ObjectId) {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *PatchedBookmarkRequest) HasObjectId() bool {
	if o != nil && !IsNil(o.ObjectId) {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given int64 and assigns it to the ObjectId field.
func (o *PatchedBookmarkRequest) SetObjectId(v int64) {
	o.ObjectId = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *PatchedBookmarkRequest) GetUser() BriefUserRequest {
	if o == nil || IsNil(o.User) {
		var ret BriefUserRequest
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBookmarkRequest) GetUserOk() (*BriefUserRequest, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedBookmarkRequest) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given BriefUserRequest and assigns it to the User field.
func (o *PatchedBookmarkRequest) SetUser(v BriefUserRequest) {
	o.User = &v
}

func (o PatchedBookmarkRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBookmarkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectType) {
		toSerialize["object_type"] = o.ObjectType
	}
	if !IsNil(o.ObjectId) {
		toSerialize["object_id"] = o.ObjectId
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedBookmarkRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedBookmarkRequest := _PatchedBookmarkRequest{}

	err = json.Unmarshal(data, &varPatchedBookmarkRequest)

	if err != nil {
		return err
	}

	*o = PatchedBookmarkRequest(varPatchedBookmarkRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBookmarkRequest struct {
	value *PatchedBookmarkRequest
	isSet bool
}

func (v NullablePatchedBookmarkRequest) Get() *PatchedBookmarkRequest {
	return v.value
}

func (v *NullablePatchedBookmarkRequest) Set(val *PatchedBookmarkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBookmarkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBookmarkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBookmarkRequest(val *PatchedBookmarkRequest) *NullablePatchedBookmarkRequest {
	return &NullablePatchedBookmarkRequest{value: val, isSet: true}
}

func (v NullablePatchedBookmarkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBookmarkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
