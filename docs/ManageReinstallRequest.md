# ManageReinstallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SshKeys** | Pointer to **[]int32** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Software** | Pointer to [**ManageSoftwareInstallInfo**](ManageSoftwareInstallInfo.md) |  | [optional] 

## Methods

### NewManageReinstallRequest

`func NewManageReinstallRequest() *ManageReinstallRequest`

NewManageReinstallRequest instantiates a new ManageReinstallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageReinstallRequestWithDefaults

`func NewManageReinstallRequestWithDefaults() *ManageReinstallRequest`

NewManageReinstallRequestWithDefaults instantiates a new ManageReinstallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManageReinstallRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManageReinstallRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManageReinstallRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ManageReinstallRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSshKeys

`func (o *ManageReinstallRequest) GetSshKeys() []int32`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *ManageReinstallRequest) GetSshKeysOk() (*[]int32, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *ManageReinstallRequest) SetSshKeys(v []int32)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *ManageReinstallRequest) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetPassword

`func (o *ManageReinstallRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ManageReinstallRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ManageReinstallRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ManageReinstallRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSoftware

`func (o *ManageReinstallRequest) GetSoftware() ManageSoftwareInstallInfo`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *ManageReinstallRequest) GetSoftwareOk() (*ManageSoftwareInstallInfo, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *ManageReinstallRequest) SetSoftware(v ManageSoftwareInstallInfo)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *ManageReinstallRequest) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


