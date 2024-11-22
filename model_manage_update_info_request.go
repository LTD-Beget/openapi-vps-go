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

// ManageUpdateInfoRequest struct for ManageUpdateInfoRequest
type ManageUpdateInfoRequest struct {
	Id *string `json:"id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewManageUpdateInfoRequest instantiates a new ManageUpdateInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageUpdateInfoRequest() *ManageUpdateInfoRequest {
	this := ManageUpdateInfoRequest{}
	return &this
}

// NewManageUpdateInfoRequestWithDefaults instantiates a new ManageUpdateInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageUpdateInfoRequestWithDefaults() *ManageUpdateInfoRequest {
	this := ManageUpdateInfoRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManageUpdateInfoRequest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageUpdateInfoRequest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManageUpdateInfoRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ManageUpdateInfoRequest) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ManageUpdateInfoRequest) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageUpdateInfoRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ManageUpdateInfoRequest) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ManageUpdateInfoRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ManageUpdateInfoRequest) GetHostname() string {
	if o == nil || isNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageUpdateInfoRequest) GetHostnameOk() (*string, bool) {
	if o == nil || isNil(o.Hostname) {
    return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ManageUpdateInfoRequest) HasHostname() bool {
	if o != nil && !isNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ManageUpdateInfoRequest) SetHostname(v string) {
	o.Hostname = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ManageUpdateInfoRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageUpdateInfoRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ManageUpdateInfoRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ManageUpdateInfoRequest) SetDescription(v string) {
	o.Description = &v
}

func (o ManageUpdateInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !isNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableManageUpdateInfoRequest struct {
	value *ManageUpdateInfoRequest
	isSet bool
}

func (v NullableManageUpdateInfoRequest) Get() *ManageUpdateInfoRequest {
	return v.value
}

func (v *NullableManageUpdateInfoRequest) Set(val *ManageUpdateInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableManageUpdateInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableManageUpdateInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageUpdateInfoRequest(val *ManageUpdateInfoRequest) *NullableManageUpdateInfoRequest {
	return &NullableManageUpdateInfoRequest{value: val, isSet: true}
}

func (v NullableManageUpdateInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageUpdateInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


