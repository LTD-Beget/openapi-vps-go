/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// NetworkSuggestPrivateAddressResponse struct for NetworkSuggestPrivateAddressResponse
type NetworkSuggestPrivateAddressResponse struct {
	Ip *string `json:"ip,omitempty"`
}

// NewNetworkSuggestPrivateAddressResponse instantiates a new NetworkSuggestPrivateAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkSuggestPrivateAddressResponse() *NetworkSuggestPrivateAddressResponse {
	this := NetworkSuggestPrivateAddressResponse{}
	return &this
}

// NewNetworkSuggestPrivateAddressResponseWithDefaults instantiates a new NetworkSuggestPrivateAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkSuggestPrivateAddressResponseWithDefaults() *NetworkSuggestPrivateAddressResponse {
	this := NetworkSuggestPrivateAddressResponse{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *NetworkSuggestPrivateAddressResponse) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSuggestPrivateAddressResponse) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *NetworkSuggestPrivateAddressResponse) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *NetworkSuggestPrivateAddressResponse) SetIp(v string) {
	o.Ip = &v
}

func (o NetworkSuggestPrivateAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkSuggestPrivateAddressResponse struct {
	value *NetworkSuggestPrivateAddressResponse
	isSet bool
}

func (v NullableNetworkSuggestPrivateAddressResponse) Get() *NetworkSuggestPrivateAddressResponse {
	return v.value
}

func (v *NullableNetworkSuggestPrivateAddressResponse) Set(val *NetworkSuggestPrivateAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSuggestPrivateAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSuggestPrivateAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSuggestPrivateAddressResponse(val *NetworkSuggestPrivateAddressResponse) *NullableNetworkSuggestPrivateAddressResponse {
	return &NullableNetworkSuggestPrivateAddressResponse{value: val, isSet: true}
}

func (v NullableNetworkSuggestPrivateAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSuggestPrivateAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


