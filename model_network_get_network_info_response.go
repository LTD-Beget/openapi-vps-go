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

// NetworkGetNetworkInfoResponse struct for NetworkGetNetworkInfoResponse
type NetworkGetNetworkInfoResponse struct {
	Ip []StructuresIpInfo `json:"ip,omitempty"`
	AdditionalIp []StructuresAdditionalIpInfo `json:"additional_ip,omitempty"`
	PrivateNetwork []StructuresPrivateNetwork `json:"private_network,omitempty"`
}

// NewNetworkGetNetworkInfoResponse instantiates a new NetworkGetNetworkInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkGetNetworkInfoResponse() *NetworkGetNetworkInfoResponse {
	this := NetworkGetNetworkInfoResponse{}
	return &this
}

// NewNetworkGetNetworkInfoResponseWithDefaults instantiates a new NetworkGetNetworkInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkGetNetworkInfoResponseWithDefaults() *NetworkGetNetworkInfoResponse {
	this := NetworkGetNetworkInfoResponse{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *NetworkGetNetworkInfoResponse) GetIp() []StructuresIpInfo {
	if o == nil || isNil(o.Ip) {
		var ret []StructuresIpInfo
		return ret
	}
	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkGetNetworkInfoResponse) GetIpOk() ([]StructuresIpInfo, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *NetworkGetNetworkInfoResponse) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given []StructuresIpInfo and assigns it to the Ip field.
func (o *NetworkGetNetworkInfoResponse) SetIp(v []StructuresIpInfo) {
	o.Ip = v
}

// GetAdditionalIp returns the AdditionalIp field value if set, zero value otherwise.
func (o *NetworkGetNetworkInfoResponse) GetAdditionalIp() []StructuresAdditionalIpInfo {
	if o == nil || isNil(o.AdditionalIp) {
		var ret []StructuresAdditionalIpInfo
		return ret
	}
	return o.AdditionalIp
}

// GetAdditionalIpOk returns a tuple with the AdditionalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkGetNetworkInfoResponse) GetAdditionalIpOk() ([]StructuresAdditionalIpInfo, bool) {
	if o == nil || isNil(o.AdditionalIp) {
    return nil, false
	}
	return o.AdditionalIp, true
}

// HasAdditionalIp returns a boolean if a field has been set.
func (o *NetworkGetNetworkInfoResponse) HasAdditionalIp() bool {
	if o != nil && !isNil(o.AdditionalIp) {
		return true
	}

	return false
}

// SetAdditionalIp gets a reference to the given []StructuresAdditionalIpInfo and assigns it to the AdditionalIp field.
func (o *NetworkGetNetworkInfoResponse) SetAdditionalIp(v []StructuresAdditionalIpInfo) {
	o.AdditionalIp = v
}

// GetPrivateNetwork returns the PrivateNetwork field value if set, zero value otherwise.
func (o *NetworkGetNetworkInfoResponse) GetPrivateNetwork() []StructuresPrivateNetwork {
	if o == nil || isNil(o.PrivateNetwork) {
		var ret []StructuresPrivateNetwork
		return ret
	}
	return o.PrivateNetwork
}

// GetPrivateNetworkOk returns a tuple with the PrivateNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkGetNetworkInfoResponse) GetPrivateNetworkOk() ([]StructuresPrivateNetwork, bool) {
	if o == nil || isNil(o.PrivateNetwork) {
    return nil, false
	}
	return o.PrivateNetwork, true
}

// HasPrivateNetwork returns a boolean if a field has been set.
func (o *NetworkGetNetworkInfoResponse) HasPrivateNetwork() bool {
	if o != nil && !isNil(o.PrivateNetwork) {
		return true
	}

	return false
}

// SetPrivateNetwork gets a reference to the given []StructuresPrivateNetwork and assigns it to the PrivateNetwork field.
func (o *NetworkGetNetworkInfoResponse) SetPrivateNetwork(v []StructuresPrivateNetwork) {
	o.PrivateNetwork = v
}

func (o NetworkGetNetworkInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.AdditionalIp) {
		toSerialize["additional_ip"] = o.AdditionalIp
	}
	if !isNil(o.PrivateNetwork) {
		toSerialize["private_network"] = o.PrivateNetwork
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkGetNetworkInfoResponse struct {
	value *NetworkGetNetworkInfoResponse
	isSet bool
}

func (v NullableNetworkGetNetworkInfoResponse) Get() *NetworkGetNetworkInfoResponse {
	return v.value
}

func (v *NullableNetworkGetNetworkInfoResponse) Set(val *NetworkGetNetworkInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkGetNetworkInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkGetNetworkInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkGetNetworkInfoResponse(val *NetworkGetNetworkInfoResponse) *NullableNetworkGetNetworkInfoResponse {
	return &NullableNetworkGetNetworkInfoResponse{value: val, isSet: true}
}

func (v NullableNetworkGetNetworkInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkGetNetworkInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


