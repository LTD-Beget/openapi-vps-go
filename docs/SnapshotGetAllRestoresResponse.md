# SnapshotGetAllRestoresResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Restore** | Pointer to [**[]SnapshotRestore**](SnapshotRestore.md) |  | [optional] 

## Methods

### NewSnapshotGetAllRestoresResponse

`func NewSnapshotGetAllRestoresResponse() *SnapshotGetAllRestoresResponse`

NewSnapshotGetAllRestoresResponse instantiates a new SnapshotGetAllRestoresResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotGetAllRestoresResponseWithDefaults

`func NewSnapshotGetAllRestoresResponseWithDefaults() *SnapshotGetAllRestoresResponse`

NewSnapshotGetAllRestoresResponseWithDefaults instantiates a new SnapshotGetAllRestoresResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestore

`func (o *SnapshotGetAllRestoresResponse) GetRestore() []SnapshotRestore`

GetRestore returns the Restore field if non-nil, zero value otherwise.

### GetRestoreOk

`func (o *SnapshotGetAllRestoresResponse) GetRestoreOk() (*[]SnapshotRestore, bool)`

GetRestoreOk returns a tuple with the Restore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestore

`func (o *SnapshotGetAllRestoresResponse) SetRestore(v []SnapshotRestore)`

SetRestore sets Restore field to given value.

### HasRestore

`func (o *SnapshotGetAllRestoresResponse) HasRestore() bool`

HasRestore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


