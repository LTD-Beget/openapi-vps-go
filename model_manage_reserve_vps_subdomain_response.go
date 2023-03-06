/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.5.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// ManageReserveVpsSubdomainResponse struct for ManageReserveVpsSubdomainResponse
type ManageReserveVpsSubdomainResponse struct {
	VpsSubdomain *string `json:"vps_subdomain,omitempty"`
	Error *ManageReserveVpsSubdomainResponseError `json:"error,omitempty"`
}

// NewManageReserveVpsSubdomainResponse instantiates a new ManageReserveVpsSubdomainResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageReserveVpsSubdomainResponse() *ManageReserveVpsSubdomainResponse {
	this := ManageReserveVpsSubdomainResponse{}
	return &this
}

// NewManageReserveVpsSubdomainResponseWithDefaults instantiates a new ManageReserveVpsSubdomainResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageReserveVpsSubdomainResponseWithDefaults() *ManageReserveVpsSubdomainResponse {
	this := ManageReserveVpsSubdomainResponse{}
	return &this
}

// GetVpsSubdomain returns the VpsSubdomain field value if set, zero value otherwise.
func (o *ManageReserveVpsSubdomainResponse) GetVpsSubdomain() string {
	if o == nil || isNil(o.VpsSubdomain) {
		var ret string
		return ret
	}
	return *o.VpsSubdomain
}

// GetVpsSubdomainOk returns a tuple with the VpsSubdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageReserveVpsSubdomainResponse) GetVpsSubdomainOk() (*string, bool) {
	if o == nil || isNil(o.VpsSubdomain) {
    return nil, false
	}
	return o.VpsSubdomain, true
}

// HasVpsSubdomain returns a boolean if a field has been set.
func (o *ManageReserveVpsSubdomainResponse) HasVpsSubdomain() bool {
	if o != nil && !isNil(o.VpsSubdomain) {
		return true
	}

	return false
}

// SetVpsSubdomain gets a reference to the given string and assigns it to the VpsSubdomain field.
func (o *ManageReserveVpsSubdomainResponse) SetVpsSubdomain(v string) {
	o.VpsSubdomain = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageReserveVpsSubdomainResponse) GetError() ManageReserveVpsSubdomainResponseError {
	if o == nil || isNil(o.Error) {
		var ret ManageReserveVpsSubdomainResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageReserveVpsSubdomainResponse) GetErrorOk() (*ManageReserveVpsSubdomainResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageReserveVpsSubdomainResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ManageReserveVpsSubdomainResponseError and assigns it to the Error field.
func (o *ManageReserveVpsSubdomainResponse) SetError(v ManageReserveVpsSubdomainResponseError) {
	o.Error = &v
}

func (o ManageReserveVpsSubdomainResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.VpsSubdomain) {
		toSerialize["vps_subdomain"] = o.VpsSubdomain
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageReserveVpsSubdomainResponse struct {
	value *ManageReserveVpsSubdomainResponse
	isSet bool
}

func (v NullableManageReserveVpsSubdomainResponse) Get() *ManageReserveVpsSubdomainResponse {
	return v.value
}

func (v *NullableManageReserveVpsSubdomainResponse) Set(val *ManageReserveVpsSubdomainResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageReserveVpsSubdomainResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageReserveVpsSubdomainResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageReserveVpsSubdomainResponse(val *ManageReserveVpsSubdomainResponse) *NullableManageReserveVpsSubdomainResponse {
	return &NullableManageReserveVpsSubdomainResponse{value: val, isSet: true}
}

func (v NullableManageReserveVpsSubdomainResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageReserveVpsSubdomainResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


