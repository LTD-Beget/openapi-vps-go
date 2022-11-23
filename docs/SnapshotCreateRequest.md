# SnapshotCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpsId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Consistent** | Pointer to **bool** |  | [optional] 

## Methods

### NewSnapshotCreateRequest

`func NewSnapshotCreateRequest() *SnapshotCreateRequest`

NewSnapshotCreateRequest instantiates a new SnapshotCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotCreateRequestWithDefaults

`func NewSnapshotCreateRequestWithDefaults() *SnapshotCreateRequest`

NewSnapshotCreateRequestWithDefaults instantiates a new SnapshotCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpsId

`func (o *SnapshotCreateRequest) GetVpsId() string`

GetVpsId returns the VpsId field if non-nil, zero value otherwise.

### GetVpsIdOk

`func (o *SnapshotCreateRequest) GetVpsIdOk() (*string, bool)`

GetVpsIdOk returns a tuple with the VpsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpsId

`func (o *SnapshotCreateRequest) SetVpsId(v string)`

SetVpsId sets VpsId field to given value.

### HasVpsId

`func (o *SnapshotCreateRequest) HasVpsId() bool`

HasVpsId returns a boolean if a field has been set.

### GetDescription

`func (o *SnapshotCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnapshotCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnapshotCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SnapshotCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConsistent

`func (o *SnapshotCreateRequest) GetConsistent() bool`

GetConsistent returns the Consistent field if non-nil, zero value otherwise.

### GetConsistentOk

`func (o *SnapshotCreateRequest) GetConsistentOk() (*bool, bool)`

GetConsistentOk returns a tuple with the Consistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistent

`func (o *SnapshotCreateRequest) SetConsistent(v bool)`

SetConsistent sets Consistent field to given value.

### HasConsistent

`func (o *SnapshotCreateRequest) HasConsistent() bool`

HasConsistent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


