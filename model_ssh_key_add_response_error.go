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

// SshKeyAddResponseError struct for SshKeyAddResponseError
type SshKeyAddResponseError struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewSshKeyAddResponseError instantiates a new SshKeyAddResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshKeyAddResponseError() *SshKeyAddResponseError {
	this := SshKeyAddResponseError{}
	return &this
}

// NewSshKeyAddResponseErrorWithDefaults instantiates a new SshKeyAddResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshKeyAddResponseErrorWithDefaults() *SshKeyAddResponseError {
	this := SshKeyAddResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SshKeyAddResponseError) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyAddResponseError) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SshKeyAddResponseError) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SshKeyAddResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SshKeyAddResponseError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyAddResponseError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SshKeyAddResponseError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SshKeyAddResponseError) SetMessage(v string) {
	o.Message = &v
}

func (o SshKeyAddResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableSshKeyAddResponseError struct {
	value *SshKeyAddResponseError
	isSet bool
}

func (v NullableSshKeyAddResponseError) Get() *SshKeyAddResponseError {
	return v.value
}

func (v *NullableSshKeyAddResponseError) Set(val *SshKeyAddResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableSshKeyAddResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableSshKeyAddResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshKeyAddResponseError(val *SshKeyAddResponseError) *NullableSshKeyAddResponseError {
	return &NullableSshKeyAddResponseError{value: val, isSet: true}
}

func (v NullableSshKeyAddResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshKeyAddResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


