# BackupRestoreFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CopyId** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **[]string** |  | [optional] 
**AffectLive** | Pointer to **bool** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 

## Methods

### NewBackupRestoreFileRequest

`func NewBackupRestoreFileRequest() *BackupRestoreFileRequest`

NewBackupRestoreFileRequest instantiates a new BackupRestoreFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreFileRequestWithDefaults

`func NewBackupRestoreFileRequestWithDefaults() *BackupRestoreFileRequest`

NewBackupRestoreFileRequestWithDefaults instantiates a new BackupRestoreFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupRestoreFileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupRestoreFileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupRestoreFileRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackupRestoreFileRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCopyId

`func (o *BackupRestoreFileRequest) GetCopyId() string`

GetCopyId returns the CopyId field if non-nil, zero value otherwise.

### GetCopyIdOk

`func (o *BackupRestoreFileRequest) GetCopyIdOk() (*string, bool)`

GetCopyIdOk returns a tuple with the CopyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyId

`func (o *BackupRestoreFileRequest) SetCopyId(v string)`

SetCopyId sets CopyId field to given value.

### HasCopyId

`func (o *BackupRestoreFileRequest) HasCopyId() bool`

HasCopyId returns a boolean if a field has been set.

### GetPath

`func (o *BackupRestoreFileRequest) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BackupRestoreFileRequest) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BackupRestoreFileRequest) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BackupRestoreFileRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetAffectLive

`func (o *BackupRestoreFileRequest) GetAffectLive() bool`

GetAffectLive returns the AffectLive field if non-nil, zero value otherwise.

### GetAffectLiveOk

`func (o *BackupRestoreFileRequest) GetAffectLiveOk() (*bool, bool)`

GetAffectLiveOk returns a tuple with the AffectLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectLive

`func (o *BackupRestoreFileRequest) SetAffectLive(v bool)`

SetAffectLive sets AffectLive field to given value.

### HasAffectLive

`func (o *BackupRestoreFileRequest) HasAffectLive() bool`

HasAffectLive returns a boolean if a field has been set.

### GetTarget

`func (o *BackupRestoreFileRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *BackupRestoreFileRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *BackupRestoreFileRequest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *BackupRestoreFileRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


