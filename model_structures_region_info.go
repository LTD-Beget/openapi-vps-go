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

// StructuresRegionInfo struct for StructuresRegionInfo
type StructuresRegionInfo struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	NameEn *string `json:"name_en,omitempty"`
	Country *string `json:"country,omitempty"`
	Available *bool `json:"available,omitempty"`
	Priority *int32 `json:"priority,omitempty"`
}

// NewStructuresRegionInfo instantiates a new StructuresRegionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresRegionInfo() *StructuresRegionInfo {
	this := StructuresRegionInfo{}
	return &this
}

// NewStructuresRegionInfoWithDefaults instantiates a new StructuresRegionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresRegionInfoWithDefaults() *StructuresRegionInfo {
	this := StructuresRegionInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StructuresRegionInfo) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresRegionInfo) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StructuresRegionInfo) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StructuresRegionInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StructuresRegionInfo) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresRegionInfo) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StructuresRegionInfo) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StructuresRegionInfo) SetName(v string) {
	o.Name = &v
}

// GetNameEn returns the NameEn field value if set, zero value otherwise.
func (o *StructuresRegionInfo) GetNameEn() string {
	if o == nil || isNil(o.NameEn) {
		var ret string
		return ret
	}
	return *o.NameEn
}

// GetNameEnOk returns a tuple with the NameEn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresRegionInfo) GetNameEnOk() (*string, bool) {
	if o == nil || isNil(o.NameEn) {
    return nil, false
	}
	return o.NameEn, true
}

// HasNameEn returns a boolean if a field has been set.
func (o *StructuresRegionInfo) HasNameEn() bool {
	if o != nil && !isNil(o.NameEn) {
		return true
	}

	return false
}

// SetNameEn gets a reference to the given string and assigns it to the NameEn field.
func (o *StructuresRegionInfo) SetNameEn(v string) {
	o.NameEn = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *StructuresRegionInfo) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresRegionInfo) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
    return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *StructuresRegionInfo) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *StructuresRegionInfo) SetCountry(v string) {
	o.Country = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *StructuresRegionInfo) GetAvailable() bool {
	if o == nil || isNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresRegionInfo) GetAvailableOk() (*bool, bool) {
	if o == nil || isNil(o.Available) {
    return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *StructuresRegionInfo) HasAvailable() bool {
	if o != nil && !isNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *StructuresRegionInfo) SetAvailable(v bool) {
	o.Available = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *StructuresRegionInfo) GetPriority() int32 {
	if o == nil || isNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresRegionInfo) GetPriorityOk() (*int32, bool) {
	if o == nil || isNil(o.Priority) {
    return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *StructuresRegionInfo) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *StructuresRegionInfo) SetPriority(v int32) {
	o.Priority = &v
}

func (o StructuresRegionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.NameEn) {
		toSerialize["name_en"] = o.NameEn
	}
	if !isNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !isNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresRegionInfo struct {
	value *StructuresRegionInfo
	isSet bool
}

func (v NullableStructuresRegionInfo) Get() *StructuresRegionInfo {
	return v.value
}

func (v *NullableStructuresRegionInfo) Set(val *StructuresRegionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresRegionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresRegionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresRegionInfo(val *StructuresRegionInfo) *NullableStructuresRegionInfo {
	return &NullableStructuresRegionInfo{value: val, isSet: true}
}

func (v NullableStructuresRegionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresRegionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

