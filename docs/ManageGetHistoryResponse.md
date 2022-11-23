# ManageGetHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**History** | Pointer to [**[]ManageHistoryItem**](ManageHistoryItem.md) |  | [optional] 

## Methods

### NewManageGetHistoryResponse

`func NewManageGetHistoryResponse() *ManageGetHistoryResponse`

NewManageGetHistoryResponse instantiates a new ManageGetHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageGetHistoryResponseWithDefaults

`func NewManageGetHistoryResponseWithDefaults() *ManageGetHistoryResponse`

NewManageGetHistoryResponseWithDefaults instantiates a new ManageGetHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistory

`func (o *ManageGetHistoryResponse) GetHistory() []ManageHistoryItem`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *ManageGetHistoryResponse) GetHistoryOk() (*[]ManageHistoryItem, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *ManageGetHistoryResponse) SetHistory(v []ManageHistoryItem)`

SetHistory sets History field to given value.

### HasHistory

`func (o *ManageGetHistoryResponse) HasHistory() bool`

HasHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


