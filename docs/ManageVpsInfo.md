# ManageVpsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Configuration** | Pointer to [**ManageVpsConfiguration**](ManageVpsConfiguration.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SshKeys** | Pointer to [**[]StructuresSshKeyInfo**](StructuresSshKeyInfo.md) |  | [optional] 
**HasPassword** | Pointer to **bool** |  | [optional] 
**ManageEnabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DateCreate** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**RescueMode** | Pointer to **bool** |  | [optional] 
**Migrating** | Pointer to **bool** |  | [optional] 
**HostUnavailable** | Pointer to **bool** |  | [optional] 
**Unblocking** | Pointer to **bool** |  | [optional] 
**Restoring** | Pointer to **bool** |  | [optional] 
**DiskUsed** | Pointer to **string** |  | [optional] 
**DiskLeft** | Pointer to **string** |  | [optional] 
**AdditionalIpAddress** | Pointer to **[]string** |  | [optional] 
**BegetSshAccessAllowed** | Pointer to **bool** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**Unarchiving** | Pointer to **bool** |  | [optional] 
**PrivateNetwork** | Pointer to [**[]StructuresAttachedPrivateNetwork**](StructuresAttachedPrivateNetwork.md) |  | [optional] 
**TechnicalDomain** | Pointer to **string** |  | [optional] 
**SoftwareDomain** | Pointer to **string** |  | [optional] 
**Software** | Pointer to [**StructuresInstalledSoftwareInfo**](StructuresInstalledSoftwareInfo.md) |  | [optional] 

## Methods

### NewManageVpsInfo

`func NewManageVpsInfo() *ManageVpsInfo`

NewManageVpsInfo instantiates a new ManageVpsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageVpsInfoWithDefaults

`func NewManageVpsInfoWithDefaults() *ManageVpsInfo`

NewManageVpsInfoWithDefaults instantiates a new ManageVpsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManageVpsInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManageVpsInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManageVpsInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ManageVpsInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *ManageVpsInfo) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ManageVpsInfo) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ManageVpsInfo) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ManageVpsInfo) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDisplayName

`func (o *ManageVpsInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ManageVpsInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ManageVpsInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ManageVpsInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHostname

`func (o *ManageVpsInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ManageVpsInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ManageVpsInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ManageVpsInfo) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetConfiguration

`func (o *ManageVpsInfo) GetConfiguration() ManageVpsConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ManageVpsInfo) GetConfigurationOk() (*ManageVpsConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ManageVpsInfo) SetConfiguration(v ManageVpsConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ManageVpsInfo) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetStatus

`func (o *ManageVpsInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManageVpsInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManageVpsInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ManageVpsInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSshKeys

`func (o *ManageVpsInfo) GetSshKeys() []StructuresSshKeyInfo`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *ManageVpsInfo) GetSshKeysOk() (*[]StructuresSshKeyInfo, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *ManageVpsInfo) SetSshKeys(v []StructuresSshKeyInfo)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *ManageVpsInfo) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetHasPassword

`func (o *ManageVpsInfo) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *ManageVpsInfo) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *ManageVpsInfo) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *ManageVpsInfo) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetManageEnabled

`func (o *ManageVpsInfo) GetManageEnabled() bool`

GetManageEnabled returns the ManageEnabled field if non-nil, zero value otherwise.

### GetManageEnabledOk

`func (o *ManageVpsInfo) GetManageEnabledOk() (*bool, bool)`

GetManageEnabledOk returns a tuple with the ManageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageEnabled

`func (o *ManageVpsInfo) SetManageEnabled(v bool)`

SetManageEnabled sets ManageEnabled field to given value.

### HasManageEnabled

`func (o *ManageVpsInfo) HasManageEnabled() bool`

HasManageEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *ManageVpsInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManageVpsInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManageVpsInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManageVpsInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDateCreate

`func (o *ManageVpsInfo) GetDateCreate() string`

GetDateCreate returns the DateCreate field if non-nil, zero value otherwise.

### GetDateCreateOk

`func (o *ManageVpsInfo) GetDateCreateOk() (*string, bool)`

GetDateCreateOk returns a tuple with the DateCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreate

`func (o *ManageVpsInfo) SetDateCreate(v string)`

SetDateCreate sets DateCreate field to given value.

### HasDateCreate

`func (o *ManageVpsInfo) HasDateCreate() bool`

HasDateCreate returns a boolean if a field has been set.

### GetIpAddress

`func (o *ManageVpsInfo) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ManageVpsInfo) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ManageVpsInfo) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ManageVpsInfo) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetRescueMode

