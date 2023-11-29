/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// ConfiguratorConfiguratorSettings struct for ConfiguratorConfiguratorSettings
type ConfiguratorConfiguratorSettings struct {
	CpuSettings *ConfiguratorCpuSettings `json:"cpu_settings,omitempty"`
	DiskSettings *ConfiguratorDiskSettings `json:"disk_settings,omitempty"`
	MemorySettings *ConfiguratorMemorySettings `json:"memory_settings,omitempty"`
}

// NewConfiguratorConfiguratorSettings instantiates a new ConfiguratorConfiguratorSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfiguratorConfiguratorSettings() *ConfiguratorConfiguratorSettings {
	this := ConfiguratorConfiguratorSettings{}
	return &this
}

// NewConfiguratorConfiguratorSettingsWithDefaults instantiates a new ConfiguratorConfiguratorSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfiguratorConfiguratorSettingsWithDefaults() *ConfiguratorConfiguratorSettings {
	this := ConfiguratorConfiguratorSettings{}
	return &this
}

// GetCpuSettings returns the CpuSettings field value if set, zero value otherwise.
func (o *ConfiguratorConfiguratorSettings) GetCpuSettings() ConfiguratorCpuSettings {
	if o == nil || isNil(o.CpuSettings) {
		var ret ConfiguratorCpuSettings
		return ret
	}
	return *o.CpuSettings
}

// GetCpuSettingsOk returns a tuple with the CpuSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguratorConfiguratorSettings) GetCpuSettingsOk() (*ConfiguratorCpuSettings, bool) {
	if o == nil || isNil(o.CpuSettings) {
    return nil, false
	}
	return o.CpuSettings, true
}

// HasCpuSettings returns a boolean if a field has been set.
func (o *ConfiguratorConfiguratorSettings) HasCpuSettings() bool {
	if o != nil && !isNil(o.CpuSettings) {
		return true
	}

	return false
}

// SetCpuSettings gets a reference to the given ConfiguratorCpuSettings and assigns it to the CpuSettings field.
func (o *ConfiguratorConfiguratorSettings) SetCpuSettings(v ConfiguratorCpuSettings) {
	o.CpuSettings = &v
}

// GetDiskSettings returns the DiskSettings field value if set, zero value otherwise.
func (o *ConfiguratorConfiguratorSettings) GetDiskSettings() ConfiguratorDiskSettings {
	if o == nil || isNil(o.DiskSettings) {
		var ret ConfiguratorDiskSettings
		return ret
	}
	return *o.DiskSettings
}

// GetDiskSettingsOk returns a tuple with the DiskSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguratorConfiguratorSettings) GetDiskSettingsOk() (*ConfiguratorDiskSettings, bool) {
	if o == nil || isNil(o.DiskSettings) {
    return nil, false
	}
	return o.DiskSettings, true
}

// HasDiskSettings returns a boolean if a field has been set.
func (o *ConfiguratorConfiguratorSettings) HasDiskSettings() bool {
	if o != nil && !isNil(o.DiskSettings) {
		return true
	}

	return false
}

// SetDiskSettings gets a reference to the given ConfiguratorDiskSettings and assigns it to the DiskSettings field.
func (o *ConfiguratorConfiguratorSettings) SetDiskSettings(v ConfiguratorDiskSettings) {
	o.DiskSettings = &v
}

// GetMemorySettings returns the MemorySettings field value if set, zero value otherwise.
func (o *ConfiguratorConfiguratorSettings) GetMemorySettings() ConfiguratorMemorySettings {
	if o == nil || isNil(o.MemorySettings) {
		var ret ConfiguratorMemorySettings
		return ret
	}
	return *o.MemorySettings
}

// GetMemorySettingsOk returns a tuple with the MemorySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguratorConfiguratorSettings) GetMemorySettingsOk() (*ConfiguratorMemorySettings, bool) {
	if o == nil || isNil(o.MemorySettings) {
    return nil, false
	}
	return o.MemorySettings, true
}

// HasMemorySettings returns a boolean if a field has been set.
func (o *ConfiguratorConfiguratorSettings) HasMemorySettings() bool {
	if o != nil && !isNil(o.MemorySettings) {
		return true
	}

	return false
}

// SetMemorySettings gets a reference to the given ConfiguratorMemorySettings and assigns it to the MemorySettings field.
func (o *ConfiguratorConfiguratorSettings) SetMemorySettings(v ConfiguratorMemorySettings) {
	o.MemorySettings = &v
}

func (o ConfiguratorConfiguratorSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CpuSettings) {
		toSerialize["cpu_settings"] = o.CpuSettings
	}
	if !isNil(o.DiskSettings) {
		toSerialize["disk_settings"] = o.DiskSettings
	}
	if !isNil(o.MemorySettings) {
		toSerialize["memory_settings"] = o.MemorySettings
	}
	return json.Marshal(toSerialize)
}

type NullableConfiguratorConfiguratorSettings struct {
	value *ConfiguratorConfiguratorSettings
	isSet bool
}

func (v NullableConfiguratorConfiguratorSettings) Get() *ConfiguratorConfiguratorSettings {
	return v.value
}

func (v *NullableConfiguratorConfiguratorSettings) Set(val *ConfiguratorConfiguratorSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableConfiguratorConfiguratorSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableConfiguratorConfiguratorSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfiguratorConfiguratorSettings(val *ConfiguratorConfiguratorSettings) *NullableConfiguratorConfiguratorSettings {
	return &NullableConfiguratorConfiguratorSettings{value: val, isSet: true}
}

func (v NullableConfiguratorConfiguratorSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfiguratorConfiguratorSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


