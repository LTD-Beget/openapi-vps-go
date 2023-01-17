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

// ManageResetPasswordResponse struct for ManageResetPasswordResponse
type ManageResetPasswordResponse struct {
	Vps *ManageVpsInfo `json:"vps,omitempty"`
	Error *ManageResetPasswordResponseError `json:"error,omitempty"`
}

// NewManageResetPasswordResponse instantiates a new ManageResetPasswordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageResetPasswordResponse() *ManageResetPasswordResponse {
	this := ManageResetPasswordResponse{}
	return &this
}

// NewManageResetPasswordResponseWithDefaults instantiates a new ManageResetPasswordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageResetPasswordResponseWithDefaults() *ManageResetPasswordResponse {
	this := ManageResetPasswordResponse{}
	return &this
}

// GetVps returns the Vps field value if set, zero value otherwise.
func (o *ManageResetPasswordResponse) GetVps() ManageVpsInfo {
	if o == nil || isNil(o.Vps) {
		var ret ManageVpsInfo
		return ret
	}
	return *o.Vps
}

// GetVpsOk returns a tuple with the Vps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageResetPasswordResponse) GetVpsOk() (*ManageVpsInfo, bool) {
	if o == nil || isNil(o.Vps) {
    return nil, false
	}
	return o.Vps, true
}

// HasVps returns a boolean if a field has been set.
func (o *ManageResetPasswordResponse) HasVps() bool {
	if o != nil && !isNil(o.Vps) {
		return true
	}

	return false
}

// SetVps gets a reference to the given ManageVpsInfo and assigns it to the Vps field.
func (o *ManageResetPasswordResponse) SetVps(v ManageVpsInfo) {
	o.Vps = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageResetPasswordResponse) GetError() ManageResetPasswordResponseError {
	if o == nil || isNil(o.Error) {
		var ret ManageResetPasswordResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageResetPasswordResponse) GetErrorOk() (*ManageResetPasswordResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageResetPasswordResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ManageResetPasswordResponseError and assigns it to the Error field.
func (o *ManageResetPasswordResponse) SetError(v ManageResetPasswordResponseError) {
	o.Error = &v
}

func (o ManageResetPasswordResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vps) {
		toSerialize["vps"] = o.Vps
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageResetPasswordResponse struct {
	value *ManageResetPasswordResponse
	isSet bool
}

func (v NullableManageResetPasswordResponse) Get() *ManageResetPasswordResponse {
	return v.value
}

func (v *NullableManageResetPasswordResponse) Set(val *ManageResetPasswordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageResetPasswordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageResetPasswordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageResetPasswordResponse(val *ManageResetPasswordResponse) *NullableManageResetPasswordResponse {
	return &NullableManageResetPasswordResponse{value: val, isSet: true}
}

func (v NullableManageResetPasswordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageResetPasswordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


