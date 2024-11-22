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

// ManageConfigurationGroup struct for ManageConfigurationGroup
type ManageConfigurationGroup struct {
	Name *string `json:"name,omitempty"`
	Group *string `json:"group,omitempty"`
}

// NewManageConfigurationGroup instantiates a new ManageConfigurationGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageConfigurationGroup() *ManageConfigurationGroup {
	this := ManageConfigurationGroup{}
	return &this
}

// NewManageConfigurationGroupWithDefaults instantiates a new ManageConfigurationGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageConfigurationGroupWithDefaults() *ManageConfigurationGroup {
	this := ManageConfigurationGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ManageConfigurationGroup) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageConfigurationGroup) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ManageConfigurationGroup) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ManageConfigurationGroup) SetName(v string) {
	o.Name = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ManageConfigurationGroup) GetGroup() string {
	if o == nil || isNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageConfigurationGroup) GetGroupOk() (*string, bool) {
	if o == nil || isNil(o.Group) {
    return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ManageConfigurationGroup) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *ManageConfigurationGroup) SetGroup(v string) {
	o.Group = &v
}

func (o ManageConfigurationGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return json.Marshal(toSerialize)
}

type NullableManageConfigurationGroup struct {
	value *ManageConfigurationGroup
	isSet bool
}

func (v NullableManageConfigurationGroup) Get() *ManageConfigurationGroup {
	return v.value
}

func (v *NullableManageConfigurationGroup) Set(val *ManageConfigurationGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableManageConfigurationGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableManageConfigurationGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageConfigurationGroup(val *ManageConfigurationGroup) *NullableManageConfigurationGroup {
	return &NullableManageConfigurationGroup{value: val, isSet: true}
}

func (v NullableManageConfigurationGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageConfigurationGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

