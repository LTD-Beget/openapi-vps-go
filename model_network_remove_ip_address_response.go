/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// NetworkRemoveIpAddressResponse struct for NetworkRemoveIpAddressResponse
type NetworkRemoveIpAddressResponse struct {
	IpAddress *string `json:"ip_address,omitempty"`
	Error *NetworkRemoveIpAddressResponseError `json:"error,omitempty"`
}

// NewNetworkRemoveIpAddressResponse instantiates a new NetworkRemoveIpAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkRemoveIpAddressResponse() *NetworkRemoveIpAddressResponse {
	this := NetworkRemoveIpAddressResponse{}
	return &this
}

// NewNetworkRemoveIpAddressResponseWithDefaults instantiates a new NetworkRemoveIpAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkRemoveIpAddressResponseWithDefaults() *NetworkRemoveIpAddressResponse {
	this := NetworkRemoveIpAddressResponse{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *NetworkRemoveIpAddressResponse) GetIpAddress() string {
	if o == nil || isNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRemoveIpAddressResponse) GetIpAddressOk() (*string, bool) {
	if o == nil || isNil(o.IpAddress) {
    return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *NetworkRemoveIpAddressResponse) HasIpAddress() bool {
	if o != nil && !isNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *NetworkRemoveIpAddressResponse) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *NetworkRemoveIpAddressResponse) GetError() NetworkRemoveIpAddressResponseError {
	if o == nil || isNil(o.Error) {
		var ret NetworkRemoveIpAddressResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRemoveIpAddressResponse) GetErrorOk() (*NetworkRemoveIpAddressResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *NetworkRemoveIpAddressResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given NetworkRemoveIpAddressResponseError and assigns it to the Error field.
func (o *NetworkRemoveIpAddressResponse) SetError(v NetworkRemoveIpAddressResponseError) {
	o.Error = &v
}

func (o NetworkRemoveIpAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkRemoveIpAddressResponse struct {
	value *NetworkRemoveIpAddressResponse
	isSet bool
}

func (v NullableNetworkRemoveIpAddressResponse) Get() *NetworkRemoveIpAddressResponse {
	return v.value
}

func (v *NullableNetworkRemoveIpAddressResponse) Set(val *NetworkRemoveIpAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkRemoveIpAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkRemoveIpAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkRemoveIpAddressResponse(val *NetworkRemoveIpAddressResponse) *NullableNetworkRemoveIpAddressResponse {
	return &NullableNetworkRemoveIpAddressResponse{value: val, isSet: true}
}

func (v NullableNetworkRemoveIpAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkRemoveIpAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


