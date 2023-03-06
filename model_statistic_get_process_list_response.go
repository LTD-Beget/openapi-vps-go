/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.5.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// StatisticGetProcessListResponse struct for StatisticGetProcessListResponse
type StatisticGetProcessListResponse struct {
	Processes *StatisticGetProcessListResponseProcessList `json:"processes,omitempty"`
	Error *StatisticGetProcessListResponseError `json:"error,omitempty"`
}

// NewStatisticGetProcessListResponse instantiates a new StatisticGetProcessListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticGetProcessListResponse() *StatisticGetProcessListResponse {
	this := StatisticGetProcessListResponse{}
	return &this
}

// NewStatisticGetProcessListResponseWithDefaults instantiates a new StatisticGetProcessListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticGetProcessListResponseWithDefaults() *StatisticGetProcessListResponse {
	this := StatisticGetProcessListResponse{}
	return &this
}

// GetProcesses returns the Processes field value if set, zero value otherwise.
func (o *StatisticGetProcessListResponse) GetProcesses() StatisticGetProcessListResponseProcessList {
	if o == nil || isNil(o.Processes) {
		var ret StatisticGetProcessListResponseProcessList
		return ret
	}
	return *o.Processes
}

// GetProcessesOk returns a tuple with the Processes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetProcessListResponse) GetProcessesOk() (*StatisticGetProcessListResponseProcessList, bool) {
	if o == nil || isNil(o.Processes) {
    return nil, false
	}
	return o.Processes, true
}

// HasProcesses returns a boolean if a field has been set.
func (o *StatisticGetProcessListResponse) HasProcesses() bool {
	if o != nil && !isNil(o.Processes) {
		return true
	}

	return false
}

// SetProcesses gets a reference to the given StatisticGetProcessListResponseProcessList and assigns it to the Processes field.
func (o *StatisticGetProcessListResponse) SetProcesses(v StatisticGetProcessListResponseProcessList) {
	o.Processes = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *StatisticGetProcessListResponse) GetError() StatisticGetProcessListResponseError {
	if o == nil || isNil(o.Error) {
		var ret StatisticGetProcessListResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetProcessListResponse) GetErrorOk() (*StatisticGetProcessListResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *StatisticGetProcessListResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given StatisticGetProcessListResponseError and assigns it to the Error field.
func (o *StatisticGetProcessListResponse) SetError(v StatisticGetProcessListResponseError) {
	o.Error = &v
}

func (o StatisticGetProcessListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Processes) {
		toSerialize["processes"] = o.Processes
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableStatisticGetProcessListResponse struct {
	value *StatisticGetProcessListResponse
	isSet bool
}

func (v NullableStatisticGetProcessListResponse) Get() *StatisticGetProcessListResponse {
	return v.value
}

func (v *NullableStatisticGetProcessListResponse) Set(val *StatisticGetProcessListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticGetProcessListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticGetProcessListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticGetProcessListResponse(val *StatisticGetProcessListResponse) *NullableStatisticGetProcessListResponse {
	return &NullableStatisticGetProcessListResponse{value: val, isSet: true}
}

func (v NullableStatisticGetProcessListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticGetProcessListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


