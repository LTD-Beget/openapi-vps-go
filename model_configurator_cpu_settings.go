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

// ConfiguratorCpuSettings struct for ConfiguratorCpuSettings
type ConfiguratorCpuSettings struct {
	Range *ConfiguratorRange `json:"range,omitempty"`
	AvailableRange *ConfiguratorRange `json:"available_range,omitempty"`
	Step *int32 `json:"step,omitempty"`
}

// NewConfiguratorCpuSettings instantiates a new ConfiguratorCpuSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfiguratorCpuSettings() *ConfiguratorCpuSettings {
	this := ConfiguratorCpuSettings{}
	return &this
}

// NewConfiguratorCpuSettingsWithDefaults instantiates a new ConfiguratorCpuSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfiguratorCpuSettingsWithDefaults() *ConfiguratorCpuSettings {
	this := ConfiguratorCpuSettings{}
	return &this
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *ConfiguratorCpuSettings) GetRange() ConfiguratorRange {
	if o == nil || isNil(o.Range) {
		var ret ConfiguratorRange
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguratorCpuSettings) GetRangeOk() (*ConfiguratorRange, bool) {
	if o == nil || isNil(o.Range) {
    return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *ConfiguratorCpuSettings) HasRange() bool {
	if o != nil && !isNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given ConfiguratorRange and assigns it to the Range field.
func (o *ConfiguratorCpuSettings) SetRange(v ConfiguratorRange) {
	o.Range = &v
}

// GetAvailableRange returns the AvailableRange field value if set, zero value otherwise.
func (o *ConfiguratorCpuSettings) GetAvailableRange() ConfiguratorRange {
	if o == nil || isNil(o.AvailableRange) {
		var ret ConfiguratorRange
		return ret
	}
	return *o.AvailableRange
}

// GetAvailableRangeOk returns a tuple with the AvailableRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguratorCpuSettings) GetAvailableRangeOk() (*ConfiguratorRange, bool) {
	if o == nil || isNil(o.AvailableRange) {
    return nil, false
	}
	return o.AvailableRange, true
}

// HasAvailableRange returns a boolean if a field has been set.
func (o *ConfiguratorCpuSettings) HasAvailableRange() bool {
	if o != nil && !isNil(o.AvailableRange) {
		return true
	}

	return false
}

// SetAvailableRange gets a reference to the given ConfiguratorRange and assigns it to the AvailableRange field.
func (o *ConfiguratorCpuSettings) SetAvailableRange(v ConfiguratorRange) {
	o.AvailableRange = &v
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *ConfiguratorCpuSettings) GetStep() int32 {
	if o == nil || isNil(o.Step) {
		var ret int32
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguratorCpuSettings) GetStepOk() (*int32, bool) {
	if o == nil || isNil(o.Step) {
    return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *ConfiguratorCpuSettings) HasStep() bool {
	if o != nil && !isNil(o.Step) {
		return true
	}

	return false
}

// SetStep gets a reference to the given int32 and assigns it to the Step field.
func (o *ConfiguratorCpuSettings) SetStep(v int32) {
	o.Step = &v
}

func (o ConfiguratorCpuSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !isNil(o.AvailableRange) {
		toSerialize["available_range"] = o.AvailableRange
	}
	if !isNil(o.Step) {
		toSerialize["step"] = o.Step
	}
	return json.Marshal(toSerialize)
}

type NullableConfiguratorCpuSettings struct {
	value *ConfiguratorCpuSettings
	isSet bool
}

func (v NullableConfiguratorCpuSettings) Get() *ConfiguratorCpuSettings {
	return v.value
}

func (v *NullableConfiguratorCpuSettings) Set(val *ConfiguratorCpuSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableConfiguratorCpuSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableConfiguratorCpuSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfiguratorCpuSettings(val *ConfiguratorCpuSettings) *NullableConfiguratorCpuSettings {
	return &NullableConfiguratorCpuSettings{value: val, isSet: true}
}

func (v NullableConfiguratorCpuSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfiguratorCpuSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


