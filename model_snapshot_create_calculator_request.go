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

// SnapshotCreateCalculatorRequest struct for SnapshotCreateCalculatorRequest
type SnapshotCreateCalculatorRequest struct {
	VpsId *string `json:"vps_id,omitempty"`
}

// NewSnapshotCreateCalculatorRequest instantiates a new SnapshotCreateCalculatorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotCreateCalculatorRequest() *SnapshotCreateCalculatorRequest {
	this := SnapshotCreateCalculatorRequest{}
	return &this
}

// NewSnapshotCreateCalculatorRequestWithDefaults instantiates a new SnapshotCreateCalculatorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotCreateCalculatorRequestWithDefaults() *SnapshotCreateCalculatorRequest {
	this := SnapshotCreateCalculatorRequest{}
	return &this
}

// GetVpsId returns the VpsId field value if set, zero value otherwise.
func (o *SnapshotCreateCalculatorRequest) GetVpsId() string {
	if o == nil || isNil(o.VpsId) {
		var ret string
		return ret
	}
	return *o.VpsId
}

// GetVpsIdOk returns a tuple with the VpsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotCreateCalculatorRequest) GetVpsIdOk() (*string, bool) {
	if o == nil || isNil(o.VpsId) {
    return nil, false
	}
	return o.VpsId, true
}

// HasVpsId returns a boolean if a field has been set.
func (o *SnapshotCreateCalculatorRequest) HasVpsId() bool {
	if o != nil && !isNil(o.VpsId) {
		return true
	}

	return false
}

// SetVpsId gets a reference to the given string and assigns it to the VpsId field.
func (o *SnapshotCreateCalculatorRequest) SetVpsId(v string) {
	o.VpsId = &v
}

func (o SnapshotCreateCalculatorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.VpsId) {
		toSerialize["vps_id"] = o.VpsId
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotCreateCalculatorRequest struct {
	value *SnapshotCreateCalculatorRequest
	isSet bool
}

func (v NullableSnapshotCreateCalculatorRequest) Get() *SnapshotCreateCalculatorRequest {
	return v.value
}

func (v *NullableSnapshotCreateCalculatorRequest) Set(val *SnapshotCreateCalculatorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotCreateCalculatorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotCreateCalculatorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotCreateCalculatorRequest(val *SnapshotCreateCalculatorRequest) *NullableSnapshotCreateCalculatorRequest {
	return &NullableSnapshotCreateCalculatorRequest{value: val, isSet: true}
}

func (v NullableSnapshotCreateCalculatorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotCreateCalculatorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


