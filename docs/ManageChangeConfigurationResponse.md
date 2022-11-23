# ManageChangeConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vps** | Pointer to [**ManageVpsInfo**](ManageVpsInfo.md) |  | [optional] 
**Error** | Pointer to [**ManageChangeConfigurationResponseError**](ManageChangeConfigurationResponseError.md) |  | [optional] 

## Methods

### NewManageChangeConfigurationResponse

`func NewManageChangeConfigurationResponse() *ManageChangeConfigurationResponse`

NewManageChangeConfigurationResponse instantiates a new ManageChangeConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageChangeConfigurationResponseWithDefaults

`func NewManageChangeConfigurationResponseWithDefaults() *ManageChangeConfigurationResponse`

NewManageChangeConfigurationResponseWithDefaults instantiates a new ManageChangeConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVps

`func (o *ManageChangeConfigurationResponse) GetVps() ManageVpsInfo`

GetVps returns the Vps field if non-nil, zero value otherwise.

### GetVpsOk

`func (o *ManageChangeConfigurationResponse) GetVpsOk() (*ManageVpsInfo, bool)`

GetVpsOk returns a tuple with the Vps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVps

`func (o *ManageChangeConfigurationResponse) SetVps(v ManageVpsInfo)`

SetVps sets Vps field to given value.

### HasVps

`func (o *ManageChangeConfigurationResponse) HasVps() bool`

HasVps returns a boolean if a field has been set.

### GetError

`func (o *ManageChangeConfigurationResponse) GetError() ManageChangeConfigurationResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ManageChangeConfigurationResponse) GetErrorOk() (*ManageChangeConfigurationResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ManageChangeConfigurationResponse) SetError(v ManageChangeConfigurationResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *ManageChangeConfigurationResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


