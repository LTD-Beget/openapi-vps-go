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

// ManageVpsInfo struct for ManageVpsInfo
type ManageVpsInfo struct {
	Id *string `json:"id,omitempty"`
	Slug *string `json:"slug,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	Configuration *ManageVpsConfiguration `json:"configuration,omitempty"`
	Status *string `json:"status,omitempty"`
	SshKeys []StructuresSshKeyInfo `json:"ssh_keys,omitempty"`
	HasPassword *bool `json:"has_password,omitempty"`
	ManageEnabled *bool `json:"manage_enabled,omitempty"`
	Description *string `json:"description,omitempty"`
	DateCreate *string `json:"date_create,omitempty"`
	IpAddress *string `json:"ip_address,omitempty"`
	RescueMode *bool `json:"rescue_mode,omitempty"`
	Migrating *bool `json:"migrating,omitempty"`
	HostUnavailable *bool `json:"host_unavailable,omitempty"`
	Unblocking *bool `json:"unblocking,omitempty"`
	Restoring *bool `json:"restoring,omitempty"`
	DiskUsed *string `json:"disk_used,omitempty"`
	DiskLeft *string `json:"disk_left,omitempty"`
	AdditionalIpAddress []string `json:"additional_ip_address,omitempty"`
	BegetSshAccessAllowed *bool `json:"beget_ssh_access_allowed,omitempty"`
	Archived *bool `json:"archived,omitempty"`
	Unarchiving *bool `json:"unarchiving,omitempty"`
	PrivateNetwork []StructuresAttachedPrivateNetwork `json:"private_network,omitempty"`
	TechnicalDomain *string `json:"technical_domain,omitempty"`
	SoftwareDomain *string `json:"software_domain,omitempty"`
	Software *StructuresInstalledSoftwareInfo `json:"software,omitempty"`
}

// NewManageVpsInfo instantiates a new ManageVpsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageVpsInfo() *ManageVpsInfo {
	this := ManageVpsInfo{}
	return &this
}

// NewManageVpsInfoWithDefaults instantiates a new ManageVpsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageVpsInfoWithDefaults() *ManageVpsInfo {
	this := ManageVpsInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ManageVpsInfo) SetId(v string) {
	o.Id = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetSlug() string {
	if o == nil || isNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetSlugOk() (*string, bool) {
	if o == nil || isNil(o.Slug) {
    return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasSlug() bool {
	if o != nil && !isNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *ManageVpsInfo) SetSlug(v string) {
	o.Slug = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ManageVpsInfo) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetHostname() string {
	if o == nil || isNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetHostnameOk() (*string, bool) {
	if o == nil || isNil(o.Hostname) {
    return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasHostname() bool {
	if o != nil && !isNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ManageVpsInfo) SetHostname(v string) {
	o.Hostname = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetConfiguration() ManageVpsConfiguration {
	if o == nil || isNil(o.Configuration) {
		var ret ManageVpsConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetConfigurationOk() (*ManageVpsConfiguration, bool) {
	if o == nil || isNil(o.Configuration) {
    return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasConfiguration() bool {
	if o != nil && !isNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given ManageVpsConfiguration and assigns it to the Configuration field.
func (o *ManageVpsInfo) SetConfiguration(v ManageVpsConfiguration) {
	o.Configuration = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ManageVpsInfo) SetStatus(v string) {
	o.Status = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetSshKeys() []StructuresSshKeyInfo {
	if o == nil || isNil(o.SshKeys) {
		var ret []StructuresSshKeyInfo
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetSshKeysOk() ([]StructuresSshKeyInfo, bool) {
	if o == nil || isNil(o.SshKeys) {
    return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasSshKeys() bool {
	if o != nil && !isNil(o.SshKeys) {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []StructuresSshKeyInfo and assigns it to the SshKeys field.
func (o *ManageVpsInfo) SetSshKeys(v []StructuresSshKeyInfo) {
	o.SshKeys = v
}

// GetHasPassword returns the HasPassword field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetHasPassword() bool {
	if o == nil || isNil(o.HasPassword) {
		var ret bool
		return ret
	}
	return *o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetHasPasswordOk() (*bool, bool) {
	if o == nil || isNil(o.HasPassword) {
    return nil, false
	}
	return o.HasPassword, true
}

// HasHasPassword returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasHasPassword() bool {
	if o != nil && !isNil(o.HasPassword) {
		return true
	}

	return false
}

// SetHasPassword gets a reference to the given bool and assigns it to the HasPassword field.
func (o *ManageVpsInfo) SetHasPassword(v bool) {
	o.HasPassword = &v
}

// GetManageEnabled returns the ManageEnabled field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetManageEnabled() bool {
	if o == nil || isNil(o.ManageEnabled) {
		var ret bool
		return ret
	}
	return *o.ManageEnabled
}

// GetManageEnabledOk returns a tuple with the ManageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetManageEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.ManageEnabled) {
    return nil, false
	}
	return o.ManageEnabled, true
}

// HasManageEnabled returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasManageEnabled() bool {
	if o != nil && !isNil(o.ManageEnabled) {
		return true
	}

	return false
}

// SetManageEnabled gets a reference to the given bool and assigns it to the ManageEnabled field.
func (o *ManageVpsInfo) SetManageEnabled(v bool) {
	o.ManageEnabled = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ManageVpsInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDateCreate returns the DateCreate field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetDateCreate() string {
	if o == nil || isNil(o.DateCreate) {
		var ret string
		return ret
	}
	return *o.DateCreate
}

// GetDateCreateOk returns a tuple with the DateCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetDateCreateOk() (*string, bool) {
	if o == nil || isNil(o.DateCreate) {
    return nil, false
	}
	return o.DateCreate, true
}

// HasDateCreate returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasDateCreate() bool {
	if o != nil && !isNil(o.DateCreate) {
		return true
	}

	return false
}

// SetDateCreate gets a reference to the given string and assigns it to the DateCreate field.
func (o *ManageVpsInfo) SetDateCreate(v string) {
	o.DateCreate = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetIpAddress() string {
	if o == nil || isNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetIpAddressOk() (*string, bool) {
	if o == nil || isNil(o.IpAddress) {
    return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasIpAddress() bool {
	if o != nil && !isNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *ManageVpsInfo) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetRescueMode returns the RescueMode field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetRescueMode() bool {
	if o == nil || isNil(o.RescueMode) {
		var ret bool
		return ret
	}
	return *o.RescueMode
}

// GetRescueModeOk returns a tuple with the RescueMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetRescueModeOk() (*bool, bool) {
	if o == nil || isNil(o.RescueMode) {
    return nil, false
	}
	return o.RescueMode, true
}

// HasRescueMode returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasRescueMode() bool {
	if o != nil && !isNil(o.RescueMode) {
		return true
	}

	return false
}

// SetRescueMode gets a reference to the given bool and assigns it to the RescueMode field.
func (o *ManageVpsInfo) SetRescueMode(v bool) {
	o.RescueMode = &v
}

// GetMigrating returns the Migrating field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetMigrating() bool {
	if o == nil || isNil(o.Migrating) {
		var ret bool
		return ret
	}
	return *o.Migrating
}

// GetMigratingOk returns a tuple with the Migrating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetMigratingOk() (*bool, bool) {
	if o == nil || isNil(o.Migrating) {
    return nil, false
	}
	return o.Migrating, true
}

// HasMigrating returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasMigrating() bool {
	if o != nil && !isNil(o.Migrating) {
		return true
	}

	return false
}

// SetMigrating gets a reference to the given bool and assigns it to the Migrating field.
func (o *ManageVpsInfo) SetMigrating(v bool) {
	o.Migrating = &v
}

// GetHostUnavailable returns the HostUnavailable field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetHostUnavailable() bool {
	if o == nil || isNil(o.HostUnavailable) {
		var ret bool
		return ret
	}
	return *o.HostUnavailable
}

// GetHostUnavailableOk returns a tuple with the HostUnavailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetHostUnavailableOk() (*bool, bool) {
	if o == nil || isNil(o.HostUnavailable) {
    return nil, false
	}
	return o.HostUnavailable, true
}

// HasHostUnavailable returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasHostUnavailable() bool {
	if o != nil && !isNil(o.HostUnavailable) {
		return true
	}

	return false
}

// SetHostUnavailable gets a reference to the given bool and assigns it to the HostUnavailable field.
func (o *ManageVpsInfo) SetHostUnavailable(v bool) {
	o.HostUnavailable = &v
}

// GetUnblocking returns the Unblocking field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetUnblocking() bool {
	if o == nil || isNil(o.Unblocking) {
		var ret bool
		return ret
	}
	return *o.Unblocking
}

// GetUnblockingOk returns a tuple with the Unblocking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetUnblockingOk() (*bool, bool) {
	if o == nil || isNil(o.Unblocking) {
    return nil, false
	}
	return o.Unblocking, true
}

// HasUnblocking returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasUnblocking() bool {
	if o != nil && !isNil(o.Unblocking) {
		return true
	}

	return false
}

// SetUnblocking gets a reference to the given bool and assigns it to the Unblocking field.
func (o *ManageVpsInfo) SetUnblocking(v bool) {
	o.Unblocking = &v
}

// GetRestoring returns the Restoring field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetRestoring() bool {
	if o == nil || isNil(o.Restoring) {
		var ret bool
		return ret
	}
	return *o.Restoring
}

// GetRestoringOk returns a tuple with the Restoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetRestoringOk() (*bool, bool) {
	if o == nil || isNil(o.Restoring) {
    return nil, false
	}
	return o.Restoring, true
}

// HasRestoring returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasRestoring() bool {
	if o != nil && !isNil(o.Restoring) {
		return true
	}

	return false
}

// SetRestoring gets a reference to the given bool and assigns it to the Restoring field.
func (o *ManageVpsInfo) SetRestoring(v bool) {
	o.Restoring = &v
}

// GetDiskUsed returns the DiskUsed field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetDiskUsed() string {
	if o == nil || isNil(o.DiskUsed) {
		var ret string
		return ret
	}
	return *o.DiskUsed
}

// GetDiskUsedOk returns a tuple with the DiskUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetDiskUsedOk() (*string, bool) {
	if o == nil || isNil(o.DiskUsed) {
    return nil, false
	}
	return o.DiskUsed, true
}

// HasDiskUsed returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasDiskUsed() bool {
	if o != nil && !isNil(o.DiskUsed) {
		return true
	}

	return false
}

// SetDiskUsed gets a reference to the given string and assigns it to the DiskUsed field.
func (o *ManageVpsInfo) SetDiskUsed(v string) {
	o.DiskUsed = &v
}

// GetDiskLeft returns the DiskLeft field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetDiskLeft() string {
	if o == nil || isNil(o.DiskLeft) {
		var ret string
		return ret
	}
	return *o.DiskLeft
}

// GetDiskLeftOk returns a tuple with the DiskLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetDiskLeftOk() (*string, bool) {
	if o == nil || isNil(o.DiskLeft) {
    return nil, false
	}
	return o.DiskLeft, true
}

// HasDiskLeft returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasDiskLeft() bool {
	if o != nil && !isNil(o.DiskLeft) {
		return true
	}

	return false
}

// SetDiskLeft gets a reference to the given string and assigns it to the DiskLeft field.
func (o *ManageVpsInfo) SetDiskLeft(v string) {
	o.DiskLeft = &v
}

// GetAdditionalIpAddress returns the AdditionalIpAddress field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetAdditionalIpAddress() []string {
	if o == nil || isNil(o.AdditionalIpAddress) {
		var ret []string
		return ret
	}
	return o.AdditionalIpAddress
}

// GetAdditionalIpAddressOk returns a tuple with the AdditionalIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetAdditionalIpAddressOk() ([]string, bool) {
	if o == nil || isNil(o.AdditionalIpAddress) {
    return nil, false
	}
	return o.AdditionalIpAddress, true
}

// HasAdditionalIpAddress returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasAdditionalIpAddress() bool {
	if o != nil && !isNil(o.AdditionalIpAddress) {
		return true
	}

	return false
}

// SetAdditionalIpAddress gets a reference to the given []string and assigns it to the AdditionalIpAddress field.
func (o *ManageVpsInfo) SetAdditionalIpAddress(v []string) {
	o.AdditionalIpAddress = v
}

// GetBegetSshAccessAllowed returns the BegetSshAccessAllowed field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetBegetSshAccessAllowed() bool {
	if o == nil || isNil(o.BegetSshAccessAllowed) {
		var ret bool
		return ret
	}
	return *o.BegetSshAccessAllowed
}

// GetBegetSshAccessAllowedOk returns a tuple with the BegetSshAccessAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetBegetSshAccessAllowedOk() (*bool, bool) {
	if o == nil || isNil(o.BegetSshAccessAllowed) {
    return nil, false
	}
	return o.BegetSshAccessAllowed, true
}

// HasBegetSshAccessAllowed returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasBegetSshAccessAllowed() bool {
	if o != nil && !isNil(o.BegetSshAccessAllowed) {
		return true
	}

	return false
}

// SetBegetSshAccessAllowed gets a reference to the given bool and assigns it to the BegetSshAccessAllowed field.
func (o *ManageVpsInfo) SetBegetSshAccessAllowed(v bool) {
	o.BegetSshAccessAllowed = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetArchived() bool {
	if o == nil || isNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetArchivedOk() (*bool, bool) {
	if o == nil || isNil(o.Archived) {
    return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasArchived() bool {
	if o != nil && !isNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *ManageVpsInfo) SetArchived(v bool) {
	o.Archived = &v
}

// GetUnarchiving returns the Unarchiving field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetUnarchiving() bool {
	if o == nil || isNil(o.Unarchiving) {
		var ret bool
		return ret
	}
	return *o.Unarchiving
}

// GetUnarchivingOk returns a tuple with the Unarchiving field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetUnarchivingOk() (*bool, bool) {
	if o == nil || isNil(o.Unarchiving) {
    return nil, false
	}
	return o.Unarchiving, true
}

// HasUnarchiving returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasUnarchiving() bool {
	if o != nil && !isNil(o.Unarchiving) {
		return true
	}

	return false
}

// SetUnarchiving gets a reference to the given bool and assigns it to the Unarchiving field.
func (o *ManageVpsInfo) SetUnarchiving(v bool) {
	o.Unarchiving = &v
}

// GetPrivateNetwork returns the PrivateNetwork field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetPrivateNetwork() []StructuresAttachedPrivateNetwork {
	if o == nil || isNil(o.PrivateNetwork) {
		var ret []StructuresAttachedPrivateNetwork
		return ret
	}
	return o.PrivateNetwork
}

// GetPrivateNetworkOk returns a tuple with the PrivateNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetPrivateNetworkOk() ([]StructuresAttachedPrivateNetwork, bool) {
	if o == nil || isNil(o.PrivateNetwork) {
    return nil, false
	}
	return o.PrivateNetwork, true
}

// HasPrivateNetwork returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasPrivateNetwork() bool {
	if o != nil && !isNil(o.PrivateNetwork) {
		return true
	}

	return false
}

// SetPrivateNetwork gets a reference to the given []StructuresAttachedPrivateNetwork and assigns it to the PrivateNetwork field.
func (o *ManageVpsInfo) SetPrivateNetwork(v []StructuresAttachedPrivateNetwork) {
	o.PrivateNetwork = v
}

// GetTechnicalDomain returns the TechnicalDomain field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetTechnicalDomain() string {
	if o == nil || isNil(o.TechnicalDomain) {
		var ret string
		return ret
	}
	return *o.TechnicalDomain
}

// GetTechnicalDomainOk returns a tuple with the TechnicalDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetTechnicalDomainOk() (*string, bool) {
	if o == nil || isNil(o.TechnicalDomain) {
    return nil, false
	}
	return o.TechnicalDomain, true
}

// HasTechnicalDomain returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasTechnicalDomain() bool {
	if o != nil && !isNil(o.TechnicalDomain) {
		return true
	}

	return false
}

// SetTechnicalDomain gets a reference to the given string and assigns it to the TechnicalDomain field.
func (o *ManageVpsInfo) SetTechnicalDomain(v string) {
	o.TechnicalDomain = &v
}

// GetSoftwareDomain returns the SoftwareDomain field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetSoftwareDomain() string {
	if o == nil || isNil(o.SoftwareDomain) {
		var ret string
		return ret
	}
	return *o.SoftwareDomain
}

// GetSoftwareDomainOk returns a tuple with the SoftwareDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetSoftwareDomainOk() (*string, bool) {
	if o == nil || isNil(o.SoftwareDomain) {
    return nil, false
	}
	return o.SoftwareDomain, true
}

// HasSoftwareDomain returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasSoftwareDomain() bool {
	if o != nil && !isNil(o.SoftwareDomain) {
		return true
	}

	return false
}

// SetSoftwareDomain gets a reference to the given string and assigns it to the SoftwareDomain field.
func (o *ManageVpsInfo) SetSoftwareDomain(v string) {
	o.SoftwareDomain = &v
}

// GetSoftware returns the Software field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetSoftware() StructuresInstalledSoftwareInfo {
	if o == nil || isNil(o.Software) {
		var ret StructuresInstalledSoftwareInfo
		return ret
	}
	return *o.Software
}

// GetSoftwareOk returns a tuple with the Software field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetSoftwareOk() (*StructuresInstalledSoftwareInfo, bool) {
	if o == nil || isNil(o.Software) {
    return nil, false
	}
	return o.Software, true
}

// HasSoftware returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasSoftware() bool {
	if o != nil && !isNil(o.Software) {
		return true
	}

	return false
}

// SetSoftware gets a reference to the given StructuresInstalledSoftwareInfo and assigns it to the Software field.
func (o *ManageVpsInfo) SetSoftware(v StructuresInstalledSoftwareInfo) {
	o.Software = &v
}

func (o ManageVpsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !isNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !isNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !isNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.SshKeys) {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	if !isNil(o.HasPassword) {
		toSerialize["has_password"] = o.HasPassword
	}
	if !isNil(o.ManageEnabled) {
		toSerialize["manage_enabled"] = o.ManageEnabled
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.DateCreate) {
		toSerialize["date_create"] = o.DateCreate
	}
	if !isNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	if !isNil(o.RescueMode) {
		toSerialize["rescue_mode"] = o.RescueMode
	}
	if !isNil(o.Migrating) {
		toSerialize["migrating"] = o.Migrating
	}
	if !isNil(o.HostUnavailable) {
		toSerialize["host_unavailable"] = o.HostUnavailable
	}
	if !isNil(o.Unblocking) {
		toSerialize["unblocking"] = o.Unblocking
	}
	if !isNil(o.Restoring) {
		toSerialize["restoring"] = o.Restoring
	}
	if !isNil(o.DiskUsed) {
		toSerialize["disk_used"] = o.DiskUsed
	}
	if !isNil(o.DiskLeft) {
		toSerialize["disk_left"] = o.DiskLeft
	}
	if !isNil(o.AdditionalIpAddress) {
		toSerialize["additional_ip_address"] = o.AdditionalIpAddress
	}
	if !isNil(o.BegetSshAccessAllowed) {
		toSerialize["beget_ssh_access_allowed"] = o.BegetSshAccessAllowed
	}
	if !isNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !isNil(o.Unarchiving) {
		toSerialize["unarchiving"] = o.Unarchiving
	}
	if !isNil(o.PrivateNetwork) {
		toSerialize["private_network"] = o.PrivateNetwork
	}
	if !isNil(o.TechnicalDomain) {
		toSerialize["technical_domain"] = o.TechnicalDomain
	}
	if !isNil(o.SoftwareDomain) {
		toSerialize["software_domain"] = o.SoftwareDomain
	}
	if !isNil(o.Software) {
		toSerialize["software"] = o.Software
	}
	return json.Marshal(toSerialize)
}

type NullableManageVpsInfo struct {
	value *ManageVpsInfo
	isSet bool
}

func (v NullableManageVpsInfo) Get() *ManageVpsInfo {
	return v.value
}

func (v *NullableManageVpsInfo) Set(val *ManageVpsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableManageVpsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableManageVpsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageVpsInfo(val *ManageVpsInfo) *NullableManageVpsInfo {
	return &NullableManageVpsInfo{value: val, isSet: true}
}

func (v NullableManageVpsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageVpsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


