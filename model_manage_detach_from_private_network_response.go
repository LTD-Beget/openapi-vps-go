/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// ManageDetachFromPrivateNetworkResponse struct for ManageDetachFromPrivateNetworkResponse
type ManageDetachFromPrivateNetworkResponse struct {
	Vps *ManageVpsInfo `json:"vps,omitempty"`
	Error *ManageDetachFromPrivateNetworkResponseError `json:"error,omitempty"`
}

// NewManageDetachFromPrivateNetworkResponse instantiates a new ManageDetachFromPrivateNetworkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageDetachFromPrivateNetworkResponse() *ManageDetachFromPrivateNetworkResponse {
	this := ManageDetachFromPrivateNetworkResponse{}
	return &this
}

// NewManageDetachFromPrivateNetworkResponseWithDefaults instantiates a new ManageDetachFromPrivateNetworkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageDetachFromPrivateNetworkResponseWithDefaults() *ManageDetachFromPrivateNetworkResponse {
	this := ManageDetachFromPrivateNetworkResponse{}
	return &this
}

// GetVps returns the Vps field value if set, zero value otherwise.
func (o *ManageDetachFromPrivateNetworkResponse) GetVps() ManageVpsInfo {
	if o == nil || isNil(o.Vps) {
		var ret ManageVpsInfo
		return ret
	}
	return *o.Vps
}

// GetVpsOk returns a tuple with the Vps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageDetachFromPrivateNetworkResponse) GetVpsOk() (*ManageVpsInfo, bool) {
	if o == nil || isNil(o.Vps) {
    return nil, false
	}
	return o.Vps, true
}

// HasVps returns a boolean if a field has been set.
func (o *ManageDetachFromPrivateNetworkResponse) HasVps() bool {
	if o != nil && !isNil(o.Vps) {
		return true
	}

	return false
}

// SetVps gets a reference to the given ManageVpsInfo and assigns it to the Vps field.
func (o *ManageDetachFromPrivateNetworkResponse) SetVps(v ManageVpsInfo) {
	o.Vps = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageDetachFromPrivateNetworkResponse) GetError() ManageDetachFromPrivateNetworkResponseError {
	if o == nil || isNil(o.Error) {
		var ret ManageDetachFromPrivateNetworkResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageDetachFromPrivateNetworkResponse) GetErrorOk() (*ManageDetachFromPrivateNetworkResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageDetachFromPrivateNetworkResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ManageDetachFromPrivateNetworkResponseError and assigns it to the Error field.
func (o *ManageDetachFromPrivateNetworkResponse) SetError(v ManageDetachFromPrivateNetworkResponseError) {
	o.Error = &v
}

func (o ManageDetachFromPrivateNetworkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vps) {
		toSerialize["vps"] = o.Vps
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageDetachFromPrivateNetworkResponse struct {
	value *ManageDetachFromPrivateNetworkResponse
	isSet bool
}

func (v NullableManageDetachFromPrivateNetworkResponse) Get() *ManageDetachFromPrivateNetworkResponse {
	return v.value
}

func (v *NullableManageDetachFromPrivateNetworkResponse) Set(val *ManageDetachFromPrivateNetworkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageDetachFromPrivateNetworkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageDetachFromPrivateNetworkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageDetachFromPrivateNetworkResponse(val *ManageDetachFromPrivateNetworkResponse) *NullableManageDetachFromPrivateNetworkResponse {
	return &NullableManageDetachFromPrivateNetworkResponse{value: val, isSet: true}
}

func (v NullableManageDetachFromPrivateNetworkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageDetachFromPrivateNetworkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


