# BackupRestoreFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**StructuresOrderInfo**](StructuresOrderInfo.md) |  | [optional] 
**Error** | Pointer to [**BackupRestoreFileResponseError**](BackupRestoreFileResponseError.md) |  | [optional] 

## Methods

### NewBackupRestoreFileResponse

`func NewBackupRestoreFileResponse() *BackupRestoreFileResponse`

NewBackupRestoreFileResponse instantiates a new BackupRestoreFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreFileResponseWithDefaults

`func NewBackupRestoreFileResponseWithDefaults() *BackupRestoreFileResponse`

NewBackupRestoreFileResponseWithDefaults instantiates a new BackupRestoreFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *BackupRestoreFileResponse) GetOrder() StructuresOrderInfo`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *BackupRestoreFileResponse) GetOrderOk() (*StructuresOrderInfo, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *BackupRestoreFileResponse) SetOrder(v StructuresOrderInfo)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *BackupRestoreFileResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetError

`func (o *BackupRestoreFileResponse) GetError() BackupRestoreFileResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BackupRestoreFileResponse) GetErrorOk() (*BackupRestoreFileResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BackupRestoreFileResponse) SetError(v BackupRestoreFileResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *BackupRestoreFileResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


