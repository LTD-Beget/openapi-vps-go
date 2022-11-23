# ManageCreateVpsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ConfigurationId** | Pointer to **string** |  | [optional] 
**Software** | Pointer to [**ManageSoftwareInstallInfo**](ManageSoftwareInstallInfo.md) |  | [optional] 
**SnapshotId** | Pointer to **string** |  | [optional] 
**SshKeys** | Pointer to **[]int32** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**BegetSshAccessAllowed** | Pointer to **bool** |  | [optional] 
**PrivateNetworks** | Pointer to [**[]ManagePrivateNetworkInfo**](ManagePrivateNetworkInfo.md) |  | [optional] 

## Methods

### NewManageCreateVpsRequest

`func NewManageCreateVpsRequest() *ManageCreateVpsRequest`

NewManageCreateVpsRequest instantiates a new ManageCreateVpsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageCreateVpsRequestWithDefaults

`func NewManageCreateVpsRequestWithDefaults() *ManageCreateVpsRequest`

NewManageCreateVpsRequestWithDefaults instantiates a new ManageCreateVpsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ManageCreateVpsRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ManageCreateVpsRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ManageCreateVpsRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ManageCreateVpsRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHostname

`func (o *ManageCreateVpsRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ManageCreateVpsRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ManageCreateVpsRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ManageCreateVpsRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetDescription

`func (o *ManageCreateVpsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManageCreateVpsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManageCreateVpsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManageCreateVpsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfigurationId

`func (o *ManageCreateVpsRequest) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *ManageCreateVpsRequest) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *ManageCreateVpsRequest) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *ManageCreateVpsRequest) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### GetSoftware

`func (o *ManageCreateVpsRequest) GetSoftware() ManageSoftwareInstallInfo`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *ManageCreateVpsRequest) GetSoftwareOk() (*ManageSoftwareInstallInfo, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *ManageCreateVpsRequest) SetSoftware(v ManageSoftwareInstallInfo)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *ManageCreateVpsRequest) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### GetSnapshotId

`func (o *ManageCreateVpsRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ManageCreateVpsRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ManageCreateVpsRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ManageCreateVpsRequest) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetSshKeys

`func (o *ManageCreateVpsRequest) GetSshKeys() []int32`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *ManageCreateVpsRequest) GetSshKeysOk() (*[]int32, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *ManageCreateVpsRequest) SetSshKeys(v []int32)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *ManageCreateVpsRequest) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetPassword

`func (o *ManageCreateVpsRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ManageCreateVpsRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ManageCreateVpsRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ManageCreateVpsRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetBegetSshAccessAllowed

`func (o *ManageCreateVpsRequest) GetBegetSshAccessAllowed() bool`

GetBegetSshAccessAllowed returns the BegetSshAccessAllowed field if non-nil, zero value otherwise.

### GetBegetSshAccessAllowedOk

`func (o *ManageCreateVpsRequest) GetBegetSshAccessAllowedOk() (*bool, bool)`

GetBegetSshAccessAllowedOk returns a tuple with the BegetSshAccessAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBegetSshAccessAllowed

`func (o *ManageCreateVpsRequest) SetBegetSshAccessAllowed(v bool)`

SetBegetSshAccessAllowed sets BegetSshAccessAllowed field to given value.

### HasBegetSshAccessAllowed

`func (o *ManageCreateVpsRequest) HasBegetSshAccessAllowed() bool`

HasBegetSshAccessAllowed returns a boolean if a field has been set.

### GetPrivateNetworks

`func (o *ManageCreateVpsRequest) GetPrivateNetworks() []ManagePrivateNetworkInfo`

GetPrivateNetworks returns the PrivateNetworks field if non-nil, zero value otherwise.

### GetPrivateNetworksOk

`func (o *ManageCreateVpsRequest) GetPrivateNetworksOk() (*[]ManagePrivateNetworkInfo, bool)`

GetPrivateNetworksOk returns a tuple with the PrivateNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworks

`func (o *ManageCreateVpsRequest) SetPrivateNetworks(v []ManagePrivateNetworkInfo)`

SetPrivateNetworks sets PrivateNetworks field to given value.

### HasPrivateNetworks

`func (o *ManageCreateVpsRequest) HasPrivateNetworks() bool`

HasPrivateNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


