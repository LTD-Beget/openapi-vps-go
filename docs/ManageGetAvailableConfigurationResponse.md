# ManageGetAvailableConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configurations** | Pointer to [**[]ManageVpsConfiguration**](ManageVpsConfiguration.md) |  | [optional] 
**ConfigurationGroups** | Pointer to [**[]ManageConfigurationGroup**](ManageConfigurationGroup.md) |  | [optional] 

## Methods

### NewManageGetAvailableConfigurationResponse

`func NewManageGetAvailableConfigurationResponse() *ManageGetAvailableConfigurationResponse`

NewManageGetAvailableConfigurationResponse instantiates a new ManageGetAvailableConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageGetAvailableConfigurationResponseWithDefaults

`func NewManageGetAvailableConfigurationResponseWithDefaults() *ManageGetAvailableConfigurationResponse`

NewManageGetAvailableConfigurationResponseWithDefaults instantiates a new ManageGetAvailableConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurations

`func (o *ManageGetAvailableConfigurationResponse) GetConfigurations() []ManageVpsConfiguration`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *ManageGetAvailableConfigurationResponse) GetConfigurationsOk() (*[]ManageVpsConfiguration, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *ManageGetAvailableConfigurationResponse) SetConfigurations(v []ManageVpsConfiguration)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *ManageGetAvailableConfigurationResponse) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetConfigurationGroups

`func (o *ManageGetAvailableConfigurationResponse) GetConfigurationGroups() []ManageConfigurationGroup`

GetConfigurationGroups returns the ConfigurationGroups field if non-nil, zero value otherwise.

### GetConfigurationGroupsOk

`func (o *ManageGetAvailableConfigurationResponse) GetConfigurationGroupsOk() (*[]ManageConfigurationGroup, bool)`

GetConfigurationGroupsOk returns a tuple with the ConfigurationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationGroups

`func (o *ManageGetAvailableConfigurationResponse) SetConfigurationGroups(v []ManageConfigurationGroup)`

SetConfigurationGroups sets ConfigurationGroups field to given value.

### HasConfigurationGroups

`func (o *ManageGetAvailableConfigurationResponse) HasConfigurationGroups() bool`

HasConfigurationGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


