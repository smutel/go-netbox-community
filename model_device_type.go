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
	"time"
)

// checks if the DeviceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceType{}

// DeviceType Adds support for custom fields and tags.
type DeviceType struct {
	Id              int32                 `json:"id"`
	Url             string                `json:"url"`
	Display         string                `json:"display"`
	Manufacturer    BriefManufacturer     `json:"manufacturer"`
	DefaultPlatform NullableBriefPlatform `json:"default_platform,omitempty"`
	Model           string                `json:"model"`
	Slug            string                `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	// Discrete part number (optional)
	PartNumber *string  `json:"part_number,omitempty"`
	UHeight    *float64 `json:"u_height,omitempty"`
	// Devices of this type are excluded when calculating rack utilization.
	ExcludeFromUtilization *bool `json:"exclude_from_utilization,omitempty"`
	// Device consumes both front and rear rack faces.
	IsFullDepth                    *bool                           `json:"is_full_depth,omitempty"`
	SubdeviceRole                  NullableDeviceTypeSubdeviceRole `json:"subdevice_role,omitempty"`
	Airflow                        NullableDeviceTypeAirflow       `json:"airflow,omitempty"`
	Weight                         NullableFloat64                 `json:"weight,omitempty"`
	WeightUnit                     NullableDeviceTypeWeightUnit    `json:"weight_unit,omitempty"`
	FrontImage                     NullableString                  `json:"front_image,omitempty"`
	RearImage                      NullableString                  `json:"rear_image,omitempty"`
	Description                    *string                         `json:"description,omitempty"`
	Comments                       *string                         `json:"comments,omitempty"`
	Tags                           []NestedTag                     `json:"tags,omitempty"`
	CustomFields                   map[string]interface{}          `json:"custom_fields,omitempty"`
	Created                        NullableTime                    `json:"created"`
	LastUpdated                    NullableTime                    `json:"last_updated"`
	DeviceCount                    int64                           `json:"device_count"`
	ConsolePortTemplateCount       int32                           `json:"console_port_template_count"`
	ConsoleServerPortTemplateCount int32                           `json:"console_server_port_template_count"`
	PowerPortTemplateCount         int32                           `json:"power_port_template_count"`
	PowerOutletTemplateCount       int32                           `json:"power_outlet_template_count"`
	InterfaceTemplateCount         int32                           `json:"interface_template_count"`
	FrontPortTemplateCount         int32                           `json:"front_port_template_count"`
	RearPortTemplateCount          int32                           `json:"rear_port_template_count"`
	DeviceBayTemplateCount         int32                           `json:"device_bay_template_count"`
	ModuleBayTemplateCount         int32                           `json:"module_bay_template_count"`
	InventoryItemTemplateCount     int32                           `json:"inventory_item_template_count"`
	AdditionalProperties           map[string]interface{}
}

type _DeviceType DeviceType

// NewDeviceType instantiates a new DeviceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceType(id int32, url string, display string, manufacturer BriefManufacturer, model string, slug string, created NullableTime, lastUpdated NullableTime, deviceCount int64, consolePortTemplateCount int32, consoleServerPortTemplateCount int32, powerPortTemplateCount int32, powerOutletTemplateCount int32, interfaceTemplateCount int32, frontPortTemplateCount int32, rearPortTemplateCount int32, deviceBayTemplateCount int32, moduleBayTemplateCount int32, inventoryItemTemplateCount int32) *DeviceType {
	this := DeviceType{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Manufacturer = manufacturer
	this.Model = model
	this.Slug = slug
	var uHeight float64 = 1.0
	this.UHeight = &uHeight
	this.Created = created
	this.LastUpdated = lastUpdated
	this.DeviceCount = deviceCount
	this.ConsolePortTemplateCount = consolePortTemplateCount
	this.ConsoleServerPortTemplateCount = consoleServerPortTemplateCount
	this.PowerPortTemplateCount = powerPortTemplateCount
	this.PowerOutletTemplateCount = powerOutletTemplateCount
	this.InterfaceTemplateCount = interfaceTemplateCount
	this.FrontPortTemplateCount = frontPortTemplateCount
	this.RearPortTemplateCount = rearPortTemplateCount
	this.DeviceBayTemplateCount = deviceBayTemplateCount
	this.ModuleBayTemplateCount = moduleBayTemplateCount
	this.InventoryItemTemplateCount = inventoryItemTemplateCount
	return &this
}

// NewDeviceTypeWithDefaults instantiates a new DeviceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceTypeWithDefaults() *DeviceType {
	this := DeviceType{}
	var uHeight float64 = 1.0
	this.UHeight = &uHeight
	return &this
}

// GetId returns the Id field value
func (o *DeviceType) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeviceType) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *DeviceType) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DeviceType) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *DeviceType) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *DeviceType) SetDisplay(v string) {
	o.Display = v
}

// GetManufacturer returns the Manufacturer field value
func (o *DeviceType) GetManufacturer() BriefManufacturer {
	if o == nil {
		var ret BriefManufacturer
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetManufacturerOk() (*BriefManufacturer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *DeviceType) SetManufacturer(v BriefManufacturer) {
	o.Manufacturer = v
}

// GetDefaultPlatform returns the DefaultPlatform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceType) GetDefaultPlatform() BriefPlatform {
	if o == nil || IsNil(o.DefaultPlatform.Get()) {
		var ret BriefPlatform
		return ret
	}
	return *o.DefaultPlatform.Get()
}

// GetDefaultPlatformOk returns a tuple with the DefaultPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceType) GetDefaultPlatformOk() (*BriefPlatform, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultPlatform.Get(), o.DefaultPlatform.IsSet()
}

// HasDefaultPlatform returns a boolean if a field has been set.
func (o *DeviceType) HasDefaultPlatform() bool {
	if o != nil && o.DefaultPlatform.IsSet() {
		return true
	}

	return false
}

// SetDefaultPlatform gets a reference to the given NullableBriefPlatform and assigns it to the DefaultPlatform field.
func (o *DeviceType) SetDefaultPlatform(v BriefPlatform) {
	o.DefaultPlatform.Set(&v)
}

// SetDefaultPlatformNil sets the value for DefaultPlatform to be an explicit nil
func (o *DeviceType) SetDefaultPlatformNil() {
	o.DefaultPlatform.Set(nil)
}

// UnsetDefaultPlatform ensures that no value is present for DefaultPlatform, not even an explicit nil
func (o *DeviceType) UnsetDefaultPlatform() {
	o.DefaultPlatform.Unset()
}

// GetModel returns the Model field value
func (o *DeviceType) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *DeviceType) SetModel(v string) {
	o.Model = v
}

// GetSlug returns the Slug field value
func (o *DeviceType) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *DeviceType) SetSlug(v string) {
	o.Slug = v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *DeviceType) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *DeviceType) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *DeviceType) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetUHeight returns the UHeight field value if set, zero value otherwise.
func (o *DeviceType) GetUHeight() float64 {
	if o == nil || IsNil(o.UHeight) {
		var ret float64
		return ret
	}
	return *o.UHeight
}

// GetUHeightOk returns a tuple with the UHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetUHeightOk() (*float64, bool) {
	if o == nil || IsNil(o.UHeight) {
		return nil, false
	}
	return o.UHeight, true
}

// HasUHeight returns a boolean if a field has been set.
func (o *DeviceType) HasUHeight() bool {
	if o != nil && !IsNil(o.UHeight) {
		return true
	}

	return false
}

// SetUHeight gets a reference to the given float64 and assigns it to the UHeight field.
func (o *DeviceType) SetUHeight(v float64) {
	o.UHeight = &v
}

// GetExcludeFromUtilization returns the ExcludeFromUtilization field value if set, zero value otherwise.
func (o *DeviceType) GetExcludeFromUtilization() bool {
	if o == nil || IsNil(o.ExcludeFromUtilization) {
		var ret bool
		return ret
	}
	return *o.ExcludeFromUtilization
}

// GetExcludeFromUtilizationOk returns a tuple with the ExcludeFromUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetExcludeFromUtilizationOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeFromUtilization) {
		return nil, false
	}
	return o.ExcludeFromUtilization, true
}

// HasExcludeFromUtilization returns a boolean if a field has been set.
func (o *DeviceType) HasExcludeFromUtilization() bool {
	if o != nil && !IsNil(o.ExcludeFromUtilization) {
		return true
	}

	return false
}

// SetExcludeFromUtilization gets a reference to the given bool and assigns it to the ExcludeFromUtilization field.
func (o *DeviceType) SetExcludeFromUtilization(v bool) {
	o.ExcludeFromUtilization = &v
}

// GetIsFullDepth returns the IsFullDepth field value if set, zero value otherwise.
func (o *DeviceType) GetIsFullDepth() bool {
	if o == nil || IsNil(o.IsFullDepth) {
		var ret bool
		return ret
	}
	return *o.IsFullDepth
}

// GetIsFullDepthOk returns a tuple with the IsFullDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetIsFullDepthOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFullDepth) {
		return nil, false
	}
	return o.IsFullDepth, true
}

// HasIsFullDepth returns a boolean if a field has been set.
func (o *DeviceType) HasIsFullDepth() bool {
	if o != nil && !IsNil(o.IsFullDepth) {
		return true
	}

	return false
}

// SetIsFullDepth gets a reference to the given bool and assigns it to the IsFullDepth field.
func (o *DeviceType) SetIsFullDepth(v bool) {
	o.IsFullDepth = &v
}

// GetSubdeviceRole returns the SubdeviceRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceType) GetSubdeviceRole() DeviceTypeSubdeviceRole {
	if o == nil || IsNil(o.SubdeviceRole.Get()) {
		var ret DeviceTypeSubdeviceRole
		return ret
	}
	return *o.SubdeviceRole.Get()
}

// GetSubdeviceRoleOk returns a tuple with the SubdeviceRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceType) GetSubdeviceRoleOk() (*DeviceTypeSubdeviceRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubdeviceRole.Get(), o.SubdeviceRole.IsSet()
}

// HasSubdeviceRole returns a boolean if a field has been set.
func (o *DeviceType) HasSubdeviceRole() bool {
	if o != nil && o.SubdeviceRole.IsSet() {
		return true
	}

	return false
}

// SetSubdeviceRole gets a reference to the given NullableDeviceTypeSubdeviceRole and assigns it to the SubdeviceRole field.
func (o *DeviceType) SetSubdeviceRole(v DeviceTypeSubdeviceRole) {
	o.SubdeviceRole.Set(&v)
}

// SetSubdeviceRoleNil sets the value for SubdeviceRole to be an explicit nil
func (o *DeviceType) SetSubdeviceRoleNil() {
	o.SubdeviceRole.Set(nil)
}

// UnsetSubdeviceRole ensures that no value is present for SubdeviceRole, not even an explicit nil
func (o *DeviceType) UnsetSubdeviceRole() {
	o.SubdeviceRole.Unset()
}

// GetAirflow returns the Airflow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceType) GetAirflow() DeviceTypeAirflow {
	if o == nil || IsNil(o.Airflow.Get()) {
		var ret DeviceTypeAirflow
		return ret
	}
	return *o.Airflow.Get()
}

// GetAirflowOk returns a tuple with the Airflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceType) GetAirflowOk() (*DeviceTypeAirflow, bool) {
	if o == nil {
		return nil, false
	}
	return o.Airflow.Get(), o.Airflow.IsSet()
}

// HasAirflow returns a boolean if a field has been set.
func (o *DeviceType) HasAirflow() bool {
	if o != nil && o.Airflow.IsSet() {
		return true
	}

	return false
}

// SetAirflow gets a reference to the given NullableDeviceTypeAirflow and assigns it to the Airflow field.
func (o *DeviceType) SetAirflow(v DeviceTypeAirflow) {
	o.Airflow.Set(&v)
}

// SetAirflowNil sets the value for Airflow to be an explicit nil
func (o *DeviceType) SetAirflowNil() {
	o.Airflow.Set(nil)
}

// UnsetAirflow ensures that no value is present for Airflow, not even an explicit nil
func (o *DeviceType) UnsetAirflow() {
	o.Airflow.Unset()
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceType) GetWeight() float64 {
	if o == nil || IsNil(o.Weight.Get()) {
		var ret float64
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceType) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *DeviceType) HasWeight() bool {
	if o != nil && o.Weight.IsSet() {
		return true
	}

	return false
}

// SetWeight gets a reference to the given NullableFloat64 and assigns it to the Weight field.
func (o *DeviceType) SetWeight(v float64) {
	o.Weight.Set(&v)
}

// SetWeightNil sets the value for Weight to be an explicit nil
func (o *DeviceType) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil
func (o *DeviceType) UnsetWeight() {
	o.Weight.Unset()
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceType) GetWeightUnit() DeviceTypeWeightUnit {
	if o == nil || IsNil(o.WeightUnit.Get()) {
		var ret DeviceTypeWeightUnit
		return ret
	}
	return *o.WeightUnit.Get()
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceType) GetWeightUnitOk() (*DeviceTypeWeightUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.WeightUnit.Get(), o.WeightUnit.IsSet()
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *DeviceType) HasWeightUnit() bool {
	if o != nil && o.WeightUnit.IsSet() {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given NullableDeviceTypeWeightUnit and assigns it to the WeightUnit field.
func (o *DeviceType) SetWeightUnit(v DeviceTypeWeightUnit) {
	o.WeightUnit.Set(&v)
}

// SetWeightUnitNil sets the value for WeightUnit to be an explicit nil
func (o *DeviceType) SetWeightUnitNil() {
	o.WeightUnit.Set(nil)
}

// UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
func (o *DeviceType) UnsetWeightUnit() {
	o.WeightUnit.Unset()
}

// GetFrontImage returns the FrontImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceType) GetFrontImage() string {
	if o == nil || IsNil(o.FrontImage.Get()) {
		var ret string
		return ret
	}
	return *o.FrontImage.Get()
}

// GetFrontImageOk returns a tuple with the FrontImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceType) GetFrontImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrontImage.Get(), o.FrontImage.IsSet()
}

// HasFrontImage returns a boolean if a field has been set.
func (o *DeviceType) HasFrontImage() bool {
	if o != nil && o.FrontImage.IsSet() {
		return true
	}

	return false
}

// SetFrontImage gets a reference to the given NullableString and assigns it to the FrontImage field.
func (o *DeviceType) SetFrontImage(v string) {
	o.FrontImage.Set(&v)
}

// SetFrontImageNil sets the value for FrontImage to be an explicit nil
func (o *DeviceType) SetFrontImageNil() {
	o.FrontImage.Set(nil)
}

// UnsetFrontImage ensures that no value is present for FrontImage, not even an explicit nil
func (o *DeviceType) UnsetFrontImage() {
	o.FrontImage.Unset()
}

// GetRearImage returns the RearImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceType) GetRearImage() string {
	if o == nil || IsNil(o.RearImage.Get()) {
		var ret string
		return ret
	}
	return *o.RearImage.Get()
}

// GetRearImageOk returns a tuple with the RearImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceType) GetRearImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RearImage.Get(), o.RearImage.IsSet()
}

// HasRearImage returns a boolean if a field has been set.
func (o *DeviceType) HasRearImage() bool {
	if o != nil && o.RearImage.IsSet() {
		return true
	}

	return false
}

// SetRearImage gets a reference to the given NullableString and assigns it to the RearImage field.
func (o *DeviceType) SetRearImage(v string) {
	o.RearImage.Set(&v)
}

// SetRearImageNil sets the value for RearImage to be an explicit nil
func (o *DeviceType) SetRearImageNil() {
	o.RearImage.Set(nil)
}

// UnsetRearImage ensures that no value is present for RearImage, not even an explicit nil
func (o *DeviceType) UnsetRearImage() {
	o.RearImage.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceType) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *DeviceType) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *DeviceType) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *DeviceType) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DeviceType) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeviceType) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *DeviceType) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *DeviceType) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *DeviceType) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *DeviceType) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *DeviceType) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceType) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *DeviceType) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *DeviceType) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceType) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *DeviceType) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetDeviceCount returns the DeviceCount field value
func (o *DeviceType) GetDeviceCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetDeviceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceCount, true
}

// SetDeviceCount sets field value
func (o *DeviceType) SetDeviceCount(v int64) {
	o.DeviceCount = v
}

// GetConsolePortTemplateCount returns the ConsolePortTemplateCount field value
func (o *DeviceType) GetConsolePortTemplateCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConsolePortTemplateCount
}

// GetConsolePortTemplateCountOk returns a tuple with the ConsolePortTemplateCount field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetConsolePortTemplateCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsolePortTemplateCount, true
}

// SetConsolePortTemplateCount sets field value
func (o *DeviceType) SetConsolePortTemplateCount(v int32) {
	o.ConsolePortTemplateCount = v
}

// GetConsoleServerPortTemplateCount returns the ConsoleServerPortTemplateCount field value
func (o *DeviceType) GetConsoleServerPortTemplateCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConsoleServerPortTemplateCount
}

// GetConsoleServerPortTemplateCountOk returns a tuple with the ConsoleServerPortTemplateCount field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetConsoleServerPortTemplateCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsoleServerPortTemplateCount, true
}

// SetConsoleServerPortTemplateCount sets field value
func (o *DeviceType) SetConsoleServerPortTemplateCount(v int32) {
	o.ConsoleServerPortTemplateCount = v
}

// GetPowerPortTemplateCount returns the PowerPortTemplateCount field value
func (o *DeviceType) GetPowerPortTemplateCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PowerPortTemplateCount
}

// GetPowerPortTemplateCountOk returns a tuple with the PowerPortTemplateCount field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetPowerPortTemplateCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerPortTemplateCount, true
}

// SetPowerPortTemplateCount sets field value
func (o *DeviceType) SetPowerPortTemplateCount(v int32) {
	o.PowerPortTemplateCount = v
}

// GetPowerOutletTemplateCount returns the PowerOutletTemplateCount field value
func (o *DeviceType) GetPowerOutletTemplateCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PowerOutletTemplateCount
}

// GetPowerOutletTemplateCountOk returns a tuple with the PowerOutletTemplateCount field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetPowerOutletTemplateCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerOutletTemplateCount, true
}

// SetPowerOutletTemplateCount sets field value
func (o *DeviceType) SetPowerOutletTemplateCount(v int32) {
	o.PowerOutletTemplateCount = v
}

// GetInterfaceTemplateCount returns the InterfaceTemplateCount field value
func (o *DeviceType) GetInterfaceTemplateCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InterfaceTemplateCount
}

// GetInterfaceTemplateCountOk returns a tuple with the InterfaceTemplateCount field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetInterfaceTemplateCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterfaceTemplateCount, true
}

// SetInterfaceTemplateCount sets field value
func (o *DeviceType) SetInterfaceTemplateCount(v int32) {
	o.InterfaceTemplateCount = v
}

// GetFrontPortTemplateCount returns the FrontPortTemplateCount field value
func (o *DeviceType) GetFrontPortTemplateCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FrontPortTemplateCount
}

// GetFrontPortTemplateCountOk returns a tuple with the FrontPortTemplateCount field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetFrontPortTemplateCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FrontPortTemplateCount, true
}

// SetFrontPortTemplateCount sets field value
func (o *DeviceType) SetFrontPortTemplateCount(v int32) {
	o.FrontPortTemplateCount = v
}

// GetRearPortTemplateCount returns the RearPortTemplateCount field value
func (o *DeviceType) GetRearPortTemplateCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RearPortTemplateCount
}

// GetRearPortTemplateCountOk returns a tuple with the RearPortTemplateCount field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetRearPortTemplateCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RearPortTemplateCount, true
}

// SetRearPortTemplateCount sets field value
func (o *DeviceType) SetRearPortTemplateCount(v int32) {
	o.RearPortTemplateCount = v
}

// GetDeviceBayTemplateCount returns the DeviceBayTemplateCount field value
func (o *DeviceType) GetDeviceBayTemplateCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DeviceBayTemplateCount
}

// GetDeviceBayTemplateCountOk returns a tuple with the DeviceBayTemplateCount field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetDeviceBayTemplateCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceBayTemplateCount, true
}

// SetDeviceBayTemplateCount sets field value
func (o *DeviceType) SetDeviceBayTemplateCount(v int32) {
	o.DeviceBayTemplateCount = v
}

// GetModuleBayTemplateCount returns the ModuleBayTemplateCount field value
func (o *DeviceType) GetModuleBayTemplateCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ModuleBayTemplateCount
}

// GetModuleBayTemplateCountOk returns a tuple with the ModuleBayTemplateCount field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetModuleBayTemplateCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleBayTemplateCount, true
}

// SetModuleBayTemplateCount sets field value
func (o *DeviceType) SetModuleBayTemplateCount(v int32) {
	o.ModuleBayTemplateCount = v
}

// GetInventoryItemTemplateCount returns the InventoryItemTemplateCount field value
func (o *DeviceType) GetInventoryItemTemplateCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InventoryItemTemplateCount
}

// GetInventoryItemTemplateCountOk returns a tuple with the InventoryItemTemplateCount field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetInventoryItemTemplateCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InventoryItemTemplateCount, true
}

// SetInventoryItemTemplateCount sets field value
func (o *DeviceType) SetInventoryItemTemplateCount(v int32) {
	o.InventoryItemTemplateCount = v
}

func (o DeviceType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["manufacturer"] = o.Manufacturer
	if o.DefaultPlatform.IsSet() {
		toSerialize["default_platform"] = o.DefaultPlatform.Get()
	}
	toSerialize["model"] = o.Model
	toSerialize["slug"] = o.Slug
	if !IsNil(o.PartNumber) {
		toSerialize["part_number"] = o.PartNumber
	}
	if !IsNil(o.UHeight) {
		toSerialize["u_height"] = o.UHeight
	}
	if !IsNil(o.ExcludeFromUtilization) {
		toSerialize["exclude_from_utilization"] = o.ExcludeFromUtilization
	}
	if !IsNil(o.IsFullDepth) {
		toSerialize["is_full_depth"] = o.IsFullDepth
	}
	if o.SubdeviceRole.IsSet() {
		toSerialize["subdevice_role"] = o.SubdeviceRole.Get()
	}
	if o.Airflow.IsSet() {
		toSerialize["airflow"] = o.Airflow.Get()
	}
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
	}
	if o.WeightUnit.IsSet() {
		toSerialize["weight_unit"] = o.WeightUnit.Get()
	}
	if o.FrontImage.IsSet() {
		toSerialize["front_image"] = o.FrontImage.Get()
	}
	if o.RearImage.IsSet() {
		toSerialize["rear_image"] = o.RearImage.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["device_count"] = o.DeviceCount
	toSerialize["console_port_template_count"] = o.ConsolePortTemplateCount
	toSerialize["console_server_port_template_count"] = o.ConsoleServerPortTemplateCount
	toSerialize["power_port_template_count"] = o.PowerPortTemplateCount
	toSerialize["power_outlet_template_count"] = o.PowerOutletTemplateCount
	toSerialize["interface_template_count"] = o.InterfaceTemplateCount
	toSerialize["front_port_template_count"] = o.FrontPortTemplateCount
	toSerialize["rear_port_template_count"] = o.RearPortTemplateCount
	toSerialize["device_bay_template_count"] = o.DeviceBayTemplateCount
	toSerialize["module_bay_template_count"] = o.ModuleBayTemplateCount
	toSerialize["inventory_item_template_count"] = o.InventoryItemTemplateCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"manufacturer",
		"model",
		"slug",
		"created",
		"last_updated",
		"device_count",
		"console_port_template_count",
		"console_server_port_template_count",
		"power_port_template_count",
		"power_outlet_template_count",
		"interface_template_count",
		"front_port_template_count",
		"rear_port_template_count",
		"device_bay_template_count",
		"module_bay_template_count",
		"inventory_item_template_count",
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

	varDeviceType := _DeviceType{}

	err = json.Unmarshal(data, &varDeviceType)

	if err != nil {
		return err
	}

	*o = DeviceType(varDeviceType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "default_platform")
		delete(additionalProperties, "model")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "part_number")
		delete(additionalProperties, "u_height")
		delete(additionalProperties, "exclude_from_utilization")
		delete(additionalProperties, "is_full_depth")
		delete(additionalProperties, "subdevice_role")
		delete(additionalProperties, "airflow")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "weight_unit")
		delete(additionalProperties, "front_image")
		delete(additionalProperties, "rear_image")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "device_count")
		delete(additionalProperties, "console_port_template_count")
		delete(additionalProperties, "console_server_port_template_count")
		delete(additionalProperties, "power_port_template_count")
		delete(additionalProperties, "power_outlet_template_count")
		delete(additionalProperties, "interface_template_count")
		delete(additionalProperties, "front_port_template_count")
		delete(additionalProperties, "rear_port_template_count")
		delete(additionalProperties, "device_bay_template_count")
		delete(additionalProperties, "module_bay_template_count")
		delete(additionalProperties, "inventory_item_template_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceType struct {
	value *DeviceType
	isSet bool
}

func (v NullableDeviceType) Get() *DeviceType {
	return v.value
}

func (v *NullableDeviceType) Set(val *DeviceType) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceType) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceType(val *DeviceType) *NullableDeviceType {
	return &NullableDeviceType{value: val, isSet: true}
}

func (v NullableDeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
