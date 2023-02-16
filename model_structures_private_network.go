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

// StructuresPrivateNetwork struct for StructuresPrivateNetwork
type StructuresPrivateNetwork struct {
	Id *string `json:"id,omitempty"`
	Subnet *string `json:"subnet,omitempty"`
	Mask *int32 `json:"mask,omitempty"`
}

// NewStructuresPrivateNetwork instantiates a new StructuresPrivateNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresPrivateNetwork() *StructuresPrivateNetwork {
	this := StructuresPrivateNetwork{}
	return &this
}

// NewStructuresPrivateNetworkWithDefaults instantiates a new StructuresPrivateNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresPrivateNetworkWithDefaults() *StructuresPrivateNetwork {
	this := StructuresPrivateNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StructuresPrivateNetwork) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresPrivateNetwork) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StructuresPrivateNetwork) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StructuresPrivateNetwork) SetId(v string) {
	o.Id = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *StructuresPrivateNetwork) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresPrivateNetwork) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *StructuresPrivateNetwork) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *StructuresPrivateNetwork) SetSubnet(v string) {
	o.Subnet = &v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *StructuresPrivateNetwork) GetMask() int32 {
	if o == nil || isNil(o.Mask) {
		var ret int32
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresPrivateNetwork) GetMaskOk() (*int32, bool) {
	if o == nil || isNil(o.Mask) {
    return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *StructuresPrivateNetwork) HasMask() bool {
	if o != nil && !isNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given int32 and assigns it to the Mask field.
func (o *StructuresPrivateNetwork) SetMask(v int32) {
	o.Mask = &v
}

func (o StructuresPrivateNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !isNil(o.Mask) {
		toSerialize["mask"] = o.Mask
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresPrivateNetwork struct {
	value *StructuresPrivateNetwork
	isSet bool
}

func (v NullableStructuresPrivateNetwork) Get() *StructuresPrivateNetwork {
	return v.value
}

func (v *NullableStructuresPrivateNetwork) Set(val *StructuresPrivateNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresPrivateNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresPrivateNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresPrivateNetwork(val *StructuresPrivateNetwork) *NullableStructuresPrivateNetwork {
	return &NullableStructuresPrivateNetwork{value: val, isSet: true}
}

func (v NullableStructuresPrivateNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresPrivateNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


