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

// SnapshotEditResponse struct for SnapshotEditResponse
type SnapshotEditResponse struct {
	Snapshot *SnapshotSnapshot `json:"snapshot,omitempty"`
}

// NewSnapshotEditResponse instantiates a new SnapshotEditResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotEditResponse() *SnapshotEditResponse {
	this := SnapshotEditResponse{}
	return &this
}

// NewSnapshotEditResponseWithDefaults instantiates a new SnapshotEditResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotEditResponseWithDefaults() *SnapshotEditResponse {
	this := SnapshotEditResponse{}
	return &this
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *SnapshotEditResponse) GetSnapshot() SnapshotSnapshot {
	if o == nil || isNil(o.Snapshot) {
		var ret SnapshotSnapshot
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotEditResponse) GetSnapshotOk() (*SnapshotSnapshot, bool) {
	if o == nil || isNil(o.Snapshot) {
    return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *SnapshotEditResponse) HasSnapshot() bool {
	if o != nil && !isNil(o.Snapshot) {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given SnapshotSnapshot and assigns it to the Snapshot field.
func (o *SnapshotEditResponse) SetSnapshot(v SnapshotSnapshot) {
	o.Snapshot = &v
}

func (o SnapshotEditResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Snapshot) {
		toSerialize["snapshot"] = o.Snapshot
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotEditResponse struct {
	value *SnapshotEditResponse
	isSet bool
}

func (v NullableSnapshotEditResponse) Get() *SnapshotEditResponse {
	return v.value
}

func (v *NullableSnapshotEditResponse) Set(val *SnapshotEditResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotEditResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotEditResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotEditResponse(val *SnapshotEditResponse) *NullableSnapshotEditResponse {
	return &NullableSnapshotEditResponse{value: val, isSet: true}
}

func (v NullableSnapshotEditResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotEditResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


