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

// StructuresSoftwareLicenseBillingTypeMonthly struct for StructuresSoftwareLicenseBillingTypeMonthly
type StructuresSoftwareLicenseBillingTypeMonthly struct {
	PriceMonth *int32 `json:"price_month,omitempty"`
}

// NewStructuresSoftwareLicenseBillingTypeMonthly instantiates a new StructuresSoftwareLicenseBillingTypeMonthly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresSoftwareLicenseBillingTypeMonthly() *StructuresSoftwareLicenseBillingTypeMonthly {
	this := StructuresSoftwareLicenseBillingTypeMonthly{}
	return &this
}

// NewStructuresSoftwareLicenseBillingTypeMonthlyWithDefaults instantiates a new StructuresSoftwareLicenseBillingTypeMonthly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresSoftwareLicenseBillingTypeMonthlyWithDefaults() *StructuresSoftwareLicenseBillingTypeMonthly {
	this := StructuresSoftwareLicenseBillingTypeMonthly{}
	return &this
}

// GetPriceMonth returns the PriceMonth field value if set, zero value otherwise.
func (o *StructuresSoftwareLicenseBillingTypeMonthly) GetPriceMonth() int32 {
	if o == nil || isNil(o.PriceMonth) {
		var ret int32
		return ret
	}
	return *o.PriceMonth
}

// GetPriceMonthOk returns a tuple with the PriceMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSoftwareLicenseBillingTypeMonthly) GetPriceMonthOk() (*int32, bool) {
	if o == nil || isNil(o.PriceMonth) {
    return nil, false
	}
	return o.PriceMonth, true
}

// HasPriceMonth returns a boolean if a field has been set.
func (o *StructuresSoftwareLicenseBillingTypeMonthly) HasPriceMonth() bool {
	if o != nil && !isNil(o.PriceMonth) {
		return true
	}

	return false
}

// SetPriceMonth gets a reference to the given int32 and assigns it to the PriceMonth field.
func (o *StructuresSoftwareLicenseBillingTypeMonthly) SetPriceMonth(v int32) {
	o.PriceMonth = &v
}

func (o StructuresSoftwareLicenseBillingTypeMonthly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PriceMonth) {
		toSerialize["price_month"] = o.PriceMonth
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresSoftwareLicenseBillingTypeMonthly struct {
	value *StructuresSoftwareLicenseBillingTypeMonthly
	isSet bool
}

func (v NullableStructuresSoftwareLicenseBillingTypeMonthly) Get() *StructuresSoftwareLicenseBillingTypeMonthly {
	return v.value
}

func (v *NullableStructuresSoftwareLicenseBillingTypeMonthly) Set(val *StructuresSoftwareLicenseBillingTypeMonthly) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresSoftwareLicenseBillingTypeMonthly) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresSoftwareLicenseBillingTypeMonthly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresSoftwareLicenseBillingTypeMonthly(val *StructuresSoftwareLicenseBillingTypeMonthly) *NullableStructuresSoftwareLicenseBillingTypeMonthly {
	return &NullableStructuresSoftwareLicenseBillingTypeMonthly{value: val, isSet: true}
}

func (v NullableStructuresSoftwareLicenseBillingTypeMonthly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresSoftwareLicenseBillingTypeMonthly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


