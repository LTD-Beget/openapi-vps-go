# ManageStartRescueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vps** | Pointer to [**ManageVpsInfo**](ManageVpsInfo.md) |  | [optional] 
**Error** | Pointer to [**ManageStartRescueResponseError**](ManageStartRescueResponseError.md) |  | [optional] 

## Methods

### NewManageStartRescueResponse

`func NewManageStartRescueResponse() *ManageStartRescueResponse`

NewManageStartRescueResponse instantiates a new ManageStartRescueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageStartRescueResponseWithDefaults

`func NewManageStartRescueResponseWithDefaults() *ManageStartRescueResponse`

NewManageStartRescueResponseWithDefaults instantiates a new ManageStartRescueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVps

`func (o *ManageStartRescueResponse) GetVps() ManageVpsInfo`

GetVps returns the Vps field if non-nil, zero value otherwise.

### GetVpsOk

`func (o *ManageStartRescueResponse) GetVpsOk() (*ManageVpsInfo, bool)`

GetVpsOk returns a tuple with the Vps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVps

`func (o *ManageStartRescueResponse) SetVps(v ManageVpsInfo)`

SetVps sets Vps field to given value.

### HasVps

`func (o *ManageStartRescueResponse) HasVps() bool`

HasVps returns a boolean if a field has been set.

### GetError

`func (o *ManageStartRescueResponse) GetError() ManageStartRescueResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ManageStartRescueResponse) GetErrorOk() (*ManageStartRescueResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ManageStartRescueResponse) SetError(v ManageStartRescueResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *ManageStartRescueResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


