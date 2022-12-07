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

// StructuresFileInfo struct for StructuresFileInfo
type StructuresFileInfo struct {
	IsLink *bool `json:"is_link,omitempty"`
	IsDir *bool `json:"is_dir,omitempty"`
	Name *string `json:"name,omitempty"`
	Ext *string `json:"ext,omitempty"`
	Path *string `json:"path,omitempty"`
	Owner *string `json:"owner,omitempty"`
	Group *string `json:"group,omitempty"`
	Size *string `json:"size,omitempty"`
	Mtime *float64 `json:"mtime,omitempty"`
	MtimeStr *string `json:"mtime_str,omitempty"`
	Mode *string `json:"mode,omitempty"`
}

// NewStructuresFileInfo instantiates a new StructuresFileInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresFileInfo() *StructuresFileInfo {
	this := StructuresFileInfo{}
	return &this
}

// NewStructuresFileInfoWithDefaults instantiates a new StructuresFileInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresFileInfoWithDefaults() *StructuresFileInfo {
	this := StructuresFileInfo{}
	return &this
}

// GetIsLink returns the IsLink field value if set, zero value otherwise.
func (o *StructuresFileInfo) GetIsLink() bool {
	if o == nil || isNil(o.IsLink) {
		var ret bool
		return ret
	}
	return *o.IsLink
}

// GetIsLinkOk returns a tuple with the IsLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresFileInfo) GetIsLinkOk() (*bool, bool) {
	if o == nil || isNil(o.IsLink) {
    return nil, false
	}
	return o.IsLink, true
}

// HasIsLink returns a boolean if a field has been set.
func (o *StructuresFileInfo) HasIsLink() bool {
	if o != nil && !isNil(o.IsLink) {
		return true
	}

	return false
}

// SetIsLink gets a reference to the given bool and assigns it to the IsLink field.
func (o *StructuresFileInfo) SetIsLink(v bool) {
	o.IsLink = &v
}

// GetIsDir returns the IsDir field value if set, zero value otherwise.
func (o *StructuresFileInfo) GetIsDir() bool {
	if o == nil || isNil(o.IsDir) {
		var ret bool
		return ret
	}
	return *o.IsDir
}

// GetIsDirOk returns a tuple with the IsDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresFileInfo) GetIsDirOk() (*bool, bool) {
	if o == nil || isNil(o.IsDir) {
    return nil, false
	}
	return o.IsDir, true
}

// HasIsDir returns a boolean if a field has been set.
func (o *StructuresFileInfo) HasIsDir() bool {
	if o != nil && !isNil(o.IsDir) {
		return true
	}

	return false
}

// SetIsDir gets a reference to the given bool and assigns it to the IsDir field.
func (o *StructuresFileInfo) SetIsDir(v bool) {
	o.IsDir = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StructuresFileInfo) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresFileInfo) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StructuresFileInfo) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StructuresFileInfo) SetName(v string) {
	o.Name = &v
}

// GetExt returns the Ext field value if set, zero value otherwise.
func (o *StructuresFileInfo) GetExt() string {
	if o == nil || isNil(o.Ext) {
		var ret string
		return ret
	}
	return *o.Ext
}

// GetExtOk returns a tuple with the Ext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresFileInfo) GetExtOk() (*string, bool) {
	if o == nil || isNil(o.Ext) {
    return nil, false
	}
	return o.Ext, true
}

// HasExt returns a boolean if a field has been set.
func (o *StructuresFileInfo) HasExt() bool {
	if o != nil && !isNil(o.Ext) {
		return true
	}

	return false
}

// SetExt gets a reference to the given string and assigns it to the Ext field.
func (o *StructuresFileInfo) SetExt(v string) {
	o.Ext = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *StructuresFileInfo) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresFileInfo) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
    return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *StructuresFileInfo) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *StructuresFileInfo) SetPath(v string) {
	o.Path = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *StructuresFileInfo) GetOwner() string {
	if o == nil || isNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresFileInfo) GetOwnerOk() (*string, bool) {
	if o == nil || isNil(o.Owner) {
    return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *StructuresFileInfo) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *StructuresFileInfo) SetOwner(v string) {
	o.Owner = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *StructuresFileInfo) GetGroup() string {
	if o == nil || isNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresFileInfo) GetGroupOk() (*string, bool) {
	if o == nil || isNil(o.Group) {
    return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *StructuresFileInfo) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *StructuresFileInfo) SetGroup(v string) {
	o.Group = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StructuresFileInfo) GetSize() string {
	if o == nil || isNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresFileInfo) GetSizeOk() (*string, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StructuresFileInfo) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *StructuresFileInfo) SetSize(v string) {
	o.Size = &v
}

// GetMtime returns the Mtime field value if set, zero value otherwise.
func (o *StructuresFileInfo) GetMtime() float64 {
	if o == nil || isNil(o.Mtime) {
		var ret float64
		return ret
	}
	return *o.Mtime
}

// GetMtimeOk returns a tuple with the Mtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresFileInfo) GetMtimeOk() (*float64, bool) {
	if o == nil || isNil(o.Mtime) {
    return nil, false
	}
	return o.Mtime, true
}

// HasMtime returns a boolean if a field has been set.
func (o *StructuresFileInfo) HasMtime() bool {
	if o != nil && !isNil(o.Mtime) {
		return true
	}

	return false
}

// SetMtime gets a reference to the given float64 and assigns it to the Mtime field.
func (o *StructuresFileInfo) SetMtime(v float64) {
	o.Mtime = &v
}

// GetMtimeStr returns the MtimeStr field value if set, zero value otherwise.
func (o *StructuresFileInfo) GetMtimeStr() string {
	if o == nil || isNil(o.MtimeStr) {
		var ret string
		return ret
	}
	return *o.MtimeStr
}

// GetMtimeStrOk returns a tuple with the MtimeStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresFileInfo) GetMtimeStrOk() (*string, bool) {
	if o == nil || isNil(o.MtimeStr) {
    return nil, false
	}
	return o.MtimeStr, true
}

// HasMtimeStr returns a boolean if a field has been set.
func (o *StructuresFileInfo) HasMtimeStr() bool {
	if o != nil && !isNil(o.MtimeStr) {
		return true
	}

	return false
}

// SetMtimeStr gets a reference to the given string and assigns it to the MtimeStr field.
func (o *StructuresFileInfo) SetMtimeStr(v string) {
	o.MtimeStr = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *StructuresFileInfo) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresFileInfo) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *StructuresFileInfo) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *StructuresFileInfo) SetMode(v string) {
	o.Mode = &v
}

func (o StructuresFileInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsLink) {
		toSerialize["is_link"] = o.IsLink
	}
	if !isNil(o.IsDir) {
		toSerialize["is_dir"] = o.IsDir
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Ext) {
		toSerialize["ext"] = o.Ext
	}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.Mtime) {
		toSerialize["mtime"] = o.Mtime
	}
	if !isNil(o.MtimeStr) {
		toSerialize["mtime_str"] = o.MtimeStr
	}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresFileInfo struct {
	value *StructuresFileInfo
	isSet bool
}

func (v NullableStructuresFileInfo) Get() *StructuresFileInfo {
	return v.value
}

func (v *NullableStructuresFileInfo) Set(val *StructuresFileInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresFileInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresFileInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresFileInfo(val *StructuresFileInfo) *NullableStructuresFileInfo {
	return &NullableStructuresFileInfo{value: val, isSet: true}
}

func (v NullableStructuresFileInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresFileInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


