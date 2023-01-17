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

// NetworkCreatePrivateNetworkResponse struct for NetworkCreatePrivateNetworkResponse
type NetworkCreatePrivateNetworkResponse struct {
	Network *StructuresPrivateNetwork `json:"network,omitempty"`
	Error *NetworkCreatePrivateNetworkResponseError `json:"error,omitempty"`
}

// NewNetworkCreatePrivateNetworkResponse instantiates a new NetworkCreatePrivateNetworkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkCreatePrivateNetworkResponse() *NetworkCreatePrivateNetworkResponse {
	this := NetworkCreatePrivateNetworkResponse{}
	return &this
}

// NewNetworkCreatePrivateNetworkResponseWithDefaults instantiates a new NetworkCreatePrivateNetworkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkCreatePrivateNetworkResponseWithDefaults() *NetworkCreatePrivateNetworkResponse {
	this := NetworkCreatePrivateNetworkResponse{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *NetworkCreatePrivateNetworkResponse) GetNetwork() StructuresPrivateNetwork {
	if o == nil || isNil(o.Network) {
		var ret StructuresPrivateNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkCreatePrivateNetworkResponse) GetNetworkOk() (*StructuresPrivateNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *NetworkCreatePrivateNetworkResponse) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given StructuresPrivateNetwork and assigns it to the Network field.
func (o *NetworkCreatePrivateNetworkResponse) SetNetwork(v StructuresPrivateNetwork) {
	o.Network = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *NetworkCreatePrivateNetworkResponse) GetError() NetworkCreatePrivateNetworkResponseError {
	if o == nil || isNil(o.Error) {
		var ret NetworkCreatePrivateNetworkResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkCreatePrivateNetworkResponse) GetErrorOk() (*NetworkCreatePrivateNetworkResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *NetworkCreatePrivateNetworkResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given NetworkCreatePrivateNetworkResponseError and assigns it to the Error field.
func (o *NetworkCreatePrivateNetworkResponse) SetError(v NetworkCreatePrivateNetworkResponseError) {
	o.Error = &v
}

func (o NetworkCreatePrivateNetworkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkCreatePrivateNetworkResponse struct {
	value *NetworkCreatePrivateNetworkResponse
	isSet bool
}

func (v NullableNetworkCreatePrivateNetworkResponse) Get() *NetworkCreatePrivateNetworkResponse {
	return v.value
}

func (v *NullableNetworkCreatePrivateNetworkResponse) Set(val *NetworkCreatePrivateNetworkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkCreatePrivateNetworkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkCreatePrivateNetworkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkCreatePrivateNetworkResponse(val *NetworkCreatePrivateNetworkResponse) *NullableNetworkCreatePrivateNetworkResponse {
	return &NullableNetworkCreatePrivateNetworkResponse{value: val, isSet: true}
}

func (v NullableNetworkCreatePrivateNetworkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkCreatePrivateNetworkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


