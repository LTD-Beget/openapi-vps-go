# BackupGetAvailableCopiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Copy** | Pointer to [**[]StructuresCopyInfo**](StructuresCopyInfo.md) |  | [optional] 

## Methods

### NewBackupGetAvailableCopiesResponse

`func NewBackupGetAvailableCopiesResponse() *BackupGetAvailableCopiesResponse`

NewBackupGetAvailableCopiesResponse instantiates a new BackupGetAvailableCopiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupGetAvailableCopiesResponseWithDefaults

`func NewBackupGetAvailableCopiesResponseWithDefaults() *BackupGetAvailableCopiesResponse`

NewBackupGetAvailableCopiesResponseWithDefaults instantiates a new BackupGetAvailableCopiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopy

`func (o *BackupGetAvailableCopiesResponse) GetCopy() []StructuresCopyInfo`

GetCopy returns the Copy field if non-nil, zero value otherwise.

### GetCopyOk

`func (o *BackupGetAvailableCopiesResponse) GetCopyOk() (*[]StructuresCopyInfo, bool)`

GetCopyOk returns a tuple with the Copy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopy

`func (o *BackupGetAvailableCopiesResponse) SetCopy(v []StructuresCopyInfo)`

SetCopy sets Copy field to given value.

### HasCopy

`func (o *BackupGetAvailableCopiesResponse) HasCopy() bool`

HasCopy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


