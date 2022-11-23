# ManageGetListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vps** | Pointer to [**[]ManageVpsInfo**](ManageVpsInfo.md) |  | [optional] 

## Methods

### NewManageGetListResponse

`func NewManageGetListResponse() *ManageGetListResponse`

NewManageGetListResponse instantiates a new ManageGetListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageGetListResponseWithDefaults

`func NewManageGetListResponseWithDefaults() *ManageGetListResponse`

NewManageGetListResponseWithDefaults instantiates a new ManageGetListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVps

`func (o *ManageGetListResponse) GetVps() []ManageVpsInfo`

GetVps returns the Vps field if non-nil, zero value otherwise.

### GetVpsOk

`func (o *ManageGetListResponse) GetVpsOk() (*[]ManageVpsInfo, bool)`

GetVpsOk returns a tuple with the Vps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVps

`func (o *ManageGetListResponse) SetVps(v []ManageVpsInfo)`

SetVps sets Vps field to given value.

### HasVps

`func (o *ManageGetListResponse) HasVps() bool`

HasVps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


