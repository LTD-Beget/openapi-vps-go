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

// SshKeyGetAllResponse struct for SshKeyGetAllResponse
type SshKeyGetAllResponse struct {
	Keys []StructuresSshKeyInfo `json:"keys,omitempty"`
}

// NewSshKeyGetAllResponse instantiates a new SshKeyGetAllResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshKeyGetAllResponse() *SshKeyGetAllResponse {
	this := SshKeyGetAllResponse{}
	return &this
}

// NewSshKeyGetAllResponseWithDefaults instantiates a new SshKeyGetAllResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshKeyGetAllResponseWithDefaults() *SshKeyGetAllResponse {
	this := SshKeyGetAllResponse{}
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *SshKeyGetAllResponse) GetKeys() []StructuresSshKeyInfo {
	if o == nil || isNil(o.Keys) {
		var ret []StructuresSshKeyInfo
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyGetAllResponse) GetKeysOk() ([]StructuresSshKeyInfo, bool) {
	if o == nil || isNil(o.Keys) {
    return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *SshKeyGetAllResponse) HasKeys() bool {
	if o != nil && !isNil(o.Keys) {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []StructuresSshKeyInfo and assigns it to the Keys field.
func (o *SshKeyGetAllResponse) SetKeys(v []StructuresSshKeyInfo) {
	o.Keys = v
}

func (o SshKeyGetAllResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Keys) {
		toSerialize["keys"] = o.Keys
	}
	return json.Marshal(toSerialize)
}

type NullableSshKeyGetAllResponse struct {
	value *SshKeyGetAllResponse
	isSet bool
}

func (v NullableSshKeyGetAllResponse) Get() *SshKeyGetAllResponse {
	return v.value
}

func (v *NullableSshKeyGetAllResponse) Set(val *SshKeyGetAllResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSshKeyGetAllResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSshKeyGetAllResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshKeyGetAllResponse(val *SshKeyGetAllResponse) *NullableSshKeyGetAllResponse {
	return &NullableSshKeyGetAllResponse{value: val, isSet: true}
}

func (v NullableSshKeyGetAllResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshKeyGetAllResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


