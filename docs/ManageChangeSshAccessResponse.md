# ManageChangeSshAccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vps** | Pointer to [**ManageVpsInfo**](ManageVpsInfo.md) |  | [optional] 
**Error** | Pointer to [**ManageChangeSshAccessResponseError**](ManageChangeSshAccessResponseError.md) |  | [optional] 

## Methods

### NewManageChangeSshAccessResponse

`func NewManageChangeSshAccessResponse() *ManageChangeSshAccessResponse`

NewManageChangeSshAccessResponse instantiates a new ManageChangeSshAccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageChangeSshAccessResponseWithDefaults

`func NewManageChangeSshAccessResponseWithDefaults() *ManageChangeSshAccessResponse`

NewManageChangeSshAccessResponseWithDefaults instantiates a new ManageChangeSshAccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVps

`func (o *ManageChangeSshAccessResponse) GetVps() ManageVpsInfo`

GetVps returns the Vps field if non-nil, zero value otherwise.

### GetVpsOk

`func (o *ManageChangeSshAccessResponse) GetVpsOk() (*ManageVpsInfo, bool)`

GetVpsOk returns a tuple with the Vps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVps

`func (o *ManageChangeSshAccessResponse) SetVps(v ManageVpsInfo)`

SetVps sets Vps field to given value.

### HasVps

`func (o *ManageChangeSshAccessResponse) HasVps() bool`

HasVps returns a boolean if a field has been set.

### GetError

`func (o *ManageChangeSshAccessResponse) GetError() ManageChangeSshAccessResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ManageChangeSshAccessResponse) GetErrorOk() (*ManageChangeSshAccessResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ManageChangeSshAccessResponse) SetError(v ManageChangeSshAccessResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *ManageChangeSshAccessResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


