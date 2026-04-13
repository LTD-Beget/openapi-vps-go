# ManagePrivateInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** |  | [optional] 
**AllocateNew** | Pointer to **map[string]interface{}** |  | [optional] 
**Specific** | Pointer to [**ManagePrivateInterfaceAllocateSpecificPrivateIp**](ManagePrivateInterfaceAllocateSpecificPrivateIp.md) |  | [optional] 

## Methods

### NewManagePrivateInterface

`func NewManagePrivateInterface() *ManagePrivateInterface`

NewManagePrivateInterface instantiates a new ManagePrivateInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagePrivateInterfaceWithDefaults

`func NewManagePrivateInterfaceWithDefaults() *ManagePrivateInterface`

NewManagePrivateInterfaceWithDefaults instantiates a new ManagePrivateInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *ManagePrivateInterface) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *ManagePrivateInterface) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *ManagePrivateInterface) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *ManagePrivateInterface) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetAllocateNew

`func (o *ManagePrivateInterface) GetAllocateNew() map[string]interface{}`

GetAllocateNew returns the AllocateNew field if non-nil, zero value otherwise.

### GetAllocateNewOk

`func (o *ManagePrivateInterface) GetAllocateNewOk() (*map[string]interface{}, bool)`

GetAllocateNewOk returns a tuple with the AllocateNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocateNew

`func (o *ManagePrivateInterface) SetAllocateNew(v map[string]interface{})`

SetAllocateNew sets AllocateNew field to given value.

### HasAllocateNew

`func (o *ManagePrivateInterface) HasAllocateNew() bool`

HasAllocateNew returns a boolean if a field has been set.

### GetSpecific

`func (o *ManagePrivateInterface) GetSpecific() ManagePrivateInterfaceAllocateSpecificPrivateIp`

GetSpecific returns the Specific field if non-nil, zero value otherwise.

### GetSpecificOk

`func (o *ManagePrivateInterface) GetSpecificOk() (*ManagePrivateInterfaceAllocateSpecificPrivateIp, bool)`

GetSpecificOk returns a tuple with the Specific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecific

`func (o *ManagePrivateInterface) SetSpecific(v ManagePrivateInterfaceAllocateSpecificPrivateIp)`

SetSpecific sets Specific field to given value.

### HasSpecific

`func (o *ManagePrivateInterface) HasSpecific() bool`

HasSpecific returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


