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

// SnapshotRemoveResponse struct for SnapshotRemoveResponse
type SnapshotRemoveResponse struct {
	Snapshot *SnapshotSnapshot `json:"snapshot,omitempty"`
	Error *SnapshotRemoveResponseError `json:"error,omitempty"`
}

// NewSnapshotRemoveResponse instantiates a new SnapshotRemoveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotRemoveResponse() *SnapshotRemoveResponse {
	this := SnapshotRemoveResponse{}
	return &this
}

// NewSnapshotRemoveResponseWithDefaults instantiates a new SnapshotRemoveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotRemoveResponseWithDefaults() *SnapshotRemoveResponse {
	this := SnapshotRemoveResponse{}
	return &this
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *SnapshotRemoveResponse) GetSnapshot() SnapshotSnapshot {
	if o == nil || isNil(o.Snapshot) {
		var ret SnapshotSnapshot
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRemoveResponse) GetSnapshotOk() (*SnapshotSnapshot, bool) {
	if o == nil || isNil(o.Snapshot) {
    return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *SnapshotRemoveResponse) HasSnapshot() bool {
	if o != nil && !isNil(o.Snapshot) {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given SnapshotSnapshot and assigns it to the Snapshot field.
func (o *SnapshotRemoveResponse) SetSnapshot(v SnapshotSnapshot) {
	o.Snapshot = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SnapshotRemoveResponse) GetError() SnapshotRemoveResponseError {
	if o == nil || isNil(o.Error) {
		var ret SnapshotRemoveResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRemoveResponse) GetErrorOk() (*SnapshotRemoveResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SnapshotRemoveResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given SnapshotRemoveResponseError and assigns it to the Error field.
func (o *SnapshotRemoveResponse) SetError(v SnapshotRemoveResponseError) {
	o.Error = &v
}

func (o SnapshotRemoveResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Snapshot) {
		toSerialize["snapshot"] = o.Snapshot
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotRemoveResponse struct {
	value *SnapshotRemoveResponse
	isSet bool
}

func (v NullableSnapshotRemoveResponse) Get() *SnapshotRemoveResponse {
	return v.value
}

func (v *NullableSnapshotRemoveResponse) Set(val *SnapshotRemoveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotRemoveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotRemoveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotRemoveResponse(val *SnapshotRemoveResponse) *NullableSnapshotRemoveResponse {
	return &NullableSnapshotRemoveResponse{value: val, isSet: true}
}

func (v NullableSnapshotRemoveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotRemoveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


