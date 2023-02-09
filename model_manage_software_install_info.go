/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// ManageSoftwareInstallInfo struct for ManageSoftwareInstallInfo
type ManageSoftwareInstallInfo struct {
	Id *int32 `json:"id,omitempty"`
	Variable *map[string]string `json:"variable,omitempty"`
}

// NewManageSoftwareInstallInfo instantiates a new ManageSoftwareInstallInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageSoftwareInstallInfo() *ManageSoftwareInstallInfo {
	this := ManageSoftwareInstallInfo{}
	return &this
}

// NewManageSoftwareInstallInfoWithDefaults instantiates a new ManageSoftwareInstallInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageSoftwareInstallInfoWithDefaults() *ManageSoftwareInstallInfo {
	this := ManageSoftwareInstallInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManageSoftwareInstallInfo) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageSoftwareInstallInfo) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManageSoftwareInstallInfo) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ManageSoftwareInstallInfo) SetId(v int32) {
	o.Id = &v
}

// GetVariable returns the Variable field value if set, zero value otherwise.
func (o *ManageSoftwareInstallInfo) GetVariable() map[string]string {
	if o == nil || isNil(o.Variable) {
		var ret map[string]string
		return ret
	}
	return *o.Variable
}

// GetVariableOk returns a tuple with the Variable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageSoftwareInstallInfo) GetVariableOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Variable) {
    return nil, false
	}
	return o.Variable, true
}

// HasVariable returns a boolean if a field has been set.
func (o *ManageSoftwareInstallInfo) HasVariable() bool {
	if o != nil && !isNil(o.Variable) {
		return true
	}

	return false
}

// SetVariable gets a reference to the given map[string]string and assigns it to the Variable field.
func (o *ManageSoftwareInstallInfo) SetVariable(v map[string]string) {
	o.Variable = &v
}

func (o ManageSoftwareInstallInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Variable) {
		toSerialize["variable"] = o.Variable
	}
	return json.Marshal(toSerialize)
}

type NullableManageSoftwareInstallInfo struct {
	value *ManageSoftwareInstallInfo
	isSet bool
}

func (v NullableManageSoftwareInstallInfo) Get() *ManageSoftwareInstallInfo {
	return v.value
}

func (v *NullableManageSoftwareInstallInfo) Set(val *ManageSoftwareInstallInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableManageSoftwareInstallInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableManageSoftwareInstallInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageSoftwareInstallInfo(val *ManageSoftwareInstallInfo) *NullableManageSoftwareInstallInfo {
	return &NullableManageSoftwareInstallInfo{value: val, isSet: true}
}

func (v NullableManageSoftwareInstallInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageSoftwareInstallInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


