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

// StructuresConfigurationParams struct for StructuresConfigurationParams
type StructuresConfigurationParams struct {
	CpuCount *int32 `json:"cpu_count,omitempty"`
	DiskSize *int32 `json:"disk_size,omitempty"`
	Memory *int32 `json:"memory,omitempty"`
}

// NewStructuresConfigurationParams instantiates a new StructuresConfigurationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresConfigurationParams() *StructuresConfigurationParams {
	this := StructuresConfigurationParams{}
	return &this
}

// NewStructuresConfigurationParamsWithDefaults instantiates a new StructuresConfigurationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresConfigurationParamsWithDefaults() *StructuresConfigurationParams {
	this := StructuresConfigurationParams{}
	return &this
}

// GetCpuCount returns the CpuCount field value if set, zero value otherwise.
func (o *StructuresConfigurationParams) GetCpuCount() int32 {
	if o == nil || isNil(o.CpuCount) {
		var ret int32
		return ret
	}
	return *o.CpuCount
}

// GetCpuCountOk returns a tuple with the CpuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresConfigurationParams) GetCpuCountOk() (*int32, bool) {
	if o == nil || isNil(o.CpuCount) {
    return nil, false
	}
	return o.CpuCount, true
}

// HasCpuCount returns a boolean if a field has been set.
func (o *StructuresConfigurationParams) HasCpuCount() bool {
	if o != nil && !isNil(o.CpuCount) {
		return true
	}

	return false
}

// SetCpuCount gets a reference to the given int32 and assigns it to the CpuCount field.
func (o *StructuresConfigurationParams) SetCpuCount(v int32) {
	o.CpuCount = &v
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *StructuresConfigurationParams) GetDiskSize() int32 {
	if o == nil || isNil(o.DiskSize) {
		var ret int32
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresConfigurationParams) GetDiskSizeOk() (*int32, bool) {
	if o == nil || isNil(o.DiskSize) {
    return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *StructuresConfigurationParams) HasDiskSize() bool {
	if o != nil && !isNil(o.DiskSize) {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given int32 and assigns it to the DiskSize field.
func (o *StructuresConfigurationParams) SetDiskSize(v int32) {
	o.DiskSize = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *StructuresConfigurationParams) GetMemory() int32 {
	if o == nil || isNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresConfigurationParams) GetMemoryOk() (*int32, bool) {
	if o == nil || isNil(o.Memory) {
    return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *StructuresConfigurationParams) HasMemory() bool {
	if o != nil && !isNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *StructuresConfigurationParams) SetMemory(v int32) {
	o.Memory = &v
}

func (o StructuresConfigurationParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CpuCount) {
		toSerialize["cpu_count"] = o.CpuCount
	}
	if !isNil(o.DiskSize) {
		toSerialize["disk_size"] = o.DiskSize
	}
	if !isNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresConfigurationParams struct {
	value *StructuresConfigurationParams
	isSet bool
}

func (v NullableStructuresConfigurationParams) Get() *StructuresConfigurationParams {
	return v.value
}

func (v *NullableStructuresConfigurationParams) Set(val *StructuresConfigurationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresConfigurationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresConfigurationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresConfigurationParams(val *StructuresConfigurationParams) *NullableStructuresConfigurationParams {
	return &NullableStructuresConfigurationParams{value: val, isSet: true}
}

func (v NullableStructuresConfigurationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresConfigurationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