`func (o *ManageVpsInfo) GetRescueMode() bool`

GetRescueMode returns the RescueMode field if non-nil, zero value otherwise.

### GetRescueModeOk

`func (o *ManageVpsInfo) GetRescueModeOk() (*bool, bool)`

GetRescueModeOk returns a tuple with the RescueMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRescueMode

`func (o *ManageVpsInfo) SetRescueMode(v bool)`

SetRescueMode sets RescueMode field to given value.

### HasRescueMode

`func (o *ManageVpsInfo) HasRescueMode() bool`

HasRescueMode returns a boolean if a field has been set.

### GetMigrating

`func (o *ManageVpsInfo) GetMigrating() bool`

GetMigrating returns the Migrating field if non-nil, zero value otherwise.

### GetMigratingOk

`func (o *ManageVpsInfo) GetMigratingOk() (*bool, bool)`

GetMigratingOk returns a tuple with the Migrating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrating

`func (o *ManageVpsInfo) SetMigrating(v bool)`

SetMigrating sets Migrating field to given value.

### HasMigrating

`func (o *ManageVpsInfo) HasMigrating() bool`

HasMigrating returns a boolean if a field has been set.

### GetHostUnavailable

`func (o *ManageVpsInfo) GetHostUnavailable() bool`

GetHostUnavailable returns the HostUnavailable field if non-nil, zero value otherwise.

### GetHostUnavailableOk

`func (o *ManageVpsInfo) GetHostUnavailableOk() (*bool, bool)`

GetHostUnavailableOk returns a tuple with the HostUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUnavailable

`func (o *ManageVpsInfo) SetHostUnavailable(v bool)`

SetHostUnavailable sets HostUnavailable field to given value.

### HasHostUnavailable

`func (o *ManageVpsInfo) HasHostUnavailable() bool`

HasHostUnavailable returns a boolean if a field has been set.

### GetUnblocking

`func (o *ManageVpsInfo) GetUnblocking() bool`

GetUnblocking returns the Unblocking field if non-nil, zero value otherwise.

### GetUnblockingOk

`func (o *ManageVpsInfo) GetUnblockingOk() (*bool, bool)`

GetUnblockingOk returns a tuple with the Unblocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnblocking

`func (o *ManageVpsInfo) SetUnblocking(v bool)`

SetUnblocking sets Unblocking field to given value.

### HasUnblocking

`func (o *ManageVpsInfo) HasUnblocking() bool`

HasUnblocking returns a boolean if a field has been set.

### GetRestoring

`func (o *ManageVpsInfo) GetRestoring() bool`

GetRestoring returns the Restoring field if non-nil, zero value otherwise.

### GetRestoringOk

`func (o *ManageVpsInfo) GetRestoringOk() (*bool, bool)`

GetRestoringOk returns a tuple with the Restoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoring

`func (o *ManageVpsInfo) SetRestoring(v bool)`

SetRestoring sets Restoring field to given value.

### HasRestoring

`func (o *ManageVpsInfo) HasRestoring() bool`

HasRestoring returns a boolean if a field has been set.

### GetDiskUsed

`func (o *ManageVpsInfo) GetDiskUsed() string`

GetDiskUsed returns the DiskUsed field if non-nil, zero value otherwise.

### GetDiskUsedOk

`func (o *ManageVpsInfo) GetDiskUsedOk() (*string, bool)`

GetDiskUsedOk returns a tuple with the DiskUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsed

`func (o *ManageVpsInfo) SetDiskUsed(v string)`

SetDiskUsed sets DiskUsed field to given value.

### HasDiskUsed

`func (o *ManageVpsInfo) HasDiskUsed() bool`

HasDiskUsed returns a boolean if a field has been set.

### GetDiskLeft

`func (o *ManageVpsInfo) GetDiskLeft() string`

GetDiskLeft returns the DiskLeft field if non-nil, zero value otherwise.

### GetDiskLeftOk

`func (o *ManageVpsInfo) GetDiskLeftOk() (*string, bool)`

GetDiskLeftOk returns a tuple with the DiskLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskLeft

`func (o *ManageVpsInfo) SetDiskLeft(v string)`

SetDiskLeft sets DiskLeft field to given value.

### HasDiskLeft

`func (o *ManageVpsInfo) HasDiskLeft() bool`

HasDiskLeft returns a boolean if a field has been set.

### GetAdditionalIpAddress

`func (o *ManageVpsInfo) GetAdditionalIpAddress() []string`

GetAdditionalIpAddress returns the AdditionalIpAddress field if non-nil, zero value otherwise.

