/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.6.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// StructuresSoftwareCategory struct for StructuresSoftwareCategory
type StructuresSoftwareCategory struct {
	SysName *string `json:"sys_name,omitempty"`
	Name *string `json:"name,omitempty"`
	NameEn *string `json:"name_en,omitempty"`
	IsMain *bool `json:"is_main,omitempty"`
	IconSvg *string `json:"icon_svg,omitempty"`
}

// NewStructuresSoftwareCategory instantiates a new StructuresSoftwareCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresSoftwareCategory() *StructuresSoftwareCategory {
	this := StructuresSoftwareCategory{}
	return &this
}

// NewStructuresSoftwareCategoryWithDefaults instantiates a new StructuresSoftwareCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresSoftwareCategoryWithDefaults() *StructuresSoftwareCategory {
	this := StructuresSoftwareCategory{}
	return &this
}

// GetSysName returns the SysName field value if set, zero value otherwise.
func (o *StructuresSoftwareCategory) GetSysName() string {
	if o == nil || isNil(o.SysName) {
		var ret string
		return ret
	}
	return *o.SysName
}

// GetSysNameOk returns a tuple with the SysName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSoftwareCategory) GetSysNameOk() (*string, bool) {
	if o == nil || isNil(o.SysName) {
    return nil, false
	}
	return o.SysName, true
}

// HasSysName returns a boolean if a field has been set.
func (o *StructuresSoftwareCategory) HasSysName() bool {
	if o != nil && !isNil(o.SysName) {
		return true
	}

	return false
}

// SetSysName gets a reference to the given string and assigns it to the SysName field.
func (o *StructuresSoftwareCategory) SetSysName(v string) {
	o.SysName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StructuresSoftwareCategory) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSoftwareCategory) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StructuresSoftwareCategory) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StructuresSoftwareCategory) SetName(v string) {
	o.Name = &v
}

// GetNameEn returns the NameEn field value if set, zero value otherwise.
func (o *StructuresSoftwareCategory) GetNameEn() string {
	if o == nil || isNil(o.NameEn) {
		var ret string
		return ret
	}
	return *o.NameEn
}

// GetNameEnOk returns a tuple with the NameEn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSoftwareCategory) GetNameEnOk() (*string, bool) {
	if o == nil || isNil(o.NameEn) {
    return nil, false
	}
	return o.NameEn, true
}

// HasNameEn returns a boolean if a field has been set.
func (o *StructuresSoftwareCategory) HasNameEn() bool {
	if o != nil && !isNil(o.NameEn) {
		return true
	}

	return false
}

// SetNameEn gets a reference to the given string and assigns it to the NameEn field.
func (o *StructuresSoftwareCategory) SetNameEn(v string) {
	o.NameEn = &v
}

// GetIsMain returns the IsMain field value if set, zero value otherwise.
func (o *StructuresSoftwareCategory) GetIsMain() bool {
	if o == nil || isNil(o.IsMain) {
		var ret bool
		return ret
	}
	return *o.IsMain
}

// GetIsMainOk returns a tuple with the IsMain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSoftwareCategory) GetIsMainOk() (*bool, bool) {
	if o == nil || isNil(o.IsMain) {
    return nil, false
	}
	return o.IsMain, true
}

// HasIsMain returns a boolean if a field has been set.
func (o *StructuresSoftwareCategory) HasIsMain() bool {
	if o != nil && !isNil(o.IsMain) {
		return true
	}

	return false
}

// SetIsMain gets a reference to the given bool and assigns it to the IsMain field.
func (o *StructuresSoftwareCategory) SetIsMain(v bool) {
	o.IsMain = &v
}

// GetIconSvg returns the IconSvg field value if set, zero value otherwise.
func (o *StructuresSoftwareCategory) GetIconSvg() string {
	if o == nil || isNil(o.IconSvg) {
		var ret string
		return ret
	}
	return *o.IconSvg
}

// GetIconSvgOk returns a tuple with the IconSvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSoftwareCategory) GetIconSvgOk() (*string, bool) {
	if o == nil || isNil(o.IconSvg) {
    return nil, false
	}
	return o.IconSvg, true
}

// HasIconSvg returns a boolean if a field has been set.
func (o *StructuresSoftwareCategory) HasIconSvg() bool {
	if o != nil && !isNil(o.IconSvg) {
		return true
	}

	return false
}

// SetIconSvg gets a reference to the given string and assigns it to the IconSvg field.
func (o *StructuresSoftwareCategory) SetIconSvg(v string) {
	o.IconSvg = &v
}

func (o StructuresSoftwareCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SysName) {
		toSerialize["sys_name"] = o.SysName
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.NameEn) {
		toSerialize["name_en"] = o.NameEn
	}
	if !isNil(o.IsMain) {
		toSerialize["is_main"] = o.IsMain
	}
	if !isNil(o.IconSvg) {
		toSerialize["icon_svg"] = o.IconSvg
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresSoftwareCategory struct {
	value *StructuresSoftwareCategory
	isSet bool
}

func (v NullableStructuresSoftwareCategory) Get() *StructuresSoftwareCategory {
	return v.value
}

func (v *NullableStructuresSoftwareCategory) Set(val *StructuresSoftwareCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresSoftwareCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresSoftwareCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresSoftwareCategory(val *StructuresSoftwareCategory) *NullableStructuresSoftwareCategory {
	return &NullableStructuresSoftwareCategory{value: val, isSet: true}
}

func (v NullableStructuresSoftwareCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresSoftwareCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


