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

// ManageResetVpsResponseError struct for ManageResetVpsResponseError
type ManageResetVpsResponseError struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewManageResetVpsResponseError instantiates a new ManageResetVpsResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageResetVpsResponseError() *ManageResetVpsResponseError {
	this := ManageResetVpsResponseError{}
	return &this
}

// NewManageResetVpsResponseErrorWithDefaults instantiates a new ManageResetVpsResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageResetVpsResponseErrorWithDefaults() *ManageResetVpsResponseError {
	this := ManageResetVpsResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ManageResetVpsResponseError) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageResetVpsResponseError) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ManageResetVpsResponseError) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ManageResetVpsResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ManageResetVpsResponseError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageResetVpsResponseError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ManageResetVpsResponseError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ManageResetVpsResponseError) SetMessage(v string) {
	o.Message = &v
}

func (o ManageResetVpsResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableManageResetVpsResponseError struct {
	value *ManageResetVpsResponseError
	isSet bool
}

func (v NullableManageResetVpsResponseError) Get() *ManageResetVpsResponseError {
	return v.value
}

func (v *NullableManageResetVpsResponseError) Set(val *ManageResetVpsResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableManageResetVpsResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableManageResetVpsResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageResetVpsResponseError(val *ManageResetVpsResponseError) *NullableManageResetVpsResponseError {
	return &NullableManageResetVpsResponseError{value: val, isSet: true}
}

func (v NullableManageResetVpsResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageResetVpsResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


