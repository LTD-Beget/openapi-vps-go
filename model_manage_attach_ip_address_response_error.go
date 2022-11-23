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

// ManageAttachIpAddressResponseError struct for ManageAttachIpAddressResponseError
type ManageAttachIpAddressResponseError struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewManageAttachIpAddressResponseError instantiates a new ManageAttachIpAddressResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageAttachIpAddressResponseError() *ManageAttachIpAddressResponseError {
	this := ManageAttachIpAddressResponseError{}
	return &this
}

// NewManageAttachIpAddressResponseErrorWithDefaults instantiates a new ManageAttachIpAddressResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageAttachIpAddressResponseErrorWithDefaults() *ManageAttachIpAddressResponseError {
	this := ManageAttachIpAddressResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ManageAttachIpAddressResponseError) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageAttachIpAddressResponseError) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ManageAttachIpAddressResponseError) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ManageAttachIpAddressResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ManageAttachIpAddressResponseError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageAttachIpAddressResponseError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ManageAttachIpAddressResponseError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ManageAttachIpAddressResponseError) SetMessage(v string) {
	o.Message = &v
}

func (o ManageAttachIpAddressResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableManageAttachIpAddressResponseError struct {
	value *ManageAttachIpAddressResponseError
	isSet bool
}

func (v NullableManageAttachIpAddressResponseError) Get() *ManageAttachIpAddressResponseError {
	return v.value
}

func (v *NullableManageAttachIpAddressResponseError) Set(val *ManageAttachIpAddressResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableManageAttachIpAddressResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableManageAttachIpAddressResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageAttachIpAddressResponseError(val *ManageAttachIpAddressResponseError) *NullableManageAttachIpAddressResponseError {
	return &NullableManageAttachIpAddressResponseError{value: val, isSet: true}
}

func (v NullableManageAttachIpAddressResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageAttachIpAddressResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


