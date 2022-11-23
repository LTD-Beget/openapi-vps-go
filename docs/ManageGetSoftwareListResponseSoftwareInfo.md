# ManageGetSoftwareListResponseSoftwareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DescriptionVersion** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **[]int32** |  | [optional] 
**Variable** | Pointer to **[]string** |  | [optional] 
**Requirements** | Pointer to [**ManageGetSoftwareListResponseSoftwareInfoRequirements**](ManageGetSoftwareListResponseSoftwareInfoRequirements.md) |  | [optional] 

## Methods

### NewManageGetSoftwareListResponseSoftwareInfo

`func NewManageGetSoftwareListResponseSoftwareInfo() *ManageGetSoftwareListResponseSoftwareInfo`

NewManageGetSoftwareListResponseSoftwareInfo instantiates a new ManageGetSoftwareListResponseSoftwareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageGetSoftwareListResponseSoftwareInfoWithDefaults

`func NewManageGetSoftwareListResponseSoftwareInfoWithDefaults() *ManageGetSoftwareListResponseSoftwareInfo`

NewManageGetSoftwareListResponseSoftwareInfoWithDefaults instantiates a new ManageGetSoftwareListResponseSoftwareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManageGetSoftwareListResponseSoftwareInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManageGetSoftwareListResponseSoftwareInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManageGetSoftwareListResponseSoftwareInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManageGetSoftwareListResponseSoftwareInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ManageGetSoftwareListResponseSoftwareInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ManageGetSoftwareListResponseSoftwareInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetVersion

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ManageGetSoftwareListResponseSoftwareInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ManageGetSoftwareListResponseSoftwareInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManageGetSoftwareListResponseSoftwareInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManageGetSoftwareListResponseSoftwareInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManageGetSoftwareListResponseSoftwareInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManageGetSoftwareListResponseSoftwareInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionVersion

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetDescriptionVersion() string`

GetDescriptionVersion returns the DescriptionVersion field if non-nil, zero value otherwise.

### GetDescriptionVersionOk

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetDescriptionVersionOk() (*string, bool)`

GetDescriptionVersionOk returns a tuple with the DescriptionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionVersion

`func (o *ManageGetSoftwareListResponseSoftwareInfo) SetDescriptionVersion(v string)`

SetDescriptionVersion sets DescriptionVersion field to given value.

### HasDescriptionVersion

`func (o *ManageGetSoftwareListResponseSoftwareInfo) HasDescriptionVersion() bool`

HasDescriptionVersion returns a boolean if a field has been set.

### GetImageId

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetImageId() []int32`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetImageIdOk() (*[]int32, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ManageGetSoftwareListResponseSoftwareInfo) SetImageId(v []int32)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *ManageGetSoftwareListResponseSoftwareInfo) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetVariable

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetVariable() []string`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetVariableOk() (*[]string, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *ManageGetSoftwareListResponseSoftwareInfo) SetVariable(v []string)`

SetVariable sets Variable field to given value.

### HasVariable

`func (o *ManageGetSoftwareListResponseSoftwareInfo) HasVariable() bool`

HasVariable returns a boolean if a field has been set.

### GetRequirements

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetRequirements() ManageGetSoftwareListResponseSoftwareInfoRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ManageGetSoftwareListResponseSoftwareInfo) GetRequirementsOk() (*ManageGetSoftwareListResponseSoftwareInfoRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ManageGetSoftwareListResponseSoftwareInfo) SetRequirements(v ManageGetSoftwareListResponseSoftwareInfoRequirements)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *ManageGetSoftwareListResponseSoftwareInfo) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


