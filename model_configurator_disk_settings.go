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

// ConfiguratorDiskSettings struct for ConfiguratorDiskSettings
type ConfiguratorDiskSettings struct {
	Range *ConfiguratorRange `json:"range,omitempty"`
	AvailableRange *ConfiguratorRange `json:"available_range,omitempty"`
	Step *int32 `json:"step,omitempty"`
}

// NewConfiguratorDiskSettings instantiates a new ConfiguratorDiskSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfiguratorDiskSettings() *ConfiguratorDiskSettings {
	this := ConfiguratorDiskSettings{}
	return &this
}

// NewConfiguratorDiskSettingsWithDefaults instantiates a new ConfiguratorDiskSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfiguratorDiskSettingsWithDefaults() *ConfiguratorDiskSettings {
	this := ConfiguratorDiskSettings{}
	return &this
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *ConfiguratorDiskSettings) GetRange() ConfiguratorRange {
	if o == nil || isNil(o.Range) {
		var ret ConfiguratorRange
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguratorDiskSettings) GetRangeOk() (*ConfiguratorRange, bool) {
	if o == nil || isNil(o.Range) {
    return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *ConfiguratorDiskSettings) HasRange() bool {
	if o != nil && !isNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given ConfiguratorRange and assigns it to the Range field.
func (o *ConfiguratorDiskSettings) SetRange(v ConfiguratorRange) {
	o.Range = &v
}

// GetAvailableRange returns the AvailableRange field value if set, zero value otherwise.
func (o *ConfiguratorDiskSettings) GetAvailableRange() ConfiguratorRange {
	if o == nil || isNil(o.AvailableRange) {
		var ret ConfiguratorRange
		return ret
	}
	return *o.AvailableRange
}

// GetAvailableRangeOk returns a tuple with the AvailableRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguratorDiskSettings) GetAvailableRangeOk() (*ConfiguratorRange, bool) {
	if o == nil || isNil(o.AvailableRange) {
    return nil, false
	}
	return o.AvailableRange, true
}

// HasAvailableRange returns a boolean if a field has been set.
func (o *ConfiguratorDiskSettings) HasAvailableRange() bool {
	if o != nil && !isNil(o.AvailableRange) {
		return true
	}

	return false
}

// SetAvailableRange gets a reference to the given ConfiguratorRange and assigns it to the AvailableRange field.
func (o *ConfiguratorDiskSettings) SetAvailableRange(v ConfiguratorRange) {
	o.AvailableRange = &v
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *ConfiguratorDiskSettings) GetStep() int32 {
	if o == nil || isNil(o.Step) {
		var ret int32
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguratorDiskSettings) GetStepOk() (*int32, bool) {
	if o == nil || isNil(o.Step) {
    return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *ConfiguratorDiskSettings) HasStep() bool {
	if o != nil && !isNil(o.Step) {
		return true
	}

	return false
}

// SetStep gets a reference to the given int32 and assigns it to the Step field.
func (o *ConfiguratorDiskSettings) SetStep(v int32) {
	o.Step = &v
}

func (o ConfiguratorDiskSettings) MarshalJSON() ([]byte, error) {
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

type NullableConfiguratorDiskSettings struct {
	value *ConfiguratorDiskSettings
	isSet bool
}

func (v NullableConfiguratorDiskSettings) Get() *ConfiguratorDiskSettings {
	return v.value
}

func (v *NullableConfiguratorDiskSettings) Set(val *ConfiguratorDiskSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableConfiguratorDiskSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableConfiguratorDiskSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfiguratorDiskSettings(val *ConfiguratorDiskSettings) *NullableConfiguratorDiskSettings {
	return &NullableConfiguratorDiskSettings{value: val, isSet: true}
}

func (v NullableConfiguratorDiskSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfiguratorDiskSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

