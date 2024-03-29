# ManageChangeConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ConfigurationId** | Pointer to **string** |  | [optional] 
**ConfigurationParams** | Pointer to [**StructuresConfigurationParams**](StructuresConfigurationParams.md) |  | [optional] 

## Methods

### NewManageChangeConfigurationRequest

`func NewManageChangeConfigurationRequest() *ManageChangeConfigurationRequest`

NewManageChangeConfigurationRequest instantiates a new ManageChangeConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageChangeConfigurationRequestWithDefaults

`func NewManageChangeConfigurationRequestWithDefaults() *ManageChangeConfigurationRequest`

NewManageChangeConfigurationRequestWithDefaults instantiates a new ManageChangeConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManageChangeConfigurationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManageChangeConfigurationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManageChangeConfigurationRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ManageChangeConfigurationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfigurationId

`func (o *ManageChangeConfigurationRequest) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *ManageChangeConfigurationRequest) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *ManageChangeConfigurationRequest) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *ManageChangeConfigurationRequest) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### GetConfigurationParams

`func (o *ManageChangeConfigurationRequest) GetConfigurationParams() StructuresConfigurationParams`

GetConfigurationParams returns the ConfigurationParams field if non-nil, zero value otherwise.

### GetConfigurationParamsOk

`func (o *ManageChangeConfigurationRequest) GetConfigurationParamsOk() (*StructuresConfigurationParams, bool)`

GetConfigurationParamsOk returns a tuple with the ConfigurationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationParams

`func (o *ManageChangeConfigurationRequest) SetConfigurationParams(v StructuresConfigurationParams)`

SetConfigurationParams sets ConfigurationParams field to given value.

### HasConfigurationParams

`func (o *ManageChangeConfigurationRequest) HasConfigurationParams() bool`

HasConfigurationParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


