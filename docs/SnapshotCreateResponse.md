# SnapshotCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshot** | Pointer to [**SnapshotSnapshot**](SnapshotSnapshot.md) |  | [optional] 
**Error** | Pointer to [**SnapshotCreateResponseError**](SnapshotCreateResponseError.md) |  | [optional] 

## Methods

### NewSnapshotCreateResponse

`func NewSnapshotCreateResponse() *SnapshotCreateResponse`

NewSnapshotCreateResponse instantiates a new SnapshotCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotCreateResponseWithDefaults

`func NewSnapshotCreateResponseWithDefaults() *SnapshotCreateResponse`

NewSnapshotCreateResponseWithDefaults instantiates a new SnapshotCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshot

`func (o *SnapshotCreateResponse) GetSnapshot() SnapshotSnapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *SnapshotCreateResponse) GetSnapshotOk() (*SnapshotSnapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *SnapshotCreateResponse) SetSnapshot(v SnapshotSnapshot)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *SnapshotCreateResponse) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetError

`func (o *SnapshotCreateResponse) GetError() SnapshotCreateResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SnapshotCreateResponse) GetErrorOk() (*SnapshotCreateResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SnapshotCreateResponse) SetError(v SnapshotCreateResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *SnapshotCreateResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


