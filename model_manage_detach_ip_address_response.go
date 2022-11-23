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

// ManageDetachIpAddressResponse struct for ManageDetachIpAddressResponse
type ManageDetachIpAddressResponse struct {
	Vps *ManageVpsInfo `json:"vps,omitempty"`
	Error *ManageDetachIpAddressResponseError `json:"error,omitempty"`
}

// NewManageDetachIpAddressResponse instantiates a new ManageDetachIpAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageDetachIpAddressResponse() *ManageDetachIpAddressResponse {
	this := ManageDetachIpAddressResponse{}
	return &this
}

// NewManageDetachIpAddressResponseWithDefaults instantiates a new ManageDetachIpAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageDetachIpAddressResponseWithDefaults() *ManageDetachIpAddressResponse {
	this := ManageDetachIpAddressResponse{}
	return &this
}

// GetVps returns the Vps field value if set, zero value otherwise.
func (o *ManageDetachIpAddressResponse) GetVps() ManageVpsInfo {
	if o == nil || isNil(o.Vps) {
		var ret ManageVpsInfo
		return ret
	}
	return *o.Vps
}

// GetVpsOk returns a tuple with the Vps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageDetachIpAddressResponse) GetVpsOk() (*ManageVpsInfo, bool) {
	if o == nil || isNil(o.Vps) {
    return nil, false
	}
	return o.Vps, true
}

// HasVps returns a boolean if a field has been set.
func (o *ManageDetachIpAddressResponse) HasVps() bool {
	if o != nil && !isNil(o.Vps) {
		return true
	}

	return false
}

// SetVps gets a reference to the given ManageVpsInfo and assigns it to the Vps field.
func (o *ManageDetachIpAddressResponse) SetVps(v ManageVpsInfo) {
	o.Vps = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageDetachIpAddressResponse) GetError() ManageDetachIpAddressResponseError {
	if o == nil || isNil(o.Error) {
		var ret ManageDetachIpAddressResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageDetachIpAddressResponse) GetErrorOk() (*ManageDetachIpAddressResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageDetachIpAddressResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ManageDetachIpAddressResponseError and assigns it to the Error field.
func (o *ManageDetachIpAddressResponse) SetError(v ManageDetachIpAddressResponseError) {
	o.Error = &v
}

func (o ManageDetachIpAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vps) {
		toSerialize["vps"] = o.Vps
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageDetachIpAddressResponse struct {
	value *ManageDetachIpAddressResponse
	isSet bool
}

func (v NullableManageDetachIpAddressResponse) Get() *ManageDetachIpAddressResponse {
	return v.value
}

func (v *NullableManageDetachIpAddressResponse) Set(val *ManageDetachIpAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageDetachIpAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageDetachIpAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageDetachIpAddressResponse(val *ManageDetachIpAddressResponse) *NullableManageDetachIpAddressResponse {
	return &NullableManageDetachIpAddressResponse{value: val, isSet: true}
}

func (v NullableManageDetachIpAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageDetachIpAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


