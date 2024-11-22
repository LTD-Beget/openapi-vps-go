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

// ManageCheckSoftwareRequirementsRequest struct for ManageCheckSoftwareRequirementsRequest
type ManageCheckSoftwareRequirementsRequest struct {
	Info *ManageSoftwareInstallInfo `json:"info,omitempty"`
}

// NewManageCheckSoftwareRequirementsRequest instantiates a new ManageCheckSoftwareRequirementsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageCheckSoftwareRequirementsRequest() *ManageCheckSoftwareRequirementsRequest {
	this := ManageCheckSoftwareRequirementsRequest{}
	return &this
}

// NewManageCheckSoftwareRequirementsRequestWithDefaults instantiates a new ManageCheckSoftwareRequirementsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageCheckSoftwareRequirementsRequestWithDefaults() *ManageCheckSoftwareRequirementsRequest {
	this := ManageCheckSoftwareRequirementsRequest{}
	return &this
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ManageCheckSoftwareRequirementsRequest) GetInfo() ManageSoftwareInstallInfo {
	if o == nil || isNil(o.Info) {
		var ret ManageSoftwareInstallInfo
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCheckSoftwareRequirementsRequest) GetInfoOk() (*ManageSoftwareInstallInfo, bool) {
	if o == nil || isNil(o.Info) {
    return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ManageCheckSoftwareRequirementsRequest) HasInfo() bool {
	if o != nil && !isNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given ManageSoftwareInstallInfo and assigns it to the Info field.
func (o *ManageCheckSoftwareRequirementsRequest) SetInfo(v ManageSoftwareInstallInfo) {
	o.Info = &v
}

func (o ManageCheckSoftwareRequirementsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	return json.Marshal(toSerialize)
}

type NullableManageCheckSoftwareRequirementsRequest struct {
	value *ManageCheckSoftwareRequirementsRequest
	isSet bool
}

func (v NullableManageCheckSoftwareRequirementsRequest) Get() *ManageCheckSoftwareRequirementsRequest {
	return v.value
}

func (v *NullableManageCheckSoftwareRequirementsRequest) Set(val *ManageCheckSoftwareRequirementsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableManageCheckSoftwareRequirementsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableManageCheckSoftwareRequirementsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageCheckSoftwareRequirementsRequest(val *ManageCheckSoftwareRequirementsRequest) *NullableManageCheckSoftwareRequirementsRequest {
	return &NullableManageCheckSoftwareRequirementsRequest{value: val, isSet: true}
}

func (v NullableManageCheckSoftwareRequirementsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageCheckSoftwareRequirementsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


