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

// ManageCheckSoftwareRequirementsResponseError struct for ManageCheckSoftwareRequirementsResponseError
type ManageCheckSoftwareRequirementsResponseError struct {
	WordpressError *string `json:"wordpress_error,omitempty"`
	DomainError *string `json:"domain_error,omitempty"`
}

// NewManageCheckSoftwareRequirementsResponseError instantiates a new ManageCheckSoftwareRequirementsResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageCheckSoftwareRequirementsResponseError() *ManageCheckSoftwareRequirementsResponseError {
	this := ManageCheckSoftwareRequirementsResponseError{}
	return &this
}

// NewManageCheckSoftwareRequirementsResponseErrorWithDefaults instantiates a new ManageCheckSoftwareRequirementsResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageCheckSoftwareRequirementsResponseErrorWithDefaults() *ManageCheckSoftwareRequirementsResponseError {
	this := ManageCheckSoftwareRequirementsResponseError{}
	return &this
}

// GetWordpressError returns the WordpressError field value if set, zero value otherwise.
func (o *ManageCheckSoftwareRequirementsResponseError) GetWordpressError() string {
	if o == nil || isNil(o.WordpressError) {
		var ret string
		return ret
	}
	return *o.WordpressError
}

// GetWordpressErrorOk returns a tuple with the WordpressError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCheckSoftwareRequirementsResponseError) GetWordpressErrorOk() (*string, bool) {
	if o == nil || isNil(o.WordpressError) {
    return nil, false
	}
	return o.WordpressError, true
}

// HasWordpressError returns a boolean if a field has been set.
func (o *ManageCheckSoftwareRequirementsResponseError) HasWordpressError() bool {
	if o != nil && !isNil(o.WordpressError) {
		return true
	}

	return false
}

// SetWordpressError gets a reference to the given string and assigns it to the WordpressError field.
func (o *ManageCheckSoftwareRequirementsResponseError) SetWordpressError(v string) {
	o.WordpressError = &v
}

// GetDomainError returns the DomainError field value if set, zero value otherwise.
func (o *ManageCheckSoftwareRequirementsResponseError) GetDomainError() string {
	if o == nil || isNil(o.DomainError) {
		var ret string
		return ret
	}
	return *o.DomainError
}

// GetDomainErrorOk returns a tuple with the DomainError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCheckSoftwareRequirementsResponseError) GetDomainErrorOk() (*string, bool) {
	if o == nil || isNil(o.DomainError) {
    return nil, false
	}
	return o.DomainError, true
}

// HasDomainError returns a boolean if a field has been set.
func (o *ManageCheckSoftwareRequirementsResponseError) HasDomainError() bool {
	if o != nil && !isNil(o.DomainError) {
		return true
	}

	return false
}

// SetDomainError gets a reference to the given string and assigns it to the DomainError field.
func (o *ManageCheckSoftwareRequirementsResponseError) SetDomainError(v string) {
	o.DomainError = &v
}

func (o ManageCheckSoftwareRequirementsResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WordpressError) {
		toSerialize["wordpress_error"] = o.WordpressError
	}
	if !isNil(o.DomainError) {
		toSerialize["domain_error"] = o.DomainError
	}
	return json.Marshal(toSerialize)
}

type NullableManageCheckSoftwareRequirementsResponseError struct {
	value *ManageCheckSoftwareRequirementsResponseError
	isSet bool
}

func (v NullableManageCheckSoftwareRequirementsResponseError) Get() *ManageCheckSoftwareRequirementsResponseError {
	return v.value
}

func (v *NullableManageCheckSoftwareRequirementsResponseError) Set(val *ManageCheckSoftwareRequirementsResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableManageCheckSoftwareRequirementsResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableManageCheckSoftwareRequirementsResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageCheckSoftwareRequirementsResponseError(val *ManageCheckSoftwareRequirementsResponseError) *NullableManageCheckSoftwareRequirementsResponseError {
	return &NullableManageCheckSoftwareRequirementsResponseError{value: val, isSet: true}
}

func (v NullableManageCheckSoftwareRequirementsResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageCheckSoftwareRequirementsResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

