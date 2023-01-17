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

// StatisticGetDiskUsageResponse struct for StatisticGetDiskUsageResponse
type StatisticGetDiskUsageResponse struct {
	DiskUsage *StatisticSeriesData `json:"disk_usage,omitempty"`
}

// NewStatisticGetDiskUsageResponse instantiates a new StatisticGetDiskUsageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticGetDiskUsageResponse() *StatisticGetDiskUsageResponse {
	this := StatisticGetDiskUsageResponse{}
	return &this
}

// NewStatisticGetDiskUsageResponseWithDefaults instantiates a new StatisticGetDiskUsageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticGetDiskUsageResponseWithDefaults() *StatisticGetDiskUsageResponse {
	this := StatisticGetDiskUsageResponse{}
	return &this
}

// GetDiskUsage returns the DiskUsage field value if set, zero value otherwise.
func (o *StatisticGetDiskUsageResponse) GetDiskUsage() StatisticSeriesData {
	if o == nil || isNil(o.DiskUsage) {
		var ret StatisticSeriesData
		return ret
	}
	return *o.DiskUsage
}

// GetDiskUsageOk returns a tuple with the DiskUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetDiskUsageResponse) GetDiskUsageOk() (*StatisticSeriesData, bool) {
	if o == nil || isNil(o.DiskUsage) {
    return nil, false
	}
	return o.DiskUsage, true
}

// HasDiskUsage returns a boolean if a field has been set.
func (o *StatisticGetDiskUsageResponse) HasDiskUsage() bool {
	if o != nil && !isNil(o.DiskUsage) {
		return true
	}

	return false
}

// SetDiskUsage gets a reference to the given StatisticSeriesData and assigns it to the DiskUsage field.
func (o *StatisticGetDiskUsageResponse) SetDiskUsage(v StatisticSeriesData) {
	o.DiskUsage = &v
}

func (o StatisticGetDiskUsageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DiskUsage) {
		toSerialize["disk_usage"] = o.DiskUsage
	}
	return json.Marshal(toSerialize)
}

type NullableStatisticGetDiskUsageResponse struct {
	value *StatisticGetDiskUsageResponse
	isSet bool
}

func (v NullableStatisticGetDiskUsageResponse) Get() *StatisticGetDiskUsageResponse {
	return v.value
}

func (v *NullableStatisticGetDiskUsageResponse) Set(val *StatisticGetDiskUsageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticGetDiskUsageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticGetDiskUsageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticGetDiskUsageResponse(val *StatisticGetDiskUsageResponse) *NullableStatisticGetDiskUsageResponse {
	return &NullableStatisticGetDiskUsageResponse{value: val, isSet: true}
}

func (v NullableStatisticGetDiskUsageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticGetDiskUsageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


