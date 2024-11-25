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

// SnapshotCreateCalculatorResponse struct for SnapshotCreateCalculatorResponse
type SnapshotCreateCalculatorResponse struct {
	PriceDay *float64 `json:"price_day,omitempty"`
	PriceDayGb *float64 `json:"price_day_gb,omitempty"`
	Size *int32 `json:"size,omitempty"`
}

// NewSnapshotCreateCalculatorResponse instantiates a new SnapshotCreateCalculatorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotCreateCalculatorResponse() *SnapshotCreateCalculatorResponse {
	this := SnapshotCreateCalculatorResponse{}
	return &this
}

// NewSnapshotCreateCalculatorResponseWithDefaults instantiates a new SnapshotCreateCalculatorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotCreateCalculatorResponseWithDefaults() *SnapshotCreateCalculatorResponse {
	this := SnapshotCreateCalculatorResponse{}
	return &this
}

// GetPriceDay returns the PriceDay field value if set, zero value otherwise.
func (o *SnapshotCreateCalculatorResponse) GetPriceDay() float64 {
	if o == nil || isNil(o.PriceDay) {
		var ret float64
		return ret
	}
	return *o.PriceDay
}

// GetPriceDayOk returns a tuple with the PriceDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotCreateCalculatorResponse) GetPriceDayOk() (*float64, bool) {
	if o == nil || isNil(o.PriceDay) {
    return nil, false
	}
	return o.PriceDay, true
}

// HasPriceDay returns a boolean if a field has been set.
func (o *SnapshotCreateCalculatorResponse) HasPriceDay() bool {
	if o != nil && !isNil(o.PriceDay) {
		return true
	}

	return false
}

// SetPriceDay gets a reference to the given float64 and assigns it to the PriceDay field.
func (o *SnapshotCreateCalculatorResponse) SetPriceDay(v float64) {
	o.PriceDay = &v
}

// GetPriceDayGb returns the PriceDayGb field value if set, zero value otherwise.
func (o *SnapshotCreateCalculatorResponse) GetPriceDayGb() float64 {
	if o == nil || isNil(o.PriceDayGb) {
		var ret float64
		return ret
	}
	return *o.PriceDayGb
}

// GetPriceDayGbOk returns a tuple with the PriceDayGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotCreateCalculatorResponse) GetPriceDayGbOk() (*float64, bool) {
	if o == nil || isNil(o.PriceDayGb) {
    return nil, false
	}
	return o.PriceDayGb, true
}

// HasPriceDayGb returns a boolean if a field has been set.
func (o *SnapshotCreateCalculatorResponse) HasPriceDayGb() bool {
	if o != nil && !isNil(o.PriceDayGb) {
		return true
	}

	return false
}

// SetPriceDayGb gets a reference to the given float64 and assigns it to the PriceDayGb field.
func (o *SnapshotCreateCalculatorResponse) SetPriceDayGb(v float64) {
	o.PriceDayGb = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SnapshotCreateCalculatorResponse) GetSize() int32 {
	if o == nil || isNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotCreateCalculatorResponse) GetSizeOk() (*int32, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SnapshotCreateCalculatorResponse) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *SnapshotCreateCalculatorResponse) SetSize(v int32) {
	o.Size = &v
}

func (o SnapshotCreateCalculatorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PriceDay) {
		toSerialize["price_day"] = o.PriceDay
	}
	if !isNil(o.PriceDayGb) {
		toSerialize["price_day_gb"] = o.PriceDayGb
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotCreateCalculatorResponse struct {
	value *SnapshotCreateCalculatorResponse
	isSet bool
}

func (v NullableSnapshotCreateCalculatorResponse) Get() *SnapshotCreateCalculatorResponse {
	return v.value
}

func (v *NullableSnapshotCreateCalculatorResponse) Set(val *SnapshotCreateCalculatorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotCreateCalculatorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotCreateCalculatorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotCreateCalculatorResponse(val *SnapshotCreateCalculatorResponse) *NullableSnapshotCreateCalculatorResponse {
	return &NullableSnapshotCreateCalculatorResponse{value: val, isSet: true}
}

func (v NullableSnapshotCreateCalculatorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotCreateCalculatorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


