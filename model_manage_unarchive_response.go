/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.5.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// ManageUnarchiveResponse struct for ManageUnarchiveResponse
type ManageUnarchiveResponse struct {
	Vps *ManageVpsInfo `json:"vps,omitempty"`
}

// NewManageUnarchiveResponse instantiates a new ManageUnarchiveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageUnarchiveResponse() *ManageUnarchiveResponse {
	this := ManageUnarchiveResponse{}
	return &this
}

// NewManageUnarchiveResponseWithDefaults instantiates a new ManageUnarchiveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageUnarchiveResponseWithDefaults() *ManageUnarchiveResponse {
	this := ManageUnarchiveResponse{}
	return &this
}

// GetVps returns the Vps field value if set, zero value otherwise.
func (o *ManageUnarchiveResponse) GetVps() ManageVpsInfo {
	if o == nil || isNil(o.Vps) {
		var ret ManageVpsInfo
		return ret
	}
	return *o.Vps
}

// GetVpsOk returns a tuple with the Vps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageUnarchiveResponse) GetVpsOk() (*ManageVpsInfo, bool) {
	if o == nil || isNil(o.Vps) {
    return nil, false
	}
	return o.Vps, true
}

// HasVps returns a boolean if a field has been set.
func (o *ManageUnarchiveResponse) HasVps() bool {
	if o != nil && !isNil(o.Vps) {
		return true
	}

	return false
}

// SetVps gets a reference to the given ManageVpsInfo and assigns it to the Vps field.
func (o *ManageUnarchiveResponse) SetVps(v ManageVpsInfo) {
	o.Vps = &v
}

func (o ManageUnarchiveResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vps) {
		toSerialize["vps"] = o.Vps
	}
	return json.Marshal(toSerialize)
}

type NullableManageUnarchiveResponse struct {
	value *ManageUnarchiveResponse
	isSet bool
}

func (v NullableManageUnarchiveResponse) Get() *ManageUnarchiveResponse {
	return v.value
}

func (v *NullableManageUnarchiveResponse) Set(val *ManageUnarchiveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageUnarchiveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageUnarchiveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageUnarchiveResponse(val *ManageUnarchiveResponse) *NullableManageUnarchiveResponse {
	return &NullableManageUnarchiveResponse{value: val, isSet: true}
}

func (v NullableManageUnarchiveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageUnarchiveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


