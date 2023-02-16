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

// SnapshotRemoveResponseError struct for SnapshotRemoveResponseError
type SnapshotRemoveResponseError struct {
	Message *string `json:"message,omitempty"`
	Code *string `json:"code,omitempty"`
}

// NewSnapshotRemoveResponseError instantiates a new SnapshotRemoveResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotRemoveResponseError() *SnapshotRemoveResponseError {
	this := SnapshotRemoveResponseError{}
	return &this
}

// NewSnapshotRemoveResponseErrorWithDefaults instantiates a new SnapshotRemoveResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotRemoveResponseErrorWithDefaults() *SnapshotRemoveResponseError {
	this := SnapshotRemoveResponseError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SnapshotRemoveResponseError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRemoveResponseError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SnapshotRemoveResponseError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SnapshotRemoveResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SnapshotRemoveResponseError) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRemoveResponseError) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SnapshotRemoveResponseError) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SnapshotRemoveResponseError) SetCode(v string) {
	o.Code = &v
}

func (o SnapshotRemoveResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotRemoveResponseError struct {
	value *SnapshotRemoveResponseError
	isSet bool
}

func (v NullableSnapshotRemoveResponseError) Get() *SnapshotRemoveResponseError {
	return v.value
}

func (v *NullableSnapshotRemoveResponseError) Set(val *SnapshotRemoveResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotRemoveResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotRemoveResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotRemoveResponseError(val *SnapshotRemoveResponseError) *NullableSnapshotRemoveResponseError {
	return &NullableSnapshotRemoveResponseError{value: val, isSet: true}
}

func (v NullableSnapshotRemoveResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotRemoveResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


