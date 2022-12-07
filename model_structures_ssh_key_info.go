/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// StructuresSshKeyInfo struct for StructuresSshKeyInfo
type StructuresSshKeyInfo struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Fingerprint *string `json:"fingerprint,omitempty"`
}

// NewStructuresSshKeyInfo instantiates a new StructuresSshKeyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresSshKeyInfo() *StructuresSshKeyInfo {
	this := StructuresSshKeyInfo{}
	return &this
}

// NewStructuresSshKeyInfoWithDefaults instantiates a new StructuresSshKeyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresSshKeyInfoWithDefaults() *StructuresSshKeyInfo {
	this := StructuresSshKeyInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StructuresSshKeyInfo) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSshKeyInfo) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StructuresSshKeyInfo) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *StructuresSshKeyInfo) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StructuresSshKeyInfo) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSshKeyInfo) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StructuresSshKeyInfo) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StructuresSshKeyInfo) SetName(v string) {
	o.Name = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *StructuresSshKeyInfo) GetFingerprint() string {
	if o == nil || isNil(o.Fingerprint) {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSshKeyInfo) GetFingerprintOk() (*string, bool) {
	if o == nil || isNil(o.Fingerprint) {
    return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *StructuresSshKeyInfo) HasFingerprint() bool {
	if o != nil && !isNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *StructuresSshKeyInfo) SetFingerprint(v string) {
	o.Fingerprint = &v
}

func (o StructuresSshKeyInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Fingerprint) {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresSshKeyInfo struct {
	value *StructuresSshKeyInfo
	isSet bool
}

func (v NullableStructuresSshKeyInfo) Get() *StructuresSshKeyInfo {
	return v.value
}

func (v *NullableStructuresSshKeyInfo) Set(val *StructuresSshKeyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresSshKeyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresSshKeyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresSshKeyInfo(val *StructuresSshKeyInfo) *NullableStructuresSshKeyInfo {
	return &NullableStructuresSshKeyInfo{value: val, isSet: true}
}

func (v NullableStructuresSshKeyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresSshKeyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


