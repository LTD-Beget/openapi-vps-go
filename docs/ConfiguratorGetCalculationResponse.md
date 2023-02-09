# ConfiguratorGetCalculationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to [**ConfiguratorGetCalculationResponseSuccess**](ConfiguratorGetCalculationResponseSuccess.md) |  | [optional] 
**Error** | Pointer to [**ConfiguratorGetCalculationResponseError**](ConfiguratorGetCalculationResponseError.md) |  | [optional] 

## Methods

### NewConfiguratorGetCalculationResponse

`func NewConfiguratorGetCalculationResponse() *ConfiguratorGetCalculationResponse`

NewConfiguratorGetCalculationResponse instantiates a new ConfiguratorGetCalculationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfiguratorGetCalculationResponseWithDefaults

`func NewConfiguratorGetCalculationResponseWithDefaults() *ConfiguratorGetCalculationResponse`

NewConfiguratorGetCalculationResponseWithDefaults instantiates a new ConfiguratorGetCalculationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ConfiguratorGetCalculationResponse) GetSuccess() ConfiguratorGetCalculationResponseSuccess`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ConfiguratorGetCalculationResponse) GetSuccessOk() (*ConfiguratorGetCalculationResponseSuccess, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ConfiguratorGetCalculationResponse) SetSuccess(v ConfiguratorGetCalculationResponseSuccess)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ConfiguratorGetCalculationResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetError

`func (o *ConfiguratorGetCalculationResponse) GetError() ConfiguratorGetCalculationResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ConfiguratorGetCalculationResponse) GetErrorOk() (*ConfiguratorGetCalculationResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ConfiguratorGetCalculationResponse) SetError(v ConfiguratorGetCalculationResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *ConfiguratorGetCalculationResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


