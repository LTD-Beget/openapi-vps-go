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

// ConfiguratorGetCalculationResponse struct for ConfiguratorGetCalculationResponse
type ConfiguratorGetCalculationResponse struct {
	Success *ConfiguratorGetCalculationResponseSuccess `json:"success,omitempty"`
	Error *ConfiguratorGetCalculationResponseError `json:"error,omitempty"`
}

// NewConfiguratorGetCalculationResponse instantiates a new ConfiguratorGetCalculationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfiguratorGetCalculationResponse() *ConfiguratorGetCalculationResponse {
	this := ConfiguratorGetCalculationResponse{}
	return &this
}

// NewConfiguratorGetCalculationResponseWithDefaults instantiates a new ConfiguratorGetCalculationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfiguratorGetCalculationResponseWithDefaults() *ConfiguratorGetCalculationResponse {
	this := ConfiguratorGetCalculationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ConfiguratorGetCalculationResponse) GetSuccess() ConfiguratorGetCalculationResponseSuccess {
	if o == nil || isNil(o.Success) {
		var ret ConfiguratorGetCalculationResponseSuccess
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguratorGetCalculationResponse) GetSuccessOk() (*ConfiguratorGetCalculationResponseSuccess, bool) {
	if o == nil || isNil(o.Success) {
    return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ConfiguratorGetCalculationResponse) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given ConfiguratorGetCalculationResponseSuccess and assigns it to the Success field.
func (o *ConfiguratorGetCalculationResponse) SetSuccess(v ConfiguratorGetCalculationResponseSuccess) {
	o.Success = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ConfiguratorGetCalculationResponse) GetError() ConfiguratorGetCalculationResponseError {
	if o == nil || isNil(o.Error) {
		var ret ConfiguratorGetCalculationResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguratorGetCalculationResponse) GetErrorOk() (*ConfiguratorGetCalculationResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ConfiguratorGetCalculationResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ConfiguratorGetCalculationResponseError and assigns it to the Error field.
func (o *ConfiguratorGetCalculationResponse) SetError(v ConfiguratorGetCalculationResponseError) {
	o.Error = &v
}

func (o ConfiguratorGetCalculationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableConfiguratorGetCalculationResponse struct {
	value *ConfiguratorGetCalculationResponse
	isSet bool
}

func (v NullableConfiguratorGetCalculationResponse) Get() *ConfiguratorGetCalculationResponse {
	return v.value
}

func (v *NullableConfiguratorGetCalculationResponse) Set(val *ConfiguratorGetCalculationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfiguratorGetCalculationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfiguratorGetCalculationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfiguratorGetCalculationResponse(val *ConfiguratorGetCalculationResponse) *NullableConfiguratorGetCalculationResponse {
	return &NullableConfiguratorGetCalculationResponse{value: val, isSet: true}
}

func (v NullableConfiguratorGetCalculationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfiguratorGetCalculationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

