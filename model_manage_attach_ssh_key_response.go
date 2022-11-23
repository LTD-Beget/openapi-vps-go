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

// ManageAttachSshKeyResponse struct for ManageAttachSshKeyResponse
type ManageAttachSshKeyResponse struct {
	Vps *ManageVpsInfo `json:"vps,omitempty"`
	Error *ManageAttachSshKeyResponseError `json:"error,omitempty"`
}

// NewManageAttachSshKeyResponse instantiates a new ManageAttachSshKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageAttachSshKeyResponse() *ManageAttachSshKeyResponse {
	this := ManageAttachSshKeyResponse{}
	return &this
}

// NewManageAttachSshKeyResponseWithDefaults instantiates a new ManageAttachSshKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageAttachSshKeyResponseWithDefaults() *ManageAttachSshKeyResponse {
	this := ManageAttachSshKeyResponse{}
	return &this
}

// GetVps returns the Vps field value if set, zero value otherwise.
func (o *ManageAttachSshKeyResponse) GetVps() ManageVpsInfo {
	if o == nil || isNil(o.Vps) {
		var ret ManageVpsInfo
		return ret
	}
	return *o.Vps
}

// GetVpsOk returns a tuple with the Vps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageAttachSshKeyResponse) GetVpsOk() (*ManageVpsInfo, bool) {
	if o == nil || isNil(o.Vps) {
    return nil, false
	}
	return o.Vps, true
}

// HasVps returns a boolean if a field has been set.
func (o *ManageAttachSshKeyResponse) HasVps() bool {
	if o != nil && !isNil(o.Vps) {
		return true
	}

	return false
}

// SetVps gets a reference to the given ManageVpsInfo and assigns it to the Vps field.
func (o *ManageAttachSshKeyResponse) SetVps(v ManageVpsInfo) {
	o.Vps = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageAttachSshKeyResponse) GetError() ManageAttachSshKeyResponseError {
	if o == nil || isNil(o.Error) {
		var ret ManageAttachSshKeyResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageAttachSshKeyResponse) GetErrorOk() (*ManageAttachSshKeyResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageAttachSshKeyResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ManageAttachSshKeyResponseError and assigns it to the Error field.
func (o *ManageAttachSshKeyResponse) SetError(v ManageAttachSshKeyResponseError) {
	o.Error = &v
}

func (o ManageAttachSshKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vps) {
		toSerialize["vps"] = o.Vps
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageAttachSshKeyResponse struct {
	value *ManageAttachSshKeyResponse
	isSet bool
}

func (v NullableManageAttachSshKeyResponse) Get() *ManageAttachSshKeyResponse {
	return v.value
}

func (v *NullableManageAttachSshKeyResponse) Set(val *ManageAttachSshKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageAttachSshKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageAttachSshKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageAttachSshKeyResponse(val *ManageAttachSshKeyResponse) *NullableManageAttachSshKeyResponse {
	return &NullableManageAttachSshKeyResponse{value: val, isSet: true}
}

func (v NullableManageAttachSshKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageAttachSshKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

