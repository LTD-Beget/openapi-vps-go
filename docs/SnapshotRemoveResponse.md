# SnapshotRemoveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshot** | Pointer to [**SnapshotSnapshot**](SnapshotSnapshot.md) |  | [optional] 
**Error** | Pointer to [**SnapshotRemoveResponseError**](SnapshotRemoveResponseError.md) |  | [optional] 

## Methods

### NewSnapshotRemoveResponse

`func NewSnapshotRemoveResponse() *SnapshotRemoveResponse`

NewSnapshotRemoveResponse instantiates a new SnapshotRemoveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotRemoveResponseWithDefaults

`func NewSnapshotRemoveResponseWithDefaults() *SnapshotRemoveResponse`

NewSnapshotRemoveResponseWithDefaults instantiates a new SnapshotRemoveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshot

`func (o *SnapshotRemoveResponse) GetSnapshot() SnapshotSnapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *SnapshotRemoveResponse) GetSnapshotOk() (*SnapshotSnapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *SnapshotRemoveResponse) SetSnapshot(v SnapshotSnapshot)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *SnapshotRemoveResponse) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetError

`func (o *SnapshotRemoveResponse) GetError() SnapshotRemoveResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SnapshotRemoveResponse) GetErrorOk() (*SnapshotRemoveResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SnapshotRemoveResponse) SetError(v SnapshotRemoveResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *SnapshotRemoveResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


