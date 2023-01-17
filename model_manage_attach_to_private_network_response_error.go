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

// ManageAttachToPrivateNetworkResponseError struct for ManageAttachToPrivateNetworkResponseError
type ManageAttachToPrivateNetworkResponseError struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewManageAttachToPrivateNetworkResponseError instantiates a new ManageAttachToPrivateNetworkResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageAttachToPrivateNetworkResponseError() *ManageAttachToPrivateNetworkResponseError {
	this := ManageAttachToPrivateNetworkResponseError{}
	return &this
}

// NewManageAttachToPrivateNetworkResponseErrorWithDefaults instantiates a new ManageAttachToPrivateNetworkResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageAttachToPrivateNetworkResponseErrorWithDefaults() *ManageAttachToPrivateNetworkResponseError {
	this := ManageAttachToPrivateNetworkResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ManageAttachToPrivateNetworkResponseError) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageAttachToPrivateNetworkResponseError) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ManageAttachToPrivateNetworkResponseError) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ManageAttachToPrivateNetworkResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ManageAttachToPrivateNetworkResponseError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageAttachToPrivateNetworkResponseError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ManageAttachToPrivateNetworkResponseError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ManageAttachToPrivateNetworkResponseError) SetMessage(v string) {
	o.Message = &v
}

func (o ManageAttachToPrivateNetworkResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableManageAttachToPrivateNetworkResponseError struct {
	value *ManageAttachToPrivateNetworkResponseError
	isSet bool
}

func (v NullableManageAttachToPrivateNetworkResponseError) Get() *ManageAttachToPrivateNetworkResponseError {
	return v.value
}

func (v *NullableManageAttachToPrivateNetworkResponseError) Set(val *ManageAttachToPrivateNetworkResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableManageAttachToPrivateNetworkResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableManageAttachToPrivateNetworkResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageAttachToPrivateNetworkResponseError(val *ManageAttachToPrivateNetworkResponseError) *NullableManageAttachToPrivateNetworkResponseError {
	return &NullableManageAttachToPrivateNetworkResponseError{value: val, isSet: true}
}

func (v NullableManageAttachToPrivateNetworkResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageAttachToPrivateNetworkResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


