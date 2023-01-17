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

// StatisticGetMemoryResponse struct for StatisticGetMemoryResponse
type StatisticGetMemoryResponse struct {
	Memory *StatisticSeriesData `json:"memory,omitempty"`
}

// NewStatisticGetMemoryResponse instantiates a new StatisticGetMemoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticGetMemoryResponse() *StatisticGetMemoryResponse {
	this := StatisticGetMemoryResponse{}
	return &this
}

// NewStatisticGetMemoryResponseWithDefaults instantiates a new StatisticGetMemoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticGetMemoryResponseWithDefaults() *StatisticGetMemoryResponse {
	this := StatisticGetMemoryResponse{}
	return &this
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *StatisticGetMemoryResponse) GetMemory() StatisticSeriesData {
	if o == nil || isNil(o.Memory) {
		var ret StatisticSeriesData
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetMemoryResponse) GetMemoryOk() (*StatisticSeriesData, bool) {
	if o == nil || isNil(o.Memory) {
    return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *StatisticGetMemoryResponse) HasMemory() bool {
	if o != nil && !isNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given StatisticSeriesData and assigns it to the Memory field.
func (o *StatisticGetMemoryResponse) SetMemory(v StatisticSeriesData) {
	o.Memory = &v
}

func (o StatisticGetMemoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	return json.Marshal(toSerialize)
}

type NullableStatisticGetMemoryResponse struct {
	value *StatisticGetMemoryResponse
	isSet bool
}

func (v NullableStatisticGetMemoryResponse) Get() *StatisticGetMemoryResponse {
	return v.value
}

func (v *NullableStatisticGetMemoryResponse) Set(val *StatisticGetMemoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticGetMemoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticGetMemoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticGetMemoryResponse(val *StatisticGetMemoryResponse) *NullableStatisticGetMemoryResponse {
	return &NullableStatisticGetMemoryResponse{value: val, isSet: true}
}

func (v NullableStatisticGetMemoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticGetMemoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


