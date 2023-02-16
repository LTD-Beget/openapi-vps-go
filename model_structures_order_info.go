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

// StructuresOrderInfo struct for StructuresOrderInfo
type StructuresOrderInfo struct {
	Id *string `json:"id,omitempty"`
	VpsId *string `json:"vps_id,omitempty"`
	VpsName *string `json:"vps_name,omitempty"`
	Type *string `json:"type,omitempty"`
	DateCreate *string `json:"date_create,omitempty"`
	DateComplete *string `json:"date_complete,omitempty"`
	Path []string `json:"path,omitempty"`
	Status *string `json:"status,omitempty"`
	CopyInfo *StructuresCopyInfo `json:"copy_info,omitempty"`
	AffectLive *bool `json:"affect_live,omitempty"`
	Target *string `json:"target,omitempty"`
	ErrorDetails *StructuresOrderInfoErrorDetails `json:"error_details,omitempty"`
}

// NewStructuresOrderInfo instantiates a new StructuresOrderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresOrderInfo() *StructuresOrderInfo {
	this := StructuresOrderInfo{}
	return &this
}

// NewStructuresOrderInfoWithDefaults instantiates a new StructuresOrderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresOrderInfoWithDefaults() *StructuresOrderInfo {
	this := StructuresOrderInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StructuresOrderInfo) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOrderInfo) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StructuresOrderInfo) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StructuresOrderInfo) SetId(v string) {
	o.Id = &v
}

// GetVpsId returns the VpsId field value if set, zero value otherwise.
func (o *StructuresOrderInfo) GetVpsId() string {
	if o == nil || isNil(o.VpsId) {
		var ret string
		return ret
	}
	return *o.VpsId
}

// GetVpsIdOk returns a tuple with the VpsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOrderInfo) GetVpsIdOk() (*string, bool) {
	if o == nil || isNil(o.VpsId) {
    return nil, false
	}
	return o.VpsId, true
}

// HasVpsId returns a boolean if a field has been set.
func (o *StructuresOrderInfo) HasVpsId() bool {
	if o != nil && !isNil(o.VpsId) {
		return true
	}

	return false
}

// SetVpsId gets a reference to the given string and assigns it to the VpsId field.
func (o *StructuresOrderInfo) SetVpsId(v string) {
	o.VpsId = &v
}

// GetVpsName returns the VpsName field value if set, zero value otherwise.
func (o *StructuresOrderInfo) GetVpsName() string {
	if o == nil || isNil(o.VpsName) {
		var ret string
		return ret
	}
	return *o.VpsName
}

// GetVpsNameOk returns a tuple with the VpsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOrderInfo) GetVpsNameOk() (*string, bool) {
	if o == nil || isNil(o.VpsName) {
    return nil, false
	}
	return o.VpsName, true
}

// HasVpsName returns a boolean if a field has been set.
func (o *StructuresOrderInfo) HasVpsName() bool {
	if o != nil && !isNil(o.VpsName) {
		return true
	}

	return false
}

// SetVpsName gets a reference to the given string and assigns it to the VpsName field.
func (o *StructuresOrderInfo) SetVpsName(v string) {
	o.VpsName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StructuresOrderInfo) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOrderInfo) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StructuresOrderInfo) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StructuresOrderInfo) SetType(v string) {
	o.Type = &v
}

// GetDateCreate returns the DateCreate field value if set, zero value otherwise.
func (o *StructuresOrderInfo) GetDateCreate() string {
	if o == nil || isNil(o.DateCreate) {
		var ret string
		return ret
	}
	return *o.DateCreate
}

// GetDateCreateOk returns a tuple with the DateCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOrderInfo) GetDateCreateOk() (*string, bool) {
	if o == nil || isNil(o.DateCreate) {
    return nil, false
	}
	return o.DateCreate, true
}

// HasDateCreate returns a boolean if a field has been set.
func (o *StructuresOrderInfo) HasDateCreate() bool {
	if o != nil && !isNil(o.DateCreate) {
		return true
	}

	return false
}

// SetDateCreate gets a reference to the given string and assigns it to the DateCreate field.
func (o *StructuresOrderInfo) SetDateCreate(v string) {
	o.DateCreate = &v
}

// GetDateComplete returns the DateComplete field value if set, zero value otherwise.
func (o *StructuresOrderInfo) GetDateComplete() string {
	if o == nil || isNil(o.DateComplete) {
		var ret string
		return ret
	}
	return *o.DateComplete
}

// GetDateCompleteOk returns a tuple with the DateComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOrderInfo) GetDateCompleteOk() (*string, bool) {
	if o == nil || isNil(o.DateComplete) {
    return nil, false
	}
	return o.DateComplete, true
}

// HasDateComplete returns a boolean if a field has been set.
func (o *StructuresOrderInfo) HasDateComplete() bool {
	if o != nil && !isNil(o.DateComplete) {
		return true
	}

	return false
}

