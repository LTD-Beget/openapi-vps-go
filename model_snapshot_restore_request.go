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

// SnapshotRestoreRequest struct for SnapshotRestoreRequest
type SnapshotRestoreRequest struct {
	Id *string `json:"id,omitempty"`
	VpsId *string `json:"vps_id,omitempty"`
}

// NewSnapshotRestoreRequest instantiates a new SnapshotRestoreRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotRestoreRequest() *SnapshotRestoreRequest {
	this := SnapshotRestoreRequest{}
	return &this
}

// NewSnapshotRestoreRequestWithDefaults instantiates a new SnapshotRestoreRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotRestoreRequestWithDefaults() *SnapshotRestoreRequest {
	this := SnapshotRestoreRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SnapshotRestoreRequest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRestoreRequest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SnapshotRestoreRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SnapshotRestoreRequest) SetId(v string) {
	o.Id = &v
}

// GetVpsId returns the VpsId field value if set, zero value otherwise.
func (o *SnapshotRestoreRequest) GetVpsId() string {
	if o == nil || isNil(o.VpsId) {
		var ret string
		return ret
	}
	return *o.VpsId
}

// GetVpsIdOk returns a tuple with the VpsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRestoreRequest) GetVpsIdOk() (*string, bool) {
	if o == nil || isNil(o.VpsId) {
    return nil, false
	}
	return o.VpsId, true
}

// HasVpsId returns a boolean if a field has been set.
func (o *SnapshotRestoreRequest) HasVpsId() bool {
	if o != nil && !isNil(o.VpsId) {
		return true
	}

	return false
}

// SetVpsId gets a reference to the given string and assigns it to the VpsId field.
func (o *SnapshotRestoreRequest) SetVpsId(v string) {
	o.VpsId = &v
}

func (o SnapshotRestoreRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.VpsId) {
		toSerialize["vps_id"] = o.VpsId
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotRestoreRequest struct {
	value *SnapshotRestoreRequest
	isSet bool
}

func (v NullableSnapshotRestoreRequest) Get() *SnapshotRestoreRequest {
	return v.value
}

func (v *NullableSnapshotRestoreRequest) Set(val *SnapshotRestoreRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotRestoreRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotRestoreRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotRestoreRequest(val *SnapshotRestoreRequest) *NullableSnapshotRestoreRequest {
	return &NullableSnapshotRestoreRequest{value: val, isSet: true}
}

func (v NullableSnapshotRestoreRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotRestoreRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

