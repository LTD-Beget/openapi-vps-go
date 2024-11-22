/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.6.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// ManageStartVpsResponse struct for ManageStartVpsResponse
type ManageStartVpsResponse struct {
	Vps *ManageVpsInfo `json:"vps,omitempty"`
	Error *ManageStartVpsResponseError `json:"error,omitempty"`
}

// NewManageStartVpsResponse instantiates a new ManageStartVpsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageStartVpsResponse() *ManageStartVpsResponse {
	this := ManageStartVpsResponse{}
	return &this
}

// NewManageStartVpsResponseWithDefaults instantiates a new ManageStartVpsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageStartVpsResponseWithDefaults() *ManageStartVpsResponse {
	this := ManageStartVpsResponse{}
	return &this
}

// GetVps returns the Vps field value if set, zero value otherwise.
func (o *ManageStartVpsResponse) GetVps() ManageVpsInfo {
	if o == nil || isNil(o.Vps) {
		var ret ManageVpsInfo
		return ret
	}
	return *o.Vps
}

// GetVpsOk returns a tuple with the Vps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStartVpsResponse) GetVpsOk() (*ManageVpsInfo, bool) {
	if o == nil || isNil(o.Vps) {
    return nil, false
	}
	return o.Vps, true
}

// HasVps returns a boolean if a field has been set.
func (o *ManageStartVpsResponse) HasVps() bool {
	if o != nil && !isNil(o.Vps) {
		return true
	}

	return false
}

// SetVps gets a reference to the given ManageVpsInfo and assigns it to the Vps field.
func (o *ManageStartVpsResponse) SetVps(v ManageVpsInfo) {
	o.Vps = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageStartVpsResponse) GetError() ManageStartVpsResponseError {
	if o == nil || isNil(o.Error) {
		var ret ManageStartVpsResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStartVpsResponse) GetErrorOk() (*ManageStartVpsResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageStartVpsResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ManageStartVpsResponseError and assigns it to the Error field.
func (o *ManageStartVpsResponse) SetError(v ManageStartVpsResponseError) {
	o.Error = &v
}

func (o ManageStartVpsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vps) {
		toSerialize["vps"] = o.Vps
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageStartVpsResponse struct {
	value *ManageStartVpsResponse
	isSet bool
}

func (v NullableManageStartVpsResponse) Get() *ManageStartVpsResponse {
	return v.value
}

func (v *NullableManageStartVpsResponse) Set(val *ManageStartVpsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageStartVpsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageStartVpsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageStartVpsResponse(val *ManageStartVpsResponse) *NullableManageStartVpsResponse {
	return &NullableManageStartVpsResponse{value: val, isSet: true}
}

func (v NullableManageStartVpsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageStartVpsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


