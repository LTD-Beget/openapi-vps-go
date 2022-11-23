# NetworkOrderIpAddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**NetworkOrderIpAddressResponseError**](NetworkOrderIpAddressResponseError.md) |  | [optional] 

## Methods

### NewNetworkOrderIpAddressResponse

`func NewNetworkOrderIpAddressResponse() *NetworkOrderIpAddressResponse`

NewNetworkOrderIpAddressResponse instantiates a new NetworkOrderIpAddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkOrderIpAddressResponseWithDefaults

`func NewNetworkOrderIpAddressResponseWithDefaults() *NetworkOrderIpAddressResponse`

NewNetworkOrderIpAddressResponseWithDefaults instantiates a new NetworkOrderIpAddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *NetworkOrderIpAddressResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkOrderIpAddressResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkOrderIpAddressResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NetworkOrderIpAddressResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetError

`func (o *NetworkOrderIpAddressResponse) GetError() NetworkOrderIpAddressResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *NetworkOrderIpAddressResponse) GetErrorOk() (*NetworkOrderIpAddressResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *NetworkOrderIpAddressResponse) SetError(v NetworkOrderIpAddressResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *NetworkOrderIpAddressResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


