/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// ManageCreateVpsRequest struct for ManageCreateVpsRequest
type ManageCreateVpsRequest struct {
	DisplayName *string `json:"display_name,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	Description *string `json:"description,omitempty"`
	ConfigurationId *string `json:"configuration_id,omitempty"`
	Software *ManageSoftwareInstallInfo `json:"software,omitempty"`
	SnapshotId *string `json:"snapshot_id,omitempty"`
	SshKeys []int32 `json:"ssh_keys,omitempty"`
	Password *string `json:"password,omitempty"`
	BegetSshAccessAllowed *bool `json:"beget_ssh_access_allowed,omitempty"`
	PrivateNetworks []ManagePrivateNetworkInfo `json:"private_networks,omitempty"`
}

// NewManageCreateVpsRequest instantiates a new ManageCreateVpsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageCreateVpsRequest() *ManageCreateVpsRequest {
	this := ManageCreateVpsRequest{}
	return &this
}

// NewManageCreateVpsRequestWithDefaults instantiates a new ManageCreateVpsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageCreateVpsRequestWithDefaults() *ManageCreateVpsRequest {
	this := ManageCreateVpsRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ManageCreateVpsRequest) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ManageCreateVpsRequest) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ManageCreateVpsRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ManageCreateVpsRequest) GetHostname() string {
	if o == nil || isNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsRequest) GetHostnameOk() (*string, bool) {
	if o == nil || isNil(o.Hostname) {
    return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ManageCreateVpsRequest) HasHostname() bool {
	if o != nil && !isNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ManageCreateVpsRequest) SetHostname(v string) {
	o.Hostname = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ManageCreateVpsRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ManageCreateVpsRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ManageCreateVpsRequest) SetDescription(v string) {
	o.Description = &v
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise.
func (o *ManageCreateVpsRequest) GetConfigurationId() string {
	if o == nil || isNil(o.ConfigurationId) {
		var ret string
		return ret
	}
	return *o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsRequest) GetConfigurationIdOk() (*string, bool) {
	if o == nil || isNil(o.ConfigurationId) {
    return nil, false
	}
	return o.ConfigurationId, true
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *ManageCreateVpsRequest) HasConfigurationId() bool {
	if o != nil && !isNil(o.ConfigurationId) {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given string and assigns it to the ConfigurationId field.
func (o *ManageCreateVpsRequest) SetConfigurationId(v string) {
	o.ConfigurationId = &v
}

// GetSoftware returns the Software field value if set, zero value otherwise.
func (o *ManageCreateVpsRequest) GetSoftware() ManageSoftwareInstallInfo {
	if o == nil || isNil(o.Software) {
		var ret ManageSoftwareInstallInfo
		return ret
	}
	return *o.Software
}

// GetSoftwareOk returns a tuple with the Software field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsRequest) GetSoftwareOk() (*ManageSoftwareInstallInfo, bool) {
	if o == nil || isNil(o.Software) {
    return nil, false
	}
	return o.Software, true
}

// HasSoftware returns a boolean if a field has been set.
func (o *ManageCreateVpsRequest) HasSoftware() bool {
	if o != nil && !isNil(o.Software) {
		return true
	}

	return false
}

// SetSoftware gets a reference to the given ManageSoftwareInstallInfo and assigns it to the Software field.
func (o *ManageCreateVpsRequest) SetSoftware(v ManageSoftwareInstallInfo) {
	o.Software = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *ManageCreateVpsRequest) GetSnapshotId() string {
	if o == nil || isNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsRequest) GetSnapshotIdOk() (*string, bool) {
	if o == nil || isNil(o.SnapshotId) {
    return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *ManageCreateVpsRequest) HasSnapshotId() bool {
	if o != nil && !isNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *ManageCreateVpsRequest) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *ManageCreateVpsRequest) GetSshKeys() []int32 {
	if o == nil || isNil(o.SshKeys) {
		var ret []int32
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsRequest) GetSshKeysOk() ([]int32, bool) {
	if o == nil || isNil(o.SshKeys) {
    return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *ManageCreateVpsRequest) HasSshKeys() bool {
	if o != nil && !isNil(o.SshKeys) {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []int32 and assigns it to the SshKeys field.
func (o *ManageCreateVpsRequest) SetSshKeys(v []int32) {
	o.SshKeys = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ManageCreateVpsRequest) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsRequest) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ManageCreateVpsRequest) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ManageCreateVpsRequest) SetPassword(v string) {
	o.Password = &v
}

// GetBegetSshAccessAllowed returns the BegetSshAccessAllowed field value if set, zero value otherwise.
func (o *ManageCreateVpsRequest) GetBegetSshAccessAllowed() bool {
	if o == nil || isNil(o.BegetSshAccessAllowed) {
		var ret bool
		return ret
	}
	return *o.BegetSshAccessAllowed
}

// GetBegetSshAccessAllowedOk returns a tuple with the BegetSshAccessAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsRequest) GetBegetSshAccessAllowedOk() (*bool, bool) {
	if o == nil || isNil(o.BegetSshAccessAllowed) {
    return nil, false
	}
	return o.BegetSshAccessAllowed, true
}

// HasBegetSshAccessAllowed returns a boolean if a field has been set.
func (o *ManageCreateVpsRequest) HasBegetSshAccessAllowed() bool {
	if o != nil && !isNil(o.BegetSshAccessAllowed) {
		return true
	}

	return false
}

// SetBegetSshAccessAllowed gets a reference to the given bool and assigns it to the BegetSshAccessAllowed field.
func (o *ManageCreateVpsRequest) SetBegetSshAccessAllowed(v bool) {
	o.BegetSshAccessAllowed = &v
}

// GetPrivateNetworks returns the PrivateNetworks field value if set, zero value otherwise.
func (o *ManageCreateVpsRequest) GetPrivateNetworks() []ManagePrivateNetworkInfo {
	if o == nil || isNil(o.PrivateNetworks) {
		var ret []ManagePrivateNetworkInfo
		return ret
	}
	return o.PrivateNetworks
}

// GetPrivateNetworksOk returns a tuple with the PrivateNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsRequest) GetPrivateNetworksOk() ([]ManagePrivateNetworkInfo, bool) {
	if o == nil || isNil(o.PrivateNetworks) {
    return nil, false
	}
	return o.PrivateNetworks, true
}

// HasPrivateNetworks returns a boolean if a field has been set.
func (o *ManageCreateVpsRequest) HasPrivateNetworks() bool {
	if o != nil && !isNil(o.PrivateNetworks) {
		return true
	}

	return false
}

// SetPrivateNetworks gets a reference to the given []ManagePrivateNetworkInfo and assigns it to the PrivateNetworks field.
func (o *ManageCreateVpsRequest) SetPrivateNetworks(v []ManagePrivateNetworkInfo) {
	o.PrivateNetworks = v
}

func (o ManageCreateVpsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !isNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.ConfigurationId) {
		toSerialize["configuration_id"] = o.ConfigurationId
	}
	if !isNil(o.Software) {
		toSerialize["software"] = o.Software
	}
	if !isNil(o.SnapshotId) {
		toSerialize["snapshot_id"] = o.SnapshotId
	}
	if !isNil(o.SshKeys) {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.BegetSshAccessAllowed) {
		toSerialize["beget_ssh_access_allowed"] = o.BegetSshAccessAllowed
	}
	if !isNil(o.PrivateNetworks) {
		toSerialize["private_networks"] = o.PrivateNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableManageCreateVpsRequest struct {
	value *ManageCreateVpsRequest
	isSet bool
}

func (v NullableManageCreateVpsRequest) Get() *ManageCreateVpsRequest {
	return v.value
}

func (v *NullableManageCreateVpsRequest) Set(val *ManageCreateVpsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableManageCreateVpsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableManageCreateVpsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageCreateVpsRequest(val *ManageCreateVpsRequest) *NullableManageCreateVpsRequest {
	return &NullableManageCreateVpsRequest{value: val, isSet: true}
}

func (v NullableManageCreateVpsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageCreateVpsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

