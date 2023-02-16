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

// ManageGetInfoResponse struct for ManageGetInfoResponse
type ManageGetInfoResponse struct {
	Vps *ManageVpsInfo `json:"vps,omitempty"`
}

// NewManageGetInfoResponse instantiates a new ManageGetInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageGetInfoResponse() *ManageGetInfoResponse {
	this := ManageGetInfoResponse{}
	return &this
}

// NewManageGetInfoResponseWithDefaults instantiates a new ManageGetInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageGetInfoResponseWithDefaults() *ManageGetInfoResponse {
	this := ManageGetInfoResponse{}
	return &this
}

// GetVps returns the Vps field value if set, zero value otherwise.
func (o *ManageGetInfoResponse) GetVps() ManageVpsInfo {
	if o == nil || isNil(o.Vps) {
		var ret ManageVpsInfo
		return ret
	}
	return *o.Vps
}

// GetVpsOk returns a tuple with the Vps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetInfoResponse) GetVpsOk() (*ManageVpsInfo, bool) {
	if o == nil || isNil(o.Vps) {
    return nil, false
	}
	return o.Vps, true
}

// HasVps returns a boolean if a field has been set.
func (o *ManageGetInfoResponse) HasVps() bool {
	if o != nil && !isNil(o.Vps) {
		return true
	}

	return false
}

// SetVps gets a reference to the given ManageVpsInfo and assigns it to the Vps field.
func (o *ManageGetInfoResponse) SetVps(v ManageVpsInfo) {
	o.Vps = &v
}

func (o ManageGetInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vps) {
		toSerialize["vps"] = o.Vps
	}
	return json.Marshal(toSerialize)
}

type NullableManageGetInfoResponse struct {
	value *ManageGetInfoResponse
	isSet bool
}

func (v NullableManageGetInfoResponse) Get() *ManageGetInfoResponse {
	return v.value
}

func (v *NullableManageGetInfoResponse) Set(val *ManageGetInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageGetInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageGetInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageGetInfoResponse(val *ManageGetInfoResponse) *NullableManageGetInfoResponse {
	return &NullableManageGetInfoResponse{value: val, isSet: true}
}

func (v NullableManageGetInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageGetInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


