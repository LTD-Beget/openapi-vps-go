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

// ManageGetListResponse struct for ManageGetListResponse
type ManageGetListResponse struct {
	Vps []ManageVpsInfo `json:"vps,omitempty"`
}

// NewManageGetListResponse instantiates a new ManageGetListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageGetListResponse() *ManageGetListResponse {
	this := ManageGetListResponse{}
	return &this
}

// NewManageGetListResponseWithDefaults instantiates a new ManageGetListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageGetListResponseWithDefaults() *ManageGetListResponse {
	this := ManageGetListResponse{}
	return &this
}

// GetVps returns the Vps field value if set, zero value otherwise.
func (o *ManageGetListResponse) GetVps() []ManageVpsInfo {
	if o == nil || isNil(o.Vps) {
		var ret []ManageVpsInfo
		return ret
	}
	return o.Vps
}

// GetVpsOk returns a tuple with the Vps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetListResponse) GetVpsOk() ([]ManageVpsInfo, bool) {
	if o == nil || isNil(o.Vps) {
    return nil, false
	}
	return o.Vps, true
}

// HasVps returns a boolean if a field has been set.
func (o *ManageGetListResponse) HasVps() bool {
	if o != nil && !isNil(o.Vps) {
		return true
	}

	return false
}

// SetVps gets a reference to the given []ManageVpsInfo and assigns it to the Vps field.
func (o *ManageGetListResponse) SetVps(v []ManageVpsInfo) {
	o.Vps = v
}

func (o ManageGetListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vps) {
		toSerialize["vps"] = o.Vps
	}
	return json.Marshal(toSerialize)
}

type NullableManageGetListResponse struct {
	value *ManageGetListResponse
	isSet bool
}

func (v NullableManageGetListResponse) Get() *ManageGetListResponse {
	return v.value
}

func (v *NullableManageGetListResponse) Set(val *ManageGetListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageGetListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageGetListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageGetListResponse(val *ManageGetListResponse) *NullableManageGetListResponse {
	return &NullableManageGetListResponse{value: val, isSet: true}
}

func (v NullableManageGetListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageGetListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


