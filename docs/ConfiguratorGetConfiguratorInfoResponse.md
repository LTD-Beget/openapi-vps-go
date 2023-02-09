# ConfiguratorGetConfiguratorInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to [**ConfiguratorConfiguratorSettings**](ConfiguratorConfiguratorSettings.md) |  | [optional] 
**IsAvailable** | Pointer to **bool** |  | [optional] 

## Methods

### NewConfiguratorGetConfiguratorInfoResponse

`func NewConfiguratorGetConfiguratorInfoResponse() *ConfiguratorGetConfiguratorInfoResponse`

NewConfiguratorGetConfiguratorInfoResponse instantiates a new ConfiguratorGetConfiguratorInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfiguratorGetConfiguratorInfoResponseWithDefaults

`func NewConfiguratorGetConfiguratorInfoResponseWithDefaults() *ConfiguratorGetConfiguratorInfoResponse`

NewConfiguratorGetConfiguratorInfoResponseWithDefaults instantiates a new ConfiguratorGetConfiguratorInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *ConfiguratorGetConfiguratorInfoResponse) GetSettings() ConfiguratorConfiguratorSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ConfiguratorGetConfiguratorInfoResponse) GetSettingsOk() (*ConfiguratorConfiguratorSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ConfiguratorGetConfiguratorInfoResponse) SetSettings(v ConfiguratorConfiguratorSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ConfiguratorGetConfiguratorInfoResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetIsAvailable

`func (o *ConfiguratorGetConfiguratorInfoResponse) GetIsAvailable() bool`

GetIsAvailable returns the IsAvailable field if non-nil, zero value otherwise.

### GetIsAvailableOk

`func (o *ConfiguratorGetConfiguratorInfoResponse) GetIsAvailableOk() (*bool, bool)`

GetIsAvailableOk returns a tuple with the IsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailable

`func (o *ConfiguratorGetConfiguratorInfoResponse) SetIsAvailable(v bool)`

SetIsAvailable sets IsAvailable field to given value.

### HasIsAvailable

`func (o *ConfiguratorGetConfiguratorInfoResponse) HasIsAvailable() bool`

HasIsAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


