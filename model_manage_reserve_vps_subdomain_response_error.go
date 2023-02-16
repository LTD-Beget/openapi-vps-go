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

// ManageReserveVpsSubdomainResponseError struct for ManageReserveVpsSubdomainResponseError
type ManageReserveVpsSubdomainResponseError struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewManageReserveVpsSubdomainResponseError instantiates a new ManageReserveVpsSubdomainResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageReserveVpsSubdomainResponseError() *ManageReserveVpsSubdomainResponseError {
	this := ManageReserveVpsSubdomainResponseError{}
	return &this
}

// NewManageReserveVpsSubdomainResponseErrorWithDefaults instantiates a new ManageReserveVpsSubdomainResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageReserveVpsSubdomainResponseErrorWithDefaults() *ManageReserveVpsSubdomainResponseError {
	this := ManageReserveVpsSubdomainResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ManageReserveVpsSubdomainResponseError) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageReserveVpsSubdomainResponseError) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ManageReserveVpsSubdomainResponseError) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ManageReserveVpsSubdomainResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ManageReserveVpsSubdomainResponseError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageReserveVpsSubdomainResponseError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ManageReserveVpsSubdomainResponseError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ManageReserveVpsSubdomainResponseError) SetMessage(v string) {
	o.Message = &v
}

func (o ManageReserveVpsSubdomainResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableManageReserveVpsSubdomainResponseError struct {
	value *ManageReserveVpsSubdomainResponseError
	isSet bool
}

func (v NullableManageReserveVpsSubdomainResponseError) Get() *ManageReserveVpsSubdomainResponseError {
	return v.value
}

func (v *NullableManageReserveVpsSubdomainResponseError) Set(val *ManageReserveVpsSubdomainResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableManageReserveVpsSubdomainResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableManageReserveVpsSubdomainResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageReserveVpsSubdomainResponseError(val *ManageReserveVpsSubdomainResponseError) *NullableManageReserveVpsSubdomainResponseError {
	return &NullableManageReserveVpsSubdomainResponseError{value: val, isSet: true}
}

func (v NullableManageReserveVpsSubdomainResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageReserveVpsSubdomainResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


