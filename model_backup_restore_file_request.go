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

// BackupRestoreFileRequest struct for BackupRestoreFileRequest
type BackupRestoreFileRequest struct {
	Id *string `json:"id,omitempty"`
	CopyId *string `json:"copy_id,omitempty"`
	Path []string `json:"path,omitempty"`
	AffectLive *bool `json:"affect_live,omitempty"`
	Target *string `json:"target,omitempty"`
}

// NewBackupRestoreFileRequest instantiates a new BackupRestoreFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupRestoreFileRequest() *BackupRestoreFileRequest {
	this := BackupRestoreFileRequest{}
	return &this
}

// NewBackupRestoreFileRequestWithDefaults instantiates a new BackupRestoreFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRestoreFileRequestWithDefaults() *BackupRestoreFileRequest {
	this := BackupRestoreFileRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BackupRestoreFileRequest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreFileRequest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BackupRestoreFileRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BackupRestoreFileRequest) SetId(v string) {
	o.Id = &v
}

// GetCopyId returns the CopyId field value if set, zero value otherwise.
func (o *BackupRestoreFileRequest) GetCopyId() string {
	if o == nil || isNil(o.CopyId) {
		var ret string
		return ret
	}
	return *o.CopyId
}

// GetCopyIdOk returns a tuple with the CopyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreFileRequest) GetCopyIdOk() (*string, bool) {
	if o == nil || isNil(o.CopyId) {
    return nil, false
	}
	return o.CopyId, true
}

// HasCopyId returns a boolean if a field has been set.
func (o *BackupRestoreFileRequest) HasCopyId() bool {
	if o != nil && !isNil(o.CopyId) {
		return true
	}

	return false
}

// SetCopyId gets a reference to the given string and assigns it to the CopyId field.
func (o *BackupRestoreFileRequest) SetCopyId(v string) {
	o.CopyId = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BackupRestoreFileRequest) GetPath() []string {
	if o == nil || isNil(o.Path) {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreFileRequest) GetPathOk() ([]string, bool) {
	if o == nil || isNil(o.Path) {
    return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BackupRestoreFileRequest) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *BackupRestoreFileRequest) SetPath(v []string) {
	o.Path = v
}

// GetAffectLive returns the AffectLive field value if set, zero value otherwise.
func (o *BackupRestoreFileRequest) GetAffectLive() bool {
	if o == nil || isNil(o.AffectLive) {
		var ret bool
		return ret
	}
	return *o.AffectLive
}

// GetAffectLiveOk returns a tuple with the AffectLive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreFileRequest) GetAffectLiveOk() (*bool, bool) {
	if o == nil || isNil(o.AffectLive) {
    return nil, false
	}
	return o.AffectLive, true
}

// HasAffectLive returns a boolean if a field has been set.
func (o *BackupRestoreFileRequest) HasAffectLive() bool {
	if o != nil && !isNil(o.AffectLive) {
		return true
	}

	return false
}

// SetAffectLive gets a reference to the given bool and assigns it to the AffectLive field.
func (o *BackupRestoreFileRequest) SetAffectLive(v bool) {
	o.AffectLive = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *BackupRestoreFileRequest) GetTarget() string {
	if o == nil || isNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreFileRequest) GetTargetOk() (*string, bool) {
	if o == nil || isNil(o.Target) {
    return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *BackupRestoreFileRequest) HasTarget() bool {
	if o != nil && !isNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *BackupRestoreFileRequest) SetTarget(v string) {
	o.Target = &v
}

func (o BackupRestoreFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.CopyId) {
		toSerialize["copy_id"] = o.CopyId
	}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.AffectLive) {
		toSerialize["affect_live"] = o.AffectLive
	}
	if !isNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableBackupRestoreFileRequest struct {
	value *BackupRestoreFileRequest
	isSet bool
}

func (v NullableBackupRestoreFileRequest) Get() *BackupRestoreFileRequest {
	return v.value
}

func (v *NullableBackupRestoreFileRequest) Set(val *BackupRestoreFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreFileRequest(val *BackupRestoreFileRequest) *NullableBackupRestoreFileRequest {
	return &NullableBackupRestoreFileRequest{value: val, isSet: true}
}

func (v NullableBackupRestoreFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


