# NetworkGetNetworkInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to [**[]StructuresIpInfo**](StructuresIpInfo.md) |  | [optional] 
**AdditionalIp** | Pointer to [**[]StructuresAdditionalIpInfo**](StructuresAdditionalIpInfo.md) |  | [optional] 
**PrivateNetwork** | Pointer to [**[]StructuresPrivateNetwork**](StructuresPrivateNetwork.md) |  | [optional] 

## Methods

### NewNetworkGetNetworkInfoResponse

`func NewNetworkGetNetworkInfoResponse() *NetworkGetNetworkInfoResponse`

NewNetworkGetNetworkInfoResponse instantiates a new NetworkGetNetworkInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkGetNetworkInfoResponseWithDefaults

`func NewNetworkGetNetworkInfoResponseWithDefaults() *NetworkGetNetworkInfoResponse`

NewNetworkGetNetworkInfoResponseWithDefaults instantiates a new NetworkGetNetworkInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *NetworkGetNetworkInfoResponse) GetIp() []StructuresIpInfo`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NetworkGetNetworkInfoResponse) GetIpOk() (*[]StructuresIpInfo, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NetworkGetNetworkInfoResponse) SetIp(v []StructuresIpInfo)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NetworkGetNetworkInfoResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetAdditionalIp

`func (o *NetworkGetNetworkInfoResponse) GetAdditionalIp() []StructuresAdditionalIpInfo`

GetAdditionalIp returns the AdditionalIp field if non-nil, zero value otherwise.

### GetAdditionalIpOk

`func (o *NetworkGetNetworkInfoResponse) GetAdditionalIpOk() (*[]StructuresAdditionalIpInfo, bool)`

GetAdditionalIpOk returns a tuple with the AdditionalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalIp

`func (o *NetworkGetNetworkInfoResponse) SetAdditionalIp(v []StructuresAdditionalIpInfo)`

SetAdditionalIp sets AdditionalIp field to given value.

### HasAdditionalIp

`func (o *NetworkGetNetworkInfoResponse) HasAdditionalIp() bool`

HasAdditionalIp returns a boolean if a field has been set.

### GetPrivateNetwork

`func (o *NetworkGetNetworkInfoResponse) GetPrivateNetwork() []StructuresPrivateNetwork`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *NetworkGetNetworkInfoResponse) GetPrivateNetworkOk() (*[]StructuresPrivateNetwork, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *NetworkGetNetworkInfoResponse) SetPrivateNetwork(v []StructuresPrivateNetwork)`

SetPrivateNetwork sets PrivateNetwork field to given value.

### HasPrivateNetwork

`func (o *NetworkGetNetworkInfoResponse) HasPrivateNetwork() bool`

HasPrivateNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


