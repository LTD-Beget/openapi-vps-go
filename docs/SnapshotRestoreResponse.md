# SnapshotRestoreResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Restore** | Pointer to [**SnapshotRestore**](SnapshotRestore.md) |  | [optional] 
**Error** | Pointer to [**SnapshotRestoreResponseError**](SnapshotRestoreResponseError.md) |  | [optional] 

## Methods

### NewSnapshotRestoreResponse

`func NewSnapshotRestoreResponse() *SnapshotRestoreResponse`

NewSnapshotRestoreResponse instantiates a new SnapshotRestoreResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotRestoreResponseWithDefaults

`func NewSnapshotRestoreResponseWithDefaults() *SnapshotRestoreResponse`

NewSnapshotRestoreResponseWithDefaults instantiates a new SnapshotRestoreResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestore

`func (o *SnapshotRestoreResponse) GetRestore() SnapshotRestore`

GetRestore returns the Restore field if non-nil, zero value otherwise.

### GetRestoreOk

`func (o *SnapshotRestoreResponse) GetRestoreOk() (*SnapshotRestore, bool)`

GetRestoreOk returns a tuple with the Restore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestore

`func (o *SnapshotRestoreResponse) SetRestore(v SnapshotRestore)`

SetRestore sets Restore field to given value.

### HasRestore

`func (o *SnapshotRestoreResponse) HasRestore() bool`

HasRestore returns a boolean if a field has been set.

### GetError

`func (o *SnapshotRestoreResponse) GetError() SnapshotRestoreResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SnapshotRestoreResponse) GetErrorOk() (*SnapshotRestoreResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SnapshotRestoreResponse) SetError(v SnapshotRestoreResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *SnapshotRestoreResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


