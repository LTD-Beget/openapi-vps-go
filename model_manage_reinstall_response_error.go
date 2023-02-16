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

// ManageReinstallResponseError struct for ManageReinstallResponseError
type ManageReinstallResponseError struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	VariableError *ManageReinstallResponseErrorSoftwareVariableError `json:"variable_error,omitempty"`
}

// NewManageReinstallResponseError instantiates a new ManageReinstallResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageReinstallResponseError() *ManageReinstallResponseError {
	this := ManageReinstallResponseError{}
	return &this
}

// NewManageReinstallResponseErrorWithDefaults instantiates a new ManageReinstallResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageReinstallResponseErrorWithDefaults() *ManageReinstallResponseError {
	this := ManageReinstallResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ManageReinstallResponseError) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageReinstallResponseError) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ManageReinstallResponseError) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ManageReinstallResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ManageReinstallResponseError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageReinstallResponseError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ManageReinstallResponseError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ManageReinstallResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetVariableError returns the VariableError field value if set, zero value otherwise.
func (o *ManageReinstallResponseError) GetVariableError() ManageReinstallResponseErrorSoftwareVariableError {
	if o == nil || isNil(o.VariableError) {
		var ret ManageReinstallResponseErrorSoftwareVariableError
		return ret
	}
	return *o.VariableError
}

// GetVariableErrorOk returns a tuple with the VariableError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageReinstallResponseError) GetVariableErrorOk() (*ManageReinstallResponseErrorSoftwareVariableError, bool) {
	if o == nil || isNil(o.VariableError) {
    return nil, false
	}
	return o.VariableError, true
}

// HasVariableError returns a boolean if a field has been set.
func (o *ManageReinstallResponseError) HasVariableError() bool {
	if o != nil && !isNil(o.VariableError) {
		return true
	}

	return false
}

// SetVariableError gets a reference to the given ManageReinstallResponseErrorSoftwareVariableError and assigns it to the VariableError field.
func (o *ManageReinstallResponseError) SetVariableError(v ManageReinstallResponseErrorSoftwareVariableError) {
	o.VariableError = &v
}

func (o ManageReinstallResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.VariableError) {
		toSerialize["variable_error"] = o.VariableError
	}
	return json.Marshal(toSerialize)
}

type NullableManageReinstallResponseError struct {
	value *ManageReinstallResponseError
	isSet bool
}

func (v NullableManageReinstallResponseError) Get() *ManageReinstallResponseError {
	return v.value
}

func (v *NullableManageReinstallResponseError) Set(val *ManageReinstallResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableManageReinstallResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableManageReinstallResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageReinstallResponseError(val *ManageReinstallResponseError) *NullableManageReinstallResponseError {
	return &NullableManageReinstallResponseError{value: val, isSet: true}
}

func (v NullableManageReinstallResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageReinstallResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


