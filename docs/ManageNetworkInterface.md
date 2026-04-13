# ManageNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Public** | Pointer to [**ManagePublicInterface**](ManagePublicInterface.md) |  | [optional] 
**Private** | Pointer to [**ManagePrivateInterface**](ManagePrivateInterface.md) |  | [optional] 

## Methods

### NewManageNetworkInterface

`func NewManageNetworkInterface() *ManageNetworkInterface`

NewManageNetworkInterface instantiates a new ManageNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageNetworkInterfaceWithDefaults

`func NewManageNetworkInterfaceWithDefaults() *ManageNetworkInterface`

NewManageNetworkInterfaceWithDefaults instantiates a new ManageNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublic

`func (o *ManageNetworkInterface) GetPublic() ManagePublicInterface`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *ManageNetworkInterface) GetPublicOk() (*ManagePublicInterface, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *ManageNetworkInterface) SetPublic(v ManagePublicInterface)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *ManageNetworkInterface) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetPrivate

`func (o *ManageNetworkInterface) GetPrivate() ManagePrivateInterface`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ManageNetworkInterface) GetPrivateOk() (*ManagePrivateInterface, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ManageNetworkInterface) SetPrivate(v ManagePrivateInterface)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ManageNetworkInterface) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


