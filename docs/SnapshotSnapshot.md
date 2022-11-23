# SnapshotSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**VpsId** | Pointer to **string** |  | [optional] 
**VpsName** | Pointer to **string** |  | [optional] 
**DateCreate** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Configuration** | Pointer to [**SnapshotRequiredConfiguration**](SnapshotRequiredConfiguration.md) |  | [optional] 
**PriceDay** | Pointer to **float64** |  | [optional] 
**InstalledSoftware** | Pointer to [**StructuresInstalledSoftwareInfo**](StructuresInstalledSoftwareInfo.md) |  | [optional] 

## Methods

### NewSnapshotSnapshot

`func NewSnapshotSnapshot() *SnapshotSnapshot`

NewSnapshotSnapshot instantiates a new SnapshotSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotSnapshotWithDefaults

`func NewSnapshotSnapshotWithDefaults() *SnapshotSnapshot`

NewSnapshotSnapshotWithDefaults instantiates a new SnapshotSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnapshotSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotSnapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SnapshotSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVpsId

`func (o *SnapshotSnapshot) GetVpsId() string`

GetVpsId returns the VpsId field if non-nil, zero value otherwise.

### GetVpsIdOk

`func (o *SnapshotSnapshot) GetVpsIdOk() (*string, bool)`

GetVpsIdOk returns a tuple with the VpsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpsId

`func (o *SnapshotSnapshot) SetVpsId(v string)`

SetVpsId sets VpsId field to given value.

### HasVpsId

`func (o *SnapshotSnapshot) HasVpsId() bool`

HasVpsId returns a boolean if a field has been set.

### GetVpsName

`func (o *SnapshotSnapshot) GetVpsName() string`

GetVpsName returns the VpsName field if non-nil, zero value otherwise.

### GetVpsNameOk

`func (o *SnapshotSnapshot) GetVpsNameOk() (*string, bool)`

GetVpsNameOk returns a tuple with the VpsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpsName

`func (o *SnapshotSnapshot) SetVpsName(v string)`

SetVpsName sets VpsName field to given value.

### HasVpsName

`func (o *SnapshotSnapshot) HasVpsName() bool`

HasVpsName returns a boolean if a field has been set.

### GetDateCreate

`func (o *SnapshotSnapshot) GetDateCreate() string`

GetDateCreate returns the DateCreate field if non-nil, zero value otherwise.

### GetDateCreateOk

`func (o *SnapshotSnapshot) GetDateCreateOk() (*string, bool)`

GetDateCreateOk returns a tuple with the DateCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreate

`func (o *SnapshotSnapshot) SetDateCreate(v string)`

SetDateCreate sets DateCreate field to given value.

### HasDateCreate

`func (o *SnapshotSnapshot) HasDateCreate() bool`

HasDateCreate returns a boolean if a field has been set.

### GetSize

`func (o *SnapshotSnapshot) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SnapshotSnapshot) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SnapshotSnapshot) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *SnapshotSnapshot) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *SnapshotSnapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SnapshotSnapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SnapshotSnapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SnapshotSnapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDescription

`func (o *SnapshotSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnapshotSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnapshotSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SnapshotSnapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfiguration

`func (o *SnapshotSnapshot) GetConfiguration() SnapshotRequiredConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SnapshotSnapshot) GetConfigurationOk() (*SnapshotRequiredConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SnapshotSnapshot) SetConfiguration(v SnapshotRequiredConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *SnapshotSnapshot) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetPriceDay

`func (o *SnapshotSnapshot) GetPriceDay() float64`

GetPriceDay returns the PriceDay field if non-nil, zero value otherwise.

### GetPriceDayOk

`func (o *SnapshotSnapshot) GetPriceDayOk() (*float64, bool)`

GetPriceDayOk returns a tuple with the PriceDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDay

`func (o *SnapshotSnapshot) SetPriceDay(v float64)`

SetPriceDay sets PriceDay field to given value.

### HasPriceDay

`func (o *SnapshotSnapshot) HasPriceDay() bool`

HasPriceDay returns a boolean if a field has been set.

### GetInstalledSoftware

`func (o *SnapshotSnapshot) GetInstalledSoftware() StructuresInstalledSoftwareInfo`

GetInstalledSoftware returns the InstalledSoftware field if non-nil, zero value otherwise.

### GetInstalledSoftwareOk

`func (o *SnapshotSnapshot) GetInstalledSoftwareOk() (*StructuresInstalledSoftwareInfo, bool)`

GetInstalledSoftwareOk returns a tuple with the InstalledSoftware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledSoftware

`func (o *SnapshotSnapshot) SetInstalledSoftware(v StructuresInstalledSoftwareInfo)`

SetInstalledSoftware sets InstalledSoftware field to given value.

### HasInstalledSoftware

`func (o *SnapshotSnapshot) HasInstalledSoftware() bool`

HasInstalledSoftware returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


