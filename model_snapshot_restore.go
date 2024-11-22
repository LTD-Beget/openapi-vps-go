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

// SnapshotRestore struct for SnapshotRestore
type SnapshotRestore struct {
	Id *string `json:"id,omitempty"`
	Snapshot *SnapshotSnapshot `json:"snapshot,omitempty"`
	VpsId *string `json:"vps_id,omitempty"`
	VpsName *string `json:"vps_name,omitempty"`
	TargetType *string `json:"target_type,omitempty"`
	DateCreate *string `json:"date_create,omitempty"`
	DateComplete *string `json:"date_complete,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewSnapshotRestore instantiates a new SnapshotRestore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotRestore() *SnapshotRestore {
	this := SnapshotRestore{}
	return &this
}

// NewSnapshotRestoreWithDefaults instantiates a new SnapshotRestore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotRestoreWithDefaults() *SnapshotRestore {
	this := SnapshotRestore{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SnapshotRestore) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRestore) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SnapshotRestore) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SnapshotRestore) SetId(v string) {
	o.Id = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *SnapshotRestore) GetSnapshot() SnapshotSnapshot {
	if o == nil || isNil(o.Snapshot) {
		var ret SnapshotSnapshot
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRestore) GetSnapshotOk() (*SnapshotSnapshot, bool) {
	if o == nil || isNil(o.Snapshot) {
    return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *SnapshotRestore) HasSnapshot() bool {
	if o != nil && !isNil(o.Snapshot) {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given SnapshotSnapshot and assigns it to the Snapshot field.
func (o *SnapshotRestore) SetSnapshot(v SnapshotSnapshot) {
	o.Snapshot = &v
}

// GetVpsId returns the VpsId field value if set, zero value otherwise.
func (o *SnapshotRestore) GetVpsId() string {
	if o == nil || isNil(o.VpsId) {
		var ret string
		return ret
	}
	return *o.VpsId
}

// GetVpsIdOk returns a tuple with the VpsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRestore) GetVpsIdOk() (*string, bool) {
	if o == nil || isNil(o.VpsId) {
    return nil, false
	}
	return o.VpsId, true
}

// HasVpsId returns a boolean if a field has been set.
func (o *SnapshotRestore) HasVpsId() bool {
	if o != nil && !isNil(o.VpsId) {
		return true
	}

	return false
}

// SetVpsId gets a reference to the given string and assigns it to the VpsId field.
func (o *SnapshotRestore) SetVpsId(v string) {
	o.VpsId = &v
}

// GetVpsName returns the VpsName field value if set, zero value otherwise.
func (o *SnapshotRestore) GetVpsName() string {
	if o == nil || isNil(o.VpsName) {
		var ret string
		return ret
	}
	return *o.VpsName
}

// GetVpsNameOk returns a tuple with the VpsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRestore) GetVpsNameOk() (*string, bool) {
	if o == nil || isNil(o.VpsName) {
    return nil, false
	}
	return o.VpsName, true
}

// HasVpsName returns a boolean if a field has been set.
func (o *SnapshotRestore) HasVpsName() bool {
	if o != nil && !isNil(o.VpsName) {
		return true
	}

	return false
}

// SetVpsName gets a reference to the given string and assigns it to the VpsName field.
func (o *SnapshotRestore) SetVpsName(v string) {
	o.VpsName = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *SnapshotRestore) GetTargetType() string {
	if o == nil || isNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRestore) GetTargetTypeOk() (*string, bool) {
	if o == nil || isNil(o.TargetType) {
    return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *SnapshotRestore) HasTargetType() bool {
	if o != nil && !isNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *SnapshotRestore) SetTargetType(v string) {
	o.TargetType = &v
}

// GetDateCreate returns the DateCreate field value if set, zero value otherwise.
func (o *SnapshotRestore) GetDateCreate() string {
	if o == nil || isNil(o.DateCreate) {
		var ret string
		return ret
	}
	return *o.DateCreate
}

// GetDateCreateOk returns a tuple with the DateCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRestore) GetDateCreateOk() (*string, bool) {
	if o == nil || isNil(o.DateCreate) {
    return nil, false
	}
	return o.DateCreate, true
}

// HasDateCreate returns a boolean if a field has been set.
func (o *SnapshotRestore) HasDateCreate() bool {
	if o != nil && !isNil(o.DateCreate) {
		return true
	}

	return false
}

// SetDateCreate gets a reference to the given string and assigns it to the DateCreate field.
func (o *SnapshotRestore) SetDateCreate(v string) {
	o.DateCreate = &v
}

// GetDateComplete returns the DateComplete field value if set, zero value otherwise.
func (o *SnapshotRestore) GetDateComplete() string {
	if o == nil || isNil(o.DateComplete) {
		var ret string
		return ret
	}
	return *o.DateComplete
}

// GetDateCompleteOk returns a tuple with the DateComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRestore) GetDateCompleteOk() (*string, bool) {
	if o == nil || isNil(o.DateComplete) {
    return nil, false
	}
	return o.DateComplete, true
}

// HasDateComplete returns a boolean if a field has been set.
func (o *SnapshotRestore) HasDateComplete() bool {
	if o != nil && !isNil(o.DateComplete) {
		return true
	}

	return false
}

// SetDateComplete gets a reference to the given string and assigns it to the DateComplete field.
func (o *SnapshotRestore) SetDateComplete(v string) {
	o.DateComplete = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SnapshotRestore) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRestore) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SnapshotRestore) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SnapshotRestore) SetStatus(v string) {
	o.Status = &v
}

func (o SnapshotRestore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Snapshot) {
		toSerialize["snapshot"] = o.Snapshot
	}
	if !isNil(o.VpsId) {
		toSerialize["vps_id"] = o.VpsId
	}
	if !isNil(o.VpsName) {
		toSerialize["vps_name"] = o.VpsName
	}
	if !isNil(o.TargetType) {
		toSerialize["target_type"] = o.TargetType
	}
	if !isNil(o.DateCreate) {
		toSerialize["date_create"] = o.DateCreate
	}
	if !isNil(o.DateComplete) {
		toSerialize["date_complete"] = o.DateComplete
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotRestore struct {
	value *SnapshotRestore
	isSet bool
}

func (v NullableSnapshotRestore) Get() *SnapshotRestore {
	return v.value
}

func (v *NullableSnapshotRestore) Set(val *SnapshotRestore) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotRestore) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotRestore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotRestore(val *SnapshotRestore) *NullableSnapshotRestore {
	return &NullableSnapshotRestore{value: val, isSet: true}
}

func (v NullableSnapshotRestore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotRestore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


