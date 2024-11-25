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

// ManageReinstallResponse struct for ManageReinstallResponse
type ManageReinstallResponse struct {
	Vps *ManageVpsInfo `json:"vps,omitempty"`
	Error *ManageReinstallResponseError `json:"error,omitempty"`
}

// NewManageReinstallResponse instantiates a new ManageReinstallResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageReinstallResponse() *ManageReinstallResponse {
	this := ManageReinstallResponse{}
	return &this
}

// NewManageReinstallResponseWithDefaults instantiates a new ManageReinstallResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageReinstallResponseWithDefaults() *ManageReinstallResponse {
	this := ManageReinstallResponse{}
	return &this
}

// GetVps returns the Vps field value if set, zero value otherwise.
func (o *ManageReinstallResponse) GetVps() ManageVpsInfo {
	if o == nil || isNil(o.Vps) {
		var ret ManageVpsInfo
		return ret
	}
	return *o.Vps
}

// GetVpsOk returns a tuple with the Vps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageReinstallResponse) GetVpsOk() (*ManageVpsInfo, bool) {
	if o == nil || isNil(o.Vps) {
    return nil, false
	}
	return o.Vps, true
}

// HasVps returns a boolean if a field has been set.
func (o *ManageReinstallResponse) HasVps() bool {
	if o != nil && !isNil(o.Vps) {
		return true
	}

	return false
}

// SetVps gets a reference to the given ManageVpsInfo and assigns it to the Vps field.
func (o *ManageReinstallResponse) SetVps(v ManageVpsInfo) {
	o.Vps = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageReinstallResponse) GetError() ManageReinstallResponseError {
	if o == nil || isNil(o.Error) {
		var ret ManageReinstallResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageReinstallResponse) GetErrorOk() (*ManageReinstallResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageReinstallResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ManageReinstallResponseError and assigns it to the Error field.
func (o *ManageReinstallResponse) SetError(v ManageReinstallResponseError) {
	o.Error = &v
}

func (o ManageReinstallResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vps) {
		toSerialize["vps"] = o.Vps
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageReinstallResponse struct {
	value *ManageReinstallResponse
	isSet bool
}

func (v NullableManageReinstallResponse) Get() *ManageReinstallResponse {
	return v.value
}

func (v *NullableManageReinstallResponse) Set(val *ManageReinstallResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageReinstallResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageReinstallResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageReinstallResponse(val *ManageReinstallResponse) *NullableManageReinstallResponse {
	return &NullableManageReinstallResponse{value: val, isSet: true}
}

func (v NullableManageReinstallResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageReinstallResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


