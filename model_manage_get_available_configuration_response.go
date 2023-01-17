/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// ManageGetAvailableConfigurationResponse struct for ManageGetAvailableConfigurationResponse
type ManageGetAvailableConfigurationResponse struct {
	Configurations []ManageVpsConfiguration `json:"configurations,omitempty"`
}

// NewManageGetAvailableConfigurationResponse instantiates a new ManageGetAvailableConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageGetAvailableConfigurationResponse() *ManageGetAvailableConfigurationResponse {
	this := ManageGetAvailableConfigurationResponse{}
	return &this
}

// NewManageGetAvailableConfigurationResponseWithDefaults instantiates a new ManageGetAvailableConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageGetAvailableConfigurationResponseWithDefaults() *ManageGetAvailableConfigurationResponse {
	this := ManageGetAvailableConfigurationResponse{}
	return &this
}

// GetConfigurations returns the Configurations field value if set, zero value otherwise.
func (o *ManageGetAvailableConfigurationResponse) GetConfigurations() []ManageVpsConfiguration {
	if o == nil || isNil(o.Configurations) {
		var ret []ManageVpsConfiguration
		return ret
	}
	return o.Configurations
}

// GetConfigurationsOk returns a tuple with the Configurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetAvailableConfigurationResponse) GetConfigurationsOk() ([]ManageVpsConfiguration, bool) {
	if o == nil || isNil(o.Configurations) {
    return nil, false
	}
	return o.Configurations, true
}

// HasConfigurations returns a boolean if a field has been set.
func (o *ManageGetAvailableConfigurationResponse) HasConfigurations() bool {
	if o != nil && !isNil(o.Configurations) {
		return true
	}

	return false
}

// SetConfigurations gets a reference to the given []ManageVpsConfiguration and assigns it to the Configurations field.
func (o *ManageGetAvailableConfigurationResponse) SetConfigurations(v []ManageVpsConfiguration) {
	o.Configurations = v
}

func (o ManageGetAvailableConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Configurations) {
		toSerialize["configurations"] = o.Configurations
	}
	return json.Marshal(toSerialize)
}

type NullableManageGetAvailableConfigurationResponse struct {
	value *ManageGetAvailableConfigurationResponse
	isSet bool
}

func (v NullableManageGetAvailableConfigurationResponse) Get() *ManageGetAvailableConfigurationResponse {
	return v.value
}

func (v *NullableManageGetAvailableConfigurationResponse) Set(val *ManageGetAvailableConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageGetAvailableConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageGetAvailableConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageGetAvailableConfigurationResponse(val *ManageGetAvailableConfigurationResponse) *NullableManageGetAvailableConfigurationResponse {
	return &NullableManageGetAvailableConfigurationResponse{value: val, isSet: true}
}

func (v NullableManageGetAvailableConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageGetAvailableConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


