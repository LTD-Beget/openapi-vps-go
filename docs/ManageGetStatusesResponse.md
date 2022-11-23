# ManageGetStatusesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | Pointer to [**[]ManageGetStatusesResponseStatusInfo**](ManageGetStatusesResponseStatusInfo.md) |  | [optional] 

## Methods

### NewManageGetStatusesResponse

`func NewManageGetStatusesResponse() *ManageGetStatusesResponse`

NewManageGetStatusesResponse instantiates a new ManageGetStatusesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageGetStatusesResponseWithDefaults

`func NewManageGetStatusesResponseWithDefaults() *ManageGetStatusesResponse`

NewManageGetStatusesResponseWithDefaults instantiates a new ManageGetStatusesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *ManageGetStatusesResponse) GetStatuses() []ManageGetStatusesResponseStatusInfo`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ManageGetStatusesResponse) GetStatusesOk() (*[]ManageGetStatusesResponseStatusInfo, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ManageGetStatusesResponse) SetStatuses(v []ManageGetStatusesResponseStatusInfo)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *ManageGetStatusesResponse) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


