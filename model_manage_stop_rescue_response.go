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

// ManageStopRescueResponse struct for ManageStopRescueResponse
type ManageStopRescueResponse struct {
	Vps *ManageVpsInfo `json:"vps,omitempty"`
	Error *ManageStopRescueResponseError `json:"error,omitempty"`
}

// NewManageStopRescueResponse instantiates a new ManageStopRescueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageStopRescueResponse() *ManageStopRescueResponse {
	this := ManageStopRescueResponse{}
	return &this
}

// NewManageStopRescueResponseWithDefaults instantiates a new ManageStopRescueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageStopRescueResponseWithDefaults() *ManageStopRescueResponse {
	this := ManageStopRescueResponse{}
	return &this
}

// GetVps returns the Vps field value if set, zero value otherwise.
func (o *ManageStopRescueResponse) GetVps() ManageVpsInfo {
	if o == nil || isNil(o.Vps) {
		var ret ManageVpsInfo
		return ret
	}
	return *o.Vps
}

// GetVpsOk returns a tuple with the Vps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStopRescueResponse) GetVpsOk() (*ManageVpsInfo, bool) {
	if o == nil || isNil(o.Vps) {
    return nil, false
	}
	return o.Vps, true
}

// HasVps returns a boolean if a field has been set.
func (o *ManageStopRescueResponse) HasVps() bool {
	if o != nil && !isNil(o.Vps) {
		return true
	}

	return false
}

// SetVps gets a reference to the given ManageVpsInfo and assigns it to the Vps field.
func (o *ManageStopRescueResponse) SetVps(v ManageVpsInfo) {
	o.Vps = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageStopRescueResponse) GetError() ManageStopRescueResponseError {
	if o == nil || isNil(o.Error) {
		var ret ManageStopRescueResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStopRescueResponse) GetErrorOk() (*ManageStopRescueResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageStopRescueResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ManageStopRescueResponseError and assigns it to the Error field.
func (o *ManageStopRescueResponse) SetError(v ManageStopRescueResponseError) {
	o.Error = &v
}

func (o ManageStopRescueResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vps) {
		toSerialize["vps"] = o.Vps
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageStopRescueResponse struct {
	value *ManageStopRescueResponse
	isSet bool
}

func (v NullableManageStopRescueResponse) Get() *ManageStopRescueResponse {
	return v.value
}

func (v *NullableManageStopRescueResponse) Set(val *ManageStopRescueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageStopRescueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageStopRescueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageStopRescueResponse(val *ManageStopRescueResponse) *NullableManageStopRescueResponse {
	return &NullableManageStopRescueResponse{value: val, isSet: true}
}

func (v NullableManageStopRescueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageStopRescueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


