# NetworkRemoveIpAddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**NetworkRemoveIpAddressResponseError**](NetworkRemoveIpAddressResponseError.md) |  | [optional] 

## Methods

### NewNetworkRemoveIpAddressResponse

`func NewNetworkRemoveIpAddressResponse() *NetworkRemoveIpAddressResponse`

NewNetworkRemoveIpAddressResponse instantiates a new NetworkRemoveIpAddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRemoveIpAddressResponseWithDefaults

`func NewNetworkRemoveIpAddressResponseWithDefaults() *NetworkRemoveIpAddressResponse`

NewNetworkRemoveIpAddressResponseWithDefaults instantiates a new NetworkRemoveIpAddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *NetworkRemoveIpAddressResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkRemoveIpAddressResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkRemoveIpAddressResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NetworkRemoveIpAddressResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetError

`func (o *NetworkRemoveIpAddressResponse) GetError() NetworkRemoveIpAddressResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *NetworkRemoveIpAddressResponse) GetErrorOk() (*NetworkRemoveIpAddressResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *NetworkRemoveIpAddressResponse) SetError(v NetworkRemoveIpAddressResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *NetworkRemoveIpAddressResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


