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

// ManageRebootVpsResponse struct for ManageRebootVpsResponse
type ManageRebootVpsResponse struct {
	Vps *ManageVpsInfo `json:"vps,omitempty"`
	Error *ManageRebootVpsResponseError `json:"error,omitempty"`
}

// NewManageRebootVpsResponse instantiates a new ManageRebootVpsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageRebootVpsResponse() *ManageRebootVpsResponse {
	this := ManageRebootVpsResponse{}
	return &this
}

// NewManageRebootVpsResponseWithDefaults instantiates a new ManageRebootVpsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageRebootVpsResponseWithDefaults() *ManageRebootVpsResponse {
	this := ManageRebootVpsResponse{}
	return &this
}

// GetVps returns the Vps field value if set, zero value otherwise.
func (o *ManageRebootVpsResponse) GetVps() ManageVpsInfo {
	if o == nil || isNil(o.Vps) {
		var ret ManageVpsInfo
		return ret
	}
	return *o.Vps
}

// GetVpsOk returns a tuple with the Vps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageRebootVpsResponse) GetVpsOk() (*ManageVpsInfo, bool) {
	if o == nil || isNil(o.Vps) {
    return nil, false
	}
	return o.Vps, true
}

// HasVps returns a boolean if a field has been set.
func (o *ManageRebootVpsResponse) HasVps() bool {
	if o != nil && !isNil(o.Vps) {
		return true
	}

	return false
}

// SetVps gets a reference to the given ManageVpsInfo and assigns it to the Vps field.
func (o *ManageRebootVpsResponse) SetVps(v ManageVpsInfo) {
	o.Vps = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageRebootVpsResponse) GetError() ManageRebootVpsResponseError {
	if o == nil || isNil(o.Error) {
		var ret ManageRebootVpsResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageRebootVpsResponse) GetErrorOk() (*ManageRebootVpsResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageRebootVpsResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ManageRebootVpsResponseError and assigns it to the Error field.
func (o *ManageRebootVpsResponse) SetError(v ManageRebootVpsResponseError) {
	o.Error = &v
}

func (o ManageRebootVpsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vps) {
		toSerialize["vps"] = o.Vps
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageRebootVpsResponse struct {
	value *ManageRebootVpsResponse
	isSet bool
}

func (v NullableManageRebootVpsResponse) Get() *ManageRebootVpsResponse {
	return v.value
}

func (v *NullableManageRebootVpsResponse) Set(val *ManageRebootVpsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageRebootVpsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageRebootVpsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageRebootVpsResponse(val *ManageRebootVpsResponse) *NullableManageRebootVpsResponse {
	return &NullableManageRebootVpsResponse{value: val, isSet: true}
}

func (v NullableManageRebootVpsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageRebootVpsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


