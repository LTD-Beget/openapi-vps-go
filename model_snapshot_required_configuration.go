/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// SnapshotRequiredConfiguration struct for SnapshotRequiredConfiguration
type SnapshotRequiredConfiguration struct {
	CpuCount *int32 `json:"cpu_count,omitempty"`
	DiskSize *int32 `json:"disk_size,omitempty"`
	Memory *int32 `json:"memory,omitempty"`
}

// NewSnapshotRequiredConfiguration instantiates a new SnapshotRequiredConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotRequiredConfiguration() *SnapshotRequiredConfiguration {
	this := SnapshotRequiredConfiguration{}
	return &this
}

// NewSnapshotRequiredConfigurationWithDefaults instantiates a new SnapshotRequiredConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotRequiredConfigurationWithDefaults() *SnapshotRequiredConfiguration {
	this := SnapshotRequiredConfiguration{}
	return &this
}

// GetCpuCount returns the CpuCount field value if set, zero value otherwise.
func (o *SnapshotRequiredConfiguration) GetCpuCount() int32 {
	if o == nil || isNil(o.CpuCount) {
		var ret int32
		return ret
	}
	return *o.CpuCount
}

// GetCpuCountOk returns a tuple with the CpuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRequiredConfiguration) GetCpuCountOk() (*int32, bool) {
	if o == nil || isNil(o.CpuCount) {
    return nil, false
	}
	return o.CpuCount, true
}

// HasCpuCount returns a boolean if a field has been set.
func (o *SnapshotRequiredConfiguration) HasCpuCount() bool {
	if o != nil && !isNil(o.CpuCount) {
		return true
	}

	return false
}

// SetCpuCount gets a reference to the given int32 and assigns it to the CpuCount field.
func (o *SnapshotRequiredConfiguration) SetCpuCount(v int32) {
	o.CpuCount = &v
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *SnapshotRequiredConfiguration) GetDiskSize() int32 {
	if o == nil || isNil(o.DiskSize) {
		var ret int32
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRequiredConfiguration) GetDiskSizeOk() (*int32, bool) {
	if o == nil || isNil(o.DiskSize) {
    return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *SnapshotRequiredConfiguration) HasDiskSize() bool {
	if o != nil && !isNil(o.DiskSize) {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given int32 and assigns it to the DiskSize field.
func (o *SnapshotRequiredConfiguration) SetDiskSize(v int32) {
	o.DiskSize = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *SnapshotRequiredConfiguration) GetMemory() int32 {
	if o == nil || isNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRequiredConfiguration) GetMemoryOk() (*int32, bool) {
	if o == nil || isNil(o.Memory) {
    return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *SnapshotRequiredConfiguration) HasMemory() bool {
	if o != nil && !isNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *SnapshotRequiredConfiguration) SetMemory(v int32) {
	o.Memory = &v
}

func (o SnapshotRequiredConfiguration) MarshalJSON() ([]byte, error) {
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

type NullableSnapshotRequiredConfiguration struct {
	value *SnapshotRequiredConfiguration
	isSet bool
}

func (v NullableSnapshotRequiredConfiguration) Get() *SnapshotRequiredConfiguration {
	return v.value
}

func (v *NullableSnapshotRequiredConfiguration) Set(val *SnapshotRequiredConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotRequiredConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotRequiredConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotRequiredConfiguration(val *SnapshotRequiredConfiguration) *NullableSnapshotRequiredConfiguration {
	return &NullableSnapshotRequiredConfiguration{value: val, isSet: true}
}

func (v NullableSnapshotRequiredConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotRequiredConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


