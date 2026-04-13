# ManagePublicInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocateNew** | Pointer to **map[string]interface{}** |  | [optional] 
**Existing** | Pointer to [**ManagePublicInterfaceUseExistingPublicIp**](ManagePublicInterfaceUseExistingPublicIp.md) |  | [optional] 

## Methods

### NewManagePublicInterface

`func NewManagePublicInterface() *ManagePublicInterface`

NewManagePublicInterface instantiates a new ManagePublicInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagePublicInterfaceWithDefaults

`func NewManagePublicInterfaceWithDefaults() *ManagePublicInterface`

NewManagePublicInterfaceWithDefaults instantiates a new ManagePublicInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocateNew

`func (o *ManagePublicInterface) GetAllocateNew() map[string]interface{}`

GetAllocateNew returns the AllocateNew field if non-nil, zero value otherwise.

### GetAllocateNewOk

`func (o *ManagePublicInterface) GetAllocateNewOk() (*map[string]interface{}, bool)`

GetAllocateNewOk returns a tuple with the AllocateNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocateNew

`func (o *ManagePublicInterface) SetAllocateNew(v map[string]interface{})`

SetAllocateNew sets AllocateNew field to given value.

### HasAllocateNew

`func (o *ManagePublicInterface) HasAllocateNew() bool`

HasAllocateNew returns a boolean if a field has been set.

### GetExisting

`func (o *ManagePublicInterface) GetExisting() ManagePublicInterfaceUseExistingPublicIp`

GetExisting returns the Existing field if non-nil, zero value otherwise.

### GetExistingOk

`func (o *ManagePublicInterface) GetExistingOk() (*ManagePublicInterfaceUseExistingPublicIp, bool)`

GetExistingOk returns a tuple with the Existing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExisting

`func (o *ManagePublicInterface) SetExisting(v ManagePublicInterfaceUseExistingPublicIp)`

SetExisting sets Existing field to given value.

### HasExisting

`func (o *ManagePublicInterface) HasExisting() bool`

HasExisting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


