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

// ManageStopVpsResponse struct for ManageStopVpsResponse
type ManageStopVpsResponse struct {
	Vps *ManageVpsInfo `json:"vps,omitempty"`
	Error *ManageStopVpsResponseError `json:"error,omitempty"`
}

// NewManageStopVpsResponse instantiates a new ManageStopVpsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageStopVpsResponse() *ManageStopVpsResponse {
	this := ManageStopVpsResponse{}
	return &this
}

// NewManageStopVpsResponseWithDefaults instantiates a new ManageStopVpsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageStopVpsResponseWithDefaults() *ManageStopVpsResponse {
	this := ManageStopVpsResponse{}
	return &this
}

// GetVps returns the Vps field value if set, zero value otherwise.
func (o *ManageStopVpsResponse) GetVps() ManageVpsInfo {
	if o == nil || isNil(o.Vps) {
		var ret ManageVpsInfo
		return ret
	}
	return *o.Vps
}

// GetVpsOk returns a tuple with the Vps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStopVpsResponse) GetVpsOk() (*ManageVpsInfo, bool) {
	if o == nil || isNil(o.Vps) {
    return nil, false
	}
	return o.Vps, true
}

// HasVps returns a boolean if a field has been set.
func (o *ManageStopVpsResponse) HasVps() bool {
	if o != nil && !isNil(o.Vps) {
		return true
	}

	return false
}

// SetVps gets a reference to the given ManageVpsInfo and assigns it to the Vps field.
func (o *ManageStopVpsResponse) SetVps(v ManageVpsInfo) {
	o.Vps = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageStopVpsResponse) GetError() ManageStopVpsResponseError {
	if o == nil || isNil(o.Error) {
		var ret ManageStopVpsResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStopVpsResponse) GetErrorOk() (*ManageStopVpsResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageStopVpsResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ManageStopVpsResponseError and assigns it to the Error field.
func (o *ManageStopVpsResponse) SetError(v ManageStopVpsResponseError) {
	o.Error = &v
}

func (o ManageStopVpsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vps) {
		toSerialize["vps"] = o.Vps
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageStopVpsResponse struct {
	value *ManageStopVpsResponse
	isSet bool
}

func (v NullableManageStopVpsResponse) Get() *ManageStopVpsResponse {
	return v.value
}

func (v *NullableManageStopVpsResponse) Set(val *ManageStopVpsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageStopVpsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageStopVpsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageStopVpsResponse(val *ManageStopVpsResponse) *NullableManageStopVpsResponse {
	return &NullableManageStopVpsResponse{value: val, isSet: true}
}

func (v NullableManageStopVpsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageStopVpsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


