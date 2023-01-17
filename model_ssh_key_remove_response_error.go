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

// SshKeyRemoveResponseError struct for SshKeyRemoveResponseError
type SshKeyRemoveResponseError struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewSshKeyRemoveResponseError instantiates a new SshKeyRemoveResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshKeyRemoveResponseError() *SshKeyRemoveResponseError {
	this := SshKeyRemoveResponseError{}
	return &this
}

// NewSshKeyRemoveResponseErrorWithDefaults instantiates a new SshKeyRemoveResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshKeyRemoveResponseErrorWithDefaults() *SshKeyRemoveResponseError {
	this := SshKeyRemoveResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SshKeyRemoveResponseError) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyRemoveResponseError) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SshKeyRemoveResponseError) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SshKeyRemoveResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SshKeyRemoveResponseError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyRemoveResponseError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SshKeyRemoveResponseError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SshKeyRemoveResponseError) SetMessage(v string) {
	o.Message = &v
}

func (o SshKeyRemoveResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableSshKeyRemoveResponseError struct {
	value *SshKeyRemoveResponseError
	isSet bool
}

func (v NullableSshKeyRemoveResponseError) Get() *SshKeyRemoveResponseError {
	return v.value
}

func (v *NullableSshKeyRemoveResponseError) Set(val *SshKeyRemoveResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableSshKeyRemoveResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableSshKeyRemoveResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshKeyRemoveResponseError(val *SshKeyRemoveResponseError) *NullableSshKeyRemoveResponseError {
	return &NullableSshKeyRemoveResponseError{value: val, isSet: true}
}

func (v NullableSshKeyRemoveResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshKeyRemoveResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


