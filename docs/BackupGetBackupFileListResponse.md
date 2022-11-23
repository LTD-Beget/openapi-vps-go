# BackupGetBackupFileListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to [**[]StructuresFileInfo**](StructuresFileInfo.md) |  | [optional] 

## Methods

### NewBackupGetBackupFileListResponse

`func NewBackupGetBackupFileListResponse() *BackupGetBackupFileListResponse`

NewBackupGetBackupFileListResponse instantiates a new BackupGetBackupFileListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupGetBackupFileListResponseWithDefaults

`func NewBackupGetBackupFileListResponseWithDefaults() *BackupGetBackupFileListResponse`

NewBackupGetBackupFileListResponseWithDefaults instantiates a new BackupGetBackupFileListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *BackupGetBackupFileListResponse) GetFile() []StructuresFileInfo`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *BackupGetBackupFileListResponse) GetFileOk() (*[]StructuresFileInfo, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *BackupGetBackupFileListResponse) SetFile(v []StructuresFileInfo)`

SetFile sets File field to given value.

### HasFile

`func (o *BackupGetBackupFileListResponse) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