### GetAdditionalIpAddressOk

`func (o *ManageVpsInfo) GetAdditionalIpAddressOk() (*[]string, bool)`

GetAdditionalIpAddressOk returns a tuple with the AdditionalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalIpAddress

`func (o *ManageVpsInfo) SetAdditionalIpAddress(v []string)`

SetAdditionalIpAddress sets AdditionalIpAddress field to given value.

### HasAdditionalIpAddress

`func (o *ManageVpsInfo) HasAdditionalIpAddress() bool`

HasAdditionalIpAddress returns a boolean if a field has been set.

### GetBegetSshAccessAllowed

`func (o *ManageVpsInfo) GetBegetSshAccessAllowed() bool`

GetBegetSshAccessAllowed returns the BegetSshAccessAllowed field if non-nil, zero value otherwise.

### GetBegetSshAccessAllowedOk

`func (o *ManageVpsInfo) GetBegetSshAccessAllowedOk() (*bool, bool)`

GetBegetSshAccessAllowedOk returns a tuple with the BegetSshAccessAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBegetSshAccessAllowed

`func (o *ManageVpsInfo) SetBegetSshAccessAllowed(v bool)`

SetBegetSshAccessAllowed sets BegetSshAccessAllowed field to given value.

### HasBegetSshAccessAllowed

`func (o *ManageVpsInfo) HasBegetSshAccessAllowed() bool`

HasBegetSshAccessAllowed returns a boolean if a field has been set.

### GetArchived

`func (o *ManageVpsInfo) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ManageVpsInfo) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ManageVpsInfo) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ManageVpsInfo) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetUnarchiving

`func (o *ManageVpsInfo) GetUnarchiving() bool`

GetUnarchiving returns the Unarchiving field if non-nil, zero value otherwise.

### GetUnarchivingOk

`func (o *ManageVpsInfo) GetUnarchivingOk() (*bool, bool)`

GetUnarchivingOk returns a tuple with the Unarchiving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnarchiving

`func (o *ManageVpsInfo) SetUnarchiving(v bool)`

SetUnarchiving sets Unarchiving field to given value.

### HasUnarchiving

`func (o *ManageVpsInfo) HasUnarchiving() bool`

HasUnarchiving returns a boolean if a field has been set.

### GetPrivateNetwork

`func (o *ManageVpsInfo) GetPrivateNetwork() []StructuresAttachedPrivateNetwork`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *ManageVpsInfo) GetPrivateNetworkOk() (*[]StructuresAttachedPrivateNetwork, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *ManageVpsInfo) SetPrivateNetwork(v []StructuresAttachedPrivateNetwork)`

SetPrivateNetwork sets PrivateNetwork field to given value.

### HasPrivateNetwork

`func (o *ManageVpsInfo) HasPrivateNetwork() bool`

HasPrivateNetwork returns a boolean if a field has been set.

### GetTechnicalDomain

`func (o *ManageVpsInfo) GetTechnicalDomain() string`

GetTechnicalDomain returns the TechnicalDomain field if non-nil, zero value otherwise.

### GetTechnicalDomainOk

`func (o *ManageVpsInfo) GetTechnicalDomainOk() (*string, bool)`

GetTechnicalDomainOk returns a tuple with the TechnicalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalDomain

`func (o *ManageVpsInfo) SetTechnicalDomain(v string)`

SetTechnicalDomain sets TechnicalDomain field to given value.

### HasTechnicalDomain

`func (o *ManageVpsInfo) HasTechnicalDomain() bool`

HasTechnicalDomain returns a boolean if a field has been set.

### GetSoftwareDomain

`func (o *ManageVpsInfo) GetSoftwareDomain() string`

GetSoftwareDomain returns the SoftwareDomain field if non-nil, zero value otherwise.

### GetSoftwareDomainOk

`func (o *ManageVpsInfo) GetSoftwareDomainOk() (*string, bool)`

GetSoftwareDomainOk returns a tuple with the SoftwareDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareDomain

`func (o *ManageVpsInfo) SetSoftwareDomain(v string)`

SetSoftwareDomain sets SoftwareDomain field to given value.

### HasSoftwareDomain

`func (o *ManageVpsInfo) HasSoftwareDomain() bool`

HasSoftwareDomain returns a boolean if a field has been set.

### GetSoftware

`func (o *ManageVpsInfo) GetSoftware() StructuresInstalledSoftwareInfo`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *ManageVpsInfo) GetSoftwareOk() (*StructuresInstalledSoftwareInfo, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *ManageVpsInfo) SetSoftware(v StructuresInstalledSoftwareInfo)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *ManageVpsInfo) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


