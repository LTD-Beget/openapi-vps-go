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

// NetworkOrderIpAddressResponse struct for NetworkOrderIpAddressResponse
type NetworkOrderIpAddressResponse struct {
	IpAddress *string `json:"ip_address,omitempty"`
	Error *NetworkOrderIpAddressResponseError `json:"error,omitempty"`
}

// NewNetworkOrderIpAddressResponse instantiates a new NetworkOrderIpAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkOrderIpAddressResponse() *NetworkOrderIpAddressResponse {
	this := NetworkOrderIpAddressResponse{}
	return &this
}

// NewNetworkOrderIpAddressResponseWithDefaults instantiates a new NetworkOrderIpAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkOrderIpAddressResponseWithDefaults() *NetworkOrderIpAddressResponse {
	this := NetworkOrderIpAddressResponse{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *NetworkOrderIpAddressResponse) GetIpAddress() string {
	if o == nil || isNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkOrderIpAddressResponse) GetIpAddressOk() (*string, bool) {
	if o == nil || isNil(o.IpAddress) {
    return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *NetworkOrderIpAddressResponse) HasIpAddress() bool {
	if o != nil && !isNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *NetworkOrderIpAddressResponse) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *NetworkOrderIpAddressResponse) GetError() NetworkOrderIpAddressResponseError {
	if o == nil || isNil(o.Error) {
		var ret NetworkOrderIpAddressResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkOrderIpAddressResponse) GetErrorOk() (*NetworkOrderIpAddressResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *NetworkOrderIpAddressResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given NetworkOrderIpAddressResponseError and assigns it to the Error field.
func (o *NetworkOrderIpAddressResponse) SetError(v NetworkOrderIpAddressResponseError) {
	o.Error = &v
}

func (o NetworkOrderIpAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkOrderIpAddressResponse struct {
	value *NetworkOrderIpAddressResponse
	isSet bool
}

func (v NullableNetworkOrderIpAddressResponse) Get() *NetworkOrderIpAddressResponse {
	return v.value
}

func (v *NullableNetworkOrderIpAddressResponse) Set(val *NetworkOrderIpAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkOrderIpAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkOrderIpAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkOrderIpAddressResponse(val *NetworkOrderIpAddressResponse) *NullableNetworkOrderIpAddressResponse {
	return &NullableNetworkOrderIpAddressResponse{value: val, isSet: true}
}

func (v NullableNetworkOrderIpAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkOrderIpAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


