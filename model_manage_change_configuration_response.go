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

// ManageChangeConfigurationResponse struct for ManageChangeConfigurationResponse
type ManageChangeConfigurationResponse struct {
	Vps *ManageVpsInfo `json:"vps,omitempty"`
	Error *ManageChangeConfigurationResponseError `json:"error,omitempty"`
}

// NewManageChangeConfigurationResponse instantiates a new ManageChangeConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageChangeConfigurationResponse() *ManageChangeConfigurationResponse {
	this := ManageChangeConfigurationResponse{}
	return &this
}

// NewManageChangeConfigurationResponseWithDefaults instantiates a new ManageChangeConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageChangeConfigurationResponseWithDefaults() *ManageChangeConfigurationResponse {
	this := ManageChangeConfigurationResponse{}
	return &this
}

// GetVps returns the Vps field value if set, zero value otherwise.
func (o *ManageChangeConfigurationResponse) GetVps() ManageVpsInfo {
	if o == nil || isNil(o.Vps) {
		var ret ManageVpsInfo
		return ret
	}
	return *o.Vps
}

// GetVpsOk returns a tuple with the Vps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageChangeConfigurationResponse) GetVpsOk() (*ManageVpsInfo, bool) {
	if o == nil || isNil(o.Vps) {
    return nil, false
	}
	return o.Vps, true
}

// HasVps returns a boolean if a field has been set.
func (o *ManageChangeConfigurationResponse) HasVps() bool {
	if o != nil && !isNil(o.Vps) {
		return true
	}

	return false
}

// SetVps gets a reference to the given ManageVpsInfo and assigns it to the Vps field.
func (o *ManageChangeConfigurationResponse) SetVps(v ManageVpsInfo) {
	o.Vps = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageChangeConfigurationResponse) GetError() ManageChangeConfigurationResponseError {
	if o == nil || isNil(o.Error) {
		var ret ManageChangeConfigurationResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageChangeConfigurationResponse) GetErrorOk() (*ManageChangeConfigurationResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageChangeConfigurationResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ManageChangeConfigurationResponseError and assigns it to the Error field.
func (o *ManageChangeConfigurationResponse) SetError(v ManageChangeConfigurationResponseError) {
	o.Error = &v
}

func (o ManageChangeConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vps) {
		toSerialize["vps"] = o.Vps
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageChangeConfigurationResponse struct {
	value *ManageChangeConfigurationResponse
	isSet bool
}

func (v NullableManageChangeConfigurationResponse) Get() *ManageChangeConfigurationResponse {
	return v.value
}

func (v *NullableManageChangeConfigurationResponse) Set(val *ManageChangeConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageChangeConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageChangeConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageChangeConfigurationResponse(val *ManageChangeConfigurationResponse) *NullableManageChangeConfigurationResponse {
	return &NullableManageChangeConfigurationResponse{value: val, isSet: true}
}

func (v NullableManageChangeConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageChangeConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


