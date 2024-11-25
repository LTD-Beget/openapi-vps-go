/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.6.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// ManageReinstallResponseErrorSoftwareVariableError struct for ManageReinstallResponseErrorSoftwareVariableError
type ManageReinstallResponseErrorSoftwareVariableError struct {
	Id *int32 `json:"id,omitempty"`
	Error []ManageReinstallResponseErrorSoftwareVariableErrorValueError `json:"error,omitempty"`
}

// NewManageReinstallResponseErrorSoftwareVariableError instantiates a new ManageReinstallResponseErrorSoftwareVariableError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageReinstallResponseErrorSoftwareVariableError() *ManageReinstallResponseErrorSoftwareVariableError {
	this := ManageReinstallResponseErrorSoftwareVariableError{}
	return &this
}

// NewManageReinstallResponseErrorSoftwareVariableErrorWithDefaults instantiates a new ManageReinstallResponseErrorSoftwareVariableError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageReinstallResponseErrorSoftwareVariableErrorWithDefaults() *ManageReinstallResponseErrorSoftwareVariableError {
	this := ManageReinstallResponseErrorSoftwareVariableError{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManageReinstallResponseErrorSoftwareVariableError) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageReinstallResponseErrorSoftwareVariableError) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManageReinstallResponseErrorSoftwareVariableError) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ManageReinstallResponseErrorSoftwareVariableError) SetId(v int32) {
	o.Id = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageReinstallResponseErrorSoftwareVariableError) GetError() []ManageReinstallResponseErrorSoftwareVariableErrorValueError {
	if o == nil || isNil(o.Error) {
		var ret []ManageReinstallResponseErrorSoftwareVariableErrorValueError
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageReinstallResponseErrorSoftwareVariableError) GetErrorOk() ([]ManageReinstallResponseErrorSoftwareVariableErrorValueError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageReinstallResponseErrorSoftwareVariableError) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []ManageReinstallResponseErrorSoftwareVariableErrorValueError and assigns it to the Error field.
func (o *ManageReinstallResponseErrorSoftwareVariableError) SetError(v []ManageReinstallResponseErrorSoftwareVariableErrorValueError) {
	o.Error = v
}

func (o ManageReinstallResponseErrorSoftwareVariableError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageReinstallResponseErrorSoftwareVariableError struct {
	value *ManageReinstallResponseErrorSoftwareVariableError
	isSet bool
}

func (v NullableManageReinstallResponseErrorSoftwareVariableError) Get() *ManageReinstallResponseErrorSoftwareVariableError {
	return v.value
}

func (v *NullableManageReinstallResponseErrorSoftwareVariableError) Set(val *ManageReinstallResponseErrorSoftwareVariableError) {
	v.value = val
	v.isSet = true
}

func (v NullableManageReinstallResponseErrorSoftwareVariableError) IsSet() bool {
	return v.isSet
}

func (v *NullableManageReinstallResponseErrorSoftwareVariableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageReinstallResponseErrorSoftwareVariableError(val *ManageReinstallResponseErrorSoftwareVariableError) *NullableManageReinstallResponseErrorSoftwareVariableError {
	return &NullableManageReinstallResponseErrorSoftwareVariableError{value: val, isSet: true}
}

func (v NullableManageReinstallResponseErrorSoftwareVariableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageReinstallResponseErrorSoftwareVariableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


