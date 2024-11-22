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

// ManageGetFileManagerSettingsResponseCredentials struct for ManageGetFileManagerSettingsResponseCredentials
type ManageGetFileManagerSettingsResponseCredentials struct {
	Type *string `json:"type,omitempty"`
	User *string `json:"user,omitempty"`
	Host *string `json:"host,omitempty"`
	ConnectionId *int32 `json:"connection_id,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewManageGetFileManagerSettingsResponseCredentials instantiates a new ManageGetFileManagerSettingsResponseCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageGetFileManagerSettingsResponseCredentials() *ManageGetFileManagerSettingsResponseCredentials {
	this := ManageGetFileManagerSettingsResponseCredentials{}
	return &this
}

// NewManageGetFileManagerSettingsResponseCredentialsWithDefaults instantiates a new ManageGetFileManagerSettingsResponseCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageGetFileManagerSettingsResponseCredentialsWithDefaults() *ManageGetFileManagerSettingsResponseCredentials {
	this := ManageGetFileManagerSettingsResponseCredentials{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ManageGetFileManagerSettingsResponseCredentials) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetFileManagerSettingsResponseCredentials) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ManageGetFileManagerSettingsResponseCredentials) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ManageGetFileManagerSettingsResponseCredentials) SetType(v string) {
	o.Type = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ManageGetFileManagerSettingsResponseCredentials) GetUser() string {
	if o == nil || isNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetFileManagerSettingsResponseCredentials) GetUserOk() (*string, bool) {
	if o == nil || isNil(o.User) {
    return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ManageGetFileManagerSettingsResponseCredentials) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ManageGetFileManagerSettingsResponseCredentials) SetUser(v string) {
	o.User = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ManageGetFileManagerSettingsResponseCredentials) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetFileManagerSettingsResponseCredentials) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ManageGetFileManagerSettingsResponseCredentials) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ManageGetFileManagerSettingsResponseCredentials) SetHost(v string) {
	o.Host = &v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *ManageGetFileManagerSettingsResponseCredentials) GetConnectionId() int32 {
	if o == nil || isNil(o.ConnectionId) {
		var ret int32
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetFileManagerSettingsResponseCredentials) GetConnectionIdOk() (*int32, bool) {
	if o == nil || isNil(o.ConnectionId) {
    return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *ManageGetFileManagerSettingsResponseCredentials) HasConnectionId() bool {
	if o != nil && !isNil(o.ConnectionId) {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given int32 and assigns it to the ConnectionId field.
func (o *ManageGetFileManagerSettingsResponseCredentials) SetConnectionId(v int32) {
	o.ConnectionId = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ManageGetFileManagerSettingsResponseCredentials) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetFileManagerSettingsResponseCredentials) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
    return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ManageGetFileManagerSettingsResponseCredentials) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ManageGetFileManagerSettingsResponseCredentials) SetPath(v string) {
	o.Path = &v
}

func (o ManageGetFileManagerSettingsResponseCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.ConnectionId) {
		toSerialize["connection_id"] = o.ConnectionId
	}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableManageGetFileManagerSettingsResponseCredentials struct {
	value *ManageGetFileManagerSettingsResponseCredentials
	isSet bool
}

func (v NullableManageGetFileManagerSettingsResponseCredentials) Get() *ManageGetFileManagerSettingsResponseCredentials {
	return v.value
}

func (v *NullableManageGetFileManagerSettingsResponseCredentials) Set(val *ManageGetFileManagerSettingsResponseCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableManageGetFileManagerSettingsResponseCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableManageGetFileManagerSettingsResponseCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageGetFileManagerSettingsResponseCredentials(val *ManageGetFileManagerSettingsResponseCredentials) *NullableManageGetFileManagerSettingsResponseCredentials {
	return &NullableManageGetFileManagerSettingsResponseCredentials{value: val, isSet: true}
}

func (v NullableManageGetFileManagerSettingsResponseCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageGetFileManagerSettingsResponseCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