// SetDateComplete gets a reference to the given string and assigns it to the DateComplete field.
func (o *StructuresOrderInfo) SetDateComplete(v string) {
	o.DateComplete = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *StructuresOrderInfo) GetPath() []string {
	if o == nil || isNil(o.Path) {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOrderInfo) GetPathOk() ([]string, bool) {
	if o == nil || isNil(o.Path) {
    return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *StructuresOrderInfo) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *StructuresOrderInfo) SetPath(v []string) {
	o.Path = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StructuresOrderInfo) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOrderInfo) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StructuresOrderInfo) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StructuresOrderInfo) SetStatus(v string) {
	o.Status = &v
}

// GetCopyInfo returns the CopyInfo field value if set, zero value otherwise.
func (o *StructuresOrderInfo) GetCopyInfo() StructuresCopyInfo {
	if o == nil || isNil(o.CopyInfo) {
		var ret StructuresCopyInfo
		return ret
	}
	return *o.CopyInfo
}

// GetCopyInfoOk returns a tuple with the CopyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOrderInfo) GetCopyInfoOk() (*StructuresCopyInfo, bool) {
	if o == nil || isNil(o.CopyInfo) {
    return nil, false
	}
	return o.CopyInfo, true
}

// HasCopyInfo returns a boolean if a field has been set.
func (o *StructuresOrderInfo) HasCopyInfo() bool {
	if o != nil && !isNil(o.CopyInfo) {
		return true
	}

	return false
}

// SetCopyInfo gets a reference to the given StructuresCopyInfo and assigns it to the CopyInfo field.
func (o *StructuresOrderInfo) SetCopyInfo(v StructuresCopyInfo) {
	o.CopyInfo = &v
}

// GetAffectLive returns the AffectLive field value if set, zero value otherwise.
func (o *StructuresOrderInfo) GetAffectLive() bool {
	if o == nil || isNil(o.AffectLive) {
		var ret bool
		return ret
	}
	return *o.AffectLive
}

// GetAffectLiveOk returns a tuple with the AffectLive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOrderInfo) GetAffectLiveOk() (*bool, bool) {
	if o == nil || isNil(o.AffectLive) {
    return nil, false
	}
	return o.AffectLive, true
}

// HasAffectLive returns a boolean if a field has been set.
func (o *StructuresOrderInfo) HasAffectLive() bool {
	if o != nil && !isNil(o.AffectLive) {
		return true
	}

	return false
}

// SetAffectLive gets a reference to the given bool and assigns it to the AffectLive field.
func (o *StructuresOrderInfo) SetAffectLive(v bool) {
	o.AffectLive = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *StructuresOrderInfo) GetTarget() string {
	if o == nil || isNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOrderInfo) GetTargetOk() (*string, bool) {
	if o == nil || isNil(o.Target) {
    return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *StructuresOrderInfo) HasTarget() bool {
	if o != nil && !isNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *StructuresOrderInfo) SetTarget(v string) {
	o.Target = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *StructuresOrderInfo) GetErrorDetails() StructuresOrderInfoErrorDetails {
	if o == nil || isNil(o.ErrorDetails) {
		var ret StructuresOrderInfoErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOrderInfo) GetErrorDetailsOk() (*StructuresOrderInfoErrorDetails, bool) {
	if o == nil || isNil(o.ErrorDetails) {
    return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *StructuresOrderInfo) HasErrorDetails() bool {
	if o != nil && !isNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given StructuresOrderInfoErrorDetails and assigns it to the ErrorDetails field.
func (o *StructuresOrderInfo) SetErrorDetails(v StructuresOrderInfoErrorDetails) {
	o.ErrorDetails = &v
}

func (o StructuresOrderInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.VpsId) {
		toSerialize["vps_id"] = o.VpsId
	}
	if !isNil(o.VpsName) {
		toSerialize["vps_name"] = o.VpsName
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.DateCreate) {
		toSerialize["date_create"] = o.DateCreate
	}
	if !isNil(o.DateComplete) {
		toSerialize["date_complete"] = o.DateComplete
	}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.CopyInfo) {
		toSerialize["copy_info"] = o.CopyInfo
	}
	if !isNil(o.AffectLive) {
		toSerialize["affect_live"] = o.AffectLive
	}
	if !isNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !isNil(o.ErrorDetails) {
		toSerialize["error_details"] = o.ErrorDetails
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresOrderInfo struct {
	value *StructuresOrderInfo
	isSet bool
}

func (v NullableStructuresOrderInfo) Get() *StructuresOrderInfo {
	return v.value
}

func (v *NullableStructuresOrderInfo) Set(val *StructuresOrderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresOrderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresOrderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresOrderInfo(val *StructuresOrderInfo) *NullableStructuresOrderInfo {
	return &NullableStructuresOrderInfo{value: val, isSet: true}
}

func (v NullableStructuresOrderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresOrderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


