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

// StructuresOrderInfoErrorDetails struct for StructuresOrderInfoErrorDetails
type StructuresOrderInfoErrorDetails struct {
	FileError *StructuresOrderInfoErrorDetailsFileError `json:"file_error,omitempty"`
}

// NewStructuresOrderInfoErrorDetails instantiates a new StructuresOrderInfoErrorDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresOrderInfoErrorDetails() *StructuresOrderInfoErrorDetails {
	this := StructuresOrderInfoErrorDetails{}
	return &this
}

// NewStructuresOrderInfoErrorDetailsWithDefaults instantiates a new StructuresOrderInfoErrorDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresOrderInfoErrorDetailsWithDefaults() *StructuresOrderInfoErrorDetails {
	this := StructuresOrderInfoErrorDetails{}
	return &this
}

// GetFileError returns the FileError field value if set, zero value otherwise.
func (o *StructuresOrderInfoErrorDetails) GetFileError() StructuresOrderInfoErrorDetailsFileError {
	if o == nil || isNil(o.FileError) {
		var ret StructuresOrderInfoErrorDetailsFileError
		return ret
	}
	return *o.FileError
}

// GetFileErrorOk returns a tuple with the FileError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOrderInfoErrorDetails) GetFileErrorOk() (*StructuresOrderInfoErrorDetailsFileError, bool) {
	if o == nil || isNil(o.FileError) {
    return nil, false
	}
	return o.FileError, true
}

// HasFileError returns a boolean if a field has been set.
func (o *StructuresOrderInfoErrorDetails) HasFileError() bool {
	if o != nil && !isNil(o.FileError) {
		return true
	}

	return false
}

// SetFileError gets a reference to the given StructuresOrderInfoErrorDetailsFileError and assigns it to the FileError field.
func (o *StructuresOrderInfoErrorDetails) SetFileError(v StructuresOrderInfoErrorDetailsFileError) {
	o.FileError = &v
}

func (o StructuresOrderInfoErrorDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FileError) {
		toSerialize["file_error"] = o.FileError
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresOrderInfoErrorDetails struct {
	value *StructuresOrderInfoErrorDetails
	isSet bool
}

func (v NullableStructuresOrderInfoErrorDetails) Get() *StructuresOrderInfoErrorDetails {
	return v.value
}

func (v *NullableStructuresOrderInfoErrorDetails) Set(val *StructuresOrderInfoErrorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresOrderInfoErrorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresOrderInfoErrorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresOrderInfoErrorDetails(val *StructuresOrderInfoErrorDetails) *NullableStructuresOrderInfoErrorDetails {
	return &NullableStructuresOrderInfoErrorDetails{value: val, isSet: true}
}

func (v NullableStructuresOrderInfoErrorDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresOrderInfoErrorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


