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

// ManageChangeSshAccessResponse struct for ManageChangeSshAccessResponse
type ManageChangeSshAccessResponse struct {
	Vps *ManageVpsInfo `json:"vps,omitempty"`
	Error *ManageChangeSshAccessResponseError `json:"error,omitempty"`
}

// NewManageChangeSshAccessResponse instantiates a new ManageChangeSshAccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageChangeSshAccessResponse() *ManageChangeSshAccessResponse {
	this := ManageChangeSshAccessResponse{}
	return &this
}

// NewManageChangeSshAccessResponseWithDefaults instantiates a new ManageChangeSshAccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageChangeSshAccessResponseWithDefaults() *ManageChangeSshAccessResponse {
	this := ManageChangeSshAccessResponse{}
	return &this
}

// GetVps returns the Vps field value if set, zero value otherwise.
func (o *ManageChangeSshAccessResponse) GetVps() ManageVpsInfo {
	if o == nil || isNil(o.Vps) {
		var ret ManageVpsInfo
		return ret
	}
	return *o.Vps
}

// GetVpsOk returns a tuple with the Vps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageChangeSshAccessResponse) GetVpsOk() (*ManageVpsInfo, bool) {
	if o == nil || isNil(o.Vps) {
    return nil, false
	}
	return o.Vps, true
}

// HasVps returns a boolean if a field has been set.
func (o *ManageChangeSshAccessResponse) HasVps() bool {
	if o != nil && !isNil(o.Vps) {
		return true
	}

	return false
}

// SetVps gets a reference to the given ManageVpsInfo and assigns it to the Vps field.
func (o *ManageChangeSshAccessResponse) SetVps(v ManageVpsInfo) {
	o.Vps = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageChangeSshAccessResponse) GetError() ManageChangeSshAccessResponseError {
	if o == nil || isNil(o.Error) {
		var ret ManageChangeSshAccessResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageChangeSshAccessResponse) GetErrorOk() (*ManageChangeSshAccessResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageChangeSshAccessResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ManageChangeSshAccessResponseError and assigns it to the Error field.
func (o *ManageChangeSshAccessResponse) SetError(v ManageChangeSshAccessResponseError) {
	o.Error = &v
}

func (o ManageChangeSshAccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vps) {
		toSerialize["vps"] = o.Vps
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageChangeSshAccessResponse struct {
	value *ManageChangeSshAccessResponse
	isSet bool
}

func (v NullableManageChangeSshAccessResponse) Get() *ManageChangeSshAccessResponse {
	return v.value
}

func (v *NullableManageChangeSshAccessResponse) Set(val *ManageChangeSshAccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageChangeSshAccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageChangeSshAccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageChangeSshAccessResponse(val *ManageChangeSshAccessResponse) *NullableManageChangeSshAccessResponse {
	return &NullableManageChangeSshAccessResponse{value: val, isSet: true}
}

func (v NullableManageChangeSshAccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageChangeSshAccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


