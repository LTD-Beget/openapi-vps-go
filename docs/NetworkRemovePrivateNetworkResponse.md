# NetworkRemovePrivateNetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**StructuresPrivateNetwork**](StructuresPrivateNetwork.md) |  | [optional] 
**Error** | Pointer to [**NetworkRemovePrivateNetworkResponseError**](NetworkRemovePrivateNetworkResponseError.md) |  | [optional] 

## Methods

### NewNetworkRemovePrivateNetworkResponse

`func NewNetworkRemovePrivateNetworkResponse() *NetworkRemovePrivateNetworkResponse`

NewNetworkRemovePrivateNetworkResponse instantiates a new NetworkRemovePrivateNetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRemovePrivateNetworkResponseWithDefaults

`func NewNetworkRemovePrivateNetworkResponseWithDefaults() *NetworkRemovePrivateNetworkResponse`

NewNetworkRemovePrivateNetworkResponseWithDefaults instantiates a new NetworkRemovePrivateNetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *NetworkRemovePrivateNetworkResponse) GetNetwork() StructuresPrivateNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkRemovePrivateNetworkResponse) GetNetworkOk() (*StructuresPrivateNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkRemovePrivateNetworkResponse) SetNetwork(v StructuresPrivateNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NetworkRemovePrivateNetworkResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetError

`func (o *NetworkRemovePrivateNetworkResponse) GetError() NetworkRemovePrivateNetworkResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *NetworkRemovePrivateNetworkResponse) GetErrorOk() (*NetworkRemovePrivateNetworkResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *NetworkRemovePrivateNetworkResponse) SetError(v NetworkRemovePrivateNetworkResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *NetworkRemovePrivateNetworkResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


