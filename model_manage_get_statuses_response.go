/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// ManageGetStatusesResponse struct for ManageGetStatusesResponse
type ManageGetStatusesResponse struct {
	Statuses []ManageGetStatusesResponseStatusInfo `json:"statuses,omitempty"`
}

// NewManageGetStatusesResponse instantiates a new ManageGetStatusesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageGetStatusesResponse() *ManageGetStatusesResponse {
	this := ManageGetStatusesResponse{}
	return &this
}

// NewManageGetStatusesResponseWithDefaults instantiates a new ManageGetStatusesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageGetStatusesResponseWithDefaults() *ManageGetStatusesResponse {
	this := ManageGetStatusesResponse{}
	return &this
}

// GetStatuses returns the Statuses field value if set, zero value otherwise.
func (o *ManageGetStatusesResponse) GetStatuses() []ManageGetStatusesResponseStatusInfo {
	if o == nil || isNil(o.Statuses) {
		var ret []ManageGetStatusesResponseStatusInfo
		return ret
	}
	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetStatusesResponse) GetStatusesOk() ([]ManageGetStatusesResponseStatusInfo, bool) {
	if o == nil || isNil(o.Statuses) {
    return nil, false
	}
	return o.Statuses, true
}

// HasStatuses returns a boolean if a field has been set.
func (o *ManageGetStatusesResponse) HasStatuses() bool {
	if o != nil && !isNil(o.Statuses) {
		return true
	}

	return false
}

// SetStatuses gets a reference to the given []ManageGetStatusesResponseStatusInfo and assigns it to the Statuses field.
func (o *ManageGetStatusesResponse) SetStatuses(v []ManageGetStatusesResponseStatusInfo) {
	o.Statuses = v
}

func (o ManageGetStatusesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Statuses) {
		toSerialize["statuses"] = o.Statuses
	}
	return json.Marshal(toSerialize)
}

type NullableManageGetStatusesResponse struct {
	value *ManageGetStatusesResponse
	isSet bool
}

func (v NullableManageGetStatusesResponse) Get() *ManageGetStatusesResponse {
	return v.value
}

func (v *NullableManageGetStatusesResponse) Set(val *ManageGetStatusesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageGetStatusesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageGetStatusesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageGetStatusesResponse(val *ManageGetStatusesResponse) *NullableManageGetStatusesResponse {
	return &NullableManageGetStatusesResponse{value: val, isSet: true}
}

func (v NullableManageGetStatusesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageGetStatusesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


