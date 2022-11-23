# BackupRestoreServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**StructuresOrderInfo**](StructuresOrderInfo.md) |  | [optional] 
**Error** | Pointer to [**BackupRestoreServerResponseError**](BackupRestoreServerResponseError.md) |  | [optional] 

## Methods

### NewBackupRestoreServerResponse

`func NewBackupRestoreServerResponse() *BackupRestoreServerResponse`

NewBackupRestoreServerResponse instantiates a new BackupRestoreServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreServerResponseWithDefaults

`func NewBackupRestoreServerResponseWithDefaults() *BackupRestoreServerResponse`

NewBackupRestoreServerResponseWithDefaults instantiates a new BackupRestoreServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *BackupRestoreServerResponse) GetOrder() StructuresOrderInfo`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *BackupRestoreServerResponse) GetOrderOk() (*StructuresOrderInfo, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *BackupRestoreServerResponse) SetOrder(v StructuresOrderInfo)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *BackupRestoreServerResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetError

`func (o *BackupRestoreServerResponse) GetError() BackupRestoreServerResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BackupRestoreServerResponse) GetErrorOk() (*BackupRestoreServerResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BackupRestoreServerResponse) SetError(v BackupRestoreServerResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *BackupRestoreServerResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


