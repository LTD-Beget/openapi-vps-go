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

// MarketplacePasswordField struct for MarketplacePasswordField
type MarketplacePasswordField struct {
	Common *MarketplaceFieldCommon `json:"common,omitempty"`
	AutoGenerated *bool `json:"auto_generated,omitempty"`
}

// NewMarketplacePasswordField instantiates a new MarketplacePasswordField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplacePasswordField() *MarketplacePasswordField {
	this := MarketplacePasswordField{}
	return &this
}

// NewMarketplacePasswordFieldWithDefaults instantiates a new MarketplacePasswordField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplacePasswordFieldWithDefaults() *MarketplacePasswordField {
	this := MarketplacePasswordField{}
	return &this
}

// GetCommon returns the Common field value if set, zero value otherwise.
func (o *MarketplacePasswordField) GetCommon() MarketplaceFieldCommon {
	if o == nil || isNil(o.Common) {
		var ret MarketplaceFieldCommon
		return ret
	}
	return *o.Common
}

// GetCommonOk returns a tuple with the Common field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplacePasswordField) GetCommonOk() (*MarketplaceFieldCommon, bool) {
	if o == nil || isNil(o.Common) {
    return nil, false
	}
	return o.Common, true
}

// HasCommon returns a boolean if a field has been set.
func (o *MarketplacePasswordField) HasCommon() bool {
	if o != nil && !isNil(o.Common) {
		return true
	}

	return false
}

// SetCommon gets a reference to the given MarketplaceFieldCommon and assigns it to the Common field.
func (o *MarketplacePasswordField) SetCommon(v MarketplaceFieldCommon) {
	o.Common = &v
}

// GetAutoGenerated returns the AutoGenerated field value if set, zero value otherwise.
func (o *MarketplacePasswordField) GetAutoGenerated() bool {
	if o == nil || isNil(o.AutoGenerated) {
		var ret bool
		return ret
	}
	return *o.AutoGenerated
}

// GetAutoGeneratedOk returns a tuple with the AutoGenerated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplacePasswordField) GetAutoGeneratedOk() (*bool, bool) {
	if o == nil || isNil(o.AutoGenerated) {
    return nil, false
	}
	return o.AutoGenerated, true
}

// HasAutoGenerated returns a boolean if a field has been set.
func (o *MarketplacePasswordField) HasAutoGenerated() bool {
	if o != nil && !isNil(o.AutoGenerated) {
		return true
	}

	return false
}

// SetAutoGenerated gets a reference to the given bool and assigns it to the AutoGenerated field.
func (o *MarketplacePasswordField) SetAutoGenerated(v bool) {
	o.AutoGenerated = &v
}

func (o MarketplacePasswordField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Common) {
		toSerialize["common"] = o.Common
	}
	if !isNil(o.AutoGenerated) {
		toSerialize["auto_generated"] = o.AutoGenerated
	}
	return json.Marshal(toSerialize)
}

type NullableMarketplacePasswordField struct {
	value *MarketplacePasswordField
	isSet bool
}

func (v NullableMarketplacePasswordField) Get() *MarketplacePasswordField {
	return v.value
}

func (v *NullableMarketplacePasswordField) Set(val *MarketplacePasswordField) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplacePasswordField) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplacePasswordField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplacePasswordField(val *MarketplacePasswordField) *NullableMarketplacePasswordField {
	return &NullableMarketplacePasswordField{value: val, isSet: true}
}

func (v NullableMarketplacePasswordField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplacePasswordField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


