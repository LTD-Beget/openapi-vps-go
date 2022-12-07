/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// StructuresIpInfo struct for StructuresIpInfo
type StructuresIpInfo struct {
	Ip *string `json:"ip,omitempty"`
	Mac *string `json:"mac,omitempty"`
	VpsId *string `json:"vps_id,omitempty"`
}

// NewStructuresIpInfo instantiates a new StructuresIpInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresIpInfo() *StructuresIpInfo {
	this := StructuresIpInfo{}
	return &this
}

// NewStructuresIpInfoWithDefaults instantiates a new StructuresIpInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresIpInfoWithDefaults() *StructuresIpInfo {
	this := StructuresIpInfo{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *StructuresIpInfo) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresIpInfo) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *StructuresIpInfo) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *StructuresIpInfo) SetIp(v string) {
	o.Ip = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *StructuresIpInfo) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresIpInfo) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *StructuresIpInfo) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *StructuresIpInfo) SetMac(v string) {
	o.Mac = &v
}

// GetVpsId returns the VpsId field value if set, zero value otherwise.
func (o *StructuresIpInfo) GetVpsId() string {
	if o == nil || isNil(o.VpsId) {
		var ret string
		return ret
	}
	return *o.VpsId
}

// GetVpsIdOk returns a tuple with the VpsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresIpInfo) GetVpsIdOk() (*string, bool) {
	if o == nil || isNil(o.VpsId) {
    return nil, false
	}
	return o.VpsId, true
}

// HasVpsId returns a boolean if a field has been set.
func (o *StructuresIpInfo) HasVpsId() bool {
	if o != nil && !isNil(o.VpsId) {
		return true
	}

	return false
}

// SetVpsId gets a reference to the given string and assigns it to the VpsId field.
func (o *StructuresIpInfo) SetVpsId(v string) {
	o.VpsId = &v
}

func (o StructuresIpInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.VpsId) {
		toSerialize["vps_id"] = o.VpsId
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresIpInfo struct {
	value *StructuresIpInfo
	isSet bool
}

func (v NullableStructuresIpInfo) Get() *StructuresIpInfo {
	return v.value
}

func (v *NullableStructuresIpInfo) Set(val *StructuresIpInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresIpInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresIpInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresIpInfo(val *StructuresIpInfo) *NullableStructuresIpInfo {
	return &NullableStructuresIpInfo{value: val, isSet: true}
}

func (v NullableStructuresIpInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresIpInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


