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

// StructuresCopyInfo struct for StructuresCopyInfo
type StructuresCopyInfo struct {
	Id *string `json:"id,omitempty"`
	VpsId *string `json:"vps_id,omitempty"`
	VpsName *string `json:"vps_name,omitempty"`
	Date *string `json:"date,omitempty"`
	Size *string `json:"size,omitempty"`
	Region *string `json:"region,omitempty"`
	Configuration *StructuresCopyInfoConfiguration `json:"configuration,omitempty"`
	InstalledSoftware *StructuresInstalledSoftwareInfo `json:"installed_software,omitempty"`
}

// NewStructuresCopyInfo instantiates a new StructuresCopyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresCopyInfo() *StructuresCopyInfo {
	this := StructuresCopyInfo{}
	return &this
}

// NewStructuresCopyInfoWithDefaults instantiates a new StructuresCopyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresCopyInfoWithDefaults() *StructuresCopyInfo {
	this := StructuresCopyInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StructuresCopyInfo) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresCopyInfo) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StructuresCopyInfo) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StructuresCopyInfo) SetId(v string) {
	o.Id = &v
}

// GetVpsId returns the VpsId field value if set, zero value otherwise.
func (o *StructuresCopyInfo) GetVpsId() string {
	if o == nil || isNil(o.VpsId) {
		var ret string
		return ret
	}
	return *o.VpsId
}

// GetVpsIdOk returns a tuple with the VpsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresCopyInfo) GetVpsIdOk() (*string, bool) {
	if o == nil || isNil(o.VpsId) {
    return nil, false
	}
	return o.VpsId, true
}

// HasVpsId returns a boolean if a field has been set.
func (o *StructuresCopyInfo) HasVpsId() bool {
	if o != nil && !isNil(o.VpsId) {
		return true
	}

	return false
}

// SetVpsId gets a reference to the given string and assigns it to the VpsId field.
func (o *StructuresCopyInfo) SetVpsId(v string) {
	o.VpsId = &v
}

// GetVpsName returns the VpsName field value if set, zero value otherwise.
func (o *StructuresCopyInfo) GetVpsName() string {
	if o == nil || isNil(o.VpsName) {
		var ret string
		return ret
	}
	return *o.VpsName
}

// GetVpsNameOk returns a tuple with the VpsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresCopyInfo) GetVpsNameOk() (*string, bool) {
	if o == nil || isNil(o.VpsName) {
    return nil, false
	}
	return o.VpsName, true
}

// HasVpsName returns a boolean if a field has been set.
func (o *StructuresCopyInfo) HasVpsName() bool {
	if o != nil && !isNil(o.VpsName) {
		return true
	}

	return false
}

// SetVpsName gets a reference to the given string and assigns it to the VpsName field.
func (o *StructuresCopyInfo) SetVpsName(v string) {
	o.VpsName = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *StructuresCopyInfo) GetDate() string {
	if o == nil || isNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresCopyInfo) GetDateOk() (*string, bool) {
	if o == nil || isNil(o.Date) {
    return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *StructuresCopyInfo) HasDate() bool {
	if o != nil && !isNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *StructuresCopyInfo) SetDate(v string) {
	o.Date = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StructuresCopyInfo) GetSize() string {
	if o == nil || isNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresCopyInfo) GetSizeOk() (*string, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StructuresCopyInfo) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *StructuresCopyInfo) SetSize(v string) {
	o.Size = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *StructuresCopyInfo) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresCopyInfo) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
    return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *StructuresCopyInfo) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *StructuresCopyInfo) SetRegion(v string) {
	o.Region = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *StructuresCopyInfo) GetConfiguration() StructuresCopyInfoConfiguration {
	if o == nil || isNil(o.Configuration) {
		var ret StructuresCopyInfoConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresCopyInfo) GetConfigurationOk() (*StructuresCopyInfoConfiguration, bool) {
	if o == nil || isNil(o.Configuration) {
    return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *StructuresCopyInfo) HasConfiguration() bool {
	if o != nil && !isNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given StructuresCopyInfoConfiguration and assigns it to the Configuration field.
func (o *StructuresCopyInfo) SetConfiguration(v StructuresCopyInfoConfiguration) {
	o.Configuration = &v
}

// GetInstalledSoftware returns the InstalledSoftware field value if set, zero value otherwise.
func (o *StructuresCopyInfo) GetInstalledSoftware() StructuresInstalledSoftwareInfo {
	if o == nil || isNil(o.InstalledSoftware) {
		var ret StructuresInstalledSoftwareInfo
		return ret
	}
	return *o.InstalledSoftware
}

// GetInstalledSoftwareOk returns a tuple with the InstalledSoftware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresCopyInfo) GetInstalledSoftwareOk() (*StructuresInstalledSoftwareInfo, bool) {
	if o == nil || isNil(o.InstalledSoftware) {
    return nil, false
	}
	return o.InstalledSoftware, true
}

// HasInstalledSoftware returns a boolean if a field has been set.
func (o *StructuresCopyInfo) HasInstalledSoftware() bool {
	if o != nil && !isNil(o.InstalledSoftware) {
		return true
	}

	return false
}

// SetInstalledSoftware gets a reference to the given StructuresInstalledSoftwareInfo and assigns it to the InstalledSoftware field.
func (o *StructuresCopyInfo) SetInstalledSoftware(v StructuresInstalledSoftwareInfo) {
	o.InstalledSoftware = &v
}

func (o StructuresCopyInfo) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !isNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !isNil(o.InstalledSoftware) {
		toSerialize["installed_software"] = o.InstalledSoftware
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresCopyInfo struct {
	value *StructuresCopyInfo
	isSet bool
}

func (v NullableStructuresCopyInfo) Get() *StructuresCopyInfo {
	return v.value
}

func (v *NullableStructuresCopyInfo) Set(val *StructuresCopyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresCopyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresCopyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresCopyInfo(val *StructuresCopyInfo) *NullableStructuresCopyInfo {
	return &NullableStructuresCopyInfo{value: val, isSet: true}
}

func (v NullableStructuresCopyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresCopyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


