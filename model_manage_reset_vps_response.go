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

// ManageResetVpsResponse struct for ManageResetVpsResponse
type ManageResetVpsResponse struct {
	Vps *ManageVpsInfo `json:"vps,omitempty"`
	Error *ManageResetVpsResponseError `json:"error,omitempty"`
}

// NewManageResetVpsResponse instantiates a new ManageResetVpsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageResetVpsResponse() *ManageResetVpsResponse {
	this := ManageResetVpsResponse{}
	return &this
}

// NewManageResetVpsResponseWithDefaults instantiates a new ManageResetVpsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageResetVpsResponseWithDefaults() *ManageResetVpsResponse {
	this := ManageResetVpsResponse{}
	return &this
}

// GetVps returns the Vps field value if set, zero value otherwise.
func (o *ManageResetVpsResponse) GetVps() ManageVpsInfo {
	if o == nil || isNil(o.Vps) {
		var ret ManageVpsInfo
		return ret
	}
	return *o.Vps
}

// GetVpsOk returns a tuple with the Vps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageResetVpsResponse) GetVpsOk() (*ManageVpsInfo, bool) {
	if o == nil || isNil(o.Vps) {
    return nil, false
	}
	return o.Vps, true
}

// HasVps returns a boolean if a field has been set.
func (o *ManageResetVpsResponse) HasVps() bool {
	if o != nil && !isNil(o.Vps) {
		return true
	}

	return false
}

// SetVps gets a reference to the given ManageVpsInfo and assigns it to the Vps field.
func (o *ManageResetVpsResponse) SetVps(v ManageVpsInfo) {
	o.Vps = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageResetVpsResponse) GetError() ManageResetVpsResponseError {
	if o == nil || isNil(o.Error) {
		var ret ManageResetVpsResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageResetVpsResponse) GetErrorOk() (*ManageResetVpsResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageResetVpsResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ManageResetVpsResponseError and assigns it to the Error field.
func (o *ManageResetVpsResponse) SetError(v ManageResetVpsResponseError) {
	o.Error = &v
}

func (o ManageResetVpsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vps) {
		toSerialize["vps"] = o.Vps
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageResetVpsResponse struct {
	value *ManageResetVpsResponse
	isSet bool
}

func (v NullableManageResetVpsResponse) Get() *ManageResetVpsResponse {
	return v.value
}

func (v *NullableManageResetVpsResponse) Set(val *ManageResetVpsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageResetVpsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageResetVpsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageResetVpsResponse(val *ManageResetVpsResponse) *NullableManageResetVpsResponse {
	return &NullableManageResetVpsResponse{value: val, isSet: true}
}

func (v NullableManageResetVpsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageResetVpsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

