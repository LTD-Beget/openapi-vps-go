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

// ManageAttachIpAddressResponse struct for ManageAttachIpAddressResponse
type ManageAttachIpAddressResponse struct {
	Vps *ManageVpsInfo `json:"vps,omitempty"`
	Error *ManageAttachIpAddressResponseError `json:"error,omitempty"`
}

// NewManageAttachIpAddressResponse instantiates a new ManageAttachIpAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageAttachIpAddressResponse() *ManageAttachIpAddressResponse {
	this := ManageAttachIpAddressResponse{}
	return &this
}

// NewManageAttachIpAddressResponseWithDefaults instantiates a new ManageAttachIpAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageAttachIpAddressResponseWithDefaults() *ManageAttachIpAddressResponse {
	this := ManageAttachIpAddressResponse{}
	return &this
}

// GetVps returns the Vps field value if set, zero value otherwise.
func (o *ManageAttachIpAddressResponse) GetVps() ManageVpsInfo {
	if o == nil || isNil(o.Vps) {
		var ret ManageVpsInfo
		return ret
	}
	return *o.Vps
}

// GetVpsOk returns a tuple with the Vps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageAttachIpAddressResponse) GetVpsOk() (*ManageVpsInfo, bool) {
	if o == nil || isNil(o.Vps) {
    return nil, false
	}
	return o.Vps, true
}

// HasVps returns a boolean if a field has been set.
func (o *ManageAttachIpAddressResponse) HasVps() bool {
	if o != nil && !isNil(o.Vps) {
		return true
	}

	return false
}

// SetVps gets a reference to the given ManageVpsInfo and assigns it to the Vps field.
func (o *ManageAttachIpAddressResponse) SetVps(v ManageVpsInfo) {
	o.Vps = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageAttachIpAddressResponse) GetError() ManageAttachIpAddressResponseError {
	if o == nil || isNil(o.Error) {
		var ret ManageAttachIpAddressResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageAttachIpAddressResponse) GetErrorOk() (*ManageAttachIpAddressResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageAttachIpAddressResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ManageAttachIpAddressResponseError and assigns it to the Error field.
func (o *ManageAttachIpAddressResponse) SetError(v ManageAttachIpAddressResponseError) {
	o.Error = &v
}

func (o ManageAttachIpAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vps) {
		toSerialize["vps"] = o.Vps
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageAttachIpAddressResponse struct {
	value *ManageAttachIpAddressResponse
	isSet bool
}

func (v NullableManageAttachIpAddressResponse) Get() *ManageAttachIpAddressResponse {
	return v.value
}

func (v *NullableManageAttachIpAddressResponse) Set(val *ManageAttachIpAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageAttachIpAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageAttachIpAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageAttachIpAddressResponse(val *ManageAttachIpAddressResponse) *NullableManageAttachIpAddressResponse {
	return &NullableManageAttachIpAddressResponse{value: val, isSet: true}
}

func (v NullableManageAttachIpAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageAttachIpAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


