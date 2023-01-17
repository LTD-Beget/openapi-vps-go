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

// BackupGetBackupFileListResponse struct for BackupGetBackupFileListResponse
type BackupGetBackupFileListResponse struct {
	File []StructuresFileInfo `json:"file,omitempty"`
}

// NewBackupGetBackupFileListResponse instantiates a new BackupGetBackupFileListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupGetBackupFileListResponse() *BackupGetBackupFileListResponse {
	this := BackupGetBackupFileListResponse{}
	return &this
}

// NewBackupGetBackupFileListResponseWithDefaults instantiates a new BackupGetBackupFileListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupGetBackupFileListResponseWithDefaults() *BackupGetBackupFileListResponse {
	this := BackupGetBackupFileListResponse{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *BackupGetBackupFileListResponse) GetFile() []StructuresFileInfo {
	if o == nil || isNil(o.File) {
		var ret []StructuresFileInfo
		return ret
	}
	return o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupGetBackupFileListResponse) GetFileOk() ([]StructuresFileInfo, bool) {
	if o == nil || isNil(o.File) {
    return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *BackupGetBackupFileListResponse) HasFile() bool {
	if o != nil && !isNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given []StructuresFileInfo and assigns it to the File field.
func (o *BackupGetBackupFileListResponse) SetFile(v []StructuresFileInfo) {
	o.File = v
}

func (o BackupGetBackupFileListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.File) {
		toSerialize["file"] = o.File
	}
	return json.Marshal(toSerialize)
}

type NullableBackupGetBackupFileListResponse struct {
	value *BackupGetBackupFileListResponse
	isSet bool
}

func (v NullableBackupGetBackupFileListResponse) Get() *BackupGetBackupFileListResponse {
	return v.value
}

func (v *NullableBackupGetBackupFileListResponse) Set(val *BackupGetBackupFileListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupGetBackupFileListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupGetBackupFileListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupGetBackupFileListResponse(val *BackupGetBackupFileListResponse) *NullableBackupGetBackupFileListResponse {
	return &NullableBackupGetBackupFileListResponse{value: val, isSet: true}
}

func (v NullableBackupGetBackupFileListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupGetBackupFileListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


