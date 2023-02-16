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

// ManageChangeSshAccessResponseError struct for ManageChangeSshAccessResponseError
type ManageChangeSshAccessResponseError struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Extra *string `json:"extra,omitempty"`
}

// NewManageChangeSshAccessResponseError instantiates a new ManageChangeSshAccessResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageChangeSshAccessResponseError() *ManageChangeSshAccessResponseError {
	this := ManageChangeSshAccessResponseError{}
	return &this
}

// NewManageChangeSshAccessResponseErrorWithDefaults instantiates a new ManageChangeSshAccessResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageChangeSshAccessResponseErrorWithDefaults() *ManageChangeSshAccessResponseError {
	this := ManageChangeSshAccessResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ManageChangeSshAccessResponseError) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageChangeSshAccessResponseError) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ManageChangeSshAccessResponseError) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ManageChangeSshAccessResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ManageChangeSshAccessResponseError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageChangeSshAccessResponseError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ManageChangeSshAccessResponseError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ManageChangeSshAccessResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *ManageChangeSshAccessResponseError) GetExtra() string {
	if o == nil || isNil(o.Extra) {
		var ret string
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageChangeSshAccessResponseError) GetExtraOk() (*string, bool) {
	if o == nil || isNil(o.Extra) {
    return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *ManageChangeSshAccessResponseError) HasExtra() bool {
	if o != nil && !isNil(o.Extra) {
		return true
	}

	return false
}

// SetExtra gets a reference to the given string and assigns it to the Extra field.
func (o *ManageChangeSshAccessResponseError) SetExtra(v string) {
	o.Extra = &v
}

func (o ManageChangeSshAccessResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Extra) {
		toSerialize["extra"] = o.Extra
	}
	return json.Marshal(toSerialize)
}

type NullableManageChangeSshAccessResponseError struct {
	value *ManageChangeSshAccessResponseError
	isSet bool
}

func (v NullableManageChangeSshAccessResponseError) Get() *ManageChangeSshAccessResponseError {
	return v.value
}

func (v *NullableManageChangeSshAccessResponseError) Set(val *ManageChangeSshAccessResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableManageChangeSshAccessResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableManageChangeSshAccessResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageChangeSshAccessResponseError(val *ManageChangeSshAccessResponseError) *NullableManageChangeSshAccessResponseError {
	return &NullableManageChangeSshAccessResponseError{value: val, isSet: true}
}

func (v NullableManageChangeSshAccessResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageChangeSshAccessResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


