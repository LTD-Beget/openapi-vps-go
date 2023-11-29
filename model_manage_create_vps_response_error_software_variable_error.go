/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// ManageCreateVpsResponseErrorSoftwareVariableError struct for ManageCreateVpsResponseErrorSoftwareVariableError
type ManageCreateVpsResponseErrorSoftwareVariableError struct {
	Id *int32 `json:"id,omitempty"`
	Error []ManageCreateVpsResponseErrorSoftwareVariableErrorValueError `json:"error,omitempty"`
}

// NewManageCreateVpsResponseErrorSoftwareVariableError instantiates a new ManageCreateVpsResponseErrorSoftwareVariableError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageCreateVpsResponseErrorSoftwareVariableError() *ManageCreateVpsResponseErrorSoftwareVariableError {
	this := ManageCreateVpsResponseErrorSoftwareVariableError{}
	return &this
}

// NewManageCreateVpsResponseErrorSoftwareVariableErrorWithDefaults instantiates a new ManageCreateVpsResponseErrorSoftwareVariableError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageCreateVpsResponseErrorSoftwareVariableErrorWithDefaults() *ManageCreateVpsResponseErrorSoftwareVariableError {
	this := ManageCreateVpsResponseErrorSoftwareVariableError{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManageCreateVpsResponseErrorSoftwareVariableError) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsResponseErrorSoftwareVariableError) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManageCreateVpsResponseErrorSoftwareVariableError) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ManageCreateVpsResponseErrorSoftwareVariableError) SetId(v int32) {
	o.Id = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageCreateVpsResponseErrorSoftwareVariableError) GetError() []ManageCreateVpsResponseErrorSoftwareVariableErrorValueError {
	if o == nil || isNil(o.Error) {
		var ret []ManageCreateVpsResponseErrorSoftwareVariableErrorValueError
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsResponseErrorSoftwareVariableError) GetErrorOk() ([]ManageCreateVpsResponseErrorSoftwareVariableErrorValueError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageCreateVpsResponseErrorSoftwareVariableError) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []ManageCreateVpsResponseErrorSoftwareVariableErrorValueError and assigns it to the Error field.
func (o *ManageCreateVpsResponseErrorSoftwareVariableError) SetError(v []ManageCreateVpsResponseErrorSoftwareVariableErrorValueError) {
	o.Error = v
}

func (o ManageCreateVpsResponseErrorSoftwareVariableError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageCreateVpsResponseErrorSoftwareVariableError struct {
	value *ManageCreateVpsResponseErrorSoftwareVariableError
	isSet bool
}

func (v NullableManageCreateVpsResponseErrorSoftwareVariableError) Get() *ManageCreateVpsResponseErrorSoftwareVariableError {
	return v.value
}

func (v *NullableManageCreateVpsResponseErrorSoftwareVariableError) Set(val *ManageCreateVpsResponseErrorSoftwareVariableError) {
	v.value = val
	v.isSet = true
}

func (v NullableManageCreateVpsResponseErrorSoftwareVariableError) IsSet() bool {
	return v.isSet
}

func (v *NullableManageCreateVpsResponseErrorSoftwareVariableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageCreateVpsResponseErrorSoftwareVariableError(val *ManageCreateVpsResponseErrorSoftwareVariableError) *NullableManageCreateVpsResponseErrorSoftwareVariableError {
	return &NullableManageCreateVpsResponseErrorSoftwareVariableError{value: val, isSet: true}
}

func (v NullableManageCreateVpsResponseErrorSoftwareVariableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageCreateVpsResponseErrorSoftwareVariableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


