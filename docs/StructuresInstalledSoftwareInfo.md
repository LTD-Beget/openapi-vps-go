# StructuresInstalledSoftwareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**FieldValue** | Pointer to [**[]StructuresInstalledSoftwareInfoFieldValue**](StructuresInstalledSoftwareInfoFieldValue.md) |  | [optional] 
**Metadata** | Pointer to [**StructuresSoftwareMetadata**](StructuresSoftwareMetadata.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DescriptionEn** | Pointer to **string** |  | [optional] 
**Category** | Pointer to [**[]StructuresSoftwareCategory**](StructuresSoftwareCategory.md) |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 

## Methods

### NewStructuresInstalledSoftwareInfo

`func NewStructuresInstalledSoftwareInfo() *StructuresInstalledSoftwareInfo`

NewStructuresInstalledSoftwareInfo instantiates a new StructuresInstalledSoftwareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuresInstalledSoftwareInfoWithDefaults

`func NewStructuresInstalledSoftwareInfoWithDefaults() *StructuresInstalledSoftwareInfo`

NewStructuresInstalledSoftwareInfoWithDefaults instantiates a new StructuresInstalledSoftwareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StructuresInstalledSoftwareInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StructuresInstalledSoftwareInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StructuresInstalledSoftwareInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StructuresInstalledSoftwareInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *StructuresInstalledSoftwareInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *StructuresInstalledSoftwareInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *StructuresInstalledSoftwareInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *StructuresInstalledSoftwareInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetVersion

`func (o *StructuresInstalledSoftwareInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StructuresInstalledSoftwareInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StructuresInstalledSoftwareInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StructuresInstalledSoftwareInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAddress

`func (o *StructuresInstalledSoftwareInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *StructuresInstalledSoftwareInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *StructuresInstalledSoftwareInfo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *StructuresInstalledSoftwareInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetStatus

`func (o *StructuresInstalledSoftwareInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StructuresInstalledSoftwareInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StructuresInstalledSoftwareInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StructuresInstalledSoftwareInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFieldValue

`func (o *StructuresInstalledSoftwareInfo) GetFieldValue() []StructuresInstalledSoftwareInfoFieldValue`

GetFieldValue returns the FieldValue field if non-nil, zero value otherwise.

### GetFieldValueOk

`func (o *StructuresInstalledSoftwareInfo) GetFieldValueOk() (*[]StructuresInstalledSoftwareInfoFieldValue, bool)`

GetFieldValueOk returns a tuple with the FieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValue

`func (o *StructuresInstalledSoftwareInfo) SetFieldValue(v []StructuresInstalledSoftwareInfoFieldValue)`

SetFieldValue sets FieldValue field to given value.

### HasFieldValue

`func (o *StructuresInstalledSoftwareInfo) HasFieldValue() bool`

HasFieldValue returns a boolean if a field has been set.

### GetMetadata

`func (o *StructuresInstalledSoftwareInfo) GetMetadata() StructuresSoftwareMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StructuresInstalledSoftwareInfo) GetMetadataOk() (*StructuresSoftwareMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StructuresInstalledSoftwareInfo) SetMetadata(v StructuresSoftwareMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StructuresInstalledSoftwareInfo) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDescription

`func (o *StructuresInstalledSoftwareInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StructuresInstalledSoftwareInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StructuresInstalledSoftwareInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StructuresInstalledSoftwareInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionEn

`func (o *StructuresInstalledSoftwareInfo) GetDescriptionEn() string`

GetDescriptionEn returns the DescriptionEn field if non-nil, zero value otherwise.

### GetDescriptionEnOk

`func (o *StructuresInstalledSoftwareInfo) GetDescriptionEnOk() (*string, bool)`

GetDescriptionEnOk returns a tuple with the DescriptionEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionEn

`func (o *StructuresInstalledSoftwareInfo) SetDescriptionEn(v string)`

SetDescriptionEn sets DescriptionEn field to given value.

### HasDescriptionEn

`func (o *StructuresInstalledSoftwareInfo) HasDescriptionEn() bool`

HasDescriptionEn returns a boolean if a field has been set.

### GetCategory

`func (o *StructuresInstalledSoftwareInfo) GetCategory() []StructuresSoftwareCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *StructuresInstalledSoftwareInfo) GetCategoryOk() (*[]StructuresSoftwareCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *StructuresInstalledSoftwareInfo) SetCategory(v []StructuresSoftwareCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *StructuresInstalledSoftwareInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetSlug

`func (o *StructuresInstalledSoftwareInfo) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *StructuresInstalledSoftwareInfo) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *StructuresInstalledSoftwareInfo) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *StructuresInstalledSoftwareInfo) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


