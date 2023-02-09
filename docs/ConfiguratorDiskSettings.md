# ConfiguratorDiskSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to [**ConfiguratorRange**](ConfiguratorRange.md) |  | [optional] 
**AvailableRange** | Pointer to [**ConfiguratorRange**](ConfiguratorRange.md) |  | [optional] 
**Step** | Pointer to **int32** |  | [optional] 

## Methods

### NewConfiguratorDiskSettings

`func NewConfiguratorDiskSettings() *ConfiguratorDiskSettings`

NewConfiguratorDiskSettings instantiates a new ConfiguratorDiskSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfiguratorDiskSettingsWithDefaults

`func NewConfiguratorDiskSettingsWithDefaults() *ConfiguratorDiskSettings`

NewConfiguratorDiskSettingsWithDefaults instantiates a new ConfiguratorDiskSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *ConfiguratorDiskSettings) GetRange() ConfiguratorRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *ConfiguratorDiskSettings) GetRangeOk() (*ConfiguratorRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *ConfiguratorDiskSettings) SetRange(v ConfiguratorRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *ConfiguratorDiskSettings) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetAvailableRange

`func (o *ConfiguratorDiskSettings) GetAvailableRange() ConfiguratorRange`

GetAvailableRange returns the AvailableRange field if non-nil, zero value otherwise.

### GetAvailableRangeOk

`func (o *ConfiguratorDiskSettings) GetAvailableRangeOk() (*ConfiguratorRange, bool)`

GetAvailableRangeOk returns a tuple with the AvailableRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableRange

`func (o *ConfiguratorDiskSettings) SetAvailableRange(v ConfiguratorRange)`

SetAvailableRange sets AvailableRange field to given value.

### HasAvailableRange

`func (o *ConfiguratorDiskSettings) HasAvailableRange() bool`

HasAvailableRange returns a boolean if a field has been set.

### GetStep

`func (o *ConfiguratorDiskSettings) GetStep() int32`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *ConfiguratorDiskSettings) GetStepOk() (*int32, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *ConfiguratorDiskSettings) SetStep(v int32)`

SetStep sets Step field to given value.

### HasStep

`func (o *ConfiguratorDiskSettings) HasStep() bool`

HasStep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


