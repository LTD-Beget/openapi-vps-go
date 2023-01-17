/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// SnapshotEditRequest struct for SnapshotEditRequest
type SnapshotEditRequest struct {
	Id *string `json:"id,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewSnapshotEditRequest instantiates a new SnapshotEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotEditRequest() *SnapshotEditRequest {
	this := SnapshotEditRequest{}
	return &this
}

// NewSnapshotEditRequestWithDefaults instantiates a new SnapshotEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotEditRequestWithDefaults() *SnapshotEditRequest {
	this := SnapshotEditRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SnapshotEditRequest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotEditRequest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SnapshotEditRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SnapshotEditRequest) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SnapshotEditRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotEditRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SnapshotEditRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SnapshotEditRequest) SetDescription(v string) {
	o.Description = &v
}

func (o SnapshotEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotEditRequest struct {
	value *SnapshotEditRequest
	isSet bool
}

func (v NullableSnapshotEditRequest) Get() *SnapshotEditRequest {
	return v.value
}

func (v *NullableSnapshotEditRequest) Set(val *SnapshotEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotEditRequest(val *SnapshotEditRequest) *NullableSnapshotEditRequest {
	return &NullableSnapshotEditRequest{value: val, isSet: true}
}

func (v NullableSnapshotEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


