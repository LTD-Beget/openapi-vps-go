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

// StatisticGetLoadAverageResponse struct for StatisticGetLoadAverageResponse
type StatisticGetLoadAverageResponse struct {
	Date []string `json:"date,omitempty"`
	La1 []float64 `json:"la1,omitempty"`
	La5 []float64 `json:"la5,omitempty"`
	La15 []float64 `json:"la15,omitempty"`
}

// NewStatisticGetLoadAverageResponse instantiates a new StatisticGetLoadAverageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticGetLoadAverageResponse() *StatisticGetLoadAverageResponse {
	this := StatisticGetLoadAverageResponse{}
	return &this
}

// NewStatisticGetLoadAverageResponseWithDefaults instantiates a new StatisticGetLoadAverageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticGetLoadAverageResponseWithDefaults() *StatisticGetLoadAverageResponse {
	this := StatisticGetLoadAverageResponse{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *StatisticGetLoadAverageResponse) GetDate() []string {
	if o == nil || isNil(o.Date) {
		var ret []string
		return ret
	}
	return o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetLoadAverageResponse) GetDateOk() ([]string, bool) {
	if o == nil || isNil(o.Date) {
    return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *StatisticGetLoadAverageResponse) HasDate() bool {
	if o != nil && !isNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given []string and assigns it to the Date field.
func (o *StatisticGetLoadAverageResponse) SetDate(v []string) {
	o.Date = v
}

// GetLa1 returns the La1 field value if set, zero value otherwise.
func (o *StatisticGetLoadAverageResponse) GetLa1() []float64 {
	if o == nil || isNil(o.La1) {
		var ret []float64
		return ret
	}
	return o.La1
}

// GetLa1Ok returns a tuple with the La1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetLoadAverageResponse) GetLa1Ok() ([]float64, bool) {
	if o == nil || isNil(o.La1) {
    return nil, false
	}
	return o.La1, true
}

// HasLa1 returns a boolean if a field has been set.
func (o *StatisticGetLoadAverageResponse) HasLa1() bool {
	if o != nil && !isNil(o.La1) {
		return true
	}

	return false
}

// SetLa1 gets a reference to the given []float64 and assigns it to the La1 field.
func (o *StatisticGetLoadAverageResponse) SetLa1(v []float64) {
	o.La1 = v
}

// GetLa5 returns the La5 field value if set, zero value otherwise.
func (o *StatisticGetLoadAverageResponse) GetLa5() []float64 {
	if o == nil || isNil(o.La5) {
		var ret []float64
		return ret
	}
	return o.La5
}

// GetLa5Ok returns a tuple with the La5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetLoadAverageResponse) GetLa5Ok() ([]float64, bool) {
	if o == nil || isNil(o.La5) {
    return nil, false
	}
	return o.La5, true
}

// HasLa5 returns a boolean if a field has been set.
func (o *StatisticGetLoadAverageResponse) HasLa5() bool {
	if o != nil && !isNil(o.La5) {
		return true
	}

	return false
}

// SetLa5 gets a reference to the given []float64 and assigns it to the La5 field.
func (o *StatisticGetLoadAverageResponse) SetLa5(v []float64) {
	o.La5 = v
}

// GetLa15 returns the La15 field value if set, zero value otherwise.
func (o *StatisticGetLoadAverageResponse) GetLa15() []float64 {
	if o == nil || isNil(o.La15) {
		var ret []float64
		return ret
	}
	return o.La15
}

// GetLa15Ok returns a tuple with the La15 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetLoadAverageResponse) GetLa15Ok() ([]float64, bool) {
	if o == nil || isNil(o.La15) {
    return nil, false
	}
	return o.La15, true
}

// HasLa15 returns a boolean if a field has been set.
func (o *StatisticGetLoadAverageResponse) HasLa15() bool {
	if o != nil && !isNil(o.La15) {
		return true
	}

	return false
}

// SetLa15 gets a reference to the given []float64 and assigns it to the La15 field.
func (o *StatisticGetLoadAverageResponse) SetLa15(v []float64) {
	o.La15 = v
}

func (o StatisticGetLoadAverageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !isNil(o.La1) {
		toSerialize["la1"] = o.La1
	}
	if !isNil(o.La5) {
		toSerialize["la5"] = o.La5
	}
	if !isNil(o.La15) {
		toSerialize["la15"] = o.La15
	}
	return json.Marshal(toSerialize)
}

type NullableStatisticGetLoadAverageResponse struct {
	value *StatisticGetLoadAverageResponse
	isSet bool
}

func (v NullableStatisticGetLoadAverageResponse) Get() *StatisticGetLoadAverageResponse {
	return v.value
}

func (v *NullableStatisticGetLoadAverageResponse) Set(val *StatisticGetLoadAverageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticGetLoadAverageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticGetLoadAverageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticGetLoadAverageResponse(val *StatisticGetLoadAverageResponse) *NullableStatisticGetLoadAverageResponse {
	return &NullableStatisticGetLoadAverageResponse{value: val, isSet: true}
}

func (v NullableStatisticGetLoadAverageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticGetLoadAverageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


