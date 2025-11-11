# ManageChangePinnedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vps** | Pointer to [**ManageVpsInfo**](ManageVpsInfo.md) |  | [optional] 
**Error** | Pointer to [**ManageChangePinnedResponseError**](ManageChangePinnedResponseError.md) |  | [optional] 

## Methods

### NewManageChangePinnedResponse

`func NewManageChangePinnedResponse() *ManageChangePinnedResponse`

NewManageChangePinnedResponse instantiates a new ManageChangePinnedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageChangePinnedResponseWithDefaults

`func NewManageChangePinnedResponseWithDefaults() *ManageChangePinnedResponse`

NewManageChangePinnedResponseWithDefaults instantiates a new ManageChangePinnedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVps

`func (o *ManageChangePinnedResponse) GetVps() ManageVpsInfo`

GetVps returns the Vps field if non-nil, zero value otherwise.

### GetVpsOk

`func (o *ManageChangePinnedResponse) GetVpsOk() (*ManageVpsInfo, bool)`

GetVpsOk returns a tuple with the Vps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVps

`func (o *ManageChangePinnedResponse) SetVps(v ManageVpsInfo)`

SetVps sets Vps field to given value.

### HasVps

`func (o *ManageChangePinnedResponse) HasVps() bool`

HasVps returns a boolean if a field has been set.

### GetError

`func (o *ManageChangePinnedResponse) GetError() ManageChangePinnedResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ManageChangePinnedResponse) GetErrorOk() (*ManageChangePinnedResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ManageChangePinnedResponse) SetError(v ManageChangePinnedResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *ManageChangePinnedResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


