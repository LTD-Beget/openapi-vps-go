/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// SshKeyAddResponse struct for SshKeyAddResponse
type SshKeyAddResponse struct {
	Key *StructuresSshKeyInfo `json:"key,omitempty"`
	Error *SshKeyAddResponseError `json:"error,omitempty"`
}

// NewSshKeyAddResponse instantiates a new SshKeyAddResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshKeyAddResponse() *SshKeyAddResponse {
	this := SshKeyAddResponse{}
	return &this
}

// NewSshKeyAddResponseWithDefaults instantiates a new SshKeyAddResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshKeyAddResponseWithDefaults() *SshKeyAddResponse {
	this := SshKeyAddResponse{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SshKeyAddResponse) GetKey() StructuresSshKeyInfo {
	if o == nil || isNil(o.Key) {
		var ret StructuresSshKeyInfo
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyAddResponse) GetKeyOk() (*StructuresSshKeyInfo, bool) {
	if o == nil || isNil(o.Key) {
    return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SshKeyAddResponse) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given StructuresSshKeyInfo and assigns it to the Key field.
func (o *SshKeyAddResponse) SetKey(v StructuresSshKeyInfo) {
	o.Key = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SshKeyAddResponse) GetError() SshKeyAddResponseError {
	if o == nil || isNil(o.Error) {
		var ret SshKeyAddResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyAddResponse) GetErrorOk() (*SshKeyAddResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SshKeyAddResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given SshKeyAddResponseError and assigns it to the Error field.
func (o *SshKeyAddResponse) SetError(v SshKeyAddResponseError) {
	o.Error = &v
}

func (o SshKeyAddResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableSshKeyAddResponse struct {
	value *SshKeyAddResponse
	isSet bool
}

func (v NullableSshKeyAddResponse) Get() *SshKeyAddResponse {
	return v.value
}

func (v *NullableSshKeyAddResponse) Set(val *SshKeyAddResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSshKeyAddResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSshKeyAddResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshKeyAddResponse(val *SshKeyAddResponse) *NullableSshKeyAddResponse {
	return &NullableSshKeyAddResponse{value: val, isSet: true}
}

func (v NullableSshKeyAddResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshKeyAddResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


