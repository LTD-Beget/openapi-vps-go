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

// StructuresAdditionalIpInfo struct for StructuresAdditionalIpInfo
type StructuresAdditionalIpInfo struct {
	Ip *string `json:"ip,omitempty"`
	VpsId *string `json:"vps_id,omitempty"`
}

// NewStructuresAdditionalIpInfo instantiates a new StructuresAdditionalIpInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresAdditionalIpInfo() *StructuresAdditionalIpInfo {
	this := StructuresAdditionalIpInfo{}
	return &this
}

// NewStructuresAdditionalIpInfoWithDefaults instantiates a new StructuresAdditionalIpInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresAdditionalIpInfoWithDefaults() *StructuresAdditionalIpInfo {
	this := StructuresAdditionalIpInfo{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *StructuresAdditionalIpInfo) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresAdditionalIpInfo) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *StructuresAdditionalIpInfo) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *StructuresAdditionalIpInfo) SetIp(v string) {
	o.Ip = &v
}

// GetVpsId returns the VpsId field value if set, zero value otherwise.
func (o *StructuresAdditionalIpInfo) GetVpsId() string {
	if o == nil || isNil(o.VpsId) {
		var ret string
		return ret
	}
	return *o.VpsId
}

// GetVpsIdOk returns a tuple with the VpsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresAdditionalIpInfo) GetVpsIdOk() (*string, bool) {
	if o == nil || isNil(o.VpsId) {
    return nil, false
	}
	return o.VpsId, true
}

// HasVpsId returns a boolean if a field has been set.
func (o *StructuresAdditionalIpInfo) HasVpsId() bool {
	if o != nil && !isNil(o.VpsId) {
		return true
	}

	return false
}

// SetVpsId gets a reference to the given string and assigns it to the VpsId field.
func (o *StructuresAdditionalIpInfo) SetVpsId(v string) {
	o.VpsId = &v
}

func (o StructuresAdditionalIpInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.VpsId) {
		toSerialize["vps_id"] = o.VpsId
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresAdditionalIpInfo struct {
	value *StructuresAdditionalIpInfo
	isSet bool
}

func (v NullableStructuresAdditionalIpInfo) Get() *StructuresAdditionalIpInfo {
	return v.value
}

func (v *NullableStructuresAdditionalIpInfo) Set(val *StructuresAdditionalIpInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresAdditionalIpInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresAdditionalIpInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresAdditionalIpInfo(val *StructuresAdditionalIpInfo) *NullableStructuresAdditionalIpInfo {
	return &NullableStructuresAdditionalIpInfo{value: val, isSet: true}
}

func (v NullableStructuresAdditionalIpInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresAdditionalIpInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


