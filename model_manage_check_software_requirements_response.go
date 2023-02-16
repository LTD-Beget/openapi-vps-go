/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// ManageCheckSoftwareRequirementsResponse struct for ManageCheckSoftwareRequirementsResponse
type ManageCheckSoftwareRequirementsResponse struct {
	Error *ManageCheckSoftwareRequirementsResponseError `json:"error,omitempty"`
}

// NewManageCheckSoftwareRequirementsResponse instantiates a new ManageCheckSoftwareRequirementsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageCheckSoftwareRequirementsResponse() *ManageCheckSoftwareRequirementsResponse {
	this := ManageCheckSoftwareRequirementsResponse{}
	return &this
}

// NewManageCheckSoftwareRequirementsResponseWithDefaults instantiates a new ManageCheckSoftwareRequirementsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageCheckSoftwareRequirementsResponseWithDefaults() *ManageCheckSoftwareRequirementsResponse {
	this := ManageCheckSoftwareRequirementsResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ManageCheckSoftwareRequirementsResponse) GetError() ManageCheckSoftwareRequirementsResponseError {
	if o == nil || isNil(o.Error) {
		var ret ManageCheckSoftwareRequirementsResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCheckSoftwareRequirementsResponse) GetErrorOk() (*ManageCheckSoftwareRequirementsResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ManageCheckSoftwareRequirementsResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ManageCheckSoftwareRequirementsResponseError and assigns it to the Error field.
func (o *ManageCheckSoftwareRequirementsResponse) SetError(v ManageCheckSoftwareRequirementsResponseError) {
	o.Error = &v
}

func (o ManageCheckSoftwareRequirementsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableManageCheckSoftwareRequirementsResponse struct {
	value *ManageCheckSoftwareRequirementsResponse
	isSet bool
}

func (v NullableManageCheckSoftwareRequirementsResponse) Get() *ManageCheckSoftwareRequirementsResponse {
	return v.value
}

func (v *NullableManageCheckSoftwareRequirementsResponse) Set(val *ManageCheckSoftwareRequirementsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageCheckSoftwareRequirementsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageCheckSoftwareRequirementsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageCheckSoftwareRequirementsResponse(val *ManageCheckSoftwareRequirementsResponse) *NullableManageCheckSoftwareRequirementsResponse {
	return &NullableManageCheckSoftwareRequirementsResponse{value: val, isSet: true}
}

func (v NullableManageCheckSoftwareRequirementsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageCheckSoftwareRequirementsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


