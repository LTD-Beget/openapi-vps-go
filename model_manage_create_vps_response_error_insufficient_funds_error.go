/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// ManageCreateVpsResponseErrorInsufficientFundsError struct for ManageCreateVpsResponseErrorInsufficientFundsError
type ManageCreateVpsResponseErrorInsufficientFundsError struct {
	CurrentBalance *float64 `json:"current_balance,omitempty"`
	NeededBalance *float64 `json:"needed_balance,omitempty"`
	ChargeOnAdd *bool `json:"charge_on_add,omitempty"`
}

// NewManageCreateVpsResponseErrorInsufficientFundsError instantiates a new ManageCreateVpsResponseErrorInsufficientFundsError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageCreateVpsResponseErrorInsufficientFundsError() *ManageCreateVpsResponseErrorInsufficientFundsError {
	this := ManageCreateVpsResponseErrorInsufficientFundsError{}
	return &this
}

// NewManageCreateVpsResponseErrorInsufficientFundsErrorWithDefaults instantiates a new ManageCreateVpsResponseErrorInsufficientFundsError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageCreateVpsResponseErrorInsufficientFundsErrorWithDefaults() *ManageCreateVpsResponseErrorInsufficientFundsError {
	this := ManageCreateVpsResponseErrorInsufficientFundsError{}
	return &this
}

// GetCurrentBalance returns the CurrentBalance field value if set, zero value otherwise.
func (o *ManageCreateVpsResponseErrorInsufficientFundsError) GetCurrentBalance() float64 {
	if o == nil || isNil(o.CurrentBalance) {
		var ret float64
		return ret
	}
	return *o.CurrentBalance
}

// GetCurrentBalanceOk returns a tuple with the CurrentBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsResponseErrorInsufficientFundsError) GetCurrentBalanceOk() (*float64, bool) {
	if o == nil || isNil(o.CurrentBalance) {
    return nil, false
	}
	return o.CurrentBalance, true
}

// HasCurrentBalance returns a boolean if a field has been set.
func (o *ManageCreateVpsResponseErrorInsufficientFundsError) HasCurrentBalance() bool {
	if o != nil && !isNil(o.CurrentBalance) {
		return true
	}

	return false
}

// SetCurrentBalance gets a reference to the given float64 and assigns it to the CurrentBalance field.
func (o *ManageCreateVpsResponseErrorInsufficientFundsError) SetCurrentBalance(v float64) {
	o.CurrentBalance = &v
}

// GetNeededBalance returns the NeededBalance field value if set, zero value otherwise.
func (o *ManageCreateVpsResponseErrorInsufficientFundsError) GetNeededBalance() float64 {
	if o == nil || isNil(o.NeededBalance) {
		var ret float64
		return ret
	}
	return *o.NeededBalance
}

// GetNeededBalanceOk returns a tuple with the NeededBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsResponseErrorInsufficientFundsError) GetNeededBalanceOk() (*float64, bool) {
	if o == nil || isNil(o.NeededBalance) {
    return nil, false
	}
	return o.NeededBalance, true
}

// HasNeededBalance returns a boolean if a field has been set.
func (o *ManageCreateVpsResponseErrorInsufficientFundsError) HasNeededBalance() bool {
	if o != nil && !isNil(o.NeededBalance) {
		return true
	}

	return false
}

// SetNeededBalance gets a reference to the given float64 and assigns it to the NeededBalance field.
func (o *ManageCreateVpsResponseErrorInsufficientFundsError) SetNeededBalance(v float64) {
	o.NeededBalance = &v
}

// GetChargeOnAdd returns the ChargeOnAdd field value if set, zero value otherwise.
func (o *ManageCreateVpsResponseErrorInsufficientFundsError) GetChargeOnAdd() bool {
	if o == nil || isNil(o.ChargeOnAdd) {
		var ret bool
		return ret
	}
	return *o.ChargeOnAdd
}

// GetChargeOnAddOk returns a tuple with the ChargeOnAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsResponseErrorInsufficientFundsError) GetChargeOnAddOk() (*bool, bool) {
	if o == nil || isNil(o.ChargeOnAdd) {
    return nil, false
	}
	return o.ChargeOnAdd, true
}

// HasChargeOnAdd returns a boolean if a field has been set.
func (o *ManageCreateVpsResponseErrorInsufficientFundsError) HasChargeOnAdd() bool {
	if o != nil && !isNil(o.ChargeOnAdd) {
		return true
	}

	return false
}

// SetChargeOnAdd gets a reference to the given bool and assigns it to the ChargeOnAdd field.
func (o *ManageCreateVpsResponseErrorInsufficientFundsError) SetChargeOnAdd(v bool) {
	o.ChargeOnAdd = &v
}

func (o ManageCreateVpsResponseErrorInsufficientFundsError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CurrentBalance) {
		toSerialize["current_balance"] = o.CurrentBalance
	}
	if !isNil(o.NeededBalance) {
		toSerialize["needed_balance"] = o.NeededBalance
	}
	if !isNil(o.ChargeOnAdd) {
		toSerialize["charge_on_add"] = o.ChargeOnAdd
	}
	return json.Marshal(toSerialize)
}

type NullableManageCreateVpsResponseErrorInsufficientFundsError struct {
	value *ManageCreateVpsResponseErrorInsufficientFundsError
	isSet bool
}

func (v NullableManageCreateVpsResponseErrorInsufficientFundsError) Get() *ManageCreateVpsResponseErrorInsufficientFundsError {
	return v.value
}

func (v *NullableManageCreateVpsResponseErrorInsufficientFundsError) Set(val *ManageCreateVpsResponseErrorInsufficientFundsError) {
	v.value = val
	v.isSet = true
}

func (v NullableManageCreateVpsResponseErrorInsufficientFundsError) IsSet() bool {
	return v.isSet
}

func (v *NullableManageCreateVpsResponseErrorInsufficientFundsError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageCreateVpsResponseErrorInsufficientFundsError(val *ManageCreateVpsResponseErrorInsufficientFundsError) *NullableManageCreateVpsResponseErrorInsufficientFundsError {
	return &NullableManageCreateVpsResponseErrorInsufficientFundsError{value: val, isSet: true}
}

func (v NullableManageCreateVpsResponseErrorInsufficientFundsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageCreateVpsResponseErrorInsufficientFundsError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


