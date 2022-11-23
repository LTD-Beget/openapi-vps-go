/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// SshKeyRemoveResponse struct for SshKeyRemoveResponse
type SshKeyRemoveResponse struct {
	Key *StructuresSshKeyInfo `json:"key,omitempty"`
	Error *SshKeyRemoveResponseError `json:"error,omitempty"`
}

// NewSshKeyRemoveResponse instantiates a new SshKeyRemoveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshKeyRemoveResponse() *SshKeyRemoveResponse {
	this := SshKeyRemoveResponse{}
	return &this
}

// NewSshKeyRemoveResponseWithDefaults instantiates a new SshKeyRemoveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshKeyRemoveResponseWithDefaults() *SshKeyRemoveResponse {
	this := SshKeyRemoveResponse{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SshKeyRemoveResponse) GetKey() StructuresSshKeyInfo {
	if o == nil || isNil(o.Key) {
		var ret StructuresSshKeyInfo
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyRemoveResponse) GetKeyOk() (*StructuresSshKeyInfo, bool) {
	if o == nil || isNil(o.Key) {
    return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SshKeyRemoveResponse) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given StructuresSshKeyInfo and assigns it to the Key field.
func (o *SshKeyRemoveResponse) SetKey(v StructuresSshKeyInfo) {
	o.Key = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SshKeyRemoveResponse) GetError() SshKeyRemoveResponseError {
	if o == nil || isNil(o.Error) {
		var ret SshKeyRemoveResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyRemoveResponse) GetErrorOk() (*SshKeyRemoveResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SshKeyRemoveResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given SshKeyRemoveResponseError and assigns it to the Error field.
func (o *SshKeyRemoveResponse) SetError(v SshKeyRemoveResponseError) {
	o.Error = &v
}

func (o SshKeyRemoveResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableSshKeyRemoveResponse struct {
	value *SshKeyRemoveResponse
	isSet bool
}

func (v NullableSshKeyRemoveResponse) Get() *SshKeyRemoveResponse {
	return v.value
}

func (v *NullableSshKeyRemoveResponse) Set(val *SshKeyRemoveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSshKeyRemoveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSshKeyRemoveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshKeyRemoveResponse(val *SshKeyRemoveResponse) *NullableSshKeyRemoveResponse {
	return &NullableSshKeyRemoveResponse{value: val, isSet: true}
}

func (v NullableSshKeyRemoveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshKeyRemoveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


