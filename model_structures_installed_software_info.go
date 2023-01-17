/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// StructuresInstalledSoftwareInfo struct for StructuresInstalledSoftwareInfo
type StructuresInstalledSoftwareInfo struct {
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Version *string `json:"version,omitempty"`
	Address *string `json:"address,omitempty"`
	Status *string `json:"status,omitempty"`
	FieldValue []StructuresInstalledSoftwareInfoFieldValue `json:"field_value,omitempty"`
	Metadata *StructuresSoftwareMetadata `json:"metadata,omitempty"`
	Description *string `json:"description,omitempty"`
	DescriptionEn *string `json:"description_en,omitempty"`
	Category []StructuresSoftwareCategory `json:"category,omitempty"`
	Slug *string `json:"slug,omitempty"`
}

// NewStructuresInstalledSoftwareInfo instantiates a new StructuresInstalledSoftwareInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresInstalledSoftwareInfo() *StructuresInstalledSoftwareInfo {
	this := StructuresInstalledSoftwareInfo{}
	return &this
}

// NewStructuresInstalledSoftwareInfoWithDefaults instantiates a new StructuresInstalledSoftwareInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresInstalledSoftwareInfoWithDefaults() *StructuresInstalledSoftwareInfo {
	this := StructuresInstalledSoftwareInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StructuresInstalledSoftwareInfo) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresInstalledSoftwareInfo) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StructuresInstalledSoftwareInfo) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StructuresInstalledSoftwareInfo) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *StructuresInstalledSoftwareInfo) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresInstalledSoftwareInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *StructuresInstalledSoftwareInfo) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *StructuresInstalledSoftwareInfo) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *StructuresInstalledSoftwareInfo) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresInstalledSoftwareInfo) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *StructuresInstalledSoftwareInfo) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *StructuresInstalledSoftwareInfo) SetVersion(v string) {
	o.Version = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *StructuresInstalledSoftwareInfo) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresInstalledSoftwareInfo) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *StructuresInstalledSoftwareInfo) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *StructuresInstalledSoftwareInfo) SetAddress(v string) {
	o.Address = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StructuresInstalledSoftwareInfo) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresInstalledSoftwareInfo) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StructuresInstalledSoftwareInfo) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StructuresInstalledSoftwareInfo) SetStatus(v string) {
	o.Status = &v
}

// GetFieldValue returns the FieldValue field value if set, zero value otherwise.
func (o *StructuresInstalledSoftwareInfo) GetFieldValue() []StructuresInstalledSoftwareInfoFieldValue {
	if o == nil || isNil(o.FieldValue) {
		var ret []StructuresInstalledSoftwareInfoFieldValue
		return ret
	}
	return o.FieldValue
}

// GetFieldValueOk returns a tuple with the FieldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresInstalledSoftwareInfo) GetFieldValueOk() ([]StructuresInstalledSoftwareInfoFieldValue, bool) {
	if o == nil || isNil(o.FieldValue) {
    return nil, false
	}
	return o.FieldValue, true
}

// HasFieldValue returns a boolean if a field has been set.
func (o *StructuresInstalledSoftwareInfo) HasFieldValue() bool {
	if o != nil && !isNil(o.FieldValue) {
		return true
	}

	return false
}

// SetFieldValue gets a reference to the given []StructuresInstalledSoftwareInfoFieldValue and assigns it to the FieldValue field.
func (o *StructuresInstalledSoftwareInfo) SetFieldValue(v []StructuresInstalledSoftwareInfoFieldValue) {
	o.FieldValue = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *StructuresInstalledSoftwareInfo) GetMetadata() StructuresSoftwareMetadata {
	if o == nil || isNil(o.Metadata) {
		var ret StructuresSoftwareMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresInstalledSoftwareInfo) GetMetadataOk() (*StructuresSoftwareMetadata, bool) {
	if o == nil || isNil(o.Metadata) {
    return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *StructuresInstalledSoftwareInfo) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given StructuresSoftwareMetadata and assigns it to the Metadata field.
func (o *StructuresInstalledSoftwareInfo) SetMetadata(v StructuresSoftwareMetadata) {
	o.Metadata = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StructuresInstalledSoftwareInfo) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresInstalledSoftwareInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StructuresInstalledSoftwareInfo) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StructuresInstalledSoftwareInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDescriptionEn returns the DescriptionEn field value if set, zero value otherwise.
func (o *StructuresInstalledSoftwareInfo) GetDescriptionEn() string {
	if o == nil || isNil(o.DescriptionEn) {
		var ret string
		return ret
	}
	return *o.DescriptionEn
}

// GetDescriptionEnOk returns a tuple with the DescriptionEn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresInstalledSoftwareInfo) GetDescriptionEnOk() (*string, bool) {
	if o == nil || isNil(o.DescriptionEn) {
    return nil, false
	}
	return o.DescriptionEn, true
}

// HasDescriptionEn returns a boolean if a field has been set.
func (o *StructuresInstalledSoftwareInfo) HasDescriptionEn() bool {
	if o != nil && !isNil(o.DescriptionEn) {
		return true
	}

	return false
}

// SetDescriptionEn gets a reference to the given string and assigns it to the DescriptionEn field.
func (o *StructuresInstalledSoftwareInfo) SetDescriptionEn(v string) {
	o.DescriptionEn = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *StructuresInstalledSoftwareInfo) GetCategory() []StructuresSoftwareCategory {
	if o == nil || isNil(o.Category) {
		var ret []StructuresSoftwareCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresInstalledSoftwareInfo) GetCategoryOk() ([]StructuresSoftwareCategory, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *StructuresInstalledSoftwareInfo) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given []StructuresSoftwareCategory and assigns it to the Category field.
func (o *StructuresInstalledSoftwareInfo) SetCategory(v []StructuresSoftwareCategory) {
	o.Category = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *StructuresInstalledSoftwareInfo) GetSlug() string {
	if o == nil || isNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresInstalledSoftwareInfo) GetSlugOk() (*string, bool) {
	if o == nil || isNil(o.Slug) {
    return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *StructuresInstalledSoftwareInfo) HasSlug() bool {
	if o != nil && !isNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *StructuresInstalledSoftwareInfo) SetSlug(v string) {
	o.Slug = &v
}

func (o StructuresInstalledSoftwareInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.FieldValue) {
		toSerialize["field_value"] = o.FieldValue
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.DescriptionEn) {
		toSerialize["description_en"] = o.DescriptionEn
	}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresInstalledSoftwareInfo struct {
	value *StructuresInstalledSoftwareInfo
	isSet bool
}

func (v NullableStructuresInstalledSoftwareInfo) Get() *StructuresInstalledSoftwareInfo {
	return v.value
}

func (v *NullableStructuresInstalledSoftwareInfo) Set(val *StructuresInstalledSoftwareInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresInstalledSoftwareInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresInstalledSoftwareInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresInstalledSoftwareInfo(val *StructuresInstalledSoftwareInfo) *NullableStructuresInstalledSoftwareInfo {
	return &NullableStructuresInstalledSoftwareInfo{value: val, isSet: true}
}

func (v NullableStructuresInstalledSoftwareInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresInstalledSoftwareInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


