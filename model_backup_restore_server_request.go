/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// BackupRestoreServerRequest struct for BackupRestoreServerRequest
type BackupRestoreServerRequest struct {
	Id *string `json:"id,omitempty"`
	CopyId *string `json:"copy_id,omitempty"`
}

// NewBackupRestoreServerRequest instantiates a new BackupRestoreServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupRestoreServerRequest() *BackupRestoreServerRequest {
	this := BackupRestoreServerRequest{}
	return &this
}

// NewBackupRestoreServerRequestWithDefaults instantiates a new BackupRestoreServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRestoreServerRequestWithDefaults() *BackupRestoreServerRequest {
	this := BackupRestoreServerRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BackupRestoreServerRequest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreServerRequest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BackupRestoreServerRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BackupRestoreServerRequest) SetId(v string) {
	o.Id = &v
}

// GetCopyId returns the CopyId field value if set, zero value otherwise.
func (o *BackupRestoreServerRequest) GetCopyId() string {
	if o == nil || isNil(o.CopyId) {
		var ret string
		return ret
	}
	return *o.CopyId
}

// GetCopyIdOk returns a tuple with the CopyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreServerRequest) GetCopyIdOk() (*string, bool) {
	if o == nil || isNil(o.CopyId) {
    return nil, false
	}
	return o.CopyId, true
}

// HasCopyId returns a boolean if a field has been set.
func (o *BackupRestoreServerRequest) HasCopyId() bool {
	if o != nil && !isNil(o.CopyId) {
		return true
	}

	return false
}

// SetCopyId gets a reference to the given string and assigns it to the CopyId field.
func (o *BackupRestoreServerRequest) SetCopyId(v string) {
	o.CopyId = &v
}

func (o BackupRestoreServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.CopyId) {
		toSerialize["copy_id"] = o.CopyId
	}
	return json.Marshal(toSerialize)
}

type NullableBackupRestoreServerRequest struct {
	value *BackupRestoreServerRequest
	isSet bool
}

func (v NullableBackupRestoreServerRequest) Get() *BackupRestoreServerRequest {
	return v.value
}

func (v *NullableBackupRestoreServerRequest) Set(val *BackupRestoreServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreServerRequest(val *BackupRestoreServerRequest) *NullableBackupRestoreServerRequest {
	return &NullableBackupRestoreServerRequest{value: val, isSet: true}
}

func (v NullableBackupRestoreServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


